package main

import (
	"github.com/gin-backend-application/controller/author_controller_implements"
	"github.com/gin-backend-application/controller/book_controller_implements"
	"github.com/gin-backend-application/mock"
	"github.com/gin-backend-application/model"
	"github.com/gin-backend-application/repository/repo_implements"
	"github.com/gin-backend-application/service/author_service_implements"
	"github.com/gin-backend-application/service/book_service_implements"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)



var (

	authors, books = mock.GenerateMockData(100)

	authorRepo = repo_implements.NewAuthorRepo(authors)
	bookRepo = repo_implements.NewBookRepo(books)

	authorInfoService = author_service_implements.NewAuthorInfoService(authorRepo)
	authorPublishService = author_service_implements.NewAuthorPublishService(authorRepo, bookRepo)
	authorRegisterService = author_service_implements.NewAuthorRegisterService(authorRepo)
	bookInfoService = book_service_implements.NewBookInfoService(authorRepo, bookRepo)
	bookUpdateService = book_service_implements.NewBookUpdateService(bookRepo)

	authorInfoController = author_controller_implements.NewAuthorInfoController(authorInfoService)
	authorPublishController = author_controller_implements.NewAuthorPublishController(authorPublishService)
	authorRegisterController = author_controller_implements.NewAuthorRegisterController(authorRegisterService)
	bookInfoController = book_controller_implements.NewBookInfoController(bookInfoService)
	bookUpdateController = book_controller_implements.NewBookUpdateController(bookUpdateService)
)

func initServer() *gin.Engine {

	server := gin.Default()

	server.GET("/api/authors/id/:id", func(context *gin.Context) {
		id, err := strconv.ParseInt(context.Param("id"), 10, 64)
		if err != nil {
			context.JSON(400, gin.H{
				"message": err.Error(),
				"time": time.Now(),
			})
			return
		}
		author, err := authorInfoController.GetAuthorById((uint64)(id))
		if err != nil {
			context.JSON(400, gin.H{
				"message": err.Error(),
				"time": time.Now(),
			})
			return
		}
		context.JSON(200, author)
	})

	server.GET("/api/authors/name/:name", func(context *gin.Context) {
		name := context.Param("name")
		authors, err := authorInfoController.GetAuthorsByName(name)
		if err != nil {
			context.JSON(400, gin.H{
				"message": err.Error(),
				"time": time.Now(),
			})
			return
		}
		context.JSON(200, authors)
	})

	server.PUT("/api/authors/:id", func(context *gin.Context) {
		var author model.Author
		err := context.BindJSON(&author)
		if err != nil {
			context.JSON(400, gin.H{
				"message": err.Error(),
				"time": time.Now(),
			})
			return
		}
		id, err := strconv.ParseInt(context.Param("id"), 10, 64)
		if err != nil {
			context.JSON(400, gin.H{
				"message": err.Error(),
				"time": time.Now(),
			})
			return
		}
		author, err = authorInfoController.UpdateAuthorInfoById(uint64(id), author)
		if err != nil {
			context.JSON(400, gin.H{
				"message": err.Error(),
				"time": time.Now(),
			})
			return
		}
		context.JSON(200, author)
	})

	server.POST("/api/authors", func(context *gin.Context) {
		var author model.Author
		err := context.BindJSON(&author)
		if err != nil {
			context.JSON(400, gin.H{
				"message": err.Error(),
				"time": time.Now(),
			})
			return
		}
		author, err = authorRegisterController.Register(author)
		if err != nil {
			context.JSON(400, gin.H{
				"message": err.Error(),
				"time": time.Now(),
			})
			return
		}
		context.JSON(200, author)
	})

	server.POST("/api/books/:id", func(context *gin.Context) {
		id, err := strconv.ParseInt(context.Param("id"), 10, 64)
		if err != nil {
			context.JSON(400, gin.H{
				"message": err.Error(),
				"time": time.Now(),
			})
			return
		}
		var book model.Book
		err = context.BindJSON(&book)
		if err != nil {
			context.JSON(400, gin.H{
				"message": err.Error(),
				"time": time.Now(),
			})
			return
		}
		book, err = authorPublishController.Publish(uint64(id), book)
		if err != nil {
			context.JSON(400, gin.H{
				"message": err.Error(),
				"time": time.Now(),
			})
			return
		}
		context.JSON(200, book)
	})

	server.GET("/api/books/id/:id", func(context *gin.Context) {
		id, err := strconv.ParseInt(context.Param("id"), 10, 64)
		if err != nil {
			context.JSON(400, gin.H{
				"message": err.Error(),
				"time": time.Now(),
			})
			return
		}
		book, err := bookInfoController.GetBookById(uint64(id))
		if err != nil {
			context.JSON(400, gin.H{
				"message": err.Error(),
				"time": time.Now(),
			})
			return
		}
		context.JSON(200, book)
	})

	server.GET("/api/books/name/:name", func(context *gin.Context) {
		name := context.Param("name")
		book, err := bookInfoController.GetBooksByName(name)
		if err != nil {
			context.JSON(400, gin.H{
				"message": err.Error(),
				"time": time.Now(),
			})
			return
		}
		context.JSON(200, book)
	})

	server.GET("/api/books/author_id/:id", func(context *gin.Context) {
		id, err := strconv.ParseInt(context.Param("id"), 10, 64)
		if err != nil {
			context.JSON(400, gin.H{
				"message": err.Error(),
				"time": time.Now(),
			})
			return
		}
		book, err := bookInfoController.GetBooksByAuthorId(uint64(id))
		if err != nil {
			context.JSON(400, gin.H{
				"message": err.Error(),
				"time": time.Now(),
			})
			return
		}
		context.JSON(200, book)
	})

	server.GET("/api/books/author_name/:name", func(context *gin.Context) {
		name := context.Param("name")
		book, err := bookInfoController.GetBooksByAuthorName(name)
		if err != nil {
			context.JSON(400, gin.H{
				"message": err.Error(),
				"time": time.Now(),
			})
			return
		}
		context.JSON(200, book)
	})

	server.PUT("/api/books/:id", func(context *gin.Context) {
		id, err := strconv.ParseInt(context.Param("id"), 10, 64)
		if err != nil {
			context.JSON(400, gin.H{
				"message": err.Error(),
				"time": time.Now(),
			})
			return
		}
		var book model.Book
		err = context.BindJSON(&book)
		if err != nil {
			context.JSON(400, gin.H{
				"message": err.Error(),
				"time": time.Now(),
			})
			return
		}
		book, err = bookUpdateController.UpdateBookInfoById(uint64(id), book)
		if err != nil {
			context.JSON(400, gin.H{
				"message": err.Error(),
				"time": time.Now(),
			})
			return
		}
		context.JSON(200, book)
	})

	return server
}

func main()  {
	server := initServer()
	log.Fatal(server.Run(":8000"))
}
