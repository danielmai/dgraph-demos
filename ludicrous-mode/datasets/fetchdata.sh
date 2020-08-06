#!/bin/bash

# Download 21million movie data and schema from the dgraph benchmarks repo master branch.
wget -O 1million.rdf.gz https://github.com/dgraph-io/benchmarks/raw/master/data/1million.rdf.gz
wget -O 1million.schema https://github.com/dgraph-io/benchmarks/raw/master/data/1million.schema
