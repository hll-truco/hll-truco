#!/bin/bash

# make sure the logs directory `--output` and `--output` exist before execution

# when using standard args launch it as:
# `sbatch -J hllworkers ~/Workspace/facu/hll-truco/sbatch/http/workers.sbatch [<hash> <deck> <abs> <report> <limit>]`

#SBATCH --job-name=hllworkers
#SBATCH --nodes=4 # nodes to use in total, per job or subjob
#SBATCH --ntasks=64 # when doing `srun` launch `ntaks` parallel jobs across `nodes`
#SBATCH --cpus-per-task=1
#SBATCH --mem-per-cpu=128M
#SBATCH --time=0-00:15:00
#SBATCH --partition=normal
#SBATCH --qos=normal
#SBATCH --output=/clusteruy/home/juan.filevich/batches/out/hll-http/big/w64/%x.%j.out
#SBATCH --error=/clusteruy/home/juan.filevich/batches/out/hll-http/big/w64/%x.%j.out
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

task_log=$HOME/batches/out/hll-http/big/w64/%x.%j_%t.out # edit this!
project=$HOME/Workspace/facu/hll-truco/hll-truco # edit this!
main=$project/cmd/count-infosets-hll-dist-http/worker/main.go
worker_name="bot_${SLURM_ARRAY_JOB_ID}_${SLURM_JOB_ID}_${SLURM_ARRAY_TASK_ID}"

rootname=hllroot # edit this!
root=http://$(squeue -u $(whoami) --name=${rootname} --states=R -h -o "%N:%k")

hash=${1:-'sha3'}
deck=${2:-'14'}
abs=${3:-'null'}
report=${4:-'1'}
limit=${5:-'600'} # 600 seconds = 10 min

echo "using: root:${root} worker:${worker_name} logs:${task_log} hash:${hash} deck:${deck} abs:${abs} report:${report} limit:${limit}"

# srun will run this cmd `nodes` times in parallel
cd ${project}
srun -l --output=${task_log} --error=${task_log} \
  go run ${main} \
    -hash=${hash} \
    -deck=${deck} \
    -abs=${abs} \
    -report=${report} \
    -limit=${limit} \
    -root=${root}

# shutdown root server
curl -X GET ${root}/exit

hr
printf "done: $(date)\n"
