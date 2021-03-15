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
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/sevatec-jon/vl-client/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

// checkDbCmd represents the classes command
var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Check the database connection is valid",
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

		if err := viper.Unmarshal(&config); err != nil {
			fmt.Printf("Error unmarshalling config file, %s", err)
		}

		fmt.Printf("School: %s \n", config.SchoolId)


		fmt.Printf("Connecting to SQL: %v \n", config.DBConn)

		db, err := sql.Open("sqlserver", config.DBConn)
		if err != nil {
			fmt.Println(" Error open db:", err.Error())
		}
		var (
			schoolName string
		)
		rows, err := db.Query("select schoolName from School where schoolid = $1", config.SchoolId)
		if err != nil {
			log.Fatal(err)
		}
		for rows.Next() {
			err := rows.Scan(&schoolName)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Found: %s", schoolName)
		}

		defer db.Close()
	},
}

func init() {
	checkCmd.AddCommand(dbCmd)


}
