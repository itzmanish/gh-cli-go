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
	CurrentUserURL    = API_BASE_URL + "/user"
	AuthorizationsURL = func() string {
		return fmt.Sprintf(API_BASE_URL + "/authorizations")

	}
	EmailsURL = func() string {
		return fmt.Sprintf("/user/emails")

	}
	FollowersURL = func() string {
		return fmt.Sprintf("/user/followers")

	}
	FollowingURL = func() string {
		return fmt.Sprintf("/user/following{/target}")

	}
	GistsURL = func() string {
		return fmt.Sprintf("/gists{/gist_id}")

	}
	PublicGistsURL = func() string {
		return fmt.Sprintf("/gists/public")

	}
	OrganizationURL = func() string {
		return fmt.Sprintf("/orgs/{org}")

	}
	OrganizationRepositoriesURL = func() string {
		return fmt.Sprintf("/orgs/{org}/repos{?type,page,per_page,sort}")

	}
	OrganizationTeamsURL = func() string {
		return fmt.Sprintf("/orgs/{org}/teams")

	}
	RateLimitURL = func() string {
		return fmt.Sprintf("/rate_limit")

	}
	RepositoryURL = func() string {
		return fmt.Sprintf("/repos/{owner}/{repo}")

	}

	CurrentUserRepositoriesURL = func() string {
		return fmt.Sprintf("/user/repos{?type,page,per_page,sort")

	}
	StarredURL = func() string {
		return fmt.Sprintf("/user/starred{/owner}{/repo")

	}
	StarredGistsURL = func() string {
		return fmt.Sprintf("/gists/starred")

	}
	UserURL = func() string {
		return fmt.Sprintf("/users/{user}")

	}
	UserOrganizationsURL = func() string {
		return fmt.Sprintf("/user/orgs")

	}
	UserRepositoriesURL = func() string {
		return fmt.Sprintf("/users/{user}/repos{?type,page,per_page,sort}")

	}
)
