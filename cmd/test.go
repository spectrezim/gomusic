/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A test command for dev purposes",
	Long: `This is a test command for development purposes.
If you're a user, it shouldn't do anything but tell you
that it doesn't do anything.
If you're a developer modifying this command, be careful
not to commit any changes to it.`,
	Run: exec,
}

func exec(cmd *cobra.Command, args []string) {
	fmt.Println("This is a test command for dev purposes.")
	fmt.Println("It shouldn't do anything.")

	if viper.IsSet("spotify.id") {
		spotifyID := viper.GetString("spotify.id")
		fmt.Printf("Spotify Client ID: %s\n", spotifyID)
	} else {
		fmt.Println("Spotify Client ID not found")
	}

	viper.Set("spotify.id", "aywfutna")
	if viper.IsSet("spotify.id") {
		spotifyID := viper.GetString("spotify.id")
		fmt.Printf("Spotify Client ID: %s\n", spotifyID)
	} else {
		fmt.Println("Spotify Client ID not found")
	}

	viper.WriteConfig()
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
