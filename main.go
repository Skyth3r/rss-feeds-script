package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mmcdole/gofeed"
)

func main() {
	rssFeeds := []string{"https://akashgoswami.com/articles/index.xml", "https://akashgoswami.dev/index.xml", "https://bradleyjkemp.dev/index.xml"}
	feedItems := []gofeed.Item{}

	for _, v := range rssFeeds {
		feedParser := gofeed.NewParser()
		feed, err := feedParser.ParseURL(v)
		if err != nil {
			log.Fatalf("Error getting feed: %v", err)
		}

		for _, item := range feed.Items {
			feedItems = append(feedItems, *item)
		}
	}

	jsonFeedItems, err := json.Marshal(feedItems)
	if err != nil {
		log.Fatalf("Error marshaling data to JSON: %v", err)
	}
	fmt.Println("feedItems as JSON: ", string(jsonFeedItems))
}
