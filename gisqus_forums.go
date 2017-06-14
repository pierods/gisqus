package gisqus

import (
	"context"
	"errors"
	"net/url"
	"time"
)

// ForumsURLS contains the URLs of the API calls for forums on Disqus
type ForumsURLS struct {
	InterestingForumsURL string
	DetailsURL           string
	CategoriesURL        string
	ListUsersURL         string
	ListThreadsURL       string
	MostLikedUsersURL    string
	ListFollowersURL     string
	MostActiveUsersURL   string
}

var forumsUrls = ForumsURLS{
	InterestingForumsURL: "https://disqus.com/api/3.0/forums/interestingForums",
	DetailsURL:           "https://disqus.com/api/3.0/forums/details.json",
	CategoriesURL:        "https://disqus.com/api/3.0/forums/listCategories.json",
	ListUsersURL:         "https://disqus.com/api/3.0/forums/listUsers.json",
	ListThreadsURL:       "https://disqus.com/api/3.0/forums/listThreads.json",
	MostLikedUsersURL:    "https://disqus.com/api/3.0/forums/listMostLikedUsers.json",
	ListFollowersURL:     "https://disqus.com/api/3.0/forums/listFollowers.json",
	MostActiveUsersURL:   "https://disqus.com/api/3.0/forums/listMostActiveUsers.json",
}

/*
ForumMostActiveUsers wraps https://disqus.com/api/docs/forums/listMostActiveUsers/ (https://disqus.com/api/3.0/forums/listMostActiveUsers.json)
*/
func (gisqus *Gisqus) ForumMostActiveUsers(ctx context.Context, forum string, values url.Values) (*ForumUserListResponse, error) {

	if forum == "" {
		return nil, errors.New("Must provide a forum id")
	}
	values.Set("api_secret", gisqus.secret)
	values.Set("forum", forum)
	url := forumsUrls.MostActiveUsersURL + "?" + values.Encode()

	var fulr ForumUserListResponse
	err := gisqus.callAndInflate(ctx, url, &fulr)
	if err != nil {
		return nil, err
	}

	for _, user := range fulr.Response {
		user.JoinedAt, err = fromDisqusTime(user.DisqusTimeJoinedAt)
		if err != nil {
			return nil, err
		}
	}
	return &fulr, nil
}

/*
ForumFollowers wraps https://disqus.com/api/docs/forums/listFollowers/ (https://disqus.com/api/3.0/forums/listFollowers.json)
*/
func (gisqus *Gisqus) ForumFollowers(ctx context.Context, forum string, values url.Values) (*ForumUserListResponse, error) {

	if forum == "" {
		return nil, errors.New("Must provide a forum id")
	}
	values.Set("api_secret", gisqus.secret)
	values.Set("forum", forum)
	url := forumsUrls.ListFollowersURL + "?" + values.Encode()

	var fulr ForumUserListResponse
	err := gisqus.callAndInflate(ctx, url, &fulr)
	if err != nil {
		return nil, err
	}

	for _, user := range fulr.Response {
		user.JoinedAt, err = fromDisqusTime(user.DisqusTimeJoinedAt)
		if err != nil {
			return nil, err
		}
	}
	return &fulr, nil
}

/*
ForumUsers wraps https://disqus.com/api/3.0/forums/listUsers.json (https://disqus.com/api/docs/forums/listUsers/)
*/
func (gisqus *Gisqus) ForumUsers(ctx context.Context, forum string, values url.Values) (*ForumUserListResponse, error) {

	if forum == "" {
		return nil, errors.New("Must provide a forum id")
	}
	values.Set("api_secret", gisqus.secret)
	values.Set("forum", forum)
	url := forumsUrls.ListUsersURL + "?" + values.Encode()

	var fulr ForumUserListResponse
	err := gisqus.callAndInflate(ctx, url, &fulr)
	if err != nil {
		return nil, err
	}

	for _, user := range fulr.Response {
		user.JoinedAt, err = fromDisqusTime(user.DisqusTimeJoinedAt)
		if err != nil {
			return nil, err
		}
	}
	return &fulr, nil
}

