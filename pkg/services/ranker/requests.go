package ranker

type articlesRequest struct {
	ArticleIDs []string `json:"articles"`
}

type userRequest struct {
	UserID string `json:"user"`
}

type voteRequest struct {
	ArticleID string `json:"article"`
	UserID    string `json:"user"`
}
