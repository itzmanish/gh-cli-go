package internal

import (
	"testing"

	"github.com/spf13/viper"
)

func TestLoadConfig(t *testing.T) {
	err := LoadConfig("../samples")
	if err != nil {
		t.Error(err)
	}
	u := viper.Get("gh_username")
	if u == nil {
		t.Error("Config not loaded from given file path")
		return
	}
	if u.(string) != "username" {
		t.Error("Config somehow loaded but with something wrong values")
	}
}

func TestSetConfig(t *testing.T) {
	data := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}
	err := SetConfig(data, ".test-gh-cli.yaml", "../samples")
	if err != nil {
		t.Error(err)
	}
}
