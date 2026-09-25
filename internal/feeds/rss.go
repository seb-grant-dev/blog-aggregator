package feeds


import (
	"fmt"
	"io"
	"context"
	"time"
	"net/http"
	"encoding/xml"
	"database/sql"
 	"github.com/lib/pq"
	"github.com/google/uuid"
	"github.com/seb-grant-dev/blog-aggregator/models"
	"github.com/seb-grant-dev/blog-aggregator/state"
	"github.com/seb-grant-dev/blog-aggregator/internal/database"
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

func ScrapeFeeds(s *state.State) error {
	ctx := context.Background()
	feed, err := s.DB.GetNextFeedToFetch(ctx)
	if err != nil {
		return err
	}

	// convert to Nullable time for SQL
	fetchTime := time.Now()

	fmt.Println("Fetching feed from: " + feed.Url)

	feedFetchedParams := database.MarkFeedFetchedParams{
		ID: feed.ID,
		LastFetchedAt: sql.NullTime{
			Time: fetchTime,
			Valid: true,
		},
		UpdatedAt: fetchTime,
	}

	err = s.DB.MarkFeedFetched(ctx,feedFetchedParams)
	if err != nil {
		return err
	}

	fetchedFeed, err := FetchFeed(ctx, feed.Url)

	for _, item := range fetchedFeed.Channel.Item {

		parsedPubDate, err := time.Parse(time.Layout,item.PubDate)
		pubDate := sql.NullTime{
			Time: parsedPubDate,
			Valid: true,
		}

		parsedDescription := sql.NullString {
			String: item.Description,
			Valid: true,
		}

		newPost := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: parsedDescription,
			PublishedAt: pubDate,
			FeedID:      feed.ID,
		}

		_,err = s.DB.CreatePost(ctx,newPost)

		if err != nil {

			if pqErr, ok := err.(*pq.Error); ok {
				code := pqErr.Code
				switch code {
				case "23505":
					continue
				default:
					fmt.Errorf("Error saving post: %s\n",err)
				}
			}
		} else {
			fmt.Printf("Saving Post: %s\n",item.Title)
		}

	}

	return nil
}
