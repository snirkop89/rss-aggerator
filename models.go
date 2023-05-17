package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/snirkop89/rss-aggregator/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseFeedToFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
		UserID:    feed.UserID,
	}
}

func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	var feeds []Feed
	for _, f := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(f))
	}
	return feeds
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func databaseFeedFollowToFeedFollow(ff database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        ff.ID,
		CreatedAt: ff.CreatedAt,
		UpdatedAt: ff.UpdatedAt,
		UserID:    ff.UserID,
		FeedID:    ff.FeedID,
	}
}

func databaseFeedFollowsToFeedFollows(fs []database.FeedFollow) []FeedFollow {
	var follows []FeedFollow
	for _, ff := range fs {
		follows = append(follows, databaseFeedFollowToFeedFollow(ff))
	}
	return follows
}
