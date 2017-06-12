package gisqus

import (
	"context"
	"errors"
	"net/url"
	"time"
)

// PostsURLS are the URLS of the Post endpoints in Disqus' API
type PostsURLS struct {
	postDetailsURL string
	postListURL    string
}

var postsUrls = PostsURLS{
	postDetailsURL: "https://disqus.com/api/3.0/posts/details.json",
	postListURL:    "https://disqus.com/api/3.0/posts/list.json",
}

/*
PostDetails wraps https://disqus.com/api/docs/posts/details/ (https://disqus.com/api/3.0/posts/details.json)
It does not support the "related" argument.
*/
func (gisqus *Gisqus) PostDetails(ctx context.Context, postID string, values url.Values) (*PostDetailsResponse, error) {

	if postID == "" {
		return nil, errors.New("Must use post parameter")
	}
	values.Set("post", postID)
	values.Set("api_secret", gisqus.secret)
	url := postsUrls.postDetailsURL + "?" + values.Encode()

	var pdr PostDetailsResponse

	err := gisqus.callAndInflate(url, &pdr, ctx)
	if err != nil {
		return nil, err
	}
	pdr.Response.CreatedAt, err = fromDisqusTime(pdr.Response.DisqusTimeCreatedAt)
	if err != nil {
		return nil, err
	}
	pdr.Response.Author.JoinedAt, err = fromDisqusTime(pdr.Response.Author.DisqusTimeJoinedAt)
	if err != nil {
		return nil, err
	}
	return &pdr, nil
}

/*
PostList wraps https://disqus.com/api/docs/posts/list/ (https://disqus.com/api/3.0/posts/list.json)
*/
func (gisqus *Gisqus) PostList(ctx context.Context, values url.Values) (*PostListResponse, error) {

	values.Set("api_secret", gisqus.secret)
	url := postsUrls.postListURL + "?" + values.Encode()

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

// PostListResponse wraps the response of the post list endpoint
type PostListResponse struct {
	ResponseStubWithCursor
	Response []*Post `json:"response"`
}

// PostDetailsResponse wraps the response of the post details endpoint
type PostDetailsResponse struct {
	ResponseStub
	Response *Post `json:"response"`
}

// Post models a Post as returned by Disqus' API
type Post struct {
	Dislikes            int          `json:"dislikes"`
	NumReports          int          `json:"numReports"`
	Likes               int          `json:"likes"`
	Message             string       `json:"message"`
	ID                  string       `json:"id"`
	IsDeleted           bool         `json:"isDeleted"`
	Author              *PostAuthor  `json:"author"`
	Media               []*PostMedia `json:"media"`
	IsSpam              bool         `json:"isSpam"`
	IsDeletedByAuthor   bool         `json:"isDeletedByAuthor"`
	CreatedAt           time.Time    `json:"-"`
	DisqusTimeCreatedAt string       `json:"createdAt"`
	Parent              int          `json:"parent"`
	IsApproved          bool         `json:"isApproved"`
	IsFlagged           bool         `json:"isFlagged"`
	RawMessage          string       `json:"rawMessage"`
	IsHighlighted       bool         `json:"isHighlighted"`
	CanVote             bool         `json:"canVote"`
	Thread              string       `json:"thread"`
	Forum               string       `json:"forum"`
	Points              int          `json:"points"`
	ModerationLabels    []string     `json:"moderationLabels"`
	IsEdited            bool         `json:"isEdited"`
	Sb                  bool         `json:"sb"`
}

// PostMedia models the fields the media field in a Post
type PostMedia struct {
}

// PostAuthor models the fields of the author field in a Post
type PostAuthor struct {
	Username                string      `json:"username"`
	About                   string      `json:"about"`
	Name                    string      `json:"name"`
	Disable3rdPartyTrackers bool        `json:"disable3rdPartyTrackers"`
	URL                     string      `json:"url"`
	IsAnonymous             bool        `json:"isAnonymous"`
	ProfileURL              string      `json:"profileUrl"`
	IsPowerContributor      bool        `json:"isPowerContributor"`
	Location                string      `json:"location"`
	IsPrivate               bool        `json:"isPrivate"`
	SignedURL               string      `json:"signedUrl"`
	IsPrimary               bool        `json:"isPrimary"`
	JoinedAt                time.Time   `json:"-"`
	DisqusTimeJoinedAt      string      `json:"joinedAt"`
	ID                      string      `json:"id"`
	Avatar                  *UserAvatar `json:"avatar"`
}
