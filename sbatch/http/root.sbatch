#!/bin/bash

# make sure the logs directory `--output` and `--output` exist before execution

#SBATCH --job-name=hllroot
#SBATCH --nodes=1 # nodes to use in total, per job or subjob
#SBATCH --ntasks=1
#SBATCH --cpus-per-task=1
#SBATCH --mem-per-cpu=128M
#SBATCH --time=0-00:15:00
#SBATCH --partition=normal
#SBATCH --qos=normal
#SBATCH --output=/clusteruy/home/juan.filevich/batches/out/hll-http/2p/E1P20AnullIipxxlW256/%x.%j.out
#SBATCH --error=/clusteruy/home/juan.filevich/batches/out/hll-http/2p/E1P20AnullIipxxlW256/%x.%j.out
#SBATCH --mail-type=FAIL # options: NONE, BEGIN, END, FAIL, REQUEUE, ALL
#SBATCH --mail-user=juan.filevich@fing.edu.uy 

hr (){
printf '%*s\n' 80 | tr ' ' '-'
}

# dump args
printf "starts: $(date)\n"
echo "args: ${@}"
hr

# pick a free port
PORT=$(python3 -c 'import socket; s=socket.socket(); s.bind(("", 0)); print(s.getsockname()[1]); s.close()')
echo "Starting srv on port ${PORT} at host $(hostname) for jobid:${SLURM_JOB_ID}"
scontrol update JobId="${SLURM_JOB_ID}" Comment="${PORT}"

project=$HOME/Workspace/facu/hll-truco/hll-truco
main=$project/cmd/count-infosets-hll-dist-http/root/main.go

# meta args
iter=${1:-1}
checkpoint_every=${2:-2} # number of iters between each checkpoint
# args
report=${3:-'10'} # Delta (in seconds) for printing log msgs
precision=${4:-'16'} # HLL precision
resume=${5:-''} # Resume from gob/json file
save=${6:-''} # save to gob/json file

echo "using main:${main} report:${report}"

cd ${project}
go run ${main} \
  -port=${PORT} \
  -report=${report} \
  -precision=${precision} \
  -resume="${resume}" \
  -save="${save}"

# or if you have an already compiled binary you could just:
# $project/bin/http/root -port=$PORT

hr
printf "done: $(date)\n"

sbatch $HOME/Workspace/facu/hll-truco/hll-truco/sbatch/http/ctrlr.sbatch $((iter - 1)) ${checkpoint_every}

printf "called ctrlr\n"
