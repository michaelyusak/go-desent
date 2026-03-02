package server

import (
	"michaelyusak/go-desent.git/handler"
	"michaelyusak/go-desent.git/middleware"
	"michaelyusak/go-desent.git/repository/memory"
	"michaelyusak/go-desent.git/service"

	"github.com/gin-gonic/gin"
)

type routerOpts struct {
	handler struct {
		common *handler.Common
		book   *handler.Book
		auth   *handler.Auth
	}

	middleware struct {
		auth *middleware.Auth
	}
}

func newRouter() *gin.Engine {
	booksRepo := memory.NewBooks()
	usersRepo := memory.NewUsers()
	tokensRepo := memory.NewTokens()

	bookService := service.NewBook(booksRepo)
	authService := service.NewAuth(usersRepo, tokensRepo)

	bookHandler := handler.NewBook(bookService)
	authHandler := handler.NewAuth(authService)

	authMiddleware := middleware.NewAuth(authService)

	return createRouter(routerOpts{
		handler: struct {
			common *handler.Common
			book   *handler.Book
			auth   *handler.Auth
		}{
			common: &handler.Common{},
			book:   bookHandler,
			auth:   authHandler,
		},

		middleware: struct {
			auth *middleware.Auth
		}{
			auth: &authMiddleware,
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

	authGuard := opts.middleware.auth.AuthGuard()

	commonRouting(router, opts.handler.common)
	bookRouting(router, authGuard, opts.handler.book)
	authRouting(router, opts.handler.auth)

	return router
}

func commonRouting(router *gin.Engine, handler *handler.Common) {
	router.GET("/ping", handler.Ping)
	router.POST("/echo", handler.Echo)
}

func bookRouting(router *gin.Engine, authGuard gin.HandlerFunc, handler *handler.Book) {
	router.POST("/books", handler.CreateBook)
	router.GET("/books", authGuard, handler.GetAllBook)
	router.GET("/books/:id", handler.GetBookById)
	router.PUT("/books/:id", handler.UpdateBook)
	router.DELETE("/books/:id", handler.DeleteBook)
}

func authRouting(router *gin.Engine, handler *handler.Auth) {
	router.POST("/auth/token", handler.Login)
}
