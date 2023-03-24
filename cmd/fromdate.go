/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

const (
	layoutD  = "2006-01-02"
	layoutDT = "2006-01-02 15:04:05"
)

// fromdateCmd represents the fromdate command
var fromdateCmd = &cobra.Command{
	Use:   "fromdate",
	Short: "Convert Datetime to epoch seconds",
	Long: `Given: 
				a date time in yyyy-MM-dd hh:mm:ss or yyyy-MM-dd,
				IST zone or not
	 		Shows: 
				epoch seconds in IST or UTC as requested`,
	Run: func(cmd *cobra.Command, args []string) {

		datetime, _ := cmd.Flags().GetString("date-time")
		timezone, _ := cmd.Flags().GetString("timezone")
		datetimeLen := len([]rune(datetime))

		if datetimeLen == 10 {
			startandEODFromDateStr(datetime, timezone)
		} else if datetimeLen == 19 {
			timeUTC, err := time.Parse(layoutDT, datetime)
			if err != nil {
				fmt.Printf("error")
				return
			}

			// current epoch
			if timezone == "IST" {
				fmt.Printf("IST current: %d\n", utcToIST(timeUTC).UnixMilli())
			} else if timezone == "UTC" {
				fmt.Printf("UTC current: %d\n", timeUTC.UnixMilli())
			} else {
				fmt.Printf("timezone %s not supported\n", timezone)
				return
			}

			// start and eod from date part in requested timezone
			date := datetime[0:10]
			startandEODFromDateStr(date, timezone)
		} else {
			fmt.Println("Please provide some input for conversion")
		}
	},
}

func startandEODFromDateStr(d string, timezone string) {
	timeUTC, err := time.Parse(layoutD, d)
	if err != nil {
		fmt.Printf("please provide datetime in yyyy-MM-dd hh:mm:ss or yyyy-MM-dd formats")
		panic(err)
	}
	var istStart = utcToIST(timeUTC).UnixMilli()
	var istEnd = utcToIST(timeUTC).Add(time.Duration(time.Hour * 24)).UnixMilli()
	var utcStart = timeUTC.UnixMilli()
	var utcEnd = timeUTC.Add(time.Duration(time.Hour * 24)).UnixMilli()

	// IST start and end of the day
	if timezone == "IST" {
		fmt.Printf("IST start of the day: %d\n", istStart)
		fmt.Printf("IST end of the day: %d\n", istEnd)
	} else if timezone == "UTC" {
		// UTC start and end of the day
		fmt.Printf("UTC start of the day: %d\n", utcStart)
		fmt.Printf("UTC end of the day: %d\n", utcEnd)
	} else {
		fmt.Printf("timezone %s not supported\n", timezone)
	}
}

func utcToIST(t time.Time) time.Time {
	return t.Add(time.Duration(time.Minute * (-330)))
}

func init() {
	rootCmd.AddCommand(fromdateCmd)
	fromdateCmd.Flags().StringP("date-time", "d", "2006-01-02", "specify datetime to get epoch seconds")
	fromdateCmd.Flags().StringP("timezone", "t", "IST", "specify timezone (e.g. IST, UTC)")
}
