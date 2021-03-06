package domain

type Article struct {
	ID        uint    `json:"id"`
	Title     string  `json:"title"`
	Content   string  `json:"content"`
	Author    *Author `json:"author"`
	IsDeleted bool    `json:"is_deleted"`
}

type ArticleQuery struct {
	Offset          int    `json:"offset"`
	Limit           int    `json:"limit"`
	AuthorID        uint   `json:"author_id"`
	Title           string `json:"title"`
	IncludedDeleted bool   `json:"included_deleted"`
}

type ArticleCreateCommand struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required,gte=10"`
}

type ArticleUpdateCommand struct {
	ID uint `json:"id" validate:"required"`
	*ArticleCreateCommand
}
