// Copyright Piero de Salvia.
// All Rights Reserved

package gisqus

import (
	"fmt"
	"os"
	"testing"
)

func init() {

	postPopularJSON := readTestFile("postspostpopular.json")
	postDetailsJSON := readTestFile("postspostdetails.json")
	postListJSON := readTestFile("postspostlist.json")

	switchHS := func(URL, JSON string) string {
		result, err := mockServer.SwitchHostAndScheme(URL, JSON)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		return result
	}

	postsUrls.PostDetailsURL = switchHS(postsUrls.PostDetailsURL, postDetailsJSON)
	postsUrls.PostListURL = switchHS(postsUrls.PostListURL, postListJSON)
	postsUrls.PostPopularURL = switchHS(postsUrls.PostPopularURL, postPopularJSON)
}

func TestPostDetails(t *testing.T) {

	_, testErr = testGisqus.PostDetails(testCtx, "", testValues)
	if testErr == nil {
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

	posts, err := testGisqus.PostList(testCtx, testValues)
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

func TestPostPopular(t *testing.T) {

	posts, err := testGisqus.PostPopular(testCtx, testValues)
	if err != nil {
		t.Fatal("Should be able to call the post popular endpoint - ", err)
	}
	if len(posts.Response) != 25 {
		t.Fatal("Should be able to correctly parse a post list")
	}
	if posts.Response[0].ID != "3357275751" {
		t.Fatal("Should be able to retrieve a post id")
	}
	if posts.Response[0].Author.Username != "mychive-3683e7511cad5234db651099216183d0" {
		t.Fatal("Should be able to retrieve a post's author's username")
	}
	if posts.Response[0].Author.ID != "252976252" {
		t.Fatal("Should be able to retrieve a post's user's id")
	}
	if ToDisqusTime(posts.Response[0].CreatedAt) != "2017-06-13T11:40:12" {
		t.Fatal("Should be able to retrieve a post's created at")
	}
	if posts.Response[0].Thread != "5906159869" {
		t.Fatal("Should be able to retrieve a post's thread")
	}
	if posts.Response[0].Forum != "thechiverules" {
		t.Fatal("Should be able to retrieve a post's forum")
	}

}
