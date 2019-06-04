Dgraph demo to transactionality with an incrementing counter.

### Run cluster

    docker-compose up

This runs 3 Dgraph Alphas and 1 Dgraph Zero with the following port configuration:

|        | gRPC-internal | gRPC-external | HTTP-external |
|--------|---------------|---------------|---------------|
| alpha1 | 7080          | 9080          | 8080          |
| alpha2 | 7082          | 9082          | 8082          |
| alpha3 | 7083          | 9083          | 8083          |
| zero1  | 5080          | -             | 6080          |

### Run counters

Run a `dgraph increment` counter per Dgraph Alpha.

    dgraph increment --addr localhost:9080 --num=1000 --wait=1s | tee a1.log
    dgraph increment --addr localhost:9082 --num=1000 --wait=1s | tee a2.log
    dgraph increment --addr localhost:9083 --num=1000 --wait=1s | tee a3.log

### Verify

Check that the total number of counters matches the sequence from 1 to 3000.

    diff -u <(grep Counter *.log | sort -nk5 | awk '{print $5}') <(seq 1 3000)

Check that the sum of the counters matches the expected sum.

    grep Counter *.log | sort -nk5 | awk '{ sum+=$5} END { print sum; print (3000*3001)/2}'
