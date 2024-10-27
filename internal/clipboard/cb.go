package clipboard

import (
	"context"
	"fmt"
	"github.com/r4ndsen/sqlc-demo/internal/db"
	"golang.design/x/clipboard"
	"log"
	"regexp"
)

type Watcher struct {
	s db.Querier
}

func New(s db.Querier) *Watcher {
	return &Watcher{s: s}
}

var isUrl = regexp.MustCompile("^https?://").Match

func (c *Watcher) Watch() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	ch := clipboard.Watch(context.TODO(), clipboard.FmtText)

	go func() {
		for data := range ch {
			s, err := c.getString(data)
			if err != nil {
				log.Println(err)
				continue
			}

			link, err := c.s.CreateLink(context.Background(), s)
			if err != nil {
				log.Printf("failed to add record: %v %v", err, link)
			}
		}
	}()
}

func (c *Watcher) getString(data []byte) (string, error) {
	if isUrl(data) {
		return string(data), nil
	}

	return "", fmt.Errorf("not a url: %v", string(data))
}
