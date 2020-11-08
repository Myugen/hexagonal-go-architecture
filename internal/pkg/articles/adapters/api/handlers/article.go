package handlers

import (
	"net/http"

	"github.com/myugen/hexagonal-go-architecture/internal/pkg/articles/adapters/api/requests"

	"github.com/myugen/hexagonal-go-architecture/internal/pkg/articles/adapters/api/responses"

	"github.com/labstack/echo/v4"
	"github.com/myugen/hexagonal-go-architecture/internal/pkg/articles/ports/services"
)

type IArticle interface {
	Get(e echo.Context) error
}

type articleHandler struct {
	articleService services.IArticle
}

func New(articleService services.IArticle) *articleHandler {
	return &articleHandler{articleService: articleService}
}

func (h *articleHandler) Get(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	article, err := h.articleService.Get(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "An error occurs getting an article", err)
	}
	return c.JSON(http.StatusOK, responses.FromModel(article))
}

func (h *articleHandler) Find(c echo.Context) error {
	ctx := c.Request().Context()
	query := c.QueryParams()
	qpArticle := requests.NewArticleQueryParams(query)
	qArticle := qpArticle.ToArticleQuery()
	articles, err := h.articleService.Find(ctx, qArticle)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "An error occurs getting articles", err)
	}

	response := make([]*responses.Article, 0)
	for _, article := range articles {
		response = append(response, responses.FromModel(article))
	}

	return c.JSON(http.StatusOK, response)
}
