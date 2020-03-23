package server

import (
	"context"
	"fmt"
	"github.com/burhon94/alifMux/pkg/mux"
	readerJSON "github.com/burhon94/json/cmd/reader"
	writerJSON "github.com/burhon94/json/cmd/writer"
	"github.com/burhon94/postservice/core/posts"
	"log"
	"net/http"
	"strconv"
	"time"
)

type PostResp struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	DataAndTime string `json:"dataandtime"`
	Poster      string `json:"poster"`
	Category    string `json:"category"`
}

type ErrorDTO struct {
	Errors []string `json:"errors"`
}

func (receiver *Server) handlerRespHealth() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
			_, err := fmt.Fprintf(writer, "Health ok")
		if err != nil {
			log.Printf("err: %v", err)
		}
	}
}

func (receiver *Server) handlerNewPost() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var bodyRequest posts.PostNew
		err := readerJSON.ReadJSONHTTP(request, &bodyRequest)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			err := writerJSON.WriteJSONHTTP(writer, &ErrorDTO{
				[]string{"err.json_invalid"},
			})
			if err != nil {
				http.Error(writer, http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
			}
			return
		}

		ctx, _ := context.WithTimeout(request.Context(), time.Second)
		err = receiver.posts.CreatePost(ctx, bodyRequest, receiver.pool)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

func (receiver *Server) handlerGetPosts() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx, _ := context.WithTimeout(request.Context(), time.Second)
		posts, err := receiver.posts.GetPosts(ctx, receiver.pool)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		err = writerJSON.WriteJSONHTTP(writer, posts)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

func (receiver *Server) getPost() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx, _ := context.WithTimeout(request.Context(), time.Second)
		value, ok := mux.FromContext(ctx, "id")
		if !ok {
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			log.Printf("can't find from ctx: %v", ok)
			return
		}

		id, err := strconv.Atoi(value)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			log.Printf("can't convert id: %v", err)
			return
		}

		getPost, err := receiver.posts.GetPost(ctx, id, receiver.pool)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		err = writerJSON.WriteJSONHTTP(writer, getPost)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		log.Print(getPost)
	}

}

func (receiver *Server) handlerGetPostsMovies() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx, _ := context.WithTimeout(request.Context(), time.Second)
		posts, err := receiver.posts.GetPostsMovies(ctx, receiver.pool)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		err = writerJSON.WriteJSONHTTP(writer, posts)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

func (receiver *Server) handlerGetPostsGames() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx, _ := context.WithTimeout(request.Context(), time.Second)
		posts, err := receiver.posts.GetPostsGames(ctx, receiver.pool)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		err = writerJSON.WriteJSONHTTP(writer, posts)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

func (receiver *Server) handlerGetPostsProgs() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx, _ := context.WithTimeout(request.Context(), time.Second)
		posts, err := receiver.posts.GetPostsProgs(ctx, receiver.pool)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		err = writerJSON.WriteJSONHTTP(writer, posts)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

func (receiver *Server) handlerGetPostsMusics() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx, _ := context.WithTimeout(request.Context(), time.Second)
		posts, err := receiver.posts.GetPostsMusics(ctx, receiver.pool)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		err = writerJSON.WriteJSONHTTP(writer, posts)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}
