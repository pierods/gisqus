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

	postsUrls.postDetailsURL, err = mock.SwitchHostAndScheme(postsUrls.postDetailsURL, mockServer.URL)
	if err != nil {
		t.Fatal(err)
	}
	_, err = testGisqus.PostDetails(testCtx, "", testValues)
	if err == nil {
		t.Fatal("Should check for an empty post id")
	}
	details, err := testGisqus.PostDetails(testCtx, "3320987826", testValues)
	if err != nil {
		t.Fatal(err)
	}
	if details.Response.ID != "3320987826" {
		t.Fatal("Should be able to retrieve a post id")
	}
	if details.Response.Author.Username != "royalewithcrowne" {
		t.Fatal("Should be able to retrieve a post's author's username")
	}
	if details.Response.Author.ID != "209257634" {
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

	postsUrls.postListURL, err = mock.SwitchHostAndScheme(postsUrls.postListURL, mockServer.URL)
	if err != nil {
		t.Fatal(err)
	}

	values := url.Values{}

	posts, err := testGisqus.PostList(testCtx, values)
	if err != nil {
		t.Fatal("Should be able to call the post list endpoint - ", err)
	}
	if len(posts.Response) != 25 {
		t.Fatal("Should be able to correctly parse a post list")
	}
	if posts.Response[0].ID != "3324481803" {
		t.Fatal("Should be able to retrieve a post id")
	}
	if posts.Response[0].Author.Username != "bautista8190p" {
		t.Fatal("Should be able to retrieve a post's author's username")
	}
	if posts.Response[0].Author.ID != "242978772" {
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
