package posts

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"
)

type ServicePost struct {
}

func NewService() *ServicePost {
	return &ServicePost{}
}

type Post struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	DataAndTime time.Time `json:"dataandtime"`
	Poster      string    `json:"poster"`
	Category    string    `json:"category"`
	FileUrl     string    `json:"file_url"`
}

type PostNew struct {
	Title    string `json:"title"`
	Poster   string `json:"poster"`
	Category string `json:"category"`
	FileUrl  string `json:"file_url"`
}

var badRequest = errors.New("bad request")

func (receiver *ServicePost) GetPosts(ctx context.Context, pool *pgxpool.Pool) (response []Post, err error) {
	posts := make([]Post, 0)
	//conn, err := pool.Acquire(context.Background())
	//if err != nil {
	//	return posts, err
	//}
	//defer conn.Release()

	rows, err := pool.Query(ctx, "SELECT id, title, dataandtime, poster, category, urlfile FROM post;")
	if err != nil {
		err = errors.New(fmt.Sprintf("can't get posts: %v", err))
		return
	}
	defer rows.Close()

	for rows.Next() {
		post := Post{}
		err := rows.Scan(&post.Id, &post.Title, &post.DataAndTime, &post.Poster, &post.Category, &post.FileUrl)
		if err != nil {
			err = errors.New(fmt.Sprintf("can't scan row: %v", err))
			return nil, err
		}
		posts = append(posts, post)
	}

	err = rows.Err()
	if err != nil {
		return
	}

	return posts, nil
}

func (receiver *ServicePost) CreatePost(ctx context.Context, request PostNew, pool *pgxpool.Pool) (err error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	if request.Title == "" {
		return badRequest
	}
	if request.Poster == "" {
		return badRequest
	}
	if request.Category == "" {
		return badRequest
	}
	if request.FileUrl == "" {
		return badRequest
	}

	dataTime := time.Now()
	_, err = conn.Exec(ctx,
		`INSERT INTO post (title, dataandtime, poster, category, urlfile) VALUES ($1, $2, $3, $4, $5)`,
		request.Title, dataTime, request.Poster, request.Category, request.FileUrl)
	if err != nil {
		return errors.New("error while create post")
	}

	return nil
}

func (receiver *ServicePost) GetPostsMovies(ctx context.Context, pool *pgxpool.Pool) (response []Post, err error) {
	posts := make([]Post, 0)
	//conn, err := pool.Acquire(ctx)
	//if err != nil {
	//	log.Print(err)
	//	return
	//}
	//defer conn.Release()
	rows, err := pool.Query(ctx, "SELECT id, title, dataandtime, poster, category, urlfile FROM post WHERE category = 'movie';")
	if err != nil {
		err = errors.New(fmt.Sprintf("can't get posts: %v", err))
		return
	}
	defer rows.Close()

	for rows.Next() {
		post := Post{}
		err := rows.Scan(&post.Id, &post.Title, &post.DataAndTime, &post.Poster, &post.Category, &post.FileUrl)
		if err != nil {
			err = errors.New(fmt.Sprintf("can't scan row: %v", err))
			return nil, err
		}
		posts = append(posts, post)
	}

	err = rows.Err()
	if err != nil {
		return
	}

	return posts, nil
}

func (receiver *ServicePost) GetPostsGames(ctx context.Context, pool *pgxpool.Pool) (response []Post, err error) {
	posts := make([]Post, 0)
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Print(err)
		return
	}
	defer conn.Release()
	rows, err := conn.Query(ctx, "SELECT id, title, dataandtime, poster, category, urlfile FROM post WHERE category = 'game';")
	if err != nil {
		err = errors.New(fmt.Sprintf("can't get posts: %v", err))
		return
	}
	defer rows.Close()

	for rows.Next() {
		post := Post{}
		err := rows.Scan(&post.Id, &post.Title, &post.DataAndTime, &post.Poster, &post.Category, &post.FileUrl)
		if err != nil {
			err = errors.New(fmt.Sprintf("can't scan row: %v", err))
			return nil, err
		}
		posts = append(posts, post)
	}

	err = rows.Err()
	if err != nil {
		return
	}

	return posts, nil
}

func (receiver *ServicePost) GetPostsProgs(ctx context.Context, pool *pgxpool.Pool) (response []Post, err error) {
	posts := make([]Post, 0)
	//conn, err := pool.Acquire(context.Background())
	//if err != nil {
	//	log.Print(err)
	//	return
	//}
	//defer conn.Release()
	rows, err := pool.Query(ctx, "SELECT id, title, dataandtime, poster, category, urlfile FROM post WHERE category = 'prog';")
	if err != nil {
		err = errors.New(fmt.Sprintf("can't get posts: %v", err))
		return
	}
	defer rows.Close()

	for rows.Next() {
		post := Post{}
		err := rows.Scan(&post.Id, &post.Title, &post.DataAndTime, &post.Poster, &post.Category, &post.FileUrl)
		if err != nil {
			err = errors.New(fmt.Sprintf("can't scan row: %v", err))
			return nil, err
		}
		posts = append(posts, post)
	}

	err = rows.Err()
	if err != nil {
		return
	}

	return posts, nil
}

func (receiver *ServicePost) GetPostsMusics(ctx context.Context, pool *pgxpool.Pool) (response []Post, err error) {
	posts := make([]Post, 0)
	//conn, err := pool.Acquire(context.Background())
	//if err != nil {
	//	log.Print(err)
	//	return
	//}
	//defer conn.Release()
	rows, err := pool.Query(ctx, "SELECT id, title, dataandtime, poster, category, urlfile FROM post WHERE category = 'music';")
	if err != nil {
		err = errors.New(fmt.Sprintf("can't get posts: %v", err))
		return
	}
	defer rows.Close()

	for rows.Next() {
		post := Post{}
		err := rows.Scan(&post.Id, &post.Title, &post.DataAndTime, &post.Poster, &post.Category, &post.FileUrl)
		if err != nil {
			err = errors.New(fmt.Sprintf("can't scan row: %v", err))
			return nil, err
		}
		posts = append(posts, post)
	}

	err = rows.Err()
	if err != nil {
		return
	}

	return posts, nil
}

func (receiver *ServicePost) GetPost(ctx context.Context, id int, pool *pgxpool.Pool) (response Post, err error) {
	//conn, err := pool.Acquire(context.Background())
	//if err != nil {
	//	err = errors.New(fmt.Sprintf("can't execute pool: %v", err))
	//	return
	//}
	//defer conn.Release()

	var titlePost, poster, category, fileUrl string
	var dataAndTimePost time.Time
	dataAndTimePost.Format("20060102150405")
	err = pool.QueryRow(ctx, `SELECT title, dataandtime, poster, category, urlfile FROM post WHERE id = $1;`, id).Scan(&titlePost, &dataAndTimePost, &poster, &category, &fileUrl)
	if err != nil {
		err = errors.New(fmt.Sprintf("can't get post by id %d: error: %v", id, err))
		return
	}

	response = Post{
		Id:          id,
		Title:       titlePost,
		DataAndTime: dataAndTimePost,
		Poster:      poster,
		Category:    category,
		FileUrl:     fileUrl,
	}
	return
}
