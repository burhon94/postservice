package server

import (
	"github.com/burhon94/alifMux/pkg/mux"
	postsSvc "github.com/burhon94/postservice/core/posts"
	"github.com/jackc/pgx/v4/pgxpool"
	"net/http"
)

type Server struct {
	router *mux.ExactMux
	posts  *postsSvc.ServicePost
	pool *pgxpool.Pool
}

func NewServer(router *mux.ExactMux, posts *postsSvc.ServicePost, pool *pgxpool.Pool) *Server {
	return &Server{router: router, posts: posts, pool: pool}
}

func (receiver *Server) Start() {
	receiver.InitRoutes()
}

func (receiver *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	receiver.router.ServeHTTP(w, r)
}
