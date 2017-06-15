// Copyright Piero de Salvia.
// All Rights Reserved

package gisqus

import (
	"context"
	"errors"
	"net/url"
	"time"
)

// ThreadsURLS are the URLs of the thread endpoints of the Disqus' API
type ThreadsURLS struct {
	ThreadListURL       string
	ThreadDetailURL     string
	ThreadPostsURL      string
	ThreadHotURL        string
	ThreadPopularURL    string
	ThreadTrendingURL   string
	ThreadUsersVotedURL string
	ThreadSetURL        string
}

var threadsUrls = ThreadsURLS{
	ThreadListURL:       "https://disqus.com/api/3.0/threads/list.json",
	ThreadDetailURL:     "https://disqus.com/api/3.0/threads/details.json",
	ThreadPostsURL:      "https://disqus.com/api/3.0/threads/listPosts.json",
	ThreadHotURL:        "https://disqus.com/api/3.0/threads/listHot.json",
	ThreadPopularURL:    "https://disqus.com/api/3.0/threads/listPopular.json",
	ThreadTrendingURL:   "https://disqus.com/api/3.0/trends/listThreads.json",
	ThreadUsersVotedURL: "https://disqus.com/api/3.0/threads/listUsersVotedThread.json",
	ThreadSetURL:        "https://disqus.com/api/3.0/threads/set.json",
}

/*
ThreadUsersVoted wraps https://disqus.com/api/docs/threads/listUsersVotedThread/ (https://disqus.com/api/3.0/threads/listUsersVotedThread.json)
Complete users are not returned by Disqus on this call
*/
func (gisqus *Gisqus) ThreadUsersVoted(ctx context.Context, threadD string, values url.Values) (*UsersVotedResponse, error) {

	if threadD == "" {
		return nil, errors.New("Must provide a thread id")
	}
	values.Set("thread", threadD)
	values.Set("api_secret", gisqus.secret)
	url := threadsUrls.ThreadUsersVotedURL + "?" + values.Encode()

	var uvr UsersVotedResponse

	err := gisqus.callAndInflate(ctx, url, &uvr)
	if err != nil {
		return nil, err
	}

	for _, user := range uvr.Response {
		user.JoinedAt, err = fromDisqusTime(user.DisqusTimeJoinedAt)
		if err != nil {
			return nil, err
		}
	}
	return &uvr, nil
}

/*
ThreadList wraps https://disqus.com/api/docs/threads/list/ (https://disqus.com/api/3.0/threads/list.json)
It does not support the "related" argument (related fields can be gotten with calls to their respective APIS)
*/
func (gisqus *Gisqus) ThreadList(ctx context.Context, values url.Values) (*ThreadListResponse, error) {

	values.Set("api_secret", gisqus.secret)
	url := threadsUrls.ThreadListURL + "?" + values.Encode()

	var tlr ThreadListResponse

	err := gisqus.callAndInflate(ctx, url, &tlr)
	if err != nil {
		return nil, err
	}
	for _, thread := range tlr.Response {
		thread.CreatedAt, err = fromDisqusTime(thread.DisqusTimeCreatedAt)
		if err != nil {
			return nil, err
		}
	}
	return &tlr, nil

}

/*
ThreadTrending wraps https://disqus.com/api/docs/trends/listThreads/ (https://disqus.com/api/3.0/trends/listThreads.json)
It does not support the "related" argument (related fields can be gotten with calls to their respective APIS)
*/
func (gisqus *Gisqus) ThreadTrending(ctx context.Context, values url.Values) (*ThreadTrendingResponse, error) {

	values.Set("api_secret", gisqus.secret)
	url := threadsUrls.ThreadTrendingURL + "?" + values.Encode()

	var tlr ThreadTrendingResponse

	err := gisqus.callAndInflate(ctx, url, &tlr)
	if err != nil {
		return nil, err
	}
	for _, trend := range tlr.Response {
		trend.TrendingThread.CreatedAt, err = fromDisqusTime(trend.TrendingThread.DisqusTimeCreatedAt)
		if err != nil {
			return nil, err
		}
		if trend.TrendingThread.HighlightedPost != nil {
			trend.TrendingThread.HighlightedPost.CreatedAt, err = fromDisqusTime(trend.TrendingThread.HighlightedPost.DisqusTimeCreatedAt)
			if err != nil {
				return nil, err
			}
		}
		if trend.TrendingThread.HighlightedPost != nil && trend.TrendingThread.HighlightedPost.Author != nil {
			trend.TrendingThread.HighlightedPost.Author.JoinedAt, err = fromDisqusTime(trend.TrendingThread.HighlightedPost.Author.DisqusTimeJoinedAt)
			if err != nil {
				return nil, err
			}
		}
	}
	return &tlr, nil

}

