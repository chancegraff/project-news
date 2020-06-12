package collector

type allRequest struct {
	Offset int `json:"offset"`
}

type getRequest struct {
	ID int `json:"id"`
}
