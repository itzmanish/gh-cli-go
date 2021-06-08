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
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/itzmanish/gh-cli-go/internal"
	"github.com/itzmanish/gh-cli-go/internal/client"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

// userCmd represents the user command
var userCmd = &cli.Command{
	Name:  "user",
	Usage: "user specific commands",
	Flags: defaultFlags,
	Before: func(c *cli.Context) error {
		base_dir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		internal.LoadConfig(base_dir + "/.config")
		username := viper.Get("gh_username")
		token := viper.Get("gh_token")
		if username == nil || token == nil {
			return errors.New("Please initialize cli with gh-cli-go init")
		}
		return nil
	},
	Action: func(c *cli.Context) error {
		return UserExecute(c)
	},
}

func UserExecute(c *cli.Context) error {
	switch c.Args().Get(0) {
	case "followers":
		res, err := client.NewRequest(client.FollowersURL)
		if err != nil {
			return err
		}
		fmt.Println(string(res))
	case "following":
		res, err := client.NewRequest(client.FollowingURL)
		if err != nil {
			return err
		}
		fmt.Println(string(res))
	case "":
		res, err := client.NewRequest(client.CurrentUserURL)
		if err != nil {
			return err
		}
		fmt.Println(string(res))

	default:
		cli.ShowAppHelp(c)
	}

	return nil
}
