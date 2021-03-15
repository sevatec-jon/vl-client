package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"
)
func Test_RootExecuteWithError(t *testing.T) {
	var expected = ""
	bOut := bytes.NewBufferString("")
	rootCmd.SetArgs([]string{"bad-command"})
	rootCmd.SetOut(bOut)
	rootCmd.Execute()
	out, err := ioutil.ReadAll(bOut)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != expected {
		t.Fatalf("expected \"%s\" got \"%s\"", expected, string(out))
	}
}

func Test_RootExecuteWithoutFlags(t *testing.T) {
	var expected = `webta-cli is a CLI wrapper for the WebTA GraphQL API.

Usage:
  webta-cli [command]

Available Commands:
  auth        Gets an auth token from the API
  help        Help about any command

Flags:
      --config string   config file (default is $HOME/.webta-cli.yaml)
  -h, --help            help for webta-cli
  -t, --toggle          Help message for toggle
      --url string      WebTA URL (default "https://local.url")

Use "webta-cli [command] --help" for more information about a command.
`

	bOut := bytes.NewBufferString("")
	rootCmd.SetArgs([]string{"","--help"})
	rootCmd.SetOut(bOut)
	Execute()
	out, err := ioutil.ReadAll(bOut)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != expected {
		t.Fatalf("expected \"%s\" got \"%s\"", expected, string(out))
	}
}