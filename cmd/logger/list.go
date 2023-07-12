package logger

import (
	"fmt"
	"net/http"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

var LoggerListCmd = &cobra.Command{
	Use:   "list",
	Short: "list current loggers and log levels",
	Long:  "Allows to obtain a list of the currently enabled loggers and their log levels.",
	Run: func(cmd *cobra.Command, args []string) {
		for _, host := range utilities.ConnectConfiguration.Hostnames {
			var loggerListURL = fmt.Sprintf("%s/admin/loggers", host)
			fmt.Printf("--- Loggers Info for Connect worker at %s ---", host)
			//fmt.Println("making a call to", loggerListURL) // control statement print
			response, err := utilities.DoCallByHost(http.MethodGet, loggerListURL, nil)
			if err != nil {
				fmt.Printf("The HTTP request failed with error %s\n", err)
			} else {
				utilities.PrintResponseJson(response)
			}
		}
	},
}
