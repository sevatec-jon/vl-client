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
	"crypto/tls"
	"fmt"
	"github.com/sevatec-jon/vl-client/internal/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/sevatec-jon/vl-client/internal/config"
)

// classesCmd represents the classes command
var attendanceCodesCmd = &cobra.Command{
	Use:   "attendancecodes",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var config config.Configuration

		if err := viper.ReadInConfig(); err != nil {
			fmt.Printf("Error reading config file, %s", err)
		}

		err := viper.Unmarshal(&config)

		client := &http.Client{Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
		}}

		verbose,_ :=  rootCmd.PersistentFlags().GetBool("verbose")
		url,_ := rootCmd.PersistentFlags().GetString("url")

		//if verbose {
			fmt.Printf("Using token: %s", config.Token)
		//}
		req, _ := http.NewRequest("GET", url + AttendanceCodesPath, nil)
		req.Header.Add("Authorization", "Bearer "+config.Token)
		req.Header.Add("Ocp-Apim-Subscription-Key", config.OcmId)

		q := req.URL.Query()
		q.Add("leaOrSchoolInfoRefId", config.SchoolRefId)
		req.URL.RawQuery = q.Encode()

		if verbose {
			fmt.Printf("%s", req.RequestURI)
		}

		res, err := client.Do(req)

		if err != nil {
			fmt.Printf("%v", err)
		}

		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("%v", err)
		}

		fmt.Printf("data: %v", string(body))

		var data models.Response

		err = json.Unmarshal(body, &data)

		if err != nil {
			fmt.Printf("%v", err)
		}

		var results []models.AttendanceCode

		err = json.Unmarshal(*data.Result, &results)

		fmt.Printf("count: %d \n", data.Count)
		fmt.Printf("results: %v", results)

	},
}

func init() {
	getCmd.AddCommand(attendanceCodesCmd)


}
