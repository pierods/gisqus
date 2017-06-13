# gisqus - a thin wrapper over Disqus' API, written in Go.

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![](https://godoc.org/github.com/pierods/gisqus?status.svg)](http://godoc.org/github.com/pierods/gisqus)
[![Go Report Card](https://goreportcard.com/badge/github.com/pierods/gisqus)](https://goreportcard.com/report/github.com/pierods/gisqus)
[![Build Status](https://travis-ci.org/pierods/gisqus.svg?branch=master)](https://travis-ci.org/pierods/gisqus)

Gisqus is a Go wrapper over Disqus' public API (https://disqus.com/api/docs/). Its main purposes are to wrap away REST calls, http error handling and modeling
of the data returned.

Gisqus only covers endpoints that read data (GET method), not the ones writing data. It is mainly meant for reporting purposes.
For this reason: 
* it only supports authentication in the form of "Authenticating as the Account Owner" (https://disqus.com/api/docs/auth/)
* endpoints that require entity IDs (thread ID, forum ID etc) but where they can be provided implicitly by authentication have their wrappers 
  requiring those parameters explicitly in the method signature

The "related" parameter in many Disqus endpoints is not supported, since data returned through it can always be gotten with a direct call to the 
respective api. In this sense, Gisqus covers the complete hierarchy of Disqus' object model.

### Endpoints covered
https://disqus.com/api/docs/
##### Forums
* details
* interestingForums
* listCategories
 listFollowers 
 listMostActiveUsers
 listMostLikedUsers
* listThreads
* listUsers

##### Threads
* details
* list
* listHot Beta
* listPopular Beta
* listPosts
 listUsersVotedThread
 set

##### Posts
* details
* getContext 
* list
 listPopular

##### Users

* details 
* interestingUsers
* listActiveForums
 listActivity 
* listFollowers 
* listFollowing 
 listFollowingChannels 
* listFollowingForums 
 listMostActiveForums 
* listPosts


###Usage