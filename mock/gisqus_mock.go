package mock

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
)

var (
	usersListPostsJSON        string
	usersUserDetailsJSON      string
	usersInterestingUsersJSON string
	usersActiveForumsJSON     string
	usersFollowersJSON        string
	usersFollowingJSON        string
	usersForumFollowingJSON   string
)

var (
	postDetailsJSON string
	postListJSON    string
)

var (
	threadListJSON         string
	threadDetailsJSON      string
	threadPostsJSON        string
	threadListHotJSON      string
	threadListPopularJSON  string
	threadListTrendingJSON string
)

var (
	forumInterestingForumsJSON string
	forumListJSON              string
	forumDetailsJSON           string
	forumListCategoriesJSON    string
	forumThreadListJSON        string
	forumMostLikedJSON         string
)

type MockServer struct {
	baseDir string
}

func NewMockServer(baseDir string) MockServer {
	return MockServer{
		baseDir,
	}
}
func (ms *MockServer) initThreads() {

	var err error
	threadListJSON, err = ms.readFile("threadsthreadlist.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	threadDetailsJSON, err = ms.readFile("threadsthreaddetails.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	threadPostsJSON, err = ms.readFile("threadsthreadposts.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	threadListHotJSON, err = ms.readFile("threadshotlist.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	threadListPopularJSON, err = ms.readFile("threadspopular.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	threadListTrendingJSON, err = ms.readFile("threadstrending.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func (ms *MockServer) initForums() {

	var err error
	forumInterestingForumsJSON, err = ms.readFile("forumsinterestingforums.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	forumListJSON, err = ms.readFile("forumslistforumusers.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	forumDetailsJSON, err = ms.readFile("forumsforumdetails.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	forumListCategoriesJSON, err = ms.readFile("forumslistcategories.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	forumThreadListJSON, err = ms.readFile("forumsforumlistthreads.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	forumMostLikedJSON, err = ms.readFile("forumsmostlikedusers.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func (ms *MockServer) initPosts() {

	var err error
	postDetailsJSON, err = ms.readFile("postspostdetails.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	postListJSON, err = ms.readFile("postspostlist.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func (ms *MockServer) initUsers() {

	var err error
	usersListPostsJSON, err = ms.readFile("userslistposts.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	usersUserDetailsJSON, err = ms.readFile("usersuserdetail.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	usersInterestingUsersJSON, err = ms.readFile("usersinterestingusers.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	usersActiveForumsJSON, err = ms.readFile("usersactiveforums.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	usersFollowersJSON, err = ms.readFile("usersfollowers.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	usersFollowingJSON, err = ms.readFile("usersfollowing.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	usersForumFollowingJSON, err = ms.readFile("usersfollowingforums.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func (ms *MockServer) readFile(fileName string) (string, error) {

	f, err := os.Open(ms.baseDir + fileName)
	defer f.Close()

	if err != nil {
		return "", err
	}
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(bytes), nil

}

func (m *MockServer) NewServer() *httptest.Server {

	m.initForums()
	m.initPosts()
	m.initThreads()
	m.initUsers()

	f := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Ratelimit-Remaining", "999")
		w.Header().Set("X-Ratelimit-Limit", "1000")
		w.Header().Set("X-Ratelimit-Reset", "1495785600")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)

		path := r.URL.Path

		switch path {
		case "/api/3.0/forums/interestingForums":
			fmt.Fprint(w, forumInterestingForumsJSON)
		case "/api/3.0/forums/listUsers.json":
			fmt.Fprint(w, forumListJSON)
		case "/api/3.0/forums/details.json":
			fmt.Fprint(w, forumDetailsJSON)
		case "/api/3.0/forums/listCategories.json":
			fmt.Fprint(w, forumListCategoriesJSON)
		case "/api/3.0/users/listPosts.json":
			fmt.Fprint(w, usersListPostsJSON)
		case "/api/3.0/users/details.json":
			fmt.Fprint(w, usersUserDetailsJSON)
		case "/api/3.0/users/interestingUsers.json":
			fmt.Fprint(w, usersInterestingUsersJSON)
		case "/api/3.0/posts/details.json":
			fmt.Fprint(w, postDetailsJSON)
		case "/api/3.0/posts/list.json":
			fmt.Fprint(w, postListJSON)
		case "/api/3.0/threads/list.json":
			fmt.Fprint(w, threadListJSON)
		case "/api/3.0/threads/details.json":
			fmt.Fprint(w, threadDetailsJSON)
		case "/api/3.0/forums/listThreads.json":
			fmt.Fprint(w, forumThreadListJSON)
		case "/api/3.0/forums/listMostLikedUsers.json":
			fmt.Fprint(w, forumMostLikedJSON)
		case "/api/3.0/threads/listPosts.json":
			fmt.Fprint(w, threadPostsJSON)
		case "/api/3.0/threads/listHot.json":
			fmt.Fprint(w, threadListHotJSON)
		case "/api/3.0/threads/listPopular.json":
			fmt.Fprint(w, threadListPopularJSON)
		case "/api/3.0/users/listActiveForums.json":
			fmt.Fprint(w, usersActiveForumsJSON)
		case "/api/3.0/users/listFollowers.json":
			fmt.Fprint(w, usersFollowersJSON)
		case "/api/3.0/users/listFollowing.json":
			fmt.Fprint(w, usersFollowingJSON)
		case "/api/3.0/users/listFollowingForums.json":
			fmt.Fprint(w, usersForumFollowingJSON)
		case "/api/3.0/trends/listThreads.json":
			fmt.Fprint(w, threadListTrendingJSON)
		}
	}
	return httptest.NewServer(http.HandlerFunc(f))
}

func SwitchHostAndScheme(source, newValues string) (string, error) {
	newValuesS, err := url.Parse(newValues)
	if err != nil {
		return "", err
	}

	sourceS, err := url.Parse(source)
	if err != nil {
		return "", err
	}
	sourceS.Host = newValuesS.Host
	sourceS.Scheme = newValuesS.Scheme

	return sourceS.String(), nil
}
