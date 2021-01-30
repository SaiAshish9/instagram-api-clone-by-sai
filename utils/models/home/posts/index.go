package posts

type Post struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	URL           string `json:"url"`
	ProfileURL    string
	Title         string
	Description   string
	CommentsCount string
	Time          string
}
