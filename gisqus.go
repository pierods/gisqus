// A thin wrapper over the Disqus API
package gisqus

import (
	"strings"
	"time"
)

type Gisqus struct {
	secret string
	limits DisqusRateLimit
}

func NewGisqus(secret string) Gisqus {
	return Gisqus{
		secret: secret,
	}
}

func (g *Gisqus) ReadThreadsURLs() ThreadsURLs {
	return threadsUrls
}

func (g *Gisqus) SetThreadsURLs(tu ThreadsURLs) {
	threadsUrls = tu
}

func (g *Gisqus) ReadUsersURLs() UsersURLs {
	return usersUrls
}

func (g *Gisqus) SetUsersURLs(uu UsersURLs) {
	usersUrls = uu
}

func (g *Gisqus) ReadForumsURLs() ForumsURLs {
	return forumsUrls
}

func (g *Gisqus) SetForumsURLs(fu ForumsURLs) {
	forumsUrls = fu
}

func (g *Gisqus) ReadPostURLs() PostsURLs {
	return postsUrls
}

func (g *Gisqus) SetPostsURLs(pu PostsURLs) {
	postsUrls = pu
}

func (g *Gisqus) Limits() DisqusRateLimit {
	return g.limits
}

func ToDisqusTime(date time.Time) string {

	return date.Format(disqusDateFormat)
}

func ToDisqusTimeExact(date time.Time) string {

	return date.Format(disqusDateFormatExact)
}

func ExtractForumId(forumString string) string {

	parts := strings.Split(forumString, "=")
	return parts[len(parts)-1]
}

type DisqusCursor struct {
	Prev    string `json:"prev"`
	HasNext bool   `json:"hasNext"`
	Next    string `json:"next"`
	HasPrev bool   `json:"hasPrev"`
	Total   int    `json:"total"`
	Id      string `json:"id"`
	More    bool   `json:"more"`
}

type DisqusRateLimit struct {
	RatelimitRemaining int
	RatelimitLimit     int
	RatelimitReset     time.Time
}

type ResponseStub struct {
	Code int `json:"code"`
}

type ResponseStubWithCursor struct {
	ResponseStub
	Cursor *DisqusCursor `json:"cursor"`
}

type Icon struct {
	Permalink string `json:"permalink"`
	Cache     string `json:"cache"`
}

type InterestingItem struct {
	Reason string `json:"reason"`
	Id     string `json:"id"`
}

const (
	PostIsAnonymous = 1 + iota
	PostHasLink
	PostHasLow_Rep_Author
	PostHasBad_Word
	PostIsFlagged
	PostNoIssue
)

const (
	PostUnapproved  = "unapproved"
	PostApproved    = "approved"
	PostSpam        = "spam"
	PostDeleted     = "deleted"
	PostFlagged     = "flagged"
	PostHighlighted = "highlighted"
)

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

const (
	OrderAsc  = "asc"
	OrderDesc = "desc"
)

const (
	SortTypeDate     = "date"
	SortTypePriority = "priority"
)
