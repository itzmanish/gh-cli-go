package client

import "fmt"

var API_BASE_URL = "https://api.github.com"

var (
	CurrentUserURL             = API_BASE_URL + "/user"
	EmailsURL                  = API_BASE_URL + "/user/emails"
	FollowersURL               = API_BASE_URL + "/user/followers"
	FollowingURL               = API_BASE_URL + "/user/following"
	GistsURL                   = API_BASE_URL + "/gists"
	PublicGistsURL             = API_BASE_URL + "/gists/public"
	CurrentUserRepositoriesURL = API_BASE_URL + "/user/repos"
	UserOrganizationsURL       = API_BASE_URL + "/user/orgs"
	StarredGistsURL            = API_BASE_URL + "/gists/starred"

	TrendingRepoURL = func(lang string) string {
		return fmt.Sprintf("%s/search/repositories?q=language:%s&sort=stars&per_page=100&page=1&order=desc", API_BASE_URL, lang)
	}
)

var Urls = map[string]string{
	"user":          CurrentUserURL,
	"repositories":  CurrentUserRepositoriesURL,
	"emails":        EmailsURL,
	"followers":     FollowersURL,
	"following":     FollowingURL,
	"gists":         GistsURL,
	"public_gists":  PublicGistsURL,
	"organization":  UserOrganizationsURL,
	"starred_gists": StarredGistsURL,
}
