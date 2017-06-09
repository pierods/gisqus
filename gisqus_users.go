package gisqus

import (
	"context"
	"errors"
	"net/url"
	"time"
)

type UsersURLs struct {
	user_detail_url            string
	user_interesting_users_url string
	user_post_list_url         string
	user_active_forums         string
	user_followers             string
	user_following             string
	user_following_forums      string
}

// following forums
// curl "https://disqus.com/api/3.0/users/listFollowingForums.json?user=195792235&api_secret=KpVAypnhCxG27eRLRbXad0i1xfbyUsHPE7E8on5wbFJkbQcIzjB0pkJ4kMOfTRmx" |jq

var usersUrls = UsersURLs{
	user_detail_url:            "https://disqus.com/api/3.0/users/details.json",
	user_interesting_users_url: "https://disqus.com/api/3.0/users/interestingUsers.json",
	user_post_list_url:         "https://disqus.com/api/3.0/users/listPosts.json",
	user_active_forums:         "https://disqus.com/api/3.0/users/listActiveForums.json",
	user_followers:             "https://disqus.com/api/3.0/users/listFollowers.json",
	user_following:             "https://disqus.com/api/3.0/users/listFollowing.json",
	user_following_forums:      "https://disqus.com/api/3.0/users/listFollowingForums.json",
}

/*
https://disqus.com/api/docs/users/listPosts/
*/
func (gisqus *Gisqus) UserPosts(userId string, values url.Values, ctx context.Context) (*PostListResponse, error) {
	if userId == "" {
		return nil, errors.New("Must provide a user id")
	}
	values.Set("api_secret", gisqus.secret)
	values.Set("user", userId)
	url := usersUrls.user_post_list_url + "?" + values.Encode()

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

func (gisqus *Gisqus) UserDetails(userId string, values url.Values, ctx context.Context) (*UserDetailsResponse, error) {

	if userId == "" {
		return nil, errors.New("Must provide a user id")
	}
	values.Set("api_secret", gisqus.secret)
	url := usersUrls.user_detail_url + "?" + values.Encode()
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

func (gisqus *Gisqus) UserInteresting(values url.Values, ctx context.Context) (*InterestingUsersResponse, error) {

	values.Set("api_secret", gisqus.secret)
	url := usersUrls.user_interesting_users_url + "?" + values.Encode()
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

func (gisqus *Gisqus) UserActiveForums(user string, values url.Values, ctx context.Context) (*ActiveForumsResponse, error) {

	if user == "" {
		return nil, errors.New("Must provide a user id")
	}
	values.Set("user", user)
	values.Set("api_secret", gisqus.secret)
	url := usersUrls.user_active_forums + "?" + values.Encode()

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
Numlikes, NumPosts, NumFollowers are not returned
*/
func (gisqus *Gisqus) UserFollowers(userId string, values url.Values, ctx context.Context) (*UserListResponse, error) {
	if userId == "" {
		return nil, errors.New("Must provide a user id")
	}
	values.Set("user", userId)
	values.Set("api_secret", gisqus.secret)
	url := usersUrls.user_followers + "?" + values.Encode()
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
Numlikes, NumPosts, NumFollowers are not returned
*/
func (gisqus *Gisqus) UserFollowing(userId string, values url.Values, ctx context.Context) (*UserListResponse, error) {
	if userId == "" {
		return nil, errors.New("Must provide a user id")
	}
	values.Set("user", userId)
	values.Set("api_secret", gisqus.secret)
	url := usersUrls.user_following + "?" + values.Encode()
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

func (gisqus *Gisqus) UserForumFollowing(user string, values url.Values, ctx context.Context) (*UserForumFollowingResponse, error) {

	if user == "" {
		return nil, errors.New("Must provide a user id")
	}
	values.Set("user", user)
	values.Set("api_secret", gisqus.secret)
	url := usersUrls.user_following_forums + "?" + values.Encode()

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

type UserListResponse struct {
	ResponseStubWithCursor
	Response []*User `json:"response"`
}
type ActiveForumsResponse struct {
	ResponseStubWithCursor
	Response []*Forum `json:"response"`
}

type InterestingUsersResponse struct {
	ResponseStubWithCursor
	Response *InterestingUsers `json:"response"`
}

type InterestingUsers struct {
	Items   []*InterestingItem `json:"items"`
	Objects map[string]*User   `json:"objects"`
}

type UserDetailsResponse struct {
	ResponseStub
	Response *User `json:"response"`
}

type UserForumFollowingResponse struct {
	ResponseStubWithCursor
	Response []*Forum `json:"response"`
}

type User struct {
	Disable3rdPartyTrackers bool        `json:"disable3rdPartyTrackers"`
	IsPowerContributor      bool        `json:"isPowerContributor"`
	IsPrimary               bool        `json:"isPrimary"`
	Id                      string      `json:"id"`
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
	Url                     string      `json:"url"`
	NumForumsFollowing      int         `json:"numForumsFollowing"`
	ProfileUrl              string      `json:"profileUrl"`
	Reputation              float32     `json:"reputation"`
	Avatar                  *UserAvatar `json:"avatar"`
	SignedUrl               string      `json:"signedUrl"`
	IsAnonymous             bool        `json:"isAnonymous"`
}

type UserAvatar struct {
	Small *Icon `json:"small"`
	Large *Icon `json:"large"`
	Icon
	IsCustom bool `json:"isCustom"`
}
