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

	return date.Format(disqusDateFormat)
}

/*
ToDisqusTimeExact returns a string in the format of the ForumDetails CreatedAt field.
*/
func ToDisqusTimeExact(date time.Time) string {

	return date.Format(disqusDateFormatExact)
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

// Post constants are used by Disqus in API calls in the "filters" parameter
const (
	PostIsAnonymous = 1 + iota
	PostHasLink
	PostHasLowRepAuthor
	PostHasBadWord
	PostIsFlagged
	PostNoIssue
)

// Post constants are used by Disqus in API calls in the "include" parameter
const (
	PostIsUnapproved      = "unapproved"
	PostIsApproved        = "approved"
	PostIsSpam            = "spam"
	PostIsDeleted         = "deleted"
	PostIncludedIsFlagged = "flagged"
	PostIsHighlighted     = "highlighted"
)

// Intervals are used by Disqus in API calls in the "since" parameter
const (
	Interval1h  = "1h"
	Interval6h  = "6h"
	Interval12h = "12h"
	Interval1d  = "1d"
	Interval3d  = "3d"
	Interval7d  = "7d"
	Interval30d = "30d"
	Interval90d = "90d"
)

// Order constant are values used by Disqus in API calls in the "order" parameter
const (
	OrderAsc  = "asc"
	OrderDesc = "desc"
)

// Sort constants are values used by Disqus in API calls in the sortType parameter
const (
	SortTypeDate     = "date"
	SortTypePriority = "priority"
)
