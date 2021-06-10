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

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

func NewClient(timeout time.Duration) *http.Client {
	return &http.Client{
		Timeout: timeout,
	}
}

func NewRequestWithAuthentication(url string) ([]byte, error) {
	username := viper.Get("gh_username")
	token := viper.Get("gh_token")
	httpClient := NewClient(5 * time.Second)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(username.(string), token.(string))
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	dst := &bytes.Buffer{}
	if err := json.Indent(dst, resBody, "", "  "); err != nil {
		return nil, err
	}
	return dst.Bytes(), nil

}
