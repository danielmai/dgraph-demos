#!/bin/bash

# Download 21million movie data and schema from the dgraph benchmarks repo master branch.
wget -O 21million.rdf.gz https://github.com/dgraph-io/benchmarks/raw/master/data/release/21million.rdf.gz
wget -O 21million.schema https://github.com/dgraph-io/benchmarks/raw/master/data/release/release.schema
