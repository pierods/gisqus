package gisqus

import (
	"net/url"
	"testing"

	"github.com/pierods/gisqus/mock"
)

func TestUserPosts(t *testing.T) {

	mockServer = ms.NewServer()
	defer mockServer.Close()

	usersUrls.user_post_list_url, err = mock.SwitchHostAndScheme(usersUrls.user_post_list_url, mockServer.URL)
	if err != nil {
		t.Fatal(err)
	}

	testValues = url.Values{}
	_, err = testGisqus.UserPosts("", testValues, testCtx)
	if err == nil {
		t.Fatal("Should be able to reject a null user")
	}
	posts, err := testGisqus.UserPosts("79849", testValues, testCtx)
	if err != nil {
		t.Fatal("Should be able to call the forum list users endpoint - ", err)
	}
	if len(posts.Response) != 25 {
		t.Fatal("Should be able to retrieve all posts")
	}
	if posts.Response[0].Message != "<p>Looks really good so far....keeping my fingers crossed.</p>" {
		t.Fatal("Should be able to retrieve post message")
	}
	if posts.Response[0].Id != "2978710471" {
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

func TestUserDetails(t *testing.T) {

	mockServer = ms.NewServer()
	defer mockServer.Close()

	usersUrls.user_detail_url, err = mock.SwitchHostAndScheme(usersUrls.user_detail_url, mockServer.URL)
	if err != nil {
		t.Fatal(err)
	}
	testValues = url.Values{}
	_, err = testGisqus.UserDetails("", testValues, testCtx)
	if err == nil {
		t.Fatal("Should check for an empty user id")
	}
	user, err := testGisqus.UserDetails("79849", testValues, testCtx)
	if err != nil {
		t.Fatal(err)
	}
	if user.Response.Id != "79849" {
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

	mockServer = ms.NewServer()
	defer mockServer.Close()

	usersUrls.user_interesting_users_url, err = mock.SwitchHostAndScheme(usersUrls.user_interesting_users_url, mockServer.URL)
	if err != nil {
		t.Fatal(err)
	}

	testValues = url.Values{}
	users, err := testGisqus.UserInteresting(testValues, testCtx)
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
	if users.Response.Objects["auth.User?id=160076302"].ProfileUrl != "https://disqus.com/by/anticonsoleshit/" {
		t.Fatal("Should be able to retrieve a profile url")
	}
	if users.Response.Objects["auth.User?id=160076302"].Reputation != 6.920859999999999 {
		t.Fatal("Should be able to retrieve a user's reputation")
	}
	if ToDisqusTime(users.Response.Objects["auth.User?id=160076302"].JoinedAt) != "2015-06-02T14:43:19" {
		t.Fatal("Should be able to retrieve a user's joined at")
	}
	if users.Response.Objects["auth.User?id=160076302"].Id != "160076302" {
		t.Fatal("Should be able to retrieve a user's id")
	}

}

func TestUserActiveForums(t *testing.T) {

	mockServer = ms.NewServer()
	defer mockServer.Close()

	usersUrls.user_active_forums, err = mock.SwitchHostAndScheme(usersUrls.user_active_forums, mockServer.URL)
	if err != nil {
		t.Fatal(err)
	}

	testValues = url.Values{}
	_, err = testGisqus.UserActiveForums("", testValues, testCtx)
	if err == nil {
		t.Fatal("Should be able to reject an emtpy user id")
	}

	forums, err := testGisqus.UserActiveForums("46351054", testValues, testCtx)
	if err != nil {
		t.Fatal(err)
	}
	if len(forums.Response) != 25 {
		t.Fatal("Should be able to correctly unmarshal items")
	}
	if forums.Response[0].CreatedAt.Format(disqusDateFormatExact) != "2008-04-09T23:30:16.843273" {
		t.Fatal("Should be able to correctly unmarshal items")
	}
	if forums.Response[0].Founder != "847" {
		t.Fatal("Should be able to correctly retrieve a forum founder")
	}
	if forums.Response[0].Id != "tvnewser" {
		t.Fatal("Should be able to correctly retrieve a forum id")
	}
	if forums.Response[0].Name != "TVNewser" {
		t.Fatal("Should be able to correctly retrieve a forum name")
	}
	if forums.Response[0].OrganizationId != 618 {
		t.Fatal("Should be able to correctly retrieve a forum org id")
	}
}

func TestUserFollowers(t *testing.T) {

	mockServer = ms.NewServer()
	defer mockServer.Close()

	usersUrls.user_followers, err = mock.SwitchHostAndScheme(usersUrls.user_followers, mockServer.URL)
	if err != nil {
		t.Fatal(err)
	}

	testValues = url.Values{}
	_, err = testGisqus.UserFollowers("", testValues, testCtx)
	if err == nil {
		t.Fatal("Should check for an empty user id")
	}
	users, err := testGisqus.UserFollowers("46351054", testValues, testCtx)
	if err != nil {
		t.Fatal(err)
	}
	if len(users.Response) != 25 {
		t.Fatal("Should be able to parse result set entirely")
	}
	if users.Response[0].Id != "32414357" {
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

	mockServer = ms.NewServer()
	defer mockServer.Close()

	usersUrls.user_following, err = mock.SwitchHostAndScheme(usersUrls.user_following, mockServer.URL)
	if err != nil {
		t.Fatal(err)
	}

	testValues = url.Values{}
	_, err = testGisqus.UserFollowing("", testValues, testCtx)
	if err == nil {
		t.Fatal("Should check for an empty user id")
	}
	users, err := testGisqus.UserFollowing("195792235", testValues, testCtx)
	if err != nil {
		t.Fatal(err)
	}
	if len(users.Response) != 25 {
		t.Fatal("Should be able to parse result set entirely")
	}
	if users.Response[0].Id != "32078576" {
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

	mockServer = ms.NewServer()
	defer mockServer.Close()

	usersUrls.user_following_forums, err = mock.SwitchHostAndScheme(usersUrls.user_following_forums, mockServer.URL)
	if err != nil {
		t.Fatal(err)
	}

	testValues = url.Values{}
	_, err = testGisqus.UserForumFollowing("", testValues, testCtx)
	if err == nil {
		t.Fatal("Should be able to reject an emtpy user id")
	}

	forums, err := testGisqus.UserForumFollowing("46351054", testValues, testCtx)
	if err != nil {
		t.Fatal(err)
	}
	if len(forums.Response) != 16 {
		t.Fatal("Should be able to correctly unmarshal items")
	}
	if forums.Response[0].CreatedAt.Format(disqusDateFormatExact) != "2015-06-04T17:40:19.641774" {
		t.Fatal("Should be able to correctly unmarshal items")
	}
	if forums.Response[0].Founder != "172746617" {
		t.Fatal("Should be able to correctly retrieve a forum founder")
	}
	if forums.Response[0].Id != "channel-animeforthepeople" {
		t.Fatal("Should be able to correctly retrieve a forum id")
	}
	if forums.Response[0].Name != "Anime For The People" {
		t.Fatal("Should be able to correctly retrieve a forum name")
	}
	if forums.Response[0].OrganizationId != 3644738 {
		t.Fatal("Should be able to correctly retrieve a forum org id")
	}
}
