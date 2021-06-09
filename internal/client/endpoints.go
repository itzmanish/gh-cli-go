/*
Copyright © 2021 Manish itzmanish108@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
	RateLimitURL               = API_BASE_URL + "/rate_limit"
	StarredGistsURL            = API_BASE_URL + "/gists/starred"

	RepositoryURL = func(user, repo string) string {
		return fmt.Sprintf("%s/repos/%s/%s", API_BASE_URL, user, repo)

	}

	UserRepositoriesURL = func(username, repoType string, per_page int) string {
		return fmt.Sprintf("/users/%s/repos?type=%s&per_page=%d", username, repoType, per_page)

	}

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
