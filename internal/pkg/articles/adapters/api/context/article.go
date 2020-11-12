package context

import (
	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	"github.com/myugen/hexagonal-go-architecture/internal/pkg/articles/adapters/db/dao"
	"github.com/myugen/hexagonal-go-architecture/internal/pkg/articles/ports/repositories"
	"github.com/myugen/hexagonal-go-architecture/pkg/logger"
	"github.com/pkg/errors"
)

type ArticleAPIContext struct {
	echo.Context
	db                *pg.DB
	tx                *pg.Tx
	log               *logger.Logger
	articleRepository repositories.IArticle
}

func NewArticleAPIContext(c echo.Context, db *pg.DB, log *logger.Logger) *ArticleAPIContext {
	return &ArticleAPIContext{
		Context:           c,
		db:                db,
		articleRepository: dao.New(),
		log:               log,
	}
}

func (a *ArticleAPIContext) ArticleRepository() repositories.IArticle {
	return a.articleRepository
}

func (a *ArticleAPIContext) Transaction() *pg.Tx {
	return a.tx
}

func (a *ArticleAPIContext) DB() *pg.DB {
	return a.db
}

func (a *ArticleAPIContext) Log() *logger.Logger {
	return a.log
}

func (a *ArticleAPIContext) BeginTransaction() (*pg.Tx, error) {
	tx, err := a.db.Begin()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	a.tx = tx
	return a.tx, nil
}