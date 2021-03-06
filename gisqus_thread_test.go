// Copyright Piero de Salvia.
// All Rights Reserved

package gisqus

import (
	"fmt"
	"os"
	"testing"
)

func init() {

	threadSetJSON := readTestFile("threadsset.json")
	threadUsersVotedJSON := readTestFile("threadsusersvoted.json")
	threadListJSON := readTestFile("threadsthreadlist.json")
	threadDetailsJSON := readTestFile("threadsthreaddetails.json")
	threadPostsJSON := readTestFile("threadsthreadposts.json")
	threadListHotJSON := readTestFile("threadshotlist.json")
	threadListPopularJSON := readTestFile("threadspopular.json")
	threadListTrendingJSON := readTestFile("threadstrending.json")

	switchHS := func(URL, JSON string) string {
		result, err := mockServer.SwitchHostAndScheme(URL, JSON)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		return result
	}

	threadsUrls.ThreadUsersVotedURL = switchHS(threadsUrls.ThreadUsersVotedURL, threadUsersVotedJSON)
	threadsUrls.ThreadSetURL = switchHS(threadsUrls.ThreadSetURL, threadSetJSON)
	threadsUrls.ThreadListURL = switchHS(threadsUrls.ThreadListURL, threadListJSON)
	threadsUrls.ThreadDetailURL = switchHS(threadsUrls.ThreadDetailURL, threadDetailsJSON)
	threadsUrls.ThreadPostsURL = switchHS(threadsUrls.ThreadPostsURL, threadPostsJSON)
	threadsUrls.ThreadHotURL = switchHS(threadsUrls.ThreadHotURL, threadListHotJSON)
	threadsUrls.ThreadPopularURL = switchHS(threadsUrls.ThreadPopularURL, threadListPopularJSON)
	threadsUrls.ThreadTrendingURL = switchHS(threadsUrls.ThreadTrendingURL, threadListTrendingJSON)
}

func TestThreadUsersVoted(t *testing.T) {

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

	threads, err := testGisqus.ThreadPopular(testCtx, testValues)
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
