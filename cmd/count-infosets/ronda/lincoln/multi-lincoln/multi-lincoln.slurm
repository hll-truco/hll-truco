#!/bin/bash

# make sure the logs directory `--output` and `--output` exist before execution

# when using standard args launch it as:
# `sbatch -J hllworkers ~/Workspace/facu/hll-truco/sbatch/http/workers.sbatch root [<hash> <deck> <abs> <report> <limit>]`

#SBATCH --job-name=multi-lincoln
#SBATCH --nodes=1 # nodes to use in total, per job or subjob
#SBATCH --ntasks=1 # when doing `srun` launch `ntaks` parallel jobs across `nodes`
#SBATCH --cpus-per-task=2
#SBATCH --mem=499G
#SBATCH --time=5-00:00:00
#SBATCH --partition=normal
#SBATCH --qos=bigmem
#SBATCH --output=/clusteruy/home/juan.filevich/batches/out/%x.%j.out
#SBATCH --error=/clusteruy/home/juan.filevich/batches/out/%x.%j.out
#SBATCH --mail-type=BEGIN,END,FAIL # options: NONE, BEGIN, END, FAIL, REQUEUE, ALL
#SBATCH --mail-user=juan.filevich@fing.edu.uy

# every (sub)job will execute this from here
# all (sub)job will share the same args ${@}

# dump args
printf "starts: $(date)\n"
echo "args: ${@}"

base=$HOME/Workspace/facu/hll-truco/hll-truco # edit this!
cmd=cmd/count-infosets/ronda/lincoln/multi-lincoln/main.go

mark=2818000000
capture=3382000000
# srun will run this cmd `nodes` times in parallel
cd ${base}

# compile
go build -ldflags="-w -s" -o bin/multi_lincoln $cmd

# run

srun --export=ALL,GOMEMLIMIT=496GiB,TIME_LIMIT_MARK=48960,TIME_LIMIT_CAPTURE=58752 \
  bin/multi_lincoln \
    -hash=sha512 \
    -abs=null \
    -info=InfosetPartidaXXLarge \
    -report=3600 \
    -marked=$mark \
    -captured=$capture \
    -mazo=false

printf "done: $(date)\n"
