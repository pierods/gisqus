package gisqus

import (
	"context"
	"errors"
	"net/url"
	"time"
)

type PostsURLs struct {
	post_details_url string
	post_list_url    string
}

var postsUrls = PostsURLs{
	post_details_url: "https://disqus.com/api/3.0/posts/details.json",
	post_list_url:    "https://disqus.com/api/3.0/posts/list.json",
}

func (gisqus *Gisqus) PostDetails(postId string, values url.Values, ctx context.Context) (*PostDetailsResponse, error) {

	if postId == "" {
		return nil, errors.New("Must use post parameter")
	}
	values.Set("post", postId)
	values.Set("api_secret", gisqus.secret)
	url := postsUrls.post_details_url + "?" + values.Encode()

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
https://disqus.com/api/docs/posts/list/
*/
func (gisqus *Gisqus) PostList(values url.Values, ctx context.Context) (*PostListResponse, error) {

	values.Set("api_secret", gisqus.secret)
	url := postsUrls.post_list_url + "?" + values.Encode()

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

type PostListResponse struct {
	ResponseStubWithCursor
	Response []*Post `json:"response"`
}

type PostDetailsResponse struct {
	ResponseStub
	Response *Post `json:"response"`
}

type Post struct {
	Dislikes            int          `json:"dislikes"`
	NumReports          int          `json:"numReports"`
	Likes               int          `json:"likes"`
	Message             string       `json:"message"`
	Id                  string       `json:"id"`
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

type PostMedia struct {
}

type PostAuthor struct {
	Username                string      `json:"username"`
	About                   string      `json:"about"`
	Name                    string      `json:"name"`
	Disable3rdPartyTrackers bool        `json:"disable3rdPartyTrackers"`
	Url                     string      `json:"url"`
	IsAnonymous             bool        `json:"isAnonymous"`
	ProfileUrl              string      `json:"profileUrl"`
	IsPowerContributor      bool        `json:"isPowerContributor"`
	Location                string      `json:"location"`
	IsPrivate               bool        `json:"isPrivate"`
	SignedUrl               string      `json:"signedUrl"`
	IsPrimary               bool        `json:"isPrimary"`
	JoinedAt                time.Time   `json:"-"`
	DisqusTimeJoinedAt      string      `json:"joinedAt"`
	Id                      string      `json:"id"`
	Avatar                  *UserAvatar `json:"avatar"`
}
