#!/bin/bash

# make sure the logs directory `--output` and `--output` exist before execution

# when using standard args launch it as:
# `sbatch -J hllworkers ~/Workspace/facu/hll-truco/sbatch/http/workers.sbatch root [<hash> <deck> <abs> <report> <limit>]`

#SBATCH --job-name=hllworkers
#SBATCH --nodes=8 # nodes to use in total, per job or subjob
#SBATCH --ntasks=99 # when doing `srun` launch `ntaks` parallel jobs across `nodes`
#SBATCH --cpus-per-task=1
#SBATCH --mem-per-cpu=128M
#SBATCH --time=0-00:15:00
#SBATCH --partition=besteffort
#SBATCH --qos=besteffort
#SBATCH --output=/clusteruy/home/juan.filevich/batches/out/hll-http/2p/E1P40AnullIipxxlW256/%x.%j.out
#SBATCH --error=/clusteruy/home/juan.filevich/batches/out/hll-http/2p/E1P40AnullIipxxlW256/%x.%j.out
#SBATCH --mail-type=NONE # options: NONE, BEGIN, END, FAIL, REQUEUE, ALL
#SBATCH --mail-user=juan.filevich@fing.edu.uy

# every (sub)job will execute this from here
# all (sub)job will share the same args ${@}

hr (){
printf '%*s\n' 80 | tr ' ' '-'
}

# dump args
printf "starts: $(date)\n"
echo "args: ${@}"
hr

task_log=$HOME/batches/out/hll-http/2p/E1P40AnullIipxxlW256/%x.%j_%t.out # edit this!
project=$HOME/Workspace/facu/hll-truco/hll-truco # edit this!
main=$project/cmd/count-infosets-hll-dist-http/worker/main.go
worker_name="bot_${SLURM_ARRAY_JOB_ID}_${SLURM_JOB_ID}_${SLURM_ARRAY_TASK_ID}"

n=${1:-'2'} # Number of players <2,4,6>
envido=${2:-'1'} # Envido limit (default 1)
pts=${3:-'40'} # Game points limit
absId=${4:-'null'} # Abstractor ID <a1, b, a2, a3, null>
infoset=${5:-'InfosetRondaBase'} # Infoset impl. to use
hash=${6:-'sha3'} # Infoset hashing function
limit=${7:-'600'} # 600 seconds = 10 min # Run time limit (in seconds) (default 1m)
report=${8:-'10'} # Delta (in seconds) for printing log msgs
root=${9:-'http://localhost:8080'} # HTTP root server
precision=${10:-'16'} # HLL precision

echo "using: worker:${worker_name} logs:${task_log} n:${n} envido:${envido} \
pts:${pts} absId:${absId} infoset:${infoset} hash:${hash} limit:${limit} \
report:${report} root:${root}"

# srun will run this cmd `nodes` times in parallel
cd ${project}
# compile
go build -ldflags="-w -s" -o bin/http/worker cmd/count-infosets-hll-dist-http/worker/main.go
# run
srun -l --output=${task_log} --error=${task_log} \
  bin/http/worker \
    -n=${n} \
    -e=${envido} \
    -p=${pts} \
    -abs=${absId} \
    -info=${infoset} \
    -hash=${hash} \
    -limit=${limit} \
    -report=${report} \
    -root=${root} \
    -precision=${precision}

# shutdown root server
curl -X GET ${root}/exit

hr
printf "done: $(date)\n"
