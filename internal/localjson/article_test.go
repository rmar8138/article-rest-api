package localjson

import (
	"fmt"
	"os"
	"testing"

	"github.com/rmar8138/article-rest-api/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RepoTestSuite struct {
	suite.Suite
	Filepath string
}

func writeTestFile(filepath string) {
	testArticles := []Article{
		{
			ID:    "1",
			Title: "test title",
			Date:  "2020-01-01",
			Body:  "test body",
			Tags:  []string{"test tag"},
		},
	}

	os.MkdirAll("./testdata", os.ModePerm)
	saveArticles(filepath, testArticles)
}

func (suite *RepoTestSuite) SetupTest() {
	dir, _ := os.Getwd()
	filepath := fmt.Sprintf("%v/testdata/testArticles.json", dir)
	writeTestFile(filepath)
	suite.Filepath = filepath
}

func (suite *RepoTestSuite) TestGet() {
	repo := NewArticleRepository(suite.Filepath)
	id := "1"
	article, err := repo.Get(id)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), id, article.ID)

	id = "2"
	article, err = repo.Get(id)
	assert.Error(suite.T(), err)
}

func (suite *RepoTestSuite) TestCreate() {
	repo := NewArticleRepository(suite.Filepath)
	newArticle := service.CreateArticleInput{
		ID:    "2",
		Title: "test",
		Date:  "2020-01-01",
		Body:  "test",
		Tags: []string{
			"test",
		},
	}
	err := repo.Create(newArticle)
	assert.NoError(suite.T(), err)

	// error should return when inserting existing id
	err = repo.Create(newArticle)
	assert.Error(suite.T(), err)
}

func (suite *RepoTestSuite) TestGetByDate() {
	repo := NewArticleRepository(suite.Filepath)
	date := "2020-01-01"
	articles, err := repo.GetByDate(date)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), articles, 1)

	// should not error out if nothing found
	date = "2021-01-01"
	articles, err = repo.GetByDate(date)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), articles, 0)
}

func TestRepoTestSuite(t *testing.T) {
	suite.Run(t, new(RepoTestSuite))
}

func TestIDAlreadyExists(t *testing.T) {
	articles := []Article{
		{
			ID: "1",
		},
		{
			ID: "2",
		},
		{
			ID: "3",
		},
	}

	assert.True(t, idAlreadyExists(articles, "3"))
	assert.False(t, idAlreadyExists(articles, "4"))
}