/*
ForumInteresting wraps https://disqus.com/api/docs/forums/interestingForums/ (https://disqus.com/api/3.0/forums/interestingForums.json)
*/
func (gisqus *Gisqus) ForumInteresting(ctx context.Context, values url.Values) (*InterestingForumsResponse, error) {

	values.Set("api_secret", gisqus.secret)
	url := forumsUrls.InterestingForumsURL + "?" + values.Encode()

	var ifr InterestingForumsResponse
	err := gisqus.callAndInflate(ctx, url, &ifr)
	if err != nil {
		return nil, err
	}

	for _, forum := range ifr.Response.Objects {
		forum.CreatedAt, err = fromDisqusTimeExact(forum.DisqusTimeCreatedAt)
		if err != nil {
			return nil, err
		}
	}
	return &ifr, nil
}

/*
ForumDetails wraps https://disqus.com/api/docs/forums/details/ (https://disqus.com/api/3.0/forums/details.json)
It does not support the "related" url parameter (other funcs can be used for drilldown)
*/
func (gisqus *Gisqus) ForumDetails(ctx context.Context, forum string, values url.Values) (*ForumDetailsResponse, error) {

	if forum == "" {
		return nil, errors.New("Must provide a forum id")
	}
	values.Set("api_secret", gisqus.secret)
	values.Set("forum", forum)
	url := forumsUrls.DetailsURL + "?" + values.Encode()

	var fdr ForumDetailsResponse

	err := gisqus.callAndInflate(ctx, url, &fdr)
	if err != nil {
		return nil, err
	}

	fdr.Response.CreatedAt, err = fromDisqusTimeExact(fdr.Response.DisqusTimeCreatedAt)
	if err != nil {
		return nil, err
	}
	return &fdr, err
}

/*
ForumCategories wraps https://disqus.com/api/docs/forums/listCategories/ (https://disqus.com/api/3.0/forums/listCategories.json)
*/
func (gisqus *Gisqus) ForumCategories(ctx context.Context, forum string, values url.Values) (*CategoriesListResponse, error) {

	if forum == "" {
		return nil, errors.New("Must provide a forum id")
	}
	values.Set("forum", forum)
	values.Set("api_secret", gisqus.secret)
	url := forumsUrls.CategoriesURL + "?" + values.Encode()

	var clr CategoriesListResponse

	err := gisqus.callAndInflate(ctx, url, &clr)
	if err != nil {
		return nil, err
	}
	return &clr, nil
}

/*
ForumThreads wraps https://disqus.com/api/docs/forums/listThreads/ (https://disqus.com/api/3.0/forums/listThreads.json)
*/
func (gisqus *Gisqus) ForumThreads(ctx context.Context, forum string, values url.Values) (*ThreadListResponse, error) {

	if forum == "" {
		return nil, errors.New("Must provide a forum id")
	}
	values.Set("forum", forum)
	values.Set("api_secret", gisqus.secret)
	url := forumsUrls.ListThreadsURL + "?" + values.Encode()

	var tlr ThreadListResponse

	err := gisqus.callAndInflate(ctx, url, &tlr)

	for _, thread := range tlr.Response {
		thread.CreatedAt, err = fromDisqusTime(thread.DisqusTimeCreatedAt)
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}
	return &tlr, nil
}

/*
ForumMostLikedUsers wraps https://disqus.com/api/docs/forums/listMostLikedUsers/ (https://disqus.com/api/3.0/forums/listMostLikedUsers.json)
Disque does not return the # of likes with this call.
*/
func (gisqus *Gisqus) ForumMostLikedUsers(ctx context.Context, forum string, values url.Values) (*MostLikedUsersResponse, error) {

	if forum == "" {
		return nil, errors.New("Must provide a forum id")
	}
	values.Set("api_secret", gisqus.secret)
	values.Set("forum", forum)
	url := forumsUrls.MostLikedUsersURL + "?" + values.Encode()

	var mlur MostLikedUsersResponse
	err := gisqus.callAndInflate(ctx, url, &mlur)
	if err != nil {
		return nil, err
	}

	for _, user := range mlur.Response {
		user.JoinedAt, err = fromDisqusTime(user.DisqusTimeJoinedAt)
		if err != nil {
			return nil, err
		}
	}
	return &mlur, nil
}

