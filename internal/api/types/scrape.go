package types

type ScrapeRequest struct {
	URL    string `json:"url" validate:"required,url"`
	Prompt string `json:"prompt" validate:"required,min=10,max=3000"`
}

type ScrapeResponse struct {
	Results string `json:"results"`
	URL     string `json:"url"`
	Prompt  string `json:"prompt"`
}
