package gisqus

import (
	"context"
	"errors"
	"net/url"
	"time"
)

type ThreadsURLs struct {
	Thread_list       string
	Thread_detail_url string
	Thread_posts      string
	Thread_hot        string
	Thread_popular    string
	Thread_trending   string
}

var threadsUrls = ThreadsURLs{
	Thread_list:       "https://disqus.com/api/3.0/threads/list.json",
	Thread_detail_url: "https://disqus.com/api/3.0/threads/details.json",
	Thread_posts:      "https://disqus.com/api/3.0/threads/listPosts.json",
	Thread_hot:        "https://disqus.com/api/3.0/threads/listHot.json",
	Thread_popular:    "https://disqus.com/api/3.0/threads/listPopular.json",
	Thread_trending:   "https://disqus.com/api/3.0/trends/listThreads.json",
}

func (gisqus *Gisqus) ThreadList(values url.Values, ctx context.Context) (*ThreadListResponse, error) {

	values.Set("api_secret", gisqus.secret)
	url := threadsUrls.Thread_list + "?" + values.Encode()

	var tlr ThreadListResponse

	err := gisqus.callAndInflate(url, &tlr, ctx)
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

func (gisqus *Gisqus) ThreadTrending(values url.Values, ctx context.Context) (*ThreadTrendingResponse, error) {

	values.Set("api_secret", gisqus.secret)
	url := threadsUrls.Thread_trending + "?" + values.Encode()

	var tlr ThreadTrendingResponse

	err := gisqus.callAndInflate(url, &tlr, ctx)
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

func (gisqus *Gisqus) ThreadDetails(threadId string, values url.Values, ctx context.Context) (*ThreadDetailResponse, error) {

	if threadId == "" {
		return nil, errors.New("Must provide thread id")
	}
	values.Set("thread", threadId)
	values.Set("api_secret", gisqus.secret)
	url := threadsUrls.Thread_detail_url + "?" + values.Encode()

	var tdr ThreadDetailResponse
	err := gisqus.callAndInflate(url, &tdr, ctx)
	if err != nil {
		return nil, err
	}

	tdr.Response.CreatedAt, err = fromDisqusTime(tdr.Response.DisqusTimeCreatedAt)
	if err != nil {
		return nil, err
	}

	return &tdr, nil
}

func (gisqus *Gisqus) ThreadPosts(thread string, values url.Values, ctx context.Context) (*PostListResponse, error) {

	if thread == "" {
		return nil, errors.New("Must provide a thread id")
	}
	values.Set("thread", thread)
	values.Set("api_secret", gisqus.secret)
	url := threadsUrls.Thread_posts + "?" + values.Encode()

	var plr PostListResponse

	err := gisqus.callAndInflate(url, &plr, ctx)
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

func (gisqus *Gisqus) ThreadHot(values url.Values, ctx context.Context) (*ThreadListResponseNoCursor, error) {

	values.Set("api_secret", gisqus.secret)
	url := threadsUrls.Thread_hot + "?" + values.Encode()

	var tlr ThreadListResponseNoCursor

	err := gisqus.callAndInflate(url, &tlr, ctx)
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

func (gisqus *Gisqus) ThreadPopular(values url.Values, ctx context.Context) (*ThreadListResponseNoCursor, error) {

	values.Set("api_secret", gisqus.secret)
	url := threadsUrls.Thread_popular + "?" + values.Encode()

	var tlr ThreadListResponseNoCursor

	err := gisqus.callAndInflate(url, &tlr, ctx)
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

type ThreadDetailResponse struct {
	ResponseStub
	Response *ThreadDetail `json:"response"`
}

type ThreadListResponse struct {
	ResponseStubWithCursor
	Response []*Thread `json:"response"`
}

type ThreadListResponseNoCursor struct {
	ResponseStub
	Response []*Thread `json:"response"`
}

type ThreadTrendingResponse struct {
	ResponseStub
	Response []*Trend `json:"response"`
}

type Trend struct {
	TrendingThread *Thread `json:"thread"`
	PostLikes      int     `json:"postLikes"`
	Posts          int     `json:"posts"`
	Score          float32 `json:"score"`
	Link           string  `json:"link"`
	Likes          int     `json:"likes"`
}

type ThreadDetail struct {
	Thread
	CanModerate bool `json:"canModerate"`
}

type Thread struct {
	Feed                string    `json:"feed"`
	Identifiers         []string  `json:"identifiers"`
	Dislikes            int       `json:"dislikes"`
	Likes               int       `json:"likes"`
	Message             string    `json:"message"`
	Id                  string    `json:"id"`
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
