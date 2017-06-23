// Copyright Piero de Salvia.
// All Rights Reserved

// Package gisqus is a thin wrapper over the Disqus API
package gisqus

import (
	"strings"
	"time"
)

// Gisqus is lib's entry point
type Gisqus struct {
	secret string
	limits DisqusRateLimit
}

/*
NewGisqus returns a new instance of Gisqus. secret is Disqus' API key
*/
func NewGisqus(secret string) Gisqus {
	return Gisqus{
		secret: secret,
	}
}

/*
ReadThreadsURLs returns all the URLs used by Gisqus to call thread endpoints
*/
func (g *Gisqus) ReadThreadsURLs() ThreadsURLS {
	return threadsUrls
}

/*
SetThreadsURLs changes the URLs used by Gisqus to call thread endpoints
*/
func (g *Gisqus) SetThreadsURLs(tu ThreadsURLS) {
	threadsUrls = tu
}

/*
ReadUsersURLs returns all the URLs used by Gisqus to call user endpoints
*/
func (g *Gisqus) ReadUsersURLs() UsersURLS {
	return usersUrls
}

/*
SetUsersURLs changes the URLS used by Gisqus to call user endpoints.
*/
func (g *Gisqus) SetUsersURLs(uu UsersURLS) {
	usersUrls = uu
}

/*
ReadForumsURLs returns all the URLs used by Gisqus to call forum endpoints.
*/
func (g *Gisqus) ReadForumsURLs() ForumsURLS {
	return forumsUrls
}

/*
SetForumsURLs changes the URLs used by Gisqus to call forum endpoints
*/
func (g *Gisqus) SetForumsURLs(fu ForumsURLS) {
	forumsUrls = fu
}

/*
ReadPostURLs returns all the URLs used by Gisqus to call post endpoints
*/
func (g *Gisqus) ReadPostURLs() PostsURLS {
	return postsUrls
}

/*
SetPostsURLs changes the URLs used by Gisqus to call post endpoints
*/
func (g *Gisqus) SetPostsURLs(pu PostsURLS) {
	postsUrls = pu
}

/*
Limits return the current rate limits for the user account
*/
func (g *Gisqus) Limits() DisqusRateLimit {
	return g.limits
}

/*
ToDisqusTime returns a string that can be used in Disqus call for timedate parameters
*/
func ToDisqusTime(date time.Time) string {

	return date.Format(DisqusDateFormat)
}

/*
ToDisqusTimeExact returns a string in the format of the ForumDetails CreatedAt field.
*/
func ToDisqusTimeExact(date time.Time) string {

	return date.Format(DisqusDateFormatExact)
}

// ExtractForumID extracts the forum ID from the keys of the map returned by the interesting forum call
func ExtractForumID(forumString string) string {

	parts := strings.Split(forumString, "=")
	return parts[len(parts)-1]
}

// DisqusCursor models the cursor used by Disqus for pagination
type DisqusCursor struct {
	Prev    string `json:"prev"`
	HasNext bool   `json:"hasNext"`
	Next    string `json:"next"`
	HasPrev bool   `json:"hasPrev"`
	Total   int    `json:"total"`
	ID      string `json:"id"`
	More    bool   `json:"more"`
}

// DisqusRateLimit models the rate limits for an user account
type DisqusRateLimit struct {
	RatelimitRemaining int
	RatelimitLimit     int
	RatelimitReset     time.Time
}

// ResponseStub is the standard Disqus response stub
type ResponseStub struct {
	Code int `json:"code"`
}

// ResponseStubWithCursor is the standard response stub for call that support pagination
type ResponseStubWithCursor struct {
	ResponseStub
	Cursor *DisqusCursor `json:"cursor"`
}

// Icon models an icon object returned by many Disqus endpoints
type Icon struct {
	Permalink string `json:"permalink"`
	Cache     string `json:"cache"`
}

// InterestingItem models the slice returned by Interesting* Disqus endpoints.
type InterestingItem struct {
	Reason string `json:"reason"`
	ID     string `json:"id"`
}

// Filter represents the possible values for filter in calls to Disqus'API
type Filter int

// Post constants are used by Disqus in API calls in the "filters" parameter
const (
	PostIsAnonymous Filter = 1 + iota
	PostHasLink
	PostHasLowRepAuthor
	PostHasBadWord
	PostIsFlagged
	PostNoIssue
)

// Include represents the possible values for includes in calls to Disqus'API
type Include string

// Post constants are used by Disqus in API calls in the "include" parameter
const (
	PostIsUnapproved      Include = "unapproved"
	PostIsApproved        Include = "approved"
	PostIsSpam            Include = "spam"
	PostIsDeleted         Include = "deleted"
	PostIncludedIsFlagged Include = "flagged"
	PostIsHighlighted     Include = "highlighted"
)

// Interval represents the possible values for intervals in calls to Disqus'API
type Interval string

// Intervals are used by Disqus in API calls in the "since" parameter
const (
	Interval1h  Interval = "1h"
	Interval6h  Interval = "6h"
	Interval12h Interval = "12h"
	Interval1d  Interval = "1d"
	Interval3d  Interval = "3d"
	Interval7d  Interval = "7d"
	Interval30d Interval = "30d"
	Interval90d Interval = "90d"
)

// Order represents the possible values for order in calls to Disqus'API
type Order string

// Order constant are values used by Disqus in API calls in the "order" parameter
const (
	OrderAsc  Order = "asc"
	OrderDesc Order = "desc"
)

// Sort represents the possible values for sort in calls to Disqus'API
type Sort string

// Sort constants are values used by Disqus in API calls in the sortType parameter
const (
	SortTypeDate     Sort = "date"
	SortTypePriority Sort = "priority"
)
