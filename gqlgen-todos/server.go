package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/masaya0521/gqlgen-todos/graph"
	"github.com/masaya0521/gqlgen-todos/graph/generated"

	_ "github.com/lib/pq"
)

const defaultPort = "8080"
const dataSource = "user=postgres password=postgres host=localhost port=5432 dbname=postgres sslmode=disable"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// 主にここの処理
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic(err)
	}
	defer func() {
		fmt.Printf("DB close")
		if db != nil {
			if err := db.Close(); err != nil {
				panic(err)
			}
		}
	}()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
