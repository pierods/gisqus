package gisqus

import (
	"context"
	"errors"
	"net/url"
	"time"
)

// UsersURLS are the URLs used by Disqus' user endpoints
type UsersURLS struct {
	userDetailURL            string
	userInterestingIUsersURL string
	userPostListURL          string
	userActiveForums         string
	userFollowers            string
	userFollowing            string
	userFollowingForums      string
}

// following forums
// curl "https://disqus.com/api/3.0/users/listFollowingForums.json?user=195792235&api_secret=KpVAypnhCxG27eRLRbXad0i1xfbyUsHPE7E8on5wbFJkbQcIzjB0pkJ4kMOfTRmx" |jq

var usersUrls = UsersURLS{
	userDetailURL:            "https://disqus.com/api/3.0/users/details.json",
	userInterestingIUsersURL: "https://disqus.com/api/3.0/users/interestingUsers.json",
	userPostListURL:          "https://disqus.com/api/3.0/users/listPosts.json",
	userActiveForums:         "https://disqus.com/api/3.0/users/listActiveForums.json",
	userFollowers:            "https://disqus.com/api/3.0/users/listFollowers.json",
	userFollowing:            "https://disqus.com/api/3.0/users/listFollowing.json",
	userFollowingForums:      "https://disqus.com/api/3.0/users/listFollowingForums.json",
}

