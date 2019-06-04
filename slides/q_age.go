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
  var(func:allofterms(name@en, "steven spielberg")) {
    films as director.film {
      p as count(starring)
      q as count(genre)
      date as initial_release_date
      years as math(since(date)/(365*24*60*60)) // HL
      score as math(cond(years > 10, 0, ln(p)+q-ln(years)))
    }
  }

  TopMovies(func: uid(films), orderdesc: val(score)) @filter(gt(val(score), 2)){ // HL
    name@en
    val(score)
    val(date)
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
