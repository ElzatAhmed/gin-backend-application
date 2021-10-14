package main

import (
	"github.com/gin-backend-application/controller/author_controller_implements"
	"github.com/gin-backend-application/mock"
	"github.com/gin-backend-application/repository/repo_implements"
	"github.com/gin-backend-application/service/author_service_implements"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)



var (

	authors, _ = mock.GenerateMockData(100)

	authorRepo = repo_implements.NewAuthorRepo(authors)
	//bookRepo = repo_implements.NewBookRepo(books)

	authorInfoService = author_service_implements.NewAuthorInfoService(authorRepo)
	//authorPublishService = author_service_implements.NewAuthorPublishService(authorRepo, bookRepo)
	//authorRegisterService = author_service_implements.NewAuthorRegisterService(authorRepo)
	//bookInfoService = book_service_implements.NewBookInfoService(authorRepo, bookRepo)
	//bookUpdateService = book_service_implements.NewBookUpdateService(bookRepo)

	authorInfoController = author_controller_implements.NewAuthorInfoController(authorInfoService)
	//authorPublishController = author_controller_implements.NewAuthorPublishController(authorPublishService)
	//authorRegisterController = author_controller_implements.NewAuthorRegisterController(authorRegisterService)
	//bookInfoController = book_controller_implements.NewBookInfoController(bookInfoService)
	//bookUpdateController = book_controller_implements.NewBookUpdateController(bookUpdateService)
)

func initServer() *gin.Engine {

	server := gin.Default()

	server.GET("/api/authors/:id", func(context *gin.Context) {
		id, err := strconv.ParseInt(context.Param("id"), 10, 64)
		if err != nil {
			context.JSON(400, gin.H{
				"message": "request error",
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

	return server
}

func main()  {
	server := initServer()
	log.Fatal(server.Run(":8000"))
}
