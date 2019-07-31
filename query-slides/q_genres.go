package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
)

var q = `
# start q OMIT
{
  movies(func: gt(count(~genre), 30000), first: 1) @recurse(depth: 5, loop: true) {
    name@en
    ~genre (first:10) @filter(gt(count(starring), 2))
    starring (first: 2)
    performance.actor
  }
}
# end q OMIT
`

func query(dc *dgo.Dgraph, q string) (*api.Response, error) {
	ctx := context.Background()
	txn := dc.NewReadOnlyTxn()
	defer txn.Discard(ctx)
	return txn.Query(ctx, q)
}

func main() {
	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	dc := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	resp, err := query(dc, q)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", string(resp.Json))
	fmt.Println("Done.")
}
