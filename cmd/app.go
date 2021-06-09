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
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/itzmanish/gh-cli-go/internal"
	"github.com/itzmanish/gh-cli-go/internal/client"
	"github.com/itzmanish/gh-cli-go/utils"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

type Response struct {
	Name     string
	Response []byte
	Error    error
}

type Repository struct {
	Id       int    `json:"id"`
	NodeId   string `json:"node_id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	HtmlURL  string `json:"html_url"`
	Language string `json:"language"`
	Stars    int    `json:"stargazers_count"`
}

type Repositories struct {
	Items []Repository `json:"items"`
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
		return cli.ShowAppHelp(c)
	},
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

	Commands: []*cli.Command{
		{
			Name:  "init",
			Usage: "Initialize user with github username and token",
			Action: func(c *cli.Context) error {
				return initApp()
			},
		},
		{
			Name:  "download",
			Usage: "Download all information available for initialized user",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "out_dir",
					Aliases: []string{"o"},
					Usage:   "output file path",
				},
			},
			Action: func(c *cli.Context) error {
				return GetAllInfos(c, true)
			},
		},
		{
			Name:  "show",
			Usage: "show all information available for initialized user",

			Action: func(c *cli.Context) error {
				return GetAllInfos(c, false)
			},
		},
		{
			Name:  "play",
			Usage: "play guess the star game",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "language",
					Aliases: []string{"l"},
					Usage:   "language of your choice for repositories",
				},
			},
			Action: func(c *cli.Context) error {
				return GuessTheStar(c)
			},
		},
	},
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
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

func GetAllInfos(c *cli.Context, download bool) error {
	resCh := make(chan Response, len(client.Urls))
	for name, url := range client.Urls {
		go func(c chan Response, name string, url string) {
			r := Response{
				Name: name,
			}
			res, err := client.NewRequestWithAuthentication(url)
			if err != nil {
				r.Error = err
				return
			}
			r.Response = res
			c <- r
		}(resCh, name, url)

	}
	for i := 0; i < len(client.Urls); i++ {
		select {
		case res := <-resCh:
			if res.Error != nil {
				log.Println(res.Error)
				continue
			}
			err := utils.Output(res.Response, download, res.Name, c.String("out_dir"))
			if err != nil {
				log.Println(err)
			}
		}
	}
	return nil
}

func GuessTheStar(c *cli.Context) error {
	fmt.Println(`
	Hi, Welcome to the guess the star game.
	Before starting Here is the guide to play the game -
	* You are presented with a repository name and url.
	* You have to guess the total number of star that repository gathered.
	* If your guess is correct with 10% of tolerance you get a point.
	* There are total of 5 rounds and every time you get a random and unique repository.
	* If you want you can specify your choice of language so that only those
	  repository will presented to you which uses the language you specified.
	============================================================
	`)
	resp, err := utils.PromptText("If you are agreed enter YES/yes/Y/y, anything else will be considered as NO", false)
	if err != nil {
		return err
	}
	if resp == "YES" || resp == "yes" || resp == "Y" || resp == "y" {
		return StartGame(c.String("language"))
	}
	fmt.Println("You have quit the game")
	return nil
}

func StartGame(lang string) error {
	score := 0
	fmt.Printf("\nLoading repositories")
	req, err := http.NewRequest("GET", client.TrendingRepoURL(lang), nil)
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	repositories := Repositories{}
	err = json.NewDecoder(res.Body).Decode(&repositories)
	if err != nil {
		return err
	}
	for round := 0; round < 5; round++ {
		fmt.Println("Current Score: ", score)
		fmt.Printf("Selected language: %s | Round %d\n", lang, round)
		repo := repositories.Items[rand.Intn(100-0+1)+0]
		fmt.Println("Repository Name: ", repo.Name)
		fmt.Println("Repository Full Name: ", repo.FullName)
		fmt.Println("Repository url: ", repo.HtmlURL)
		fmt.Println("Repository language: ", repo.Language)
		gvalue, err := utils.PromptText("Guess the star", true)
		if err != nil {
			return err
		}
		v, err := strconv.Atoi(gvalue)
		if err != nil {
			return err
		}
		if CheckIfCorrect(v, repo.Stars) {
			score = score + 1
		} else {
			fmt.Printf("\nWrong answer\n")
		}
	}
	fmt.Println("Your score: ", score)
	if score < 4 {
		fmt.Println("Better luck next time.")

	} else {
		fmt.Println("You win")
	}
	return nil
}

func CheckIfCorrect(v int, star int) bool {
	starH := float64(star) + float64(star)*0.01
	starL := float64(star) - float64(star)*0.01
	if float64(v) >= starL && float64(v) <= starH {
		return true
	}
	return false
}
