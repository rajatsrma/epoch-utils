/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert epoch to date time",
	Long:  `Helps coverting epoch to human readable date time. Default converts to IST date time`,
	Run: func(cmd *cobra.Command, args []string) {

		seconds, _ := cmd.Flags().GetInt64("epoch-seconds")
		millis, _ := cmd.Flags().GetInt64("epoch-millis")

		if millis > 0 {
			timeT := time.UnixMilli(millis)
			fmt.Printf("IST: %s\n", timeT)
			fmt.Printf("UTC: %s\n", timeT.UTC())
		} else if seconds > 0 {
			timeT := time.Unix(seconds, 0)
			fmt.Printf("IST: %s\n", timeT)
			fmt.Printf("UTC: %s\n", timeT.UTC())
		} else {
			fmt.Println("Please provide some input for conversion")
		}
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// convertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// convertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	convertCmd.Flags().Int64P("epoch-seconds", "s", 0, "specify epoch seconds for datetime")
	convertCmd.Flags().Int64P("epoch-millis", "m", 0, "specify epoch milli-seconds for datetime")

}
