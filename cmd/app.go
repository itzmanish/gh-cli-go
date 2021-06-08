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
	"fmt"
	"os"
	"time"

	"github.com/itzmanish/gh-cli-go/utils"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

var cfgFile string

var defaultFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:    "download",
		Aliases: []string{"d"},
		Usage:   "Retrieve and save all the information to local filesystem",
	},
}

// rootCmd represents the base command when called without any subcommands
var app = &cli.App{
	Name:     "gh-cli-go",
	Usage:    "A Github CLI which provides downloading users data from github.",
	Version:  "v1.0.0",
	Compiled: time.Now(),
	Authors: []*cli.Author{
		{
			Name:  "Manish",
			Email: "itzmanish108@gmail.com",
		},
	},
	Copyright: "(c) 2021 Manish",
	HelpName:  "gh-cli-go",
	ArgsUsage: "[args and such]",
	Action: func(c *cli.Context) error {
		fmt.Println("boom! I say!")
		return nil
	},
	Flags: []cli.Flag{
		&cli.PathFlag{
			Name:  "config_path",
			Usage: "use custom config file",
		},
	},
	Commands: []*cli.Command{
		{
			Name:  "init",
			Usage: "Initialize user with github username and token",
			Action: func(c *cli.Context) error {
				return initApp()
			},
		},
		userCmd,
		repoCmd,
		orgCmd,
	},
}

func Run(args []string) (err error) {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Shows current cli version",
	}

	app.EnableBashCompletion = true
	err = app.Run(args)
	return
}

func initApp() error {
	username, err := utils.PromptText("username", true)
	if err != nil {
		return err
	}
	password, err := utils.PromptTextMasked("Token/Password", true)
	if err != nil {
		return err
	}
	viper.Set("gh_username", username)
	viper.Set("gh_token", password)
	base_dir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	err = viper.WriteConfigAs(base_dir + "/.config/.gh-cli.yaml")
	if err != nil {
		return err
	}
	fmt.Println("gh-cli successfully initialized.")
	return nil
}
