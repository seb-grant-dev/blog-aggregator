package feeds

import (
	"fmt"
	"io"
	"context"
	"net/http"
	"encoding/xml"
	"github.com/seb-grant-dev/blog-aggregator/models"
)



func FetchFeed(ctx context.Context, feedURL string) (*models.RSSFeed,error) {
	req,err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve feed from: %s",feedURL)
	}

	client := &http.Client{}

	req.Header.Set("User-Agent","gator")

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP Request failed: %s",err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Error decoding the response")
	}


	var rssContent = models.RSSFeed{}

	err = xml.Unmarshal(body,&rssContent)
	if err != nil {
		return nil, fmt.Errorf("There was a problem decoding the RSS content pulled from the server")
	}

	fmt.Printf("Retrieved %d items from the feed\n",len(rssContent.Channel.Item))

	return &rssContent, nil
}
