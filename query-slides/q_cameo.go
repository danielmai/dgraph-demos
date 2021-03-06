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
  spielberg as var(func: allofterms(name@en, "steven spielberg")) {
    uid
  }
  
  cameos(func: uid(spielberg)) @cascade @normalize {
    director : name@en
    director.film {
      film : name@en
      starring {
        performance.actor @filter(uid(spielberg)) {
          actor : name@en
        }
        performance.character {
          character : name@en
        }
      }
    }
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
