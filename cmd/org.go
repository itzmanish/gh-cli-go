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

var orgCmd = &cli.Command{
	Name:  "org",
	Usage: "organisation specific commands",
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
		return OrgExecute(c)
	},
}

func OrgExecute(c *cli.Context) error {
	res, err := client.NewRequest(client.UserOrganizationsURL)
	if err != nil {
		return err
	}
	fmt.Println(string(res))
	return nil
}
