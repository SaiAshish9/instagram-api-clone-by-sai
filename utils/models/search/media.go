package search

type Media struct {
	ID         int    `json:"id"`
	URL        string `json:"url"`
	IS_VIDEO   int    `json:"is_video"`
	IS_GALLERY int    `json:"is_gallery"`
}
