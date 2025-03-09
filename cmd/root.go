package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "compass",
	Short: "A brief description of your application",
	Long: `Welcome to Compass :The Ultimate Hackathon Scraper CLI`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// here we will be writing a sample command to echo a name 
//
// var echoName = &cobra.Command{
// 	Use: "echo [name]",
// 	Args: cobra.ExactArgs(1),
// 	Run: func(cmd *cobra.Command, args []string) {
// 		fmt.Println("helloooo",args[0])
// 	},
// }
//
var fetchHack = &cobra.Command{
	Use: "fetchHack",
	Short: "This command it to fetch the details of hackathons ",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fetching hackathon details......................")
		FetchHackathons()
	},
}


// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.web-scraper.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// rootCmd.AddCommand(echoName)
	rootCmd.AddCommand(fetchHack)
}


