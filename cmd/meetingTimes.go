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
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/sevatec-jon/vl-client/internal/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var meetingTimesCmd = &cobra.Command{
	Use:   "meetingTimes",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		if err := viper.ReadInConfig(); err != nil {
			fmt.Printf("Error reading config file, %s", err)
		}

		err := viper.Unmarshal(&Config)

		db, err := sql.Open("sqlserver", Config.DBConn)
		if err != nil {
			fmt.Println(" Error open db:", err.Error())
		}
		//defer db.Close()
		startRec,_ := cmd.Flags().GetInt("startRec")

		getClasses(db, startRec, processMeetingTimes)
	},
}

func processMeetingTimes(classId int) {

	db, err := sql.Open("sqlserver", Config.DBConn)

	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	}}

	verbose,_ :=  rootCmd.PersistentFlags().GetBool("verbose")

	url,_ := rootCmd.PersistentFlags().GetString("url")

	var fullUri = url + fmt.Sprintf("VendorLink/GBService/" + Config.GBDistrictID + "/" + SectionEnrollmentPath)
	if verbose {
		fmt.Printf("%s %s %s \n", Config.OcmId, Config.GBSchoolID, Config.GBDistrictID)
		fmt.Printf("%s \n", fullUri)
	}

	req, _ := http.NewRequest("GET", fullUri, nil)
	req.Header.Add("Authorization", "Bearer "+Config.Token)
	req.Header.Add("Ocp-Apim-Subscription-Key", Config.OcmId)

	q := req.URL.Query()
	q.Add("classId", strconv.Itoa(classId))
	req.URL.RawQuery = q.Encode()

	if verbose {
		fmt.Printf("%s \n", req.URL)
	}

	res, err := client.Do(req)

	if res.StatusCode == 401 || res.StatusCode == 403 {
		//fmt.Printf("Please generate a new auth token. %v", err)
		log.Fatalf("Please generate a new auth token.")
	}

	if res.StatusCode == 404 {
		return
	}

	if err != nil {
		fmt.Printf("%v", err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("%v", err)
	}

	//fmt.Printf("data: %v", string(body))

	var data models.Response

	err = json.Unmarshal(body, &data)

	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Printf("count: %d \n", data.Count)

	var results []models.Class

	err = json.Unmarshal(*data.Result, &results)

	//fmt.Printf("results: %v", results)
	row := results[0]

	stmt, err := db.Prepare(`INSERT INTO MeetingTerm(ClassID,	SchoolID,	Term, Room)
										VALUES(@a, @b, @c, @d); select convert(bigint, SCOPE_IDENTITY());`)
	if err != nil {
		log.Fatal(err)
	}

	for _, m := range row.ReportingPeriodMeetingTimes[0:] {
		fmt.Printf("period meet time: %v \n", m)

		for _, p := range m.ReportingPeriods[0:] {
			row := stmt.QueryRowContext(
				context.Background(),
				sql.Named("a", row.ClassID),
				sql.Named("b", Config.SchoolId),
				sql.Named("c", p.Name),
				sql.Named("d", p.Room))

			var rowID int64
			err = row.Scan(&rowID)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("inserted id: %v \n", rowID)
		}

	}
}
func init() {

	meetingTimesCmd.Flags().StringP("class", "c", "", "class id to get sections enrollments for")

	meetingTimesCmd.Flags().Int("startRec", 0, "starting class id to get sections enrollments")


	getCmd.AddCommand(meetingTimesCmd)

}