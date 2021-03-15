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
	"crypto/tls"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/sevatec-jon/vl-client/internal/config"
	"github.com/sevatec-jon/vl-client/internal/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
)

// classesCmd represents the classes command
var teacherCmd = &cobra.Command{
	Use:   "teachers",
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

		var fullUri = url + fmt.Sprintf("VendorLink/GBService/" + config.GBDistrictID + "/" + TeachersPath)
		if verbose {
			fmt.Printf("%s %s %s \n", config.OcmId, config.GBSchoolID, config.GBDistrictID)
			fmt.Printf("%s\n", fullUri)
		}

		req, _ := http.NewRequest("GET", fullUri, nil)
		req.Header.Add("Authorization", "Bearer "+config.Token)
		req.Header.Add("Ocp-Apim-Subscription-Key", config.OcmId)

		q := req.URL.Query()
		q.Add("schoolId", config.GBSchoolID)
		req.URL.RawQuery = q.Encode()

		if verbose {
			fmt.Printf("%s", req.URL)
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

		fmt.Printf("%s", string(body))

		var data models.Response

		err = json.Unmarshal(body, &data)

		if err != nil {
			fmt.Printf("%v", err)
		}

		var results []models.Teacher

		err = json.Unmarshal(*data.Result, &results)

		fmt.Printf("count: %d \n", data.Count)


		db, err := sql.Open("sqlserver", config.DBConn)
		if err != nil {
			fmt.Println(" Error open db:", err.Error())
		}
		defer db.Close()

		stmt, err := db.Prepare("INSERT INTO Teachers(schoolid,teacherid, firstname, lastname, RefId, DistrictStaffId, Active) VALUES(@a, @b, @c, @d, @e, @f, @g); select convert(bigint, SCOPE_IDENTITY());")
		if err != nil {
			log.Fatal(err)
		}

		for i, row := range results[1:] {

			fmt.Printf("row[%d]: %v \n", i, row)
			//_, err := stmt.Exec(row.TeacherID, row.LastName)
			row := stmt.QueryRowContext(
				context.Background(),
				sql.Named("a", config.SchoolId),
				sql.Named("b", row.TeacherID),
				sql.Named("c", row.FirstName),
				sql.Named("d", row.LastName),
				sql.Named("e", row.ExternalRefId),
				sql.Named("f", row.DistrictStaffID),
				sql.Named("g", row.Active))

			var newTeacherID int64
			err = row.Scan(&newTeacherID)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("inserted id: %v \n", newTeacherID)

		}

	},
}

func init() {
	getCmd.AddCommand(teacherCmd)


}
/*
 for _, row := range rows[1:] {
    _, err := stmt.Exec(row[0], row[1])
    if err != nil {
      log.Fatal(err)
    }
  }
 */