// MostLikedUsersResponse models the response of the most liked users in forum endpoint
type MostLikedUsersResponse struct {
	ResponseStubWithCursor
	Response []*User `json:"response"`
}

// ForumUserListResponse models the response of the user list in forum endpoint
type ForumUserListResponse struct {
	ResponseStubWithCursor
	Response []*User `json:"response"`
}

// CategoriesListResponse models the category list in forum endpoint
type CategoriesListResponse struct {
	ResponseStubWithCursor
	Response []*Category `json:"response"`
}

// Category models the category field in forums
type Category struct {
	IsDefault bool   `json:"isDefault"`
	Title     string `json:"title"`
	Order     int    `json:"order"`
	Forum     string `json:"forum"`
	ID        string `json:"id"`
}

// InterestingForumsResponse models the response to a call to Disqus' Interesting Forums (https://disqus.com/api/docs/forums/interestingForums/)
type InterestingForumsResponse struct {
	ResponseStubWithCursor
	Response *InterestingForums `json:"response"`
}

// InterestingForums models the actual data contained in a call to Disqus' Interesting Forums (https://disqus.com/api/docs/forums/interestingForums/)
type InterestingForums struct {
	Items   []*InterestingItem `json:"items"`
	Objects map[string]*Forum  `json:"objects"`
}

// ForumDetailsResponse modeles the response to a call to Disqus' Forum details endpoint (https://disqus.com/api/docs/forums/details/)
type ForumDetailsResponse struct {
	ResponseStub
	Response *Forum `json:"response"`
}

// Forum models the fields of a forum, as returned by Disqus' API
type Forum struct {
	RawGuidelines         string             `json:"raw_guidelines"`
	TwitterName           string             `json:"twitterName"`
	Guidelines            string             `json:"guidelines"`
	Favicon               *Icon              `json:"favicon"`
	DisableDisqusBranding bool               `json:"disableDisqusBranding"`
	ID                    string             `json:"id"`
	CreatedAt             time.Time          `json:"-"`
	DisqusTimeCreatedAt   string             `json:"createdAt"`
	Category              string             `json:"category"`
	Founder               string             `json:"founder"`
	DaysAlive             int                `json:"daysAlive"`
	InstallCompleted      bool               `json:"installCompleted"`
	Pk                    string             `json:"pk"`
	Channel               *ForumChannel      `json:"channel"`
	Description           string             `json:"description"`
	RawDescription        string             `json:"raw_description"`
	AdsReviewStatus       int                `json:"adsReviewStatus"`
	Permissions           *DisqusPermissions `json:"permissions"`
	Name                  string             `json:"name"`
	Language              string             `json:"language"`
	Settings              *ForumSettings     `json:"settings"`
	OrganizationID        int                `json:"organizationId"`
	DaysThreadAlive       int                `json:"daysThreadAlive"`
	Avatar                *ForumAvatar       `json:"avatar"`
	SignedURL             string             `json:"signedUrl"`
}

