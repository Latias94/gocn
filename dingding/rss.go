package dingding

import (
	"fmt"
	"github.com/gorilla/feeds"
	"log"
	"time"
)

var items []*feeds.Item
var feed = &feeds.Feed{
	Title:       "GOCN 每日新闻",
	Link:        &feeds.Link{Href: "https://github.com/gocn/news"},
	Description: "GOCN 每日新闻",
	Author:      &feeds.Author{Name: "Frankorz", Email: "frankorz@qq.com"},
	Created:     time.Now(),
}


func GenerateRSS(rssContent RSSContent) string {
	item := feeds.Item{
		Title:       rssContent.Message.DailyTitle,
		Link:        &feeds.Link{Href: rssContent.Message.PostUrl},
		Description: rssContent.Markdown.Content,
		Author:      &feeds.Author{Name: rssContent.Message.Author},
		Created:     time.Now(),
	}
	items = append(items, &item)
	feed.Items = items

	rss, err := feed.ToRss()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Generate new RSS item")
	return rss
}
