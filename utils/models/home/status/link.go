package status

type Link struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"url"`
	Msg   string `json:"msg"`
}
