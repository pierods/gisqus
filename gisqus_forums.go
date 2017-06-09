package gisqus

import (
	"context"
	"errors"
	"net/url"
	"time"
)

type ForumsURLs struct {
	forum_interesting_forums_url string
	forum_details_url            string
	forum_categories_url         string
	forum_list_users             string
	forum_list_threads           string
	forum_most_liked_users       string
}

var forumsUrls = ForumsURLs{
	forum_interesting_forums_url: "https://disqus.com/api/3.0/forums/interestingForums",
	forum_details_url:            "https://disqus.com/api/3.0/forums/details.json",
	forum_categories_url:         "https://disqus.com/api/3.0/forums/listCategories.json",
	forum_list_users:             "https://disqus.com/api/3.0/forums/listUsers.json",
	forum_list_threads:           "https://disqus.com/api/3.0/forums/listThreads.json",
	forum_most_liked_users:       "https://disqus.com/api/3.0/forums/listMostLikedUsers.json",
}

func (gisqus *Gisqus) ForumUsers(forum string, values url.Values, ctx context.Context) (*ForumUserListResponse, error) {

	if forum == "" {
		return nil, errors.New("Must provide a forum id")
	}
	values.Set("api_secret", gisqus.secret)
	values.Set("forum", forum)
	url := forumsUrls.forum_list_users + "?" + values.Encode()

	var fulr ForumUserListResponse
	err := gisqus.callAndInflate(url, &fulr, ctx)
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

func (gisqus *Gisqus) ForumInteresting(values url.Values, ctx context.Context) (*InterestingForumsResponse, error) {

	values.Set("api_secret", gisqus.secret)
	url := forumsUrls.forum_interesting_forums_url + "?" + values.Encode()

	var ifr InterestingForumsResponse
	err := gisqus.callAndInflate(url, &ifr, ctx)
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

func (gisqus *Gisqus) ForumDetails(forum string, values url.Values, ctx context.Context) (*ForumDetailsResponse, error) {

	if forum == "" {
		return nil, errors.New("Must provide a forum id")
	}
	values.Set("api_secret", gisqus.secret)
	values.Set("forum", forum)
	url := forumsUrls.forum_details_url + "?" + values.Encode()

	var fdr ForumDetailsResponse

	err := gisqus.callAndInflate(url, &fdr, ctx)
	if err != nil {
		return nil, err
	}

	fdr.Response.CreatedAt, err = fromDisqusTimeExact(fdr.Response.DisqusTimeCreatedAt)
	if err != nil {
		return nil, err
	}
	return &fdr, err
}

func (gisqus *Gisqus) ForumCategories(forum string, values url.Values, ctx context.Context) (*CategoriesListResponse, error) {

	if forum == "" {
		return nil, errors.New("Must provide a forum id")
	}
	values.Set("forum", forum)
	values.Set("api_secret", gisqus.secret)
	url := forumsUrls.forum_categories_url + "?" + values.Encode()

	var clr CategoriesListResponse

	err := gisqus.callAndInflate(url, &clr, ctx)
	if err != nil {
		return nil, err
	}
	return &clr, nil
}

func (gisqus *Gisqus) ForumThreads(forum string, values url.Values, ctx context.Context) (*ThreadListResponse, error) {

	if forum == "" {
		return nil, errors.New("Must provide a forum id")
	}
	values.Set("forum", forum)
	values.Set("api_secret", gisqus.secret)
	url := forumsUrls.forum_list_threads + "?" + values.Encode()

	var tlr ThreadListResponse

	err := gisqus.callAndInflate(url, &tlr, ctx)

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
https://disqus.com/api/docs/forums/listMostLikedUsers/
Will not return the # of likes
*/
func (gisqus *Gisqus) ForumMostLikedUsers(forum string, values url.Values, ctx context.Context) (*MostLikedUsersResponse, error) {

	if forum == "" {
		return nil, errors.New("Must provide a forum id")
	}
	values.Set("api_secret", gisqus.secret)
	values.Set("forum", forum)
	url := forumsUrls.forum_most_liked_users + "?" + values.Encode()

	var mlur MostLikedUsersResponse
	err := gisqus.callAndInflate(url, &mlur, ctx)
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

type MostLikedUsersResponse struct {
	ResponseStubWithCursor
	Response []*User `json:"response"`
}

type ForumUserListResponse struct {
	ResponseStubWithCursor
	Response []*User `json:"response"`
}

type CategoriesListResponse struct {
	ResponseStubWithCursor
	Response []*Category `json:"response"`
}

type Category struct {
	IsDefault bool   `json:"isDefault"`
	Title     string `json:"title"`
	Order     int    `json:"order"`
	Forum     string `json:"forum"`
	Id        string `json:"id"`
}

type InterestingForumsResponse struct {
	ResponseStubWithCursor
	Response *InterestingForums `json:"response"`
}

type InterestingForums struct {
	Items   []*InterestingItem `json:"items"`
	Objects map[string]*Forum  `json:"objects"`
}

type ForumDetailsResponse struct {
	ResponseStub
	Response *Forum `json:"response"`
}

type Forum struct {
	RawGuidelines         string             `json:"raw_guidelines"`
	TwitterName           string             `json:"twitterName"`
	Guidelines            string             `json:"guidelines"`
	Favicon               *Icon              `json:"favicon"`
	DisableDisqusBranding bool               `json:"disableDisqusBranding"`
	Id                    string             `json:"id"`
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
	OrganizationId        int                `json:"organizationId"`
	DaysThreadAlive       int                `json:"daysThreadAlive"`
	Avatar                *ForumAvatar       `json:"avatar"`
	SignedUrl             string             `json:"signedUrl"`
}

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
	OrganicDiscoveryEnabled          bool `json:"organicProductDiscovery"`
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

type DisqusPermissions struct {
}

type ForumAvatar struct {
	Small *Icon `json:"small"`
	Large *Icon `json:"large"`
}

type ForumChannel struct {
	BannerColor     string          `json:"bannerColor"`
	Slug            string          `json:"slug"`
	DateAdded       time.Time       `json:"-"`
	DisqusDateAdded string          `json:"dateAdded"`
	Name            string          `json:"name"`
	Banner          string          `json:"banner"`
	BannerColorHex  string          `json:"bannerColorHex"`
	Id              string          `json:"id"`
	Hidden          bool            `json:"hidden"`
	IsAggregation   bool            `json:"isAggregation"`
	Avatar          string          `json:"avatar"`
	EnableCuration  bool            `json:"enableCuration"`
	IsCategory      bool            `json:"isCategory"`
	AdminOnly       bool            `json:"adminOnly"`
	Options         *ChannelOptions `json:"options"`
	OwnerId         string          `json:"ownerId"`
}

type ChannelOptions struct {
	AboutUrlPath      string             `json:"aboutUrlPath"`
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

type ChannelCoverImage struct {
	Cache string
}

type ChannelTitleLogo struct {
	Small string
	Cache string
}
