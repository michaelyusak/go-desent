package server

import (
	"michaelyusak/go-desent.git/handler"
	"michaelyusak/go-desent.git/repository/memory"
	"michaelyusak/go-desent.git/service"

	"github.com/gin-gonic/gin"
)

type routerOpts struct {
	handler struct {
		common *handler.Common
		book   *handler.Book
	}
}

func newRouter() *gin.Engine {
	booksRepo := memory.NewBooks()

	bookService := service.NewBook(booksRepo)

	bookHandler := handler.NewBook(bookService)

	return createRouter(routerOpts{
		handler: struct {
			common *handler.Common
			book   *handler.Book
		}{
			common: &handler.Common{},
			book:   bookHandler,
		},
	},
		[]string{},
	)
}

func createRouter(opts routerOpts, allowedOrigins []string) *gin.Engine {
	router := gin.New()

	router.Use(
		gin.Recovery(),
	)

	commonRouting(router, opts.handler.common)
	bookRouting(router, opts.handler.book)

	return router
}

func commonRouting(router *gin.Engine, handler *handler.Common) {
	router.GET("/ping", handler.Ping)
	router.POST("/echo", handler.Echo)
}

func bookRouting(router *gin.Engine, handler *handler.Book) {
	router.POST("/books", handler.CreateBook)
	router.GET("/books", handler.GetAllBook)
	router.GET("/books/:id", handler.GetBookById)
	router.PUT("/books/:id", handler.UpdateBook)
	router.DELETE("/books/:id", handler.DeleteBook)
}