/*
ThreadSet wraps https://disqus.com/api/docs/threads/set/ (https://disqus.com/api/3.0/threads/set.json)
*/
func (gisqus *Gisqus) ThreadSet(ctx context.Context, threadsIDs []string, values url.Values) (*ThreadListResponseNoCursor, error) {

	if threadsIDs == nil || len(threadsIDs) == 0 {
		return nil, errors.New("Must provide one or more thread ids")
	}
	for _, thread := range threadsIDs {
		values.Add("thread", thread)
	}
	values.Set("api_secret", gisqus.secret)
	url := threadsUrls.ThreadSetURL + "?" + values.Encode()

	var tlr ThreadListResponseNoCursor

	err := gisqus.callAndInflate(ctx, url, &tlr)
	if err != nil {
		return nil, err
	}
	for _, thread := range tlr.Response {
		thread.CreatedAt, err = fromDisqusTime(thread.DisqusTimeCreatedAt)
		if err != nil {
			return nil, err
		}
	}
	return &tlr, nil

}

/*
ThreadDetails wraps https://disqus.com/api/docs/threads/details/ (https://disqus.com/api/3.0/threads/details.json)
It does not support the "related" argument (related fields can be gotten with calls to their respective APIS)
*/
func (gisqus *Gisqus) ThreadDetails(ctx context.Context, threadID string, values url.Values) (*ThreadDetailResponse, error) {

	if threadID == "" {
		return nil, errors.New("Must provide thread id")
	}
	values.Set("thread", threadID)
	values.Set("api_secret", gisqus.secret)
	url := threadsUrls.ThreadDetailURL + "?" + values.Encode()

	var tdr ThreadDetailResponse
	err := gisqus.callAndInflate(ctx, url, &tdr)
	if err != nil {
		return nil, err
	}

	tdr.Response.CreatedAt, err = fromDisqusTime(tdr.Response.DisqusTimeCreatedAt)
	if err != nil {
		return nil, err
	}

	return &tdr, nil
}

/*
ThreadPosts wraps https://disqus.com/api/docs/threads/listPosts/ (https://disqus.com/api/3.0/threads/listPosts.json)
It does not support the "related" argument (related fields can be gotten with calls to their respective APIS)
*/
func (gisqus *Gisqus) ThreadPosts(ctx context.Context, threadID string, values url.Values) (*PostListResponse, error) {

	if threadID == "" {
		return nil, errors.New("Must provide a thread id")
	}
	values.Set("thread", threadID)
	values.Set("api_secret", gisqus.secret)
	url := threadsUrls.ThreadPostsURL + "?" + values.Encode()

	var plr PostListResponse

	err := gisqus.callAndInflate(ctx, url, &plr)
	if err != nil {
		return nil, err
	}

	for _, post := range plr.Response {
		post.CreatedAt, err = fromDisqusTime(post.DisqusTimeCreatedAt)
		if err != nil {
			return nil, err
		}
		post.Author.JoinedAt, err = fromDisqusTime(post.Author.DisqusTimeJoinedAt)

		if err != nil {
			return nil, err
		}
	}

	return &plr, nil
}

