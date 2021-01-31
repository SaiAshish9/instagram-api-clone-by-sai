package search

type Media struct {
	ID         int    `json:"id"`
	URL        string `json:"url"`
	IS_VIDEO   string `json:"is_video"`
	IS_GALLERY string `json:"is_gallery"`
}
