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
	"fmt"
	"log"
	"os"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/itzmanish/gh-cli-go/internal/user"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var username string
var accessToken string

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "For user specific commands",
	Long:  "You must generated access token for your account with resource permission which you need to access and download.",
	Run: func(cmd *cobra.Command, args []string) {
		if value, err := cmd.Flags().GetBool("info"); err == nil && value {
			err := user.Execute(cmd, args)
			if err != nil {
				log.Println(err)
			}
			return
		}
		cmd.Help()
	},
}

var userInit = &cobra.Command{
	Use:   "init",
	Short: "Initialize user with username and access token",
	Run: func(cmd *cobra.Command, args []string) {
		username, err := promptUsername()
		if err != nil {
			log.Println(err)
		}
		password, err := promptPassword()
		if err != nil {
			log.Println(err)
		}
		f, err := os.Create(".gh-cli.json")
		if err != nil {
			log.Println(err)
		}
		defer f.Close()
		c := user.Config{
			Username: username,
			Token:    password,
		}
		err = json.NewEncoder(f).Encode(c)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("gh-cli successfully initialized.")
	},
}

func init() {
	rootCmd.AddCommand(userCmd)
	userCmd.AddCommand(userInit)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userCmd.PersistentFlags().StringP("username", "u", "", "github username")
	// userCmd.PersistentFlags().StringP("token", "t", "", "github access token")
	// userCmd.MarkPersistentFlagRequired("username")
	// userCmd.MarkPersistentFlagRequired("token")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	userCmd.Flags().BoolP("info", "i", false, "Show user info")
}

func promptUsername() (string, error) {
	usernameValidate :=
		func(input string) error {
			return validation.Validate(input,
				validation.Required, // not empty
			)
		}

	userNamePrompt := promptui.Prompt{
		Label:    "Username",
		Validate: usernameValidate,
	}

	return userNamePrompt.Run()
}

func promptPassword() (string, error) {
	passwordValidate :=
		func(input string) error {
			return validation.Validate(input,
				validation.Required, // not empty
			)
		}

	passwordPrompt := promptui.Prompt{
		Label:    "Password",
		Validate: passwordValidate,
	}
	return passwordPrompt.Run()
}
