package handler

import (
	"errors"
	"fmt"
	"michaelyusak/go-desent.git/apperror"
	"michaelyusak/go-desent.git/entity"
	"michaelyusak/go-desent.git/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Book struct {
	bookService service.Book
}

func NewBook(
	bookService service.Book,
) *Book {
	return &Book{
		bookService: bookService,
	}
}

func (h *Book) CreateBook(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	var book entity.Book
	err := ctx.ShouldBindJSON(&book)
	if err != nil {
		logrus.WithError(err).Warn("[handler][book][CreateBook][ctx.ShouldBindJSON]")

		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
		return
	}

	err = h.bookService.CreateBook(ctx, &book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, book)
}

func (h *Book) GetAllBook(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	var limit, page int

	limitStr := ctx.Query("limit")
	if limitStr != "" {
		l, err := strconv.Atoi(limitStr)
		if err != nil {
			logrus.WithError(err).WithField("raw", limitStr).Warn("[handler][book][GetAllBook][strconv.Atoi(limitStr)]")

			ctx.JSON(http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
			return
		}
		limit = l
	}

	pageStr := ctx.Query("page")
	if pageStr != "" {
		p, err := strconv.Atoi(pageStr)
		if err != nil {
			logrus.WithError(err).WithField("raw", pageStr).Warn("[handler][book][GetAllBook][strconv.Atoi(pageStr)]")

			ctx.JSON(http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
			return
		}
		page = p
	}

	var filter entity.GetBookFilter

	filter.Author = ctx.Query("author")
	filter.Limit = limit
	filter.Page = page

	logrus.WithField("path", ctx.FullPath()).WithField("queries", fmt.Sprintf("%+v", filter)).Info("[handler][book][GetAllBook] filter parsed")

	books, err := h.bookService.GetAllBook(ctx, filter)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func (h *Book) GetBookById(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	bookId := ctx.Param("id")

	book, err := h.bookService.GetBookById(ctx, bookId)
	if err != nil {
		var apperr *apperror.AppError
		if errors.As(err, &apperr) {
			ctx.JSON(apperr.Code, map[string]string{
				"message": apperr.Message,
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (h *Book) UpdateBook(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	bookId := ctx.Param("id")

	var book entity.Book
	err := ctx.ShouldBindJSON(&book)
	if err != nil {
		logrus.WithError(err).Warn("[handler][book][UpdateBook][ctx.ShouldBindJSON]")

		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
		return
	}

	book.Id = bookId

	logrus.WithField("path", ctx.FullPath()).WithField("book", fmt.Sprintf("%+v", book)).Info("[handler][book][UpdateBook] path")

	updated, err := h.bookService.UpdateBook(ctx, &book)
	if err != nil {
		var apperr *apperror.AppError
		if errors.As(err, &apperr) {
			ctx.JSON(apperr.Code, map[string]string{
				"message": apperr.Message,
			})
			return
		}

		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, updated)
}

func (h *Book) DeleteBook(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	bookId := ctx.Param("id")

	deleted, err := h.bookService.DeleteBook(ctx, bookId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, deleted)
}
