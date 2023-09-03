package model

type AddRequest struct {
	Url string `json:"url" validate:"required"`
}

type AddResponse struct {
	ShortUrl string `json:"short_url"`
}
