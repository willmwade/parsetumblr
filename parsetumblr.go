// parsetumblr project parseblogger.go
package parsetumblr

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Feed struct {
	Url        string
	Limit      int64     //Default 20 Max 50
	StartIndex int64     //0 Indexed
	XMLName    xml.Name  `xml:"tumblr"`
	Tumblelog  Tumblelog `xml:"tumblelog"`
	Entries    Entries   `xml:"posts"`
}

type Tumblelog struct {
	XMLName xml.Name `xml:"tumblelog"`
	Name    string   `xml:"name,attr"`
	Title   string   `xml:"title,attr"`
}

type Entries struct {
	XMLName xml.Name `xml:"posts"`
	Start   int64    `xml:"start,attr"`
	Total   int64    `xml:"total,attr"`
	Entries []Entry  `xml:"post"`
}

type Entry struct {
	//Raw     string   `xml:",innerxml"`
	XMLName           xml.Name `xml:"post"`
	Id                string   `xml:"id,attr"`
	Url               string   `xml:"url,attr"`
	UrlWithSlug       string   `xml:"url-with-slug,attr"`
	Type              string   `xml:"type,attr"`
	UnixTimestamp     int64    `xml:"unix-timestamp,attr"`
	published         time.Time
	ReblogKey         string           `xml:"reblog-key,attr"`
	Slug              string           `xml:"slug,attr"`
	RegularTitle      string           `xml:"regular-title"`
	RegularBody       string           `xml:"regular-body"`
	LinkText          string           `xml:"link-text"`
	LinkUrl           string           `xml:"link-url"`
	QuoteText         string           `xml:"quote-text"`
	QuoteSource       string           `xml:"quote-source"`
	PhotoCaption      string           `xml:"photo-caption"`
	Photos            []Photo          `xml:"photo-url"`
	ConversationTitle string           `xml:"conversation-title"`
	ConversationText  string           `xml:"conversation-text"`
	ConversationLines ConversationLine `xml:"line"`
	VideoCaption      string           `xml:"video-caption"`
	VideoSource       string           `xml:"video-source"`
	VideoPlayer       string           `xml:"video-player"`
	AudioCaption      string           `xml:"audio-caption"`
	AudioPlayer       string           `xml:"audio-player"`
	Question          string           `xml:"question"`
	Answer            string           `xml:"answer"`
}

type Photo struct {
	Size int64  `xml:"max-width,attr"`
	Url  string `xml:",innerxml"`
}

type ConversationLine struct {
	Name    string `xml:"name,attr"`
	Label   string `xml:"label,attr"`
	Content string `xml:",innerxml"`
}

func NewFeed(url string) *Feed {
	var f Feed
	f.Url = url
	return &f
}

func (f *Feed) FetchUrl() string {
	url := f.Url + "/api/read?"
	if f.Limit != 0 {
		url += "num=" + strconv.FormatInt(f.Limit, 10)
	} else {
		url += "num=20"
	}

	if f.StartIndex != 0 {
		url += "&start=" + strconv.FormatInt(f.StartIndex, 10)
	}

	return url
}

func (f *Feed) GetFeed(client *http.Client) error {
	xmlrsp, err := client.Get(f.FetchUrl())
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(xmlrsp.Body)
	if err != nil {
		return err
	}
	if err := xml.Unmarshal(body, f); err != nil {
		return err
	}
	return nil
}

func (e Entry) Published() time.Time {
	if e.published.IsZero() {
		e.published = time.Unix(e.UnixTimestamp, 0)
	}

	return e.published
}
