#!/bin/bash
source environment
dgraph live -f datasets/1million.rdf.gz -s datasets/1million.schema
