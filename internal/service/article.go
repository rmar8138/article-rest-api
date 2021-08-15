package service

import (
	"sort"

	"github.com/pkg/errors"
	"github.com/rmar8138/article-rest-api/internal/domain"
)

// ArticleRepository defines implementation of a repo to be passed into the service layer
type ArticleRepository interface {
	Get(id string) (domain.Article, error)
	GetByDate(date string) ([]domain.Article, error)
	Create(CreateArticleInput) error
}

// ArticleService represents the service layer of articles
type ArticleService struct {
	repo ArticleRepository
}

// NewArticleService instantiates a new article service
func NewArticleService(repo ArticleRepository) *ArticleService {
	return &ArticleService{
		repo: repo,
	}
}

// CreateArticleInput is the shape of input needed for creating an article
type CreateArticleInput struct {
	ID    string
	Title string
	Date  string
	Body  string
	Tags  []string
}

// Get retrieves a single article by its id
func (as *ArticleService) Get(id string) (domain.Article, error) {
	article, err := as.repo.Get(id)
	if err != nil {
		return domain.Article{}, errors.Wrapf(err, "unable to get article with id: %v", id)
	}

	return article, nil
}

// Create creates a new article
func (as *ArticleService) Create(newArticle CreateArticleInput) error {
	err := as.repo.Create(newArticle)
	if err != nil {
		return errors.Wrapf(err, "unable to create new article")
	}

	return nil
}

// ArticlesByTagAndDate represents the output of the business logic when getting info on articles by tag name and date
type ArticlesByTagAndDate struct {
	Tag         string
	Count       int
	ArticleIDs  []string
	RelatedTags []string
}

// GetArticlesByTagAndDate gets information on articles based on a given tagname and article date
func (as *ArticleService) GetArticlesByTagAndDate(tagName, date string) (ArticlesByTagAndDate, error) {
	// get articles by date
	articles, err := as.repo.GetByDate(date)
	if err != nil {
		return ArticlesByTagAndDate{}, errors.Wrapf(err, "unable to get articles by date")
	}

	relatedTagsMap := make(map[string]bool)
	var count int
	var articleIDs []string

	for _, a := range articles {
		// get all related tags with no duplicates
		for _, t := range a.Tags {
			if t != tagName {
				relatedTagsMap[t] = true
			}
		}
		// get count of total number of tags between articles
		count += len(a.Tags)

		articleIDs = append(articleIDs, a.ID)
	}

	// get list of last 10 article ids
	// assuming that the higher the id, the later the article was submitted
	sort.Slice(articleIDs, func(i, j int) bool {
		return articleIDs[i] > articleIDs[j]
	})
	if len(articleIDs) > 10 {
		articleIDs = articleIDs[:10]
	}

	relatedTags := make([]string, 0, len(relatedTagsMap))
	for tag := range relatedTagsMap {
		relatedTags = append(relatedTags, tag)
	}

	return ArticlesByTagAndDate{
		Tag:         tagName,
		Count:       count,
		ArticleIDs:  articleIDs,
		RelatedTags: relatedTags,
	}, nil
}
