package model

import (
	"encoding/xml"
	"time"
)

type Opml struct {
	XMLName xml.Name `xml:"opml"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Head    OpmlHead `xml:"head"`
	Body    OpmlBody `xml:"body"`
}

type OpmlHead struct {
	Text       string `xml:",chardata"`
	Title      string `xml:"title"`
	OwnerEmail string `xml:"ownerEmail"`
}

type OpmlBody struct {
	Text         string       `xml:",chardata"`
	OutlineOuter OutlineOuter `xml:"outline"`
}

type OutlineOuter struct {
	Text     string        `xml:",chardata"`
	AttrText string        `xml:"text,attr"`
	Title    string        `xml:"title,attr"`
	Outline  []OpmlOutline `xml:"outline"`
}

type OpmlOutline struct {
	Text     string `xml:",chardata"`
	Type     string `xml:"type,attr"`
	XmlUrl   string `xml:"xmlUrl,attr"`
	AttrText string `xml:"text,attr"`
	HtmlUrl  string `xml:"htmlUrl,attr"`
	Title    string `xml:"title,attr"`
}

type Feed struct {
	ID          string `json:"id" clover:"id"`
	Title       string `json:"title" clover:"title"`
	Description string `json:"description" clover:"description"`
	Text        string `json:"text" clover:"text"`
	Type        string `json:"type" clover:"type"`
	HtmlUrl     string `json:"html_url" clover:"html_url"`
	XMLUrl      string `json:"xml_url" clover:"xml_url"`
}

type BlogRoll struct {
	ID          string
	Title       string
	Tag         string
	Category    string
	DateCreated time.Time
	Feeds       []Feed
}
