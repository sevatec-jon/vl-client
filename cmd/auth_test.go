package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func Test_ExecuteAuthCommandWithoutFlags(t *testing.T) {
	var expected = "auth called for address: https://local.url"
	//cmd := NewAuthCmd()
	//cmd.SetArgs()
	rootCmd.SetArgs([]string{"auth"})
	bOut := bytes.NewBufferString("")
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

func Test_ExecuteAuthCommandWithUrlFlags(t *testing.T) {
	var expected = "auth called for address: https://test.url"
	//cmd := NewAuthCmd()
	//cmd.SetArgs()
	rootCmd.SetArgs([]string{"auth", "--url", "https://test.url"})
	bOut := bytes.NewBufferString("")
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
