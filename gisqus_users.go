// Copyright Piero de Salvia.
// All Rights Reserved

package gisqus

import (
	"context"
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
	"time"
)

// UsersURLS are the URLs used by Disqus' user endpoints
type UsersURLS struct {
	DetailURL            string
	InterestingIUsersURL string
	PostListURL          string
	ActiveForumsURL      string
	FollowersURL         string
	FollowingURL         string
	FollowingForumsURL   string
	ActivityURL          string
	MostActiveForumsURL  string
}

var usersUrls = UsersURLS{
	DetailURL:            "https://disqus.com/api/3.0/users/details.json",
	InterestingIUsersURL: "https://disqus.com/api/3.0/users/interestingUsers.json",
	PostListURL:          "https://disqus.com/api/3.0/users/listPosts.json",
	ActiveForumsURL:      "https://disqus.com/api/3.0/users/listActiveForums.json",
	FollowersURL:         "https://disqus.com/api/3.0/users/listFollowers.json",
	FollowingURL:         "https://disqus.com/api/3.0/users/listFollowing.json",
	FollowingForumsURL:   "https://disqus.com/api/3.0/users/listFollowingForums.json",
	ActivityURL:          "https://disqus.com/api/3.0/users/listActivity.json",
	MostActiveForumsURL:  "https://disqus.com/api/3.0/users/listMostActiveForums.json",
}

type activityResponseRaw struct {
	ResponseStubWithCursor
	Fragments []ActivityResponseFragment `json:"response"`
}

// ActivityResponseFragment is exported because of the json parser
type ActivityResponseFragment struct {
	FragmentType string          `json:"type"`
	FragmentData json.RawMessage `json:"object"`
}

// RawPost is exported because of the json parser
type RawPost struct {
	postBase
	Parent Post `json:"parent"`
}

/*
UserActivities wraps https://disqus.com/api/docs/users/listActivity/ (https://disqus.com/api/3.0/users/listActivity.json)
*/
func (gisqus *Gisqus) UserActivities(ctx context.Context, userID string, values url.Values) (*ActivitiesListResponse, error) {
	if userID == "" {
		return nil, errors.New("Must provide a user id")
	}
	values.Set("api_secret", gisqus.secret)
	values.Set("user", userID)
	values.Set("related", "")

	url := usersUrls.ActivityURL + "?" + values.Encode()

	var arr activityResponseRaw

	err := gisqus.callAndInflate(ctx, url, &arr)
	if err != nil {
		return nil, err
	}
	alr := ActivitiesListResponse{}
	alr.ResponseStubWithCursor = arr.ResponseStubWithCursor

	for _, fragment := range arr.Fragments {
		switch fragment.FragmentType {
		case "reply":
			fallthrough
		case "post":
			var rawPost RawPost
			err := json.Unmarshal(fragment.FragmentData, &rawPost)
			if err != nil {
				return nil, err
			}
			var post Post
			postIDInt, err := strconv.Atoi(rawPost.Parent.ID)
			if err != nil {
				return nil, err
			}
			post = Post{
				rawPost.postBase,
				postIDInt,
			}
			alr.Posts = append(alr.Posts, &post)
		default:
			return nil, errors.New(fragment.FragmentType)
		}
	}

	for _, post := range alr.Posts {
		post.CreatedAt, err = fromDisqusTime(post.DisqusTimeCreatedAt)
		if err != nil {
			return nil, err
		}
		post.Author.JoinedAt, err = fromDisqusTime(post.Author.DisqusTimeJoinedAt)

		if err != nil {
			return nil, err
		}
	}

	return &alr, nil
}

/*
UserMostActiveForums wraps https://disqus.com/api/docs/users/listMostActiveForums/ (https://disqus.com/api/3.0/users/listMostActiveForums.json)
*/
func (gisqus *Gisqus) UserMostActiveForums(ctx context.Context, userID string, values url.Values) (*MostActiveForumsResponse, error) {
	if userID == "" {
		return nil, errors.New("Must provide a user id")
	}
	values.Set("api_secret", gisqus.secret)
	values.Set("user", userID)
	url := usersUrls.MostActiveForumsURL + "?" + values.Encode()

	var mafr MostActiveForumsResponse

	err := gisqus.callAndInflate(ctx, url, &mafr)
	if err != nil {
		return nil, err
	}

	for _, forum := range mafr.Response {
		forum.CreatedAt, err = fromDisqusTimeExact(forum.DisqusTimeCreatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &mafr, nil
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
	url := usersUrls.PostListURL + "?" + values.Encode()

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
UserDetails wraps https://disqus.com/api/docs/users/details/ (https://disqus.com/api/3.0/users/details.json)
*/
func (gisqus *Gisqus) UserDetails(ctx context.Context, userID string, values url.Values) (*UserDetailsResponse, error) {

	if userID == "" {
		return nil, errors.New("Must provide a user id")
	}
	values.Set("api_secret", gisqus.secret)
	url := usersUrls.DetailURL + "?" + values.Encode()
	var udr UserDetailsResponse
	err := gisqus.callAndInflate(ctx, url, &udr)
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
	url := usersUrls.InterestingIUsersURL + "?" + values.Encode()
	var iur InterestingUsersResponse

	err := gisqus.callAndInflate(ctx, url, &iur)
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
func (gisqus *Gisqus) UserActiveForums(ctx context.Context, userID string, values url.Values) (*ActiveForumsResponse, error) {

	if userID == "" {
		return nil, errors.New("Must provide a user id")
	}
	values.Set("user", userID)
	values.Set("api_secret", gisqus.secret)
	url := usersUrls.ActiveForumsURL + "?" + values.Encode()

	var afr ActiveForumsResponse
	err := gisqus.callAndInflate(ctx, url, &afr)
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
	url := usersUrls.FollowersURL + "?" + values.Encode()
	var fr UserListResponse

	err := gisqus.callAndInflate(ctx, url, &fr)
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
	url := usersUrls.FollowingURL + "?" + values.Encode()
	var fr UserListResponse

	err := gisqus.callAndInflate(ctx, url, &fr)
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
func (gisqus *Gisqus) UserForumFollowing(ctx context.Context, userID string, values url.Values) (*UserForumFollowingResponse, error) {

	if userID == "" {
		return nil, errors.New("Must provide a user id")
	}
	values.Set("user", userID)
	values.Set("api_secret", gisqus.secret)
	url := usersUrls.FollowingForumsURL + "?" + values.Encode()

	var uffr UserForumFollowingResponse
	err := gisqus.callAndInflate(ctx, url, &uffr)
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

// ActivitiesListResponse models the response of the user activity list endpoint
type ActivitiesListResponse struct {
	ResponseStubWithCursor
	Posts []*Post `json:"response"`
}

// MostActiveForumsResponse models the responser of the user most active forums endpoint
type MostActiveForumsResponse struct {
	ResponseStub
	Response []*Forum `json:"response"`
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
