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
	usersMostActiveForumsJSON, err := readTestFile("usersmostactiveforums.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	usersActivitiesJSON, err := readTestFile("userslistactivity.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	usersListPostsJSON, err := readTestFile("userslistposts.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	usersUserDetailsJSON, err := readTestFile("usersuserdetail.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	usersInterestingUsersJSON, err := readTestFile("usersinterestingusers.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	usersActiveForumsJSON, err := readTestFile("usersactiveforums.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	usersFollowersJSON, err := readTestFile("usersfollowers.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	usersFollowingJSON, err := readTestFile("usersfollowing.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	usersForumFollowingJSON, err := readTestFile("usersfollowingforums.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	usersUrls.PostListURL, testErr = mockServer.SwitchHostAndScheme(usersUrls.PostListURL, usersListPostsJSON)
	if testErr != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	usersUrls.MostActiveForumsURL, testErr = mockServer.SwitchHostAndScheme(usersUrls.MostActiveForumsURL, usersMostActiveForumsJSON)
	if testErr != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	usersUrls.ActivityURL, testErr = mockServer.SwitchHostAndScheme(usersUrls.ActivityURL, usersActivitiesJSON)
	if testErr != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	usersUrls.DetailURL, testErr = mockServer.SwitchHostAndScheme(usersUrls.DetailURL, usersUserDetailsJSON)
	if testErr != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	usersUrls.InterestingIUsersURL, testErr = mockServer.SwitchHostAndScheme(usersUrls.InterestingIUsersURL, usersInterestingUsersJSON)
	if testErr != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	usersUrls.ActiveForumsURL, testErr = mockServer.SwitchHostAndScheme(usersUrls.ActiveForumsURL, usersActiveForumsJSON)
	if testErr != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	usersUrls.FollowersURL, testErr = mockServer.SwitchHostAndScheme(usersUrls.FollowersURL, usersFollowersJSON)
	if testErr != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	usersUrls.FollowingURL, testErr = mockServer.SwitchHostAndScheme(usersUrls.FollowingURL, usersFollowingJSON)
	if testErr != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	usersUrls.FollowingForumsURL, testErr = mockServer.SwitchHostAndScheme(usersUrls.FollowingForumsURL, usersForumFollowingJSON)
	if testErr != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func TestUserPosts(t *testing.T) {

	_, testErr = testGisqus.UserPosts(testCtx, "", testValues)
	if testErr == nil {
		t.Fatal("Should be able to reject a null user")
	}
	posts, err := testGisqus.UserPosts(testCtx, "79849", testValues)
	if err != nil {
		t.Fatal("Should be able to call the forum list users endpoint - ", err)
	}
	if len(posts.Response) != 25 {
		t.Fatal("Should be able to retrieve all posts")
	}
	if posts.Response[0].Message != "<p>Looks really good so far....keeping my fingers crossed.</p>" {
		t.Fatal("Should be able to retrieve post message")
	}
	if posts.Response[0].ID != "2978710471" {
		t.Fatal("Should be able to retrieve post id")
	}
	if posts.Response[0].Author.Username != "laross19" {
		t.Fatal("Should be able to retrieve author username")
	}
	if ToDisqusTime(posts.Response[0].CreatedAt) != "2016-11-01T02:54:43" {
		t.Fatal("Should be able to retrieve post created at")
	}
	if posts.Response[0].Parent != 2978642690 {
		t.Fatal("Should be able to retrieve post parent")
	}
	if posts.Response[0].Thread != "5268209091" {
		t.Fatal("Should be able to retrieve post thread")
	}
	if posts.Response[0].Forum != "mapleleafshotstove" {
		t.Fatal("Should be able to retrieve post forum")
	}
}

func TestMostActiveForums(t *testing.T) {

	_, testErr = testGisqus.UserMostActiveForums(testCtx, "", testValues)
	if testErr == nil {
		t.Fatal("Should check for an empty user id")
	}
	forums, err := testGisqus.UserMostActiveForums(testCtx, "253940813", testValues)
	if err != nil {
		t.Fatal(err)
	}
	if ToDisqusTimeExact(forums.Response[0].CreatedAt) != "2011-06-02T18:50:59.765656" {
		t.Fatal("Should be able to parse the created at field")
	}
	if forums.Response[0].Founder != "110449899" {
		t.Fatal("Should be able to retrieve founder")
	}

	if forums.Response[0].OrganizationID != 110 {
		t.Fatal("Should be able to retrieve an organization id")
	}
	if !forums.Response[0].Settings.OrganicDiscoveryEnabled {
		t.Fatal("Should be able to retrieve organicDiscoveryEnabled")
	}
}

func TestUserActivities(t *testing.T) {

	_, testErr = testGisqus.UserDetails(testCtx, "", testValues)
	if testErr == nil {
		t.Fatal("Should check for an empty user id")
	}
	activities, err := testGisqus.UserActivities(testCtx, "79849", testValues)
	if err != nil {
		t.Fatal(err)
	}
	if len(activities.Posts) != 25 {
		t.Fatal("Should be able to retrieve all posts")
	}
	if activities.Posts[0].ID != "3357202237" {
		t.Fatal("Should be able to retrieve a post id")
	}
	if activities.Posts[0].Author.Username != "coachbuzzcut" {
		t.Fatal("Should be able to retrieve a post's author's username")
	}
	if activities.Posts[0].Author.ID != "253940813" {
		t.Fatal("Should be able to retrieve a post's username's id")
	}
	if ToDisqusTime(activities.Posts[0].CreatedAt) != "2017-06-13T10:20:29" {
		t.Fatal("Should be able to retrieve a post's created at")
	}
	if activities.Posts[0].Parent != 3356547778 {
		t.Fatal("Should be able to retrieve a post's parent")
	}
	if activities.Posts[0].Thread != "5903840168" {
		t.Fatal("Should be able to retrieve a post's thread")
	}
	if activities.Posts[0].Forum != "tmz" {
		t.Fatal("Should be able to retrieve a post's forum")
	}
}

func TestUserDetails(t *testing.T) {

	_, testErr = testGisqus.UserDetails(testCtx, "", testValues)
	if testErr == nil {
		t.Fatal("Should check for an empty user id")
	}
	user, err := testGisqus.UserDetails(testCtx, "79849", testValues)
	if err != nil {
		t.Fatal(err)
	}
	if user.Response.ID != "79849" {
		t.Fatal("Should be able to retrieve a user id")
	}
	if user.Response.Rep != 1.2537909999999999 {
		t.Fatal("Should be able to retrieve a user's rep")
	}
	if ToDisqusTime(user.Response.JoinedAt) != "2008-08-10T02:54:57" {
		t.Fatal("Should be able to retrieve a user's joined at")
	}
	if user.Response.Username != "laross19" {
		t.Fatal("Should be able to retrieve a user's username")
	}
	if user.Response.NumLikesReceived != 14 {
		t.Fatal("Should be able to retrieve a users' likes received")
	}
	if user.Response.NumPosts != 56 {
		t.Fatal("Should be able to retrieve a user's number of posts")
	}
	if user.Response.NumFollowers != 22 {
		t.Fatal("Should be able to retrieve a user's number of followers")
	}
	if user.Response.NumFollowing != 33 {
		t.Fatal("Should be able to retrieve a user's number of following")
	}
	if user.Response.NumForumsFollowing != 44 {
		t.Fatal("Should be able to retrieve a user's number of forums following")
	}
}

func TestUserInteresting(t *testing.T) {

	users, err := testGisqus.UserInteresting(testCtx, testValues)
	if err != nil {
		t.Fatal("Should be able to call the forum list interesting users endpoint - ", err)
	}
	if len(users.Response.Items) != 5 {
		t.Fatal("Should be able to retrieve all posts")
	}
	if len(users.Response.Objects) != 5 {
		t.Fatal("Should be able to retrieve all posts")
	}
	if users.Response.Objects["auth.User?id=160076302"].Username != "anticonsoleshit" {
		t.Fatal("Should be able to retrieve a username")
	}
	if users.Response.Objects["auth.User?id=160076302"].Name != "de ja ful" {
		t.Fatal("Should be able to retrieve a user name")
	}
	if users.Response.Objects["auth.User?id=160076302"].ProfileURL != "https://disqus.com/by/anticonsoleshit/" {
		t.Fatal("Should be able to retrieve a profile url")
	}
	if users.Response.Objects["auth.User?id=160076302"].Reputation != 6.920859999999999 {
		t.Fatal("Should be able to retrieve a user's reputation")
	}
	if ToDisqusTime(users.Response.Objects["auth.User?id=160076302"].JoinedAt) != "2015-06-02T14:43:19" {
		t.Fatal("Should be able to retrieve a user's joined at")
	}
	if users.Response.Objects["auth.User?id=160076302"].ID != "160076302" {
		t.Fatal("Should be able to retrieve a user's id")
	}

}

func TestUserActiveForums(t *testing.T) {

	_, testErr = testGisqus.UserActiveForums(testCtx, "", testValues)
	if testErr == nil {
		t.Fatal("Should be able to reject an empty user id")
	}

	forums, err := testGisqus.UserActiveForums(testCtx, "46351054", testValues)
	if err != nil {
		t.Fatal(err)
	}
	if len(forums.Response) != 25 {
		t.Fatal("Should be able to correctly unmarshal items")
	}
	if forums.Response[0].CreatedAt.Format(DisqusDateFormatExact) != "2008-04-09T23:30:16.843273" {
		t.Fatal("Should be able to correctly unmarshal items")
	}
	if forums.Response[0].Founder != "847" {
		t.Fatal("Should be able to correctly retrieve a forum founder")
	}
	if forums.Response[0].ID != "tvnewser" {
		t.Fatal("Should be able to correctly retrieve a forum id")
	}
	if forums.Response[0].Name != "TVNewser" {
		t.Fatal("Should be able to correctly retrieve a forum name")
	}
	if forums.Response[0].OrganizationID != 618 {
		t.Fatal("Should be able to correctly retrieve a forum org id")
	}
}

func TestUserFollowers(t *testing.T) {

	_, testErr = testGisqus.UserFollowers(testCtx, "", testValues)
	if testErr == nil {
		t.Fatal("Should check for an empty user id")
	}
	users, err := testGisqus.UserFollowers(testCtx, "46351054", testValues)
	if err != nil {
		t.Fatal(err)
	}
	if len(users.Response) != 25 {
		t.Fatal("Should be able to parse result set entirely")
	}
	if users.Response[0].ID != "32414357" {
		t.Fatal("Should be able to retrieve a user id")
	}
	if users.Response[0].Rep != 0.4153629999999999 {
		t.Fatal("Should be able to retrieve a user's rep")
	}
	if ToDisqusTime(users.Response[0].JoinedAt) != "2012-09-18T17:29:47" {
		t.Fatal("Should be able to retrieve a user's joined at")
	}
	if users.Response[0].Username != "disqus_IpEgXB3c55" {
		t.Fatal("Should be able to retrieve a user's username")
	}

}

func TestUserFollowing(t *testing.T) {

	_, testErr = testGisqus.UserFollowing(testCtx, "", testValues)
	if testErr == nil {
		t.Fatal("Should check for an empty user id")
	}
	users, err := testGisqus.UserFollowing(testCtx, "195792235", testValues)
	if err != nil {
		t.Fatal(err)
	}
	if len(users.Response) != 25 {
		t.Fatal("Should be able to parse result set entirely")
	}
	if users.Response[0].ID != "32078576" {
		t.Fatal("Should be able to retrieve a user id")
	}
	if users.Response[0].Rep != 1.3459269999999999 {
		t.Fatal("Should be able to retrieve a user's rep")
	}
	if ToDisqusTime(users.Response[0].JoinedAt) != "2012-09-13T05:54:26" {
		t.Fatal("Should be able to retrieve a user's joined at")
	}
	if users.Response[0].Username != "flamedance58" {
		t.Fatal("Should be able to retrieve a user's username")
	}

}

func TestUserForumFollowing(t *testing.T) {

	_, testErr = testGisqus.UserForumFollowing(testCtx, "", testValues)
	if testErr == nil {
		t.Fatal("Should be able to reject an empty user id")
	}

	forums, err := testGisqus.UserForumFollowing(testCtx, "46351054", testValues)
	if err != nil {
		t.Fatal(err)
	}
	if len(forums.Response) != 16 {
		t.Fatal("Should be able to correctly unmarshal items")
	}
	if forums.Response[0].CreatedAt.Format(DisqusDateFormatExact) != "2015-06-04T17:40:19.641774" {
		t.Fatal("Should be able to correctly unmarshal items")
	}
	if forums.Response[0].Founder != "172746617" {
		t.Fatal("Should be able to correctly retrieve a forum founder")
	}
	if forums.Response[0].ID != "channel-animeforthepeople" {
		t.Fatal("Should be able to correctly retrieve a forum id")
	}
	if forums.Response[0].Name != "Anime For The People" {
		t.Fatal("Should be able to correctly retrieve a forum name")
	}
	if forums.Response[0].OrganizationID != 3644738 {
		t.Fatal("Should be able to correctly retrieve a forum org id")
	}
}
