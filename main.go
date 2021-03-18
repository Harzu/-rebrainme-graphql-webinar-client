package main

import (
	"context"
	"fmt"
	"log"

	"github.com/machinebox/graphql"
)

func main() {
	client := graphql.NewClient("http://localhost:8080/graph/query")
	req := graphql.NewRequest(`{
		me {
			id
			name
			address
			orders {
				id
				products {
					id
				}
			}
		}
	}`)

	// При каждом вызове login resolver токен изменится, поэтому нужно не забыть
	// в случае нового логина, так-же изменить токен здесь.
	req.Header.Set("X-Session", "8c80efb8-d71a-4047-ba0b-c77018cc5b33")

	var resp struct {
		Me struct {
			ID      int64
			Name    string
			Address string
			Orders  []struct {
				ID       int64
				Products []struct {
					ID int64
				}
			}
		}
	}
	if err := client.Run(context.Background(), req, &resp); err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