/*
UserPosts wraps https://disqus.com/api/docs/users/listPosts/ (https://disqus.com/api/3.0/users/listPosts.json)
It does not support the "related" argument (related fields can be gotten with calls to their respective APIS)
*/
func (gisqus *Gisqus) UserPosts(ctx context.Context, userID string, values url.Values) (*PostListResponse, error) {
	if userID == "" {
		return nil, errors.New("Must provide a user id")
	}
	values.Set("api_secret", gisqus.secret)
	values.Set("user", userID)
	url := usersUrls.userPostListURL + "?" + values.Encode()

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

/*
UserDetails wraps https://disqus.com/api/docs/users/details/ (https://disqus.com/api/3.0/users/details.json)
*/
func (gisqus *Gisqus) UserDetails(ctx context.Context, userID string, values url.Values) (*UserDetailsResponse, error) {

	if userID == "" {
		return nil, errors.New("Must provide a user id")
	}
	values.Set("api_secret", gisqus.secret)
	url := usersUrls.userDetailURL + "?" + values.Encode()
	var udr UserDetailsResponse
	err := gisqus.callAndInflate(url, &udr, ctx)
	if err != nil {
		return nil, err
	}

	udr.Response.JoinedAt, err = fromDisqusTime(udr.Response.DisqusTimeJoinedAt)
	if err != nil {
		return nil, err
	}
	return &udr, nil
}

/*
UserInteresting wraps https://disqus.com/api/docs/users/interestingUsers/ (https://disqus.com/api/3.0/users/interestingUsers.json)
*/
func (gisqus *Gisqus) UserInteresting(ctx context.Context, values url.Values) (*InterestingUsersResponse, error) {

	values.Set("api_secret", gisqus.secret)
	url := usersUrls.userInterestingIUsersURL + "?" + values.Encode()
	var iur InterestingUsersResponse

	err := gisqus.callAndInflate(url, &iur, ctx)
	if err != nil {
		return nil, err
	}

	for _, user := range iur.Response.Objects {
		user.JoinedAt, err = fromDisqusTime(user.DisqusTimeJoinedAt)
		if err != nil {
			return nil, err
		}
	}
	return &iur, nil
}

/*
UserActiveForums wraps https://disqus.com/api/docs/users/listActiveForums/ (https://disqus.com/api/3.0/users/listActiveForums.json)
*/
func (gisqus *Gisqus) UserActiveForums(ctx context.Context, user string, values url.Values) (*ActiveForumsResponse, error) {

	if user == "" {
		return nil, errors.New("Must provide a user id")
	}
	values.Set("user", user)
	values.Set("api_secret", gisqus.secret)
	url := usersUrls.userActiveForums + "?" + values.Encode()

	var afr ActiveForumsResponse
	err := gisqus.callAndInflate(url, &afr, ctx)
	if err != nil {
		return nil, err
	}

	for _, forum := range afr.Response {
		forum.CreatedAt, err = fromDisqusTimeExact(forum.DisqusTimeCreatedAt)
		if err != nil {
			return nil, err
		}
	}
	return &afr, nil
}

/*
UserFollowers wraps https://disqus.com/api/docs/users/listFollowers/ (https://disqus.com/api/3.0/users/listFollowers.json)
Numlikes, NumPosts, NumFollowers are not returned by Disqus' API
*/
func (gisqus *Gisqus) UserFollowers(ctx context.Context, userID string, values url.Values) (*UserListResponse, error) {
	if userID == "" {
		return nil, errors.New("Must provide a user id")
	}
	values.Set("user", userID)
	values.Set("api_secret", gisqus.secret)
	url := usersUrls.userFollowers + "?" + values.Encode()
	var fr UserListResponse

	err := gisqus.callAndInflate(url, &fr, ctx)
	if err != nil {
		return nil, err
	}

	for _, user := range fr.Response {
		user.JoinedAt, err = fromDisqusTime(user.DisqusTimeJoinedAt)
		if err != nil {
			return nil, err
		}
	}
	return &fr, nil
}

/*
UserFollowing wraps https://disqus.com/api/docs/users/listFollowing/ (https://disqus.com/api/3.0/users/listFollowing.json)
Numlikes, NumPosts, NumFollowers are not returned by Disqus' API
*/
func (gisqus *Gisqus) UserFollowing(ctx context.Context, userID string, values url.Values) (*UserListResponse, error) {
	if userID == "" {
		return nil, errors.New("Must provide a user id")
	}
	values.Set("user", userID)
	values.Set("api_secret", gisqus.secret)
	url := usersUrls.userFollowing + "?" + values.Encode()
	var fr UserListResponse

	err := gisqus.callAndInflate(url, &fr, ctx)
	if err != nil {
		return nil, err
	}

	for _, user := range fr.Response {
		user.JoinedAt, err = fromDisqusTime(user.DisqusTimeJoinedAt)
		if err != nil {
			return nil, err
		}
	}
	return &fr, nil
}

/*
UserForumFollowing wraps https://disqus.com/api/docs/users/listFollowingForums/ (https://disqus.com/api/3.0/users/listFollowingForums.json)
*/
func (gisqus *Gisqus) UserForumFollowing(ctx context.Context, user string, values url.Values) (*UserForumFollowingResponse, error) {

	if user == "" {
		return nil, errors.New("Must provide a user id")
	}
	values.Set("user", user)
	values.Set("api_secret", gisqus.secret)
	url := usersUrls.userFollowingForums + "?" + values.Encode()

	var uffr UserForumFollowingResponse
	err := gisqus.callAndInflate(url, &uffr, ctx)
	if err != nil {
		return nil, err
	}

	for _, forum := range uffr.Response {
		forum.CreatedAt, err = fromDisqusTimeExact(forum.DisqusTimeCreatedAt)
		if err != nil {
			return nil, err
		}
	}
	return &uffr, nil
}

//UserListResponse models the response of various user endpoints.
type UserListResponse struct {
	ResponseStubWithCursor
	Response []*User `json:"response"`
}

// ActiveForumsResponse models the response of the active forums user endpoint.
type ActiveForumsResponse struct {
	ResponseStubWithCursor
	Response []*Forum `json:"response"`
}

// InterestingUsersResponse models the response of the interesting users endpoint.
type InterestingUsersResponse struct {
	ResponseStubWithCursor
	Response *InterestingUsers `json:"response"`
}

// InterestingUsers models the objects returned by the interesting users endpoint.
type InterestingUsers struct {
	Items   []*InterestingItem `json:"items"`
	Objects map[string]*User   `json:"objects"`
}

// UserDetailsResponse models the response of the user detail endpoint
type UserDetailsResponse struct {
	ResponseStub
	Response *User `json:"response"`
}

// UserForumFollowingResponse models the response of the user forum following endpoint.
type UserForumFollowingResponse struct {
	ResponseStubWithCursor
	Response []*Forum `json:"response"`
}

// User models the user object returned by the user detail endpoint.
type User struct {
	Disable3rdPartyTrackers bool        `json:"disable3rdPartyTrackers"`
	IsPowerContributor      bool        `json:"isPowerContributor"`
	IsPrimary               bool        `json:"isPrimary"`
	ID                      string      `json:"id"`
	NumFollowers            int         `json:"numFollowers"`
	Rep                     float32     `json:"rep"`
	NumFollowing            int         `json:"numFollowing"`
	NumPosts                int         `json:"numPosts"`
	Location                string      `json:"location"`
	IsPrivate               bool        `json:"isPrivate"`
	JoinedAt                time.Time   `json:"-"`
	DisqusTimeJoinedAt      string      `json:"joinedAt"`
	Username                string      `json:"username"`
	NumLikesReceived        int         `json:"numLikesReceived"`
	ReputationLabel         string      `json:"reputationLabel"`
	About                   string      `json:"about"`
	Name                    string      `json:"name"`
	URL                     string      `json:"url"`
	NumForumsFollowing      int         `json:"numForumsFollowing"`
	ProfileURL              string      `json:"profileUrl"`
	Reputation              float32     `json:"reputation"`
	Avatar                  *UserAvatar `json:"avatar"`
	SignedURL               string      `json:"signedUrl"`
	IsAnonymous             bool        `json:"isAnonymous"`
}

// UserAvatar models the avatar field of the user object.
type UserAvatar struct {
	Small *Icon `json:"small"`
	Large *Icon `json:"large"`
	Icon
	IsCustom bool `json:"isCustom"`
}
