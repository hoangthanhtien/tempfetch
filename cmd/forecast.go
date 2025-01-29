/*
Copyright © 2025 Hoang Thanh Tien <hoangthanhtien0604@gmail.com>
*/
package cmd

import (
	// "fmt"
	"github.com/spf13/cobra"
	"tempfetch/utils"
)

// forecastCmd represents the forecast command
var forecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "Weather forecast",
	Run: func(cmd *cobra.Command, args []string) {
		currentUserLoc, _ := utils.GetCurrentUserInfo()
		// fmt.Printf("Current user loc: %#v", currentUserLoc)
		// fmt.Println("Start getting forecast")

		forecastParams := utils.ForecastParam{
			Long:     currentUserLoc.Long,
			Lat:      currentUserLoc.Lat,
			Timezone: currentUserLoc.Timezone,
			Hourly:   "temperature_2m", // TODO: Make this more flexible
		}
		forecastRepsonse, _ := utils.GetForecast(forecastParams)
		// fmt.Printf("Forecast reponse: %#v", forecastRepsonse)
		// fmt.Println("Rendering ASCII chart")
		utils.RenderASCIIChart(forecastRepsonse)
	},
}

func init() {
	rootCmd.AddCommand(forecastCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// forecastCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// forecastCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
