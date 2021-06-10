package cmd

import (
	"flag"
	"os"
	"testing"

	"github.com/itzmanish/gh-cli-go/internal"
	"github.com/urfave/cli/v2"
)

func TestCheckIfCorrect(t *testing.T) {
	inp1 := 500
	inp2 := 420
	if checkIfCorrect(inp1, inp2) {
		t.Error("expected false but got true")
	}
	inp3 := 500
	inp4 := 510
	if !checkIfCorrect(inp3, inp4) {
		t.Errorf("%d is matched with %d with 10%% tolerence but got false", inp3, inp4)
	}
}

func TestInitApp(t *testing.T) {
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	content := []byte("manish\ntoken")
	tmpfile, err := os.CreateTemp("", "")
	if err != nil {
		t.Error(err)
		return
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		t.Error(err)
		return
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		t.Error(err)
		return
	}
	os.Stdin = tmpfile
	err = initApp()
	if err != nil {
		t.Errorf("Error not expected here but got: %v", err)
	}
}

func TestRun(t *testing.T) {
	testcases := []struct {
		name        string
		flags       string
		errExpected bool
	}{
		{
			name:        "Show all information",
			flags:       "show",
			errExpected: false,
		},
		{
			name:        "Get help screen",
			flags:       "help",
			errExpected: false,
		},
		{
			name:        "Download to standard out directory",
			flags:       "download",
			errExpected: false,
		},
		{
			name:        "Download to custom directory",
			flags:       "download -o custom_out",
			errExpected: false,
		},
		{
			name:        "Wrong command",
			flags:       "wrong",
			errExpected: true,
		},
	}
	for _, test := range testcases {
		args := os.Args[0:1]            // name of the program.
		args = append(args, test.flags) // Append a flag
		t.Run(test.name, func(t *testing.T) {
			err := Run(args)
			if err != nil && !test.errExpected {
				t.Error(err)
			}
			if test.errExpected && err != nil {
				t.Error("Error expected but got no error")
			}
		})
	}
}

func TestGetAllInfos(t *testing.T) {
	testcases := []struct {
		name        string
		download    bool
		errExpected bool
	}{
		{
			name:        "Success without download",
			download:    false,
			errExpected: false,
		},
		{
			name:        "Success with download enabled",
			download:    true,
			errExpected: false,
		},
	}
	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			err := internal.LoadConfig("../samples")
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			ctx := cli.NewContext(app, &flag.FlagSet{}, nil)
			err = getAllInfos(ctx, test.download)
			if err != nil && !test.errExpected {
				t.Errorf("Expected no error but got: %v", err)
				return
			}
			if err == nil && test.errExpected {
				t.Error("Expect error but got no error")
			}

		})
	}
}

func TestStartGame(t *testing.T) {
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	content := []byte("14420\n2215\n25445\n46525\n5894")
	tmpfile, err := os.CreateTemp("", "")
	if err != nil {
		t.Error(err)
		return
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		t.Error(err)
		return
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		t.Error(err)
		return
	}
	os.Stdin = tmpfile

	err = startGame("go")
	if err != nil {
		t.Error(err)
	}
}
