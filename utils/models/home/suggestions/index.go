package suggestions

type Suggestion struct {
	ID    int    `json:"id"`
	Image string `json:"image"`
	Title string `json:"title"`
	Desc1 string `json:"desc1"`
	Desc2 string `json:"desc2"`
}
