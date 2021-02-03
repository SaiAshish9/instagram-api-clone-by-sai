package status

type Link struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"url"`
	Msg         string `json:"msg"`
	StatusImg   string `json:"status-img"`
	StatusVideo string `json:"status-video"`
	StatusMsg   string `json:"status-msg"`
	Time        string `json:"time"`
}
