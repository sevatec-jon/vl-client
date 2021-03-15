/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"context"
	"fmt"
	"github.com/spf13/viper"

	//"golang.org/x/oauth2"

	//"os"

	"github.com/spf13/cobra"
	"golang.org/x/oauth2/clientcredentials"
)


func NewAuthCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
	 	Short: "Gets an auth token from the API",
	 	Long:  `Used to request an oauth token from the api.`,
	 	Run: doAuth,
 	}

 	cmd.Flags().String("client","", "Client to get an auth token.")
 	cmd.Flags().String("secret","", "Secret to get an auth token.")

 	return cmd
}

func doAuth(cmd *cobra.Command, args []string) {
	url,_ := rootCmd.PersistentFlags().GetString("url")
	fmt.Fprintf(cmd.OutOrStdout(), "auth called for address: %s \n", url)

	client, _ := cmd.Flags().GetString("client")
	secret, _ := cmd.Flags().GetString("secret")

	config := clientcredentials.Config{
		ClientID:     client,
		ClientSecret: secret,
		TokenURL:     url + "auth/connect/token",
		Scopes:       []string{"vendorlink"},
	}

	token, err := config.Token(context.Background())

	if err != nil {
		fmt.Fprintf(cmd.OutOrStdout(),"%s", err)
	}

	viper.Set("token", token.AccessToken)
	err = viper.WriteConfig()
	if err != nil {
		fmt.Fprintf(cmd.OutOrStdout(),"%s", err)
	}
	fmt.Fprintf(cmd.OutOrStdout(), "%s", token.AccessToken)
}

func init() {
	rootCmd.AddCommand(NewAuthCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// authCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// authCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
