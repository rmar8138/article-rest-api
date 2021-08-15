package service

import (
	"testing"

	"github.com/rmar8138/article-rest-api/internal/domain"
	"github.com/stretchr/testify/assert"
)

type mockArticleRepository struct{}

func (m *mockArticleRepository) Get(id string) (domain.Article, error) {
	return domain.Article{
		ID:    "1",
		Title: "test 1",
		Date:  "2021-01-01",
		Body:  "test body",
		Tags: []string{
			"health",
			"science",
		},
	}, nil
}

func (m *mockArticleRepository) Create(newArticle CreateArticleInput) error {
	return nil
}

func (m *mockArticleRepository) GetByDate(date string) ([]domain.Article, error) {
	return []domain.Article{
		{
			ID:    "1",
			Title: "test 1",
			Date:  "2021-01-01",
			Body:  "test body",
			Tags: []string{
				"health",
				"science",
			},
		},
		{
			ID:    "2",
			Title: "test 2",
			Date:  "2021-01-01",
			Body:  "test body",
			Tags: []string{
				"health",
				"fitness",
			},
		},
	}, nil
}

func TestGet(t *testing.T) {
	repo := mockArticleRepository{}
	svc := NewArticleService(&repo)

	id := "1"
	article, err := svc.Get(id)
	assert.NoError(t, err)
	assert.Equal(t, article.ID, id)
}

func TestCreate(t *testing.T) {
	repo := mockArticleRepository{}
	svc := NewArticleService(&repo)

	newArticle := CreateArticleInput{
		ID:    "1",
		Title: "test 1",
		Date:  "2021-01-01",
		Body:  "test body",
		Tags: []string{
			"health",
			"science",
		},
	}
	err := svc.Create(newArticle)
	assert.NoError(t, err)
}

func TestGetArticlesByTagAndDate(t *testing.T) {
	tagName := "health"
	date := "2021-01-01"
	repo := mockArticleRepository{}
	svc := NewArticleService(&repo)

	articleInfo, err := svc.GetArticlesByTagAndDate(tagName, date)
	assert.NoError(t, err)
	assert.Equal(t, "health", articleInfo.Tag)
	assert.Equal(t, 4, articleInfo.Count)
	assert.Len(t, articleInfo.ArticleIDs, 2)
	assert.Contains(t, articleInfo.ArticleIDs, "1")
	assert.Contains(t, articleInfo.ArticleIDs, "2")
	assert.Len(t, articleInfo.RelatedTags, 2)
	assert.Contains(t, articleInfo.RelatedTags, "fitness")
	assert.Contains(t, articleInfo.RelatedTags, "science")
}
