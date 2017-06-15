// Copyright Piero de Salvia.
// All Rights Reserved

package gisqus

import (
	"fmt"
	"net/url"
	"os"
	"testing"
)

func init() {

	var err error
	threadSetJSON, err := readTestFile("threadsset.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	threadUsersVotedJSON, err := readTestFile("threadsusersvoted.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	threadListJSON, err := readTestFile("threadsthreadlist.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	threadDetailsJSON, err := readTestFile("threadsthreaddetails.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	threadPostsJSON, err := readTestFile("threadsthreadposts.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	threadListHotJSON, err := readTestFile("threadshotlist.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	threadListPopularJSON, err := readTestFile("threadspopular.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	threadListTrendingJSON, err := readTestFile("threadstrending.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	threadsUrls.ThreadUsersVotedURL, testErr = mockServer.SwitchHostAndScheme(threadsUrls.ThreadUsersVotedURL, threadUsersVotedJSON)
	if testErr != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	threadsUrls.ThreadSetURL, testErr = mockServer.SwitchHostAndScheme(threadsUrls.ThreadSetURL, threadSetJSON)
	if testErr != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	threadsUrls.ThreadListURL, testErr = mockServer.SwitchHostAndScheme(threadsUrls.ThreadListURL, threadListJSON)
	if testErr != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	threadsUrls.ThreadDetailURL, testErr = mockServer.SwitchHostAndScheme(threadsUrls.ThreadDetailURL, threadDetailsJSON)
	if testErr != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	threadsUrls.ThreadPostsURL, testErr = mockServer.SwitchHostAndScheme(threadsUrls.ThreadPostsURL, threadPostsJSON)
	if testErr != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	threadsUrls.ThreadHotURL, testErr = mockServer.SwitchHostAndScheme(threadsUrls.ThreadHotURL, threadListHotJSON)
	if testErr != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	threadsUrls.ThreadPopularURL, testErr = mockServer.SwitchHostAndScheme(threadsUrls.ThreadPopularURL, threadListPopularJSON)
	if testErr != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	threadsUrls.ThreadTrendingURL, testErr = mockServer.SwitchHostAndScheme(threadsUrls.ThreadTrendingURL, threadListTrendingJSON)
	if testErr != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func TestThreadUsersVoted(t *testing.T) {

	testValues = url.Values{}

	_, testErr = testGisqus.ThreadDetails(testCtx, "", testValues)
	if testErr == nil {
		t.Fatal("Should check for an empty thread id")
	}
	users, err := testGisqus.ThreadUsersVoted(testCtx, "5846923796", testValues)
	if err != nil {
		t.Fatal(err)
	}

	if len(users.Response) != 5 {
		t.Fatal("Should be able to parse result set entirely")
	}
	if users.Response[0].ID != "19365741" {
		t.Fatal("Should be able to retrieve a user id")
	}
	if users.Response[0].IsPowerContributor {
		t.Fatal("Should be able to retrieve a user's power contributor")
	}
	if ToDisqusTime(users.Response[0].JoinedAt) != "2011-11-22T10:43:15" {
		t.Fatal("Should be able to retrieve a user's joined at")
	}
	if users.Response[0].Username != "bigboss400" {
		t.Fatal("Should be able to retrieve a user's username")
	}
}

func TestThreadSet(t *testing.T) {

	testValues = url.Values{}

	_, testErr = testGisqus.ThreadSet(testCtx, []string{}, testValues)
	if testErr == nil {
		t.Fatal("Should check for an empty thread id")
	}
	_, testErr = testGisqus.ThreadSet(testCtx, nil, testValues)
	if testErr == nil {
		t.Fatal("Should check for an empty thread id")
	}
	threads, err := testGisqus.ThreadSet(testCtx, []string{"5903840168", "5850192558"}, testValues)
	if err != nil {
		t.Fatal(err)
	}
	if len(threads.Response) != 2 {
		t.Fatal("Should be able to correctly parse a thread list")
	}
	if threads.Response[0].Feed != "https://tmz.disqus.com/039bachelor_in_paradise039_star_corinne_olympios_says_she_didn039t_consent_to_sexual_contact_with_de/latest.rss" {
		t.Fatal("Should be able to retrieve a thread's feed url")
	}
	if threads.Response[0].ID != "5903840168" {
		t.Fatal("Should be able to retrieve a thread's id")
	}
	if threads.Response[0].Category != "3341905" {
		t.Fatal("Should be able to retrieve a thread's category")
	}
	if threads.Response[0].Author != "116162885" {
		t.Fatal("Should be able to retrieve a thread's author")
	}
	if ToDisqusTime(threads.Response[0].CreatedAt) != "2017-06-12T17:48:04" {
		t.Fatal("Should be able to retrieve a thread's created at")
	}
	if threads.Response[0].Forum != "tmz" {
		t.Fatal("Should be able to retrieve a thread's forum id")
	}
	if threads.Response[0].Title != "&#039;Bachelor in Paradise&#039; Star Corinne Olympios Says She Didn&#039;t Consent to Sexual Contact with DeMario Jackson" {
		t.Fatal("Should be able to retrieve a thread's title")
	}
}

func TestThreadList(t *testing.T) {

	testValues = url.Values{}
	threads, err := testGisqus.ThreadList(testCtx, testValues)
	if err != nil {
		t.Fatal("Should be able to call the thread list endpoint - ", err)
	}
	if len(threads.Response) != 25 {
		t.Fatal("Should be able to correctly parse a thread list")
	}
	if threads.Response[0].Feed != "https://babbel-magazine.disqus.com/personalidades_multilingues_ao_longo_da_historia_babbelcom_087/latest.rss" {
		t.Fatal("Should be able to retrieve a thread's feed url")
	}
	if threads.Response[0].ID != "5850192558" {
		t.Fatal("Should be able to retrieve a thread's id")
	}
	if threads.Response[0].Category != "3261556" {
		t.Fatal("Should be able to retrieve a thread's category")
	}
	if threads.Response[0].Author != "121561733" {
		t.Fatal("Should be able to retrieve a thread's author")
	}
	if ToDisqusTime(threads.Response[0].CreatedAt) != "2017-05-25T18:16:19" {
		t.Fatal("Should be able to retrieve a thread's created at")
	}
	if threads.Response[0].Forum != "babbel-magazine" {
		t.Fatal("Should be able to retrieve a thread's forum id")
	}
	if threads.Response[0].Title != "Personalidades multilíngues ao longo da História - Babbel.com" {
		t.Fatal("Should be able to retrieve a thread's title")
	}
}

func TestThreadDetails(t *testing.T) {

	testValues = url.Values{}

	_, testErr = testGisqus.ThreadDetails(testCtx, "", testValues)
	if testErr == nil {
		t.Fatal("Should check for an empty thread id")
	}
	details, err := testGisqus.ThreadDetails(testCtx, "5846923796", testValues)
	if err != nil {
		t.Fatal(err)
	}
	if details.Response.ID != "5846923796" {
		t.Fatal("Should be able to retrieve a thread id")
	}
	if details.Response.Category != "783882" {
		t.Fatal("Should be able to retrieve a thread id")
	}
	if details.Response.Author != "9408501" {
		t.Fatal("Should be able to retrieve a thread's author")
	}
	if ToDisqusTime(details.Response.CreatedAt) != "2017-05-24T16:41:44" {
		t.Fatal("Should be able to parse a thread's created at")
	}
	if details.Response.Forum != "mapleleafshotstove" {
		t.Fatal("Should be able to retrieve a thread's forum id")
	}
	if details.Response.Posts != 1927 {
		t.Fatal("Should be able to retrieve a thread's number of posts")
	}
}

func TestThreadPosts(t *testing.T) {

	testValues = url.Values{}
	_, testErr = testGisqus.ThreadPosts(testCtx, "", testValues)
	if testErr == nil {
		t.Fatal("Should check for empty thread id")
	}
	posts, err := testGisqus.ThreadPosts(testCtx, "5846923796", testValues)
	if err != nil {
		t.Fatal(err)
	}
	if len(posts.Response) != 25 {
		t.Fatal("Should be able to correctly parse a post list")
	}
	if posts.Response[0].ID != "3325943139" {
		t.Fatal("Should be able to retrieve a post id")
	}
	if posts.Response[0].Author.Username != "loovtrain" {
		t.Fatal("Should be able to retrieve a post's author's username")
	}
	if posts.Response[0].Author.ID != "163477624" {
		t.Fatal("Should be able to retrieve a post's user's id")
	}
	if ToDisqusTime(posts.Response[0].CreatedAt) != "2017-05-26T15:12:18" {
		t.Fatal("Should be able to retrieve a post's created at")
	}
	if posts.Response[0].Parent != 3325896546 {
		t.Fatal("Should be able to retrieve a post's parent")
	}
	if posts.Response[0].Thread != "5846923796" {
		t.Fatal("Should be able to retrieve a post's thread")
	}
	if posts.Response[0].Forum != "mapleleafshotstove" {
		t.Fatal("Should be able to retrieve a post's forum")
	}

}

func TestThreadListHot(t *testing.T) {

	testValues = url.Values{}
	threads, err := testGisqus.ThreadHot(testCtx, testValues)
	if err != nil {
		t.Fatal("Should be able to call the thread list endpoint - ", err)
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

func TestThreadListPopular(t *testing.T) {

	values := url.Values{}
	threads, err := testGisqus.ThreadPopular(testCtx, values)
	if err != nil {
		t.Fatal("Should be able to call the thread list popular endpoint - ", err)
	}
	if len(threads.Response) != 25 {
		t.Fatal("Should be able to correctly parse a thread list")
	}
	if threads.Response[0].Feed != "https://alloutdoor.disqus.com/sig_sauer_sued_by_new_jersey_state_police/latest.rss" {
		t.Fatal("Should be able to retrieve a thread's feed url")
	}
	if threads.Response[0].ID != "5829486853" {
		t.Fatal("Should be able to retrieve a thread's id")
	}
	if threads.Response[0].Category != "2409406" {
		t.Fatal("Should be able to retrieve a thread's category")
	}
	if threads.Response[0].Author != "37536641" {
		t.Fatal("Should be able to retrieve a thread's author")
	}
	if ToDisqusTime(threads.Response[0].CreatedAt) != "2017-05-18T19:23:54" {
		t.Fatal("Should be able to retrieve a thread's created at")
	}
	if threads.Response[0].Forum != "alloutdoor" {
		t.Fatal("Should be able to retrieve a thread's forum id")
	}
	if threads.Response[0].Title != "Sig Sauer Sued By New Jersey State Police" {
		t.Fatal("Should be able to retrieve a thread's title")
	}
}

func TestThreadListTrending(t *testing.T) {

	testValues = url.Values{}
	trends, err := testGisqus.ThreadTrending(testCtx, testValues)
	if err != nil {
		t.Fatal("Should be able to call the thread trending endpoint - ", err)
	}
	if len(trends.Response) != 10 {
		t.Fatal("Should be able to correctly parse a thread list")
	}
	if trends.Response[2].PostLikes != 1665 {
		t.Fatal("Should be able to retrieve a trend's postlikes")
	}
	if trends.Response[2].Posts != 62 {
		t.Fatal("Should be able to retrieve a trend's posts")
	}
	if trends.Response[2].Score != 1.497732426303855 {
		t.Fatal("Should be able to retrieve a trends's score")
	}
	if trends.Response[2].Likes != 90 {
		t.Fatal("Should be able to retrieve a trends' likes")
	}
	if trends.Response[2].TrendingThread.Feed != "https://kissanime.disqus.com/berserk_2017_anime_watch_berserk_2017_anime_online_in_high_quality/latest.rss" {
		t.Fatal("Should be able to retrieve a thread's feed url")
	}
	if trends.Response[2].TrendingThread.ID != "5592902940" {
		t.Fatal("Should be able to retrieve a thread's id")
	}
	if trends.Response[2].TrendingThread.Category != "3204063" {
		t.Fatal("Should be able to retrieve a thread's category")
	}
	if trends.Response[2].TrendingThread.Author != "100108732" {
		t.Fatal("Should be able to retrieve a thread's author")
	}
	if ToDisqusTime(trends.Response[2].TrendingThread.CreatedAt) != "2017-03-01T01:42:44" {
		t.Fatal("Should be able to retrieve a thread's created at")
	}
	if trends.Response[2].TrendingThread.Forum != "kissanime" {
		t.Fatal("Should be able to retrieve a thread's forum id")
	}
	if trends.Response[2].TrendingThread.Title != "Berserk (2017) anime | Watch Berserk (2017) anime online in high quality" {
		t.Fatal("Should be able to retrieve a thread's title")
	}
	if trends.Response[2].TrendingThread.HighlightedPost.ID != "3316658778" {
		t.Fatal("Should be able to retrieve a trend's highlighted post id")
	}
	if ToDisqusTime(trends.Response[2].TrendingThread.HighlightedPost.CreatedAt) != "2017-05-20T23:15:11" {
		t.Fatal("Should be able to retrieve a trend's highlighted post created at")
	}
	if trends.Response[2].TrendingThread.HighlightedPost.Author.Username != "Umbrielle" {
		t.Fatal("Should be able to retrieve a trend's highlighted post's author")
	}
	if ToDisqusTime(trends.Response[2].TrendingThread.HighlightedPost.Author.JoinedAt) != "2015-02-06T14:28:51" {
		t.Fatal("Should be able to retrieve a trend's highlighted post's joined at")
	}
	if trends.Response[2].TrendingThread.HighlightedPost.Author.ID != "143213885" {
		t.Fatal("Should be able to retrieve a trend's highlighted post's author")
	}
}
