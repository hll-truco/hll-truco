#! /bin/sh

mkdir -p ~/batches/out/hll-http/2p/E1P40AnullIipxxlW256 && \
sbatch \
    -J hllroot1 \
    --time=3-00:00:00 \
    --begin=now+0 \
    ~/Workspace/facu/hll-truco/hll-truco/sbatch/http/root.sbatch \
    10 4

sleep 1

jobname=hllworkers1 && \
rootname=hllroot1 && \
rootid=$(squeue -u $(whoami) --name=${rootname} -h -o "%i") && \
roothost=$(squeue -u $(whoami) --name=${rootname} --states=R -h -o "%N:%k") && \
sbatch \
    -J ${jobname} \
    --time=3-00:00:00 \
    --dependency=after:${rootid} \
    ~/Workspace/facu/hll-truco/hll-truco/sbatch/http/workers.sbatch \
    2 1 40 null InfosetPartidaXXLarge sha3 252000 10 http://${roothost} 4