/*
ThreadHot wraps https://disqus.com/api/docs/threads/listHot/ (https://disqus.com/api/3.0/threads/listHot.json)
It does not support the "related" argument (related fields can be gotten with calls to their respective APIS)
*/
func (gisqus *Gisqus) ThreadHot(ctx context.Context, values url.Values) (*ThreadListResponseNoCursor, error) {

	values.Set("api_secret", gisqus.secret)
	url := threadsUrls.ThreadHotURL + "?" + values.Encode()

	var tlr ThreadListResponseNoCursor

	err := gisqus.callAndInflate(ctx, url, &tlr)
	if err != nil {
		return nil, err
	}
	for _, thread := range tlr.Response {
		thread.CreatedAt, err = fromDisqusTime(thread.DisqusTimeCreatedAt)
		if err != nil {
			return nil, err
		}
	}
	return &tlr, nil
}

/*
ThreadPopular wraps https://disqus.com/api/docs/threads/listPopular/ (https://disqus.com/api/3.0/threads/listPopular.json)
It does not support the "related" argument (related fields can be gotten with calls to their respective APIS)
*/
func (gisqus *Gisqus) ThreadPopular(ctx context.Context, values url.Values) (*ThreadListResponseNoCursor, error) {

	values.Set("api_secret", gisqus.secret)
	url := threadsUrls.ThreadPopularURL + "?" + values.Encode()

	var tlr ThreadListResponseNoCursor

	err := gisqus.callAndInflate(ctx, url, &tlr)
	if err != nil {
		return nil, err
	}
	for _, thread := range tlr.Response {
		thread.CreatedAt, err = fromDisqusTime(thread.DisqusTimeCreatedAt)
		if err != nil {
			return nil, err
		}
	}
	return &tlr, nil
}

// UsersVotedResponse models the response of the users voted thread endpoint
type UsersVotedResponse struct {
	ResponseStub
	Response []*User `json:"response"`
}

// ThreadDetailResponse models the response of the thread details endpoint.
type ThreadDetailResponse struct {
	ResponseStub
	Response *ThreadDetail `json:"response"`
}

// ThreadListResponse models the response of various thread endpoints.
type ThreadListResponse struct {
	ResponseStubWithCursor
	Response []*Thread `json:"response"`
}

// ThreadListResponseNoCursor models the response of various thread endpoints.
type ThreadListResponseNoCursor struct {
	ResponseStub
	Response []*Thread `json:"response"`
}

// ThreadTrendingResponse models the response of the trending threads endpoint.
type ThreadTrendingResponse struct {
	ResponseStub
	Response []*Trend `json:"response"`
}

// Trend models the trend returned by the trending threads endpoint
type Trend struct {
	TrendingThread *Thread `json:"thread"`
	PostLikes      int     `json:"postLikes"`
	Posts          int     `json:"posts"`
	Score          float32 `json:"score"`
	Link           string  `json:"link"`
	Likes          int     `json:"likes"`
}

// ThreadDetail models the fields returned by the thread detail endpoint
type ThreadDetail struct {
	Thread
	CanModerate bool `json:"canModerate"`
}

// Thread models the Thread returned by Disqus' API calls
type Thread struct {
	Feed                string    `json:"feed"`
	Identifiers         []string  `json:"identifiers"`
	Dislikes            int       `json:"dislikes"`
	Likes               int       `json:"likes"`
	Message             string    `json:"message"`
	ID                  string    `json:"id"`
	IsDeleted           bool      `json:"isDeleted"`
	Category            string    `json:"category"`
	Author              string    `json:"author"`
	UserScore           int       `json:"userScore"`
	IsSpam              bool      `json:"isSpam"`
	SignedLink          string    `json:"signedLink"`
	CreatedAt           time.Time `json:"-"`
	DisqusTimeCreatedAt string    `json:"createdAt"`
	HasStreaming        bool      `json:"hasStreaming"`
	RawMessage          string    `json:"rawMessage"`
	IsClosed            bool      `json:"isClosed"`
	Link                string    `json:"link"`
	Slug                string    `json:"slug"`
	Forum               string    `json:"forum"`
	CleanTitle          string    `json:"clean_title"`
	Posts               int       `json:"posts"`
	UserSubscription    bool      `json:"userSubscription"`
	Title               string    `json:"title"`
	HighlightedPost     *Post     `json:"highlightedPost"`
}
