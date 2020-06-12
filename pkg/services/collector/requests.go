package collector

type allRequest struct {
	Offset int `json:"offset"`
}

type getRequest struct {
	ID int `json:"id"`
}

// Proxies

type articlesRequest struct {
	ArticleIDs []string `json:"articles"`
}
