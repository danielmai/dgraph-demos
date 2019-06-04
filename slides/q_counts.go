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
  director(func: eq(name@en, "Steven Spielberg")) {
    name@en
    director.film {
      name@en
      numGenres : g as count(genre)
    }
    totalGenres : sum(val(g))
    genresPerMovie : avg(val(g))
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
