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

package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/itzmanish/gh-cli-go/internal/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Username string `json:"username"`
	Token    string `json:"password"`
}

func Execute(cmd *cobra.Command, args []string) error {
	username := viper.Get("username")
	token := viper.Get("password")
	c := client.NewClient(5 * time.Second)
	req, err := http.NewRequest("GET", client.CurrentUserURL, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(username.(string), token.(string))
	res, err := c.Do(req)
	if err != nil {
		return err
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	dst := &bytes.Buffer{}
	if err := json.Indent(dst, resBody, "", "  "); err != nil {
		panic(err)
	}
	fmt.Println(dst.String())
	return nil
}
