#!/bin/bash

# Download 21million movie data and schema from the dgraph benchmarks repo release/v1.0 branch.
wget -O 21million.rdf.gz https://github.com/dgraph-io/benchmarks/raw/release/v1.0/data/release/21million.rdf.gz
wget -O 21million.schema https://github.com/dgraph-io/benchmarks/raw/release/v1.0/data/release/release.schema
