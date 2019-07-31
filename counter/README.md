Dgraph demo of transactions with an incrementing counter.

* Increment a counter transactionally
* Observe correctness, even when servers are getting killed and restarted.

### Run cluster

    docker-compose up

This runs 3 Dgraph Alphas and 1 Dgraph Zero with the following port configuration:

|        | `--port_offset` | gRPC-internal | HTTP-external | gRPC-external |
|--------|-----------------|---------------|---------------|---------------|
| alpha1 | 0               | 7080          | 8080          | 9080          |
| alpha2 | 2               | 7082          | 8082          | 9082          |
| alpha3 | 3               | 7083          | 8083          | 9083          |
| zero1  | 0               | 5080          | 6080          | -             |

### Run counters

Run a `dgraph increment` counter per Dgraph Alpha.

    dgraph increment --addr localhost:9080 --num=1000 --wait=1s | tee a1.log
    dgraph increment --addr localhost:9082 --num=1000 --wait=1s | tee a2.log
    dgraph increment --addr localhost:9083 --num=1000 --wait=1s | tee a3.log

### Restarting Alphas

Try stopping an Alpha while the counter tool is running. As long as the majority
of the Alpha group is up, the other running counters can continue making
progress.

    docker-compose stop alpha1

See the other two counters continue running normally. After some time, start alpha1. The counter that it sees once it has re-joined the cluster is the **latest** value of the entire cluster. This shows that Dgraph guarantees linearizable reads.

### Verify

Check that the total number of counters matches the sequence from 1 to 3000.

    diff -u <(grep Counter *.log | sort -nk5 | awk '{print $5}') <(seq 1 3000)

Check that the sum of the counters matches the expected sum.

    grep Counter *.log | sort -nk5 | awk '{ sum+=$5} END { print sum; print (3000*3001)/2}'
