package requests

import (
	"github.com/myugen/hexagonal-go-architecture/internal/articles/domain"
)

type ArticleCreateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (r *ArticleCreateRequest) ToCommand() *domain.ArticleCreateCommand {
	return &domain.ArticleCreateCommand{
		Title:   r.Title,
		Content: r.Content,
	}
}

type ArticleUpdateRequest struct {
	ID uint `json:"id"`
	*ArticleCreateRequest
}

func (r *ArticleUpdateRequest) ToCommand() *domain.ArticleUpdateCommand {
	return &domain.ArticleUpdateCommand{
		ID: r.ID,
		ArticleCreateCommand: &domain.ArticleCreateCommand{
			Title:   r.Title,
			Content: r.Content,
		},
	}
}

type ArticleQueryParams struct {
	Offset          int    `param:"offset"`
	Limit           int    `param:"limit"`
	AuthorID        uint   `param:"author_id"`
	Title           string `param:"title"`
	IncludedDeleted bool   `param:"included_deleted"`
}

func (a *ArticleQueryParams) ToArticleQuery() *domain.ArticleQuery {
	return &domain.ArticleQuery{
		Offset:          a.Offset,
		Limit:           a.Limit,
		AuthorID:        a.AuthorID,
		Title:           a.Title,
		IncludedDeleted: a.IncludedDeleted,
	}
}
