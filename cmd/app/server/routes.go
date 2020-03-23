package server

import "github.com/burhon94/postservice/core/middleware/logger"

func (receiver *Server) InitRoutes() {

	receiver.router.GET("/api/health", receiver.handlerRespHealth(), logger.Logger("GET/HEALTH"))

	receiver.router.GET("/api/posts", receiver.handlerGetPosts(), logger.Logger("GET/AllPosts"))

	receiver.router.GET("/api/posts/movies", receiver.handlerGetPostsMovies(), logger.Logger("GET/POSTS/Movies"))

	receiver.router.GET("/api/posts/games", receiver.handlerGetPostsGames(), logger.Logger("GET/POSTS/Games"))

	receiver.router.GET("/api/posts/programmes", receiver.handlerGetPostsProgs(), logger.Logger("GET/POSTS/Programms"))

	receiver.router.GET("/api/posts/musics", receiver.handlerGetPostsMusics(), logger.Logger("GET/POSTS/Musics"))

	receiver.router.GET("/api/post/{id}", receiver.getPost(), logger.Logger("GET/POSTbyId"))

	receiver.router.POST("/api/post/0", receiver.handlerNewPost(), logger.Logger("Creat/Post"))
}