// ForumSettings models the fields of the forum settings field in a Forum
type ForumSettings struct {
	SupportLevel                     int  `json:"supportLevel"`
	AdsDRNativeEnabled               bool `json:"adsDRNativeEnabled"`
	Disable3rdPartyTrackers          bool `json:"disable3rdPartyTrackers"`
	AdsVideoEnabled                  bool `json:"adsVideoEnabled"`
	AdsProductVideoEnabled           bool `json:"adsProductVideoEnabled"`
	AdsPositionTopEnabled            bool `json:"adsPositionTopEnabled"`
	AudienceSyncEnabled              bool `json:"audienceSyncEnabled"`
	UnapproveLinks                   bool `json:"unapproveLinks"`
	AdsEnabled                       bool `json:"adsEnabled"`
	AdsProductLinksThumbnailsEnabled bool `json:"adsProductLinksThumbnailsEnabled"`
	AdsProductStoriesEnabled         bool `json:"adsProductStoriesEnabled"`
	OrganicDiscoveryEnabled          bool `json:"organicDiscoveryEnabled"`
	AdsProductDisplayEnabled         bool `json:"adsProductDisplayEnabled"`
	DiscoveryLocked                  bool `json:"discoveryLocked"`
	HasCustomAvatar                  bool `json:"hasCustomAvatar"`
	LinkAffiliationEnabled           bool `json:"linkAffiliationEnabled"`
	AllowAnonPost                    bool `json:"allowAnonPost"`
	AllowMedia                       bool `json:"allowMedia"`
	AdultContent                     bool `json:"adultContent"`
	AllowAnonVotes                   bool `json:"allowAnonVotes"`
	MustVerify                       bool `json:"mustVerify"`
	MustVerifyEmail                  bool `json:"mustVerifyEmail"`
	SsoRequired                      bool `json:"ssoRequired"`
	MediaembedEnabled                bool `json:"mediaembedEnabled"`
	AdsPositionBottomEnabled         bool `json:"adsPositionBottomEnabled"`
	AdsProductLinksEnabled           bool `json:"adsProductLinksEnabled"`
	ValidateAllPosts                 bool `json:"validateAllPosts"`
	AdsSettingsLocked                bool `json:"adsSettingsLocked"`
	IsVIP                            bool `json:"isVIP"`
	AdsPositionInthreadEnabled       bool `json:"AdsPositionInthreadEnabled"`
}

// DisqusPermissions models the fields of the forum permissions field in a Forum
type DisqusPermissions struct {
}

// ForumAvatar models the fields of the forum avatar field in a forum
type ForumAvatar struct {
	Small *Icon `json:"small"`
	Large *Icon `json:"large"`
}

// ForumChannel models the fields of the forum channel field in a forum
type ForumChannel struct {
	BannerColor     string          `json:"bannerColor"`
	Slug            string          `json:"slug"`
	DateAdded       time.Time       `json:"-"`
	DisqusDateAdded string          `json:"dateAdded"`
	Name            string          `json:"name"`
	Banner          string          `json:"banner"`
	BannerColorHex  string          `json:"bannerColorHex"`
	ID              string          `json:"id"`
	Hidden          bool            `json:"hidden"`
	IsAggregation   bool            `json:"isAggregation"`
	Avatar          string          `json:"avatar"`
	EnableCuration  bool            `json:"enableCuration"`
	IsCategory      bool            `json:"isCategory"`
	AdminOnly       bool            `json:"adminOnly"`
	Options         *ChannelOptions `json:"options"`
	OwnerID         string          `json:"ownerId"`
}

// ChannelOptions models the fields of the options field in a channel
type ChannelOptions struct {
	AboutURLPath      string             `json:"aboutUrlPath"`
	Description       string             `json:"description"`
	Coverimage        *ChannelCoverImage `json:"coverImage"`
	BannerTimestamp   time.Time          `json:"bannerTimestamp"`
	ModEmail          string             `json:"modEmail"`
	AlertBackground   string             `json:"alertBackground"`
	Favicon           string             `json:"favicon"`
	Title             string             `json:"title"`
	IsCurationChannel bool               `json:"isCurationChannel"`
	TitleLogo         *ChannelTitleLogo  `json:"titleLogo"`
	BannerColor       string             `json:"bannerColor"`
}

// ChannelCoverImage models the fields of the coverImage field in a channel option
type ChannelCoverImage struct {
	Cache string
}

// ChannelTitleLogo models the fields of the titleLogo field in a channel option
type ChannelTitleLogo struct {
	Small string
	Cache string
}
