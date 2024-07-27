#!/bin/bash

#SBATCH --job-name=ctrlr
#SBATCH --nodes=1 # nodes to use in total, per job or subjob
#SBATCH --ntasks=1
#SBATCH --cpus-per-task=1
#SBATCH --mem-per-cpu=128M
#SBATCH --time=0-00:30:00
#SBATCH --partition=normal
#SBATCH --qos=normal
#SBATCH --output=/clusteruy/home/juan.filevich/batches/out/hll-http/ctrlr/%x.%j.out
#SBATCH --error=/clusteruy/home/juan.filevich/batches/out/hll-http/ctrlr/%x.%j.out
#SBATCH --mail-type=NONE # options: NONE, BEGIN, END, FAIL, REQUEUE, ALL
#SBATCH --mail-user=juan.filevich@fing.edu.uy 

# WARNING: Stop all jobs but this one (this prevents besteffor jobs from running)
scancel $(squeue -u $(whoami) -h -o "%A" | grep -v $SLURM_JOB_ID)

# horizontal rule
hr() { printf '%*s\n' 80 | tr ' ' '-'; }

# meta args
iter=${1:-1} # number of iter to run
checkpoint_every=${2:-2} # number of iters between each checkpoint

# dump args
printf "starts: $(date)\n"
echo "args: ${@}" && hr

# Check if the number of iter is greater than zero
if [[ ${iter} -le 0 ]]; then
  printf "Done with all iter: $(date)\n"
  exit 0
fi

# new state saved by root before exiting
save_dir=$HOME/batches/sal/hll-http/2p/E1P20AnullIipxxlW256
latest="${save_dir}/latest_prec6_f1024.json"

# make a checkpoint if necessary
if (( $iter % $checkpoint_every == 0 )); then
  new_checkpoint="${latest%.json}_iter${iter}.json"
  cp "${latest}" "${new_checkpoint}"
  printf "Checkpoint ${new_checkpoint} saved at $(date)\n"
else
  printf "Skipping checkpoint creation\n"
fi

# 1. submit hll::root
# 2. submit hll::workers normal
# 3. submit hll::workers besteffort

# 1. submit hll::root
mkdir -p $HOME/batches/out/hll-http/2p/E1P20AnullIipxxlW256

job_name="hllroot_${iter}" && \
script=$HOME/Workspace/facu/hll-truco/hll-truco/sbatch/http/root.sbatch && \
time_limit="5-00:00:00" && \
# args
report=10 && \
precision=6 && \
resume=${latest} && \
save=${latest} && \
job_msg=$(
  sbatch \
    -J ${job_name} \
    --time="${time_limit}" \
    --begin=now+0 \
    "${script}" \
    $iter $checkpoint_every \
    $report $precision "${resume}" "${save}"
)

root_job_id=${job_msg##* }

# Wait until the status of the job is running
while true; do
  job_status=$(squeue -j ${root_job_id} -h -o %T)
  if [[ "$job_status" == "RUNNING" ]]; then
    echo "${job_name} is now running and its ID is ${root_job_id}"
    break
  fi
  sleep 1
done

sleep 5

# hll workers shared vars
root_host=$(squeue -u $(whoami) -j ${root_job_id} --states=R -h -o "%N:%k")

# hll workers args
n=2
envido=1
pts=20
absId=null
infoset=InfosetPartidaXXLarge
hash=sha3
limit=0 # <- redeclare later
report=10
root=http://${root_host}
precision=6
mazo=false
resume=${latest}

# 2. submit normal workers
job_name=hllworkers_normal
script=$HOME/Workspace/facu/hll-truco/hll-truco/sbatch/http/workers.normal.sbatch
time_limit="5-00:00:00"
limit=428400
sbatch \
  -J ${job_name} \
  --time="${time_limit}" \
  --dependency=after:${root_job_id} \
  "${script}" \
  ${n} ${envido} ${pts} ${absId} "${infoset}" ${hash} ${limit} ${report} "${root}" ${precision} ${mazo} "${resume}"

# 3. submit besteffort workers
job_name=hllworkers_effort
script=$HOME/Workspace/facu/hll-truco/hll-truco/sbatch/http/workers.sbatch
time_limit="5-00:00:00"
limit=428400 # 5d = 432000s -> 432000s - (1h = 3600s) = 428400s
sbatch \
  -J ${job_name} \
  --time="${time_limit}" \
  --dependency=after:${root_job_id} \
  "${script}" \
  ${n} ${envido} ${pts} ${absId} "${infoset}" ${hash} ${limit} ${report} "${root}" ${precision} ${mazo} "${resume}"

hr
printf "done submitting all jobs at $(date)\n"

# todos:
# fail if not enough args
# use signals instead of curl

# upload with
# `rsync -avz -e 'ssh -p 10022 -i ~/.ssh/id_rsa' --exclude '__pycache__/' --exclude '*.out' --exclude '*.log' --exclude '.git/' --exclude '*.png' /tmp/hll-truco 'cluster.uy:Workspace/facu/hll-truco/'`

# alternatively, you can query a job ID by its name using:
# rootname=hllroot && \
# rootid=$(squeue -u $(whoami) --name=${rootname} -h -o "%i") && \
# root_host=$(squeue -u $(whoami) --name=${rootname} --states=R -h -o "%N:%k")