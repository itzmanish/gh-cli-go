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

var repoCmd = &cli.Command{
	Name:  "repo",
	Usage: "repository specific commands",
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
		return RepoExecute(c)
	},
}

func RepoExecute(c *cli.Context) error {
	user := viper.Get("gh_username")
	if c.Args().Get(0) != "" {
		res, err := client.NewRequest(client.RepositoryURL(user.(string), c.Args().First()))
		if err != nil {
			return err
		}
		fmt.Println(string(res))
		return nil
	}
	res, err := client.NewRequest(client.CurrentUserRepositoriesURL)
	if err != nil {
		return err
	}
	fmt.Println(string(res))

	return nil
}
