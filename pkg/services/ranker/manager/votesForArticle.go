package manager

import "github.com/chancegraff/project-news/internal/models"

// GetVotesForArticle will the vote count for a single article
func (m *Manager) GetVotesForArticle(articleID string) (models.ArticleVotes, error) {
	var articleVotes models.ArticleVotes

	// Build query
	query := m.Store.Database.
		Table("votes").
		Select("article_id, count(*) as votes").
		Where("article_id = ?", articleID).
		Group("article_id")

	// Commit transaction
	if err := query.Find(&articleVotes).Error; err != nil {
		return articleVotes, err
	}

	return articleVotes, nil
}

// GetVotesForArticles will return the vote count for an array of articles
func (m *Manager) GetVotesForArticles(articleIDs []string) ([]models.ArticleVotes, error) {
	var articleVotesArray []models.ArticleVotes

	// Build query
	query := m.Store.Database.
		Table("votes").
		Select("article_id, count(*) as votes").
		Where("article_id IN (?)", articleIDs).
		Group("article_id")

	// Commit transaction
	if err := query.Find(&articleVotesArray).Error; err != nil {
		return articleVotesArray, err
	}

	return articleVotesArray, nil
}
