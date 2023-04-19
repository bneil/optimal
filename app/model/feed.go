package model

type Feed struct {
	ID      string `json:"id" clover:"id"`
	Title   string `json:"title" clover:"title"`
	Text    string `json:"text" clover:"text"`
	Type    string `json:"type" clover:"type"`
	HtmlUrl string `json:"html_url" clover:"html_url"`
	XMLUrl  string `json:"xml_url" clover:"xml_url"`
}

type List struct {
	Feeds []Feed
}
