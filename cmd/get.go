/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Send a GET request",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		fmt.Println("Sending GET request to", url)
		rStr := fmt.Sprintf("Response: \n%s", getRequest(url))
		fmt.Println(rStr)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	getCmd.Flags().String("url", "u", "The url to send the request to")
}


func getRequest(url string) []byte {
    request, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		log.Printf("Error forming request: %v", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "httpcli (github.com/reuben-james/httpcli)")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Error sending request: %v", err)
	}

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
	}

	return responseBytes
}