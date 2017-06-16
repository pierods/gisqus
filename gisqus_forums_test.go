// Copyright Piero de Salvia.
// All Rights Reserved

package gisqus

import (
	"fmt"
	"os"
	"testing"
)

func init() {

	var err error

	forumInterestingForumsJSON, err := readTestFile("forumsinterestingforums.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	forumMostActiveUsersJSON, err := readTestFile("forumslistmostactive.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	forumListUsersJSON, err := readTestFile("forumslistforumusers.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	forumDetailsJSON, err := readTestFile("forumsforumdetails.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	forumListCategoriesJSON, err := readTestFile("forumslistcategories.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	forumThreadListJSON, err := readTestFile("forumsforumlistthreads.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	forumMostLikedUsersJSON, err := readTestFile("forumsmostlikedusers.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	forumFollowersJSON, err := readTestFile("forumslistfollowers.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	forumsUrls.MostActiveUsersURL, testErr = mockServer.SwitchHostAndScheme(forumsUrls.MostActiveUsersURL, forumMostActiveUsersJSON)
	if testErr != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	forumsUrls.ListFollowersURL, testErr = mockServer.SwitchHostAndScheme(forumsUrls.ListFollowersURL, forumFollowersJSON)
	if testErr != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	forumsUrls.ListUsersURL, testErr = mockServer.SwitchHostAndScheme(forumsUrls.ListUsersURL, forumListUsersJSON)
	if testErr != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	forumsUrls.InterestingForumsURL, testErr = mockServer.SwitchHostAndScheme(forumsUrls.InterestingForumsURL, forumInterestingForumsJSON)
	if testErr != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	forumsUrls.DetailsURL, testErr = mockServer.SwitchHostAndScheme(forumsUrls.DetailsURL, forumDetailsJSON)
	if testErr != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	forumsUrls.CategoriesURL, testErr = mockServer.SwitchHostAndScheme(forumsUrls.CategoriesURL, forumListCategoriesJSON)
	if testErr != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	forumsUrls.ListThreadsURL, testErr = mockServer.SwitchHostAndScheme(forumsUrls.ListThreadsURL, forumThreadListJSON)
	if testErr != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	forumsUrls.MostLikedUsersURL, testErr = mockServer.SwitchHostAndScheme(forumsUrls.MostLikedUsersURL, forumMostLikedUsersJSON)
	if testErr != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	forumsUrls.ListUsersURL, testErr = mockServer.SwitchHostAndScheme(forumsUrls.ListUsersURL, forumListUsersJSON)
	if testErr != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func TestForumMostActiveUsers(t *testing.T) {

	_, testErr = testGisqus.ForumMostActiveUsers(testCtx, "", testValues)
	if testErr == nil {
		t.Fatal("Should be able to reject a null forum")
	}
	users, err := testGisqus.ForumMostActiveUsers(testCtx, "mapleleafshotstove", testValues)
	if err != nil {
		t.Fatal("Should be able to call the forum followers endpoint - ", err)
	}

	if len(users.Response) != 24 {
		t.Fatal("Should be able to correctly parse a user list")
	}
	if users.Response[0].Username != "icechest" {
		t.Fatal("Should be able to retrieve a username")
	}
	if users.Response[0].Rep != 23.690665 {
		t.Fatal("Should be able to retrieve a reputation")
	}
	if ToDisqusTime(users.Response[0].JoinedAt) != "2015-07-06T22:57:31" {
		t.Fatal("Should be able to retrieve a joined at date")
	}
	if users.Response[0].Avatar.Small.Permalink != "https://disqus.com/api/users/avatars/icechest.jpg" {
		t.Fatal("Should be able to retrieve an avatar")
	}
	if users.Response[0].Avatar.Small.Cache != "https://c.disquscdn.com/uploads/users/16444/4895/avatar32.jpg?1461376631" {
		t.Fatal("Should be able to retrieve an avatar")
	}
}

func TestForumFollowers(t *testing.T) {

	_, testErr = testGisqus.ForumFollowers(testCtx, "", testValues)
	if testErr == nil {
		t.Fatal("Should be able to reject a null forum")
	}
	users, err := testGisqus.ForumFollowers(testCtx, "mapleleafshotstove", testValues)
	if err != nil {
		t.Fatal("Should be able to call the forum followers endpoint - ", err)
	}

	if len(users.Response) != 25 {
		t.Fatal("Should be able to correctly parse a user list")
	}
	if users.Response[0].Username != "Johnld778" {
		t.Fatal("Should be able to retrieve a username")
	}
	if users.Response[0].Rep != 1.375799 {
		t.Fatal("Should be able to retrieve a reputation")
	}
	if ToDisqusTime(users.Response[0].JoinedAt) != "2008-02-27T08:05:22" {
		t.Fatal("Should be able to retrieve a joined at date")
	}
	if users.Response[0].Avatar.Small.Permalink != "https://disqus.com/api/users/avatars/Johnld778.jpg" {
		t.Fatal("Should be able to retrieve an avatar")
	}
	if users.Response[0].Avatar.Small.Cache != "https://c.disquscdn.com/uploads/users/12235/avatar32.jpg?1395182401" {
		t.Fatal("Should be able to retrieve an avatar")
	}
}

func TestForumUsers(t *testing.T) {

	_, testErr = testGisqus.ForumUsers(testCtx, "", testValues)
	if testErr == nil {
		t.Fatal("Should be able to reject a null forum")
	}
	users, err := testGisqus.ForumUsers(testCtx, "mapleleafshotstove", testValues)
	if err != nil {
		t.Fatal("Should be able to call the forum list users endpoint - ", err)
	}
	if len(users.Response) != 25 {
		t.Fatal("Should be able to correctly parse a user list")
	}
	if users.Response[0].Username != "laross19" {
		t.Fatal("Should be able to retrieve a username")
	}
	if users.Response[0].Rep != 1.2537909999999999 {
		t.Fatal("Should be able to retrieve a reputation")
	}
	if ToDisqusTime(users.Response[0].JoinedAt) != "2008-08-10T02:54:57" {
		t.Fatal("Should be able to retrieve a joined at date")
	}
	if users.Response[0].Avatar.Small.Permalink != "https://disqus.com/api/users/avatars/laross19.jpg" {
		t.Fatal("Should be able to retrieve an avatar")
	}
	if users.Response[0].Avatar.Small.Cache != "//a.disquscdn.com/1495487563/images/noavatar32.png" {
		t.Fatal("Should be able to retrieve an avatar")
	}
}

func TestForumsInteresting(t *testing.T) {

	interestingForums, err := testGisqus.ForumInteresting(testCtx, testValues)
	if err != nil {
		t.Fatal(err)
	}
	if len(interestingForums.Response.Items) != 5 {
		t.Fatal("Should be able to correctly unmarshal items")
	}
	if interestingForums.Response.Items[0].Reason != "583 comments this week" {
		t.Fatal("Should be able to correctly unmarshal items")
	}
	if interestingForums.Response.Items[0].ID != "forums.Forum?id=2373958" {
		t.Fatal("Should be able to correctly unmarshal items")
	}
	if len(interestingForums.Response.Objects) != 5 {
		t.Fatal("Should be able to correctly unmarshal items")
	}
	if interestingForums.Response.Objects["forums.Forum?id=770598"].Favicon.Permalink != "https://disqus.com/api/forums/favicons/mapleleafshotstove.jpg" {
		t.Fatal("Should be able to correctly unmarshal items")
	}
	if interestingForums.Response.Objects["forums.Forum?id=770598"].Favicon.Cache != "https://c.disquscdn.com/uploads/forums/77/598/favicon.png" {
		t.Fatal("Should be able to correctly unmarshal items")
	}

	if interestingForums.Response.Objects["forums.Forum?id=770598"].CreatedAt.Format(DisqusDateFormatExact) != "2011-04-21T18:47:32.503946" {
		t.Fatal("Should be able to correctly unmarshal items")
	}
	if !interestingForums.Response.Objects["forums.Forum?id=770598"].Settings.AllowAnonPost {
		t.Fatal("Should be able to correctly unmarshal items")
	}
	if interestingForums.Response.Objects["forums.Forum?id=770598"].Avatar.Small.Permalink != "https://disqus.com/api/forums/avatars/mapleleafshotstove.jpg?size=32" {
		t.Fatal("Should be able to correctly unmarshal items")
	}
	if interestingForums.Response.Objects["forums.Forum?id=770598"].Avatar.Small.Cache != "https://c.disquscdn.com/uploads/forums/77/598/avatar32.jpg?1435553857" {
		t.Fatal("Should be able to correctly unmarshal items")
	}
}

func TestForumDetails(t *testing.T) {

	_, testErr = testGisqus.ForumDetails(testCtx, "", testValues)
	if testErr == nil {
		t.Fatal("Should check for an empty forum id")
	}
	details, err := testGisqus.ForumDetails(testCtx, "mapleleafshotstove", testValues)
	if err != nil {
		t.Fatal(err)
	}
	if ToDisqusTimeExact(details.Response.CreatedAt) != "2011-04-21T18:47:32.503946" {
		t.Fatal("Should be able to parse the created at field")
	}
	if details.Response.Founder != "9408501" {
		t.Fatal("Should be able to retrieve founder")
	}
	if !details.Response.Settings.AllowAnonPost {
		t.Fatal("Should be able to retrieve Allow Anon Post")
	}
	if details.Response.OrganizationID != 583182 {
		t.Fatal("Should be able to retrieve an organization id")
	}
}

func TestForumCategories(t *testing.T) {

	_, testErr = testGisqus.ForumCategories(testCtx, "", testValues)
	if testErr == nil {
		t.Fatal("Should check for an empty forum id")
	}
	categories, err := testGisqus.ForumCategories(testCtx, "mapleleafshotstove", testValues)
	if err != nil {
		t.Fatal(err)
	}
	if categories.Response[0].Title != "General" {
		t.Fatal("Should be able to retrieve a category name")
	}
	if categories.Response[0].Forum != "alloutdoor" {
		t.Fatal("Should be able to retrieve a forum id")
	}
	if categories.Response[0].ID != "2409406" {
		t.Fatal("Should be able to retrieve a category id")
	}

}

func TestForumThreads(t *testing.T) {

	_, testErr = testGisqus.ForumThreads(testCtx, "", testValues)
	if testErr == nil {
		t.Fatal("Should check for an empty forum id")
	}
	threads, err := testGisqus.ForumThreads(testCtx, "mapleleafshotstove", testValues)
	if err != nil {
		t.Fatal(err)
	}
	if len(threads.Response) != 25 {
		t.Fatal("Should be able to correctly parse a thread list")
	}
	if threads.Response[0].Feed != "https://mapleleafshotstove.disqus.com/leafs_links_bob_mckenzie_discusses_kyle_dubas_report_shoots_down_fictitious_william_nylander_trade_r/latest.rss" {
		t.Fatal("Should be able to retrieve a thread's feed url")
	}
	if threads.Response[0].ID != "5846923796" {
		t.Fatal("Should be able to retrieve a thread's id")
	}
	if threads.Response[0].Category != "783882" {
		t.Fatal("Should be able to retrieve a thread's category")
	}
	if threads.Response[0].Author != "9408501" {
		t.Fatal("Should be able to retrieve a thread's author")
	}
	if ToDisqusTime(threads.Response[0].CreatedAt) != "2017-05-24T16:41:44" {
		t.Fatal("Should be able to retrieve a thread's created at")
	}
	if threads.Response[0].Forum != "mapleleafshotstove" {
		t.Fatal("Should be able to retrieve a thread's forum id")
	}
	if threads.Response[0].Title != "Leafs Links: Bob McKenzie discusses Kyle Dubas report, shoots down fictitious William Nylander trade rumours; Sheldon Keefe on Carl Grundstrom, Kasperi Kapanen and more" {
		t.Fatal("Should be able to retrieve a thread's title")
	}
}

func TestForumMostLikedUsers(t *testing.T) {

	_, testErr = testGisqus.ForumMostLikedUsers(testCtx, "", testValues)
	if testErr == nil {
		t.Fatal("Should be able to reject a null forum")
	}
	users, err := testGisqus.ForumMostLikedUsers(testCtx, "mapleleafshotstove", testValues)
	if err != nil {
		t.Fatal("Should be able to call the forum list users endpoint - ", err)
	}
	if len(users.Response) != 25 {
		t.Fatal("Should be able to correctly parse a user list")
	}
	if users.Response[0].Username != "Burtonboy" {
		t.Fatal("Should be able to retrieve a username")
	}
	if users.Response[0].ID != "9413311" {
		t.Fatal("Should be able to retrieve a user id")
	}
	if ToDisqusTime(users.Response[0].JoinedAt) != "2011-04-22T02:22:13" {
		t.Fatal("Should be able to retrieve a joined at date")
	}
	// rest of user details are tested in user list test
}

func TestRetrieveCursor(t *testing.T) {

	users, err := testGisqus.ForumUsers(testCtx, "mapleleafshotstove", testValues)
	if err != nil {
		t.Fatal("Should be able to call the forum list users endpoint - ", err)
	}
	if !users.Cursor.HasNext {
		t.Fatal("Should be able to correctly parse a cursor")
	}
	if users.Cursor.Next != "2329875:0:0" {
		t.Fatal("Should be able to correctly parse a cursor")
	}
	if users.Cursor.ID != "2329875:0:0" {
		t.Fatal("Should be able to correctly parse a cursor")
	}

}

func TestExtractForumId(t *testing.T) {

	if ExtractForumID("forums.Forum?id=770598") != "770598" {
		t.Fatal("Should be able to correctly extract a forum id")
	}
}
