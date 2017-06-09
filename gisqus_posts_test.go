package gisqus

import (
	"net/url"
	"testing"

	"github.com/pierods/gisqus/mock"
)

func TestPostDetails(t *testing.T) {
	mockServer = ms.NewServer()
	defer mockServer.Close()
	testValues = url.Values{}

	postsUrls.post_details_url, err = mock.SwitchHostAndScheme(postsUrls.post_details_url, mockServer.URL)
	if err != nil {
		t.Fatal(err)
	}
	_, err = testGisqus.PostDetails("", testValues, testCtx)
	if err == nil {
		t.Fatal("Should check for an empty post id")
	}
	details, err := testGisqus.PostDetails("3320987826", testValues, testCtx)
	if err != nil {
		t.Fatal(err)
	}
	if details.Response.Id != "3320987826" {
		t.Fatal("Should be able to retrieve a post id")
	}
	if details.Response.Author.Username != "royalewithcrowne" {
		t.Fatal("Should be able to retrieve a post's author's username")
	}
	if details.Response.Author.Id != "209257634" {
		t.Fatal("Should be able to retrieve a post's username's id")
	}
	if ToDisqusTime(details.Response.CreatedAt) != "2017-05-23T17:57:41" {
		t.Fatal("Should be able to retrieve a post's created at")
	}
	if details.Response.Parent != 3320975377 {
		t.Fatal("Should be able to retrieve a post's parent")
	}
	if details.Response.Thread != "5843656825" {
		t.Fatal("Should be able to retrieve a post's thread")
	}
	if details.Response.Forum != "wrestlinginc" {
		t.Fatal("Should be able to retrieve a post's forum")
	}
}

func TestPostList(t *testing.T) {

	mockServer = ms.NewServer()
	defer mockServer.Close()

	postsUrls.post_list_url, err = mock.SwitchHostAndScheme(postsUrls.post_list_url, mockServer.URL)
	if err != nil {
		t.Fatal(err)
	}

	values := url.Values{}

	posts, err := testGisqus.PostList(values, testCtx)
	if err != nil {
		t.Fatal("Should be able to call the post list endpoint - ", err)
	}
	if len(posts.Response) != 25 {
		t.Fatal("Should be able to correctly parse a post list")
	}
	if posts.Response[0].Id != "3324481803" {
		t.Fatal("Should be able to retrieve a post id")
	}
	if posts.Response[0].Author.Username != "bautista8190p" {
		t.Fatal("Should be able to retrieve a post's author's username")
	}
	if posts.Response[0].Author.Id != "242978772" {
		t.Fatal("Should be able to retrieve a post's user's id")
	}
	if ToDisqusTime(posts.Response[0].CreatedAt) != "2017-05-25T17:43:47" {
		t.Fatal("Should be able to retrieve a post's created at")
	}
	if posts.Response[1].Parent != 3324406982 {
		t.Fatal("Should be able to retrieve a post's parent")
	}
	if posts.Response[0].Thread != "4978714775" {
		t.Fatal("Should be able to retrieve a post's thread")
	}
	if posts.Response[0].Forum != "pregunta2" {
		t.Fatal("Should be able to retrieve a post's forum")
	}

}
