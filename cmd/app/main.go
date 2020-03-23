package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/burhon94/alifMux/pkg/mux"
	"github.com/burhon94/bdi/pkg/di"
	"github.com/burhon94/postservice/cmd/app/server"
	"github.com/burhon94/postservice/core/posts"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net"
	"net/http"
)

// -host 0.0.0.0 -port 7777 -dsn postgres://user:pass@localhost:5455/posts
var (
	hostFlag = flag.String("host", "", "Server host")
	portFlag = flag.String("port", "", "Server host")
	dsn      = flag.String("dsn", "", "Postgres DSN")
)

type DSN string

func main() {
	flag.Parse()
	addr := net.JoinHostPort(*hostFlag, *portFlag)
	start(addr, *dsn)
}

func start(addr, dsn string) {
	container := di.NewContainer()
	err := container.Provide(
		server.NewServer,
		mux.NewExactMux,
		posts.NewService,
		func() DSN { return DSN(dsn) },
		func(dsn DSN) *pgxpool.Pool {
			pool, err := pgxpool.Connect(context.Background(), string(dsn))
			if err != nil {
				panic(fmt.Errorf("can't create pool: %w", err))
			}
			return pool
		},
	)
	if err != nil {
		panic(fmt.Sprintf("can't setting provide: %v", err))
	}

	container.Start()
	var appServer *server.Server
	container.Component(&appServer)
	log.Printf("service posts listening ... %s", addr)
	panic(http.ListenAndServe(addr, appServer))
}
