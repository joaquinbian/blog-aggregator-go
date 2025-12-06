package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func handlerFetchFeed(state *State, cmd Command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")

	if err != nil {
		return err
	}

	fmt.Printf("%v\n", feed)

	return nil
}

func fetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error) {

	var feed = &RSSFeed{}

	req, err := http.NewRequestWithContext(ctx, "GET", feedUrl, http.NoBody)

	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error generando request: %v", err)
	}

	req.Header.Set("User-Agent", "gator")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error en la request: %v", err)
	}

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error leyendo body: %v", err)
	}

	if err = xml.Unmarshal(data, feed); err != nil {
		return &RSSFeed{}, fmt.Errorf("error unmarshall data: %v", err)
	}

	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)

	for idx := range feed.Channel.Item {
		feed.Channel.Item[idx].Description = html.UnescapeString(feed.Channel.Item[idx].Description)
		feed.Channel.Item[idx].Title = html.UnescapeString(feed.Channel.Item[idx].Title)
	}

	return feed, nil

}
