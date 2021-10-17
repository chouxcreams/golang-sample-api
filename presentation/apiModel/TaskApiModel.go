package apiModel

type TaskCommand struct {
	Text *string `json:"text"`
}

type TaskResponse struct {
	Text *string `json:"text"`
	Id   int     `json:"id"`
}
