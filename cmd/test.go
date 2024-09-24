/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
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
	Run: runTestCmd,
}

func runTestCmd(cmd *cobra.Command, args []string) {
	fmt.Println("This is a test command for dev purposes.")
	fmt.Println("It shouldn't do anything.")

	// Read json from file
	data, err := os.ReadFile("./token.json")
	if err != nil {
		log.Fatalf("Error reading token file: %v\n", err)
	}

	// Unmarshal json
	var tok *oauth2.Token
	err = json.Unmarshal(data, &tok)
	if err != nil {
		log.Fatalf("Error unmarshaling data: %v\n", err)
	}

	// Try to create a client
	spotifyID := viper.GetString("spotify.id")
	spotifySecret := viper.GetString("spotify.secret")
	auth := spotifyauth.New(
		spotifyauth.WithRedirectURL(redirectURI),
		spotifyauth.WithScopes(spotifyauth.ScopePlaylistReadPrivate),
		spotifyauth.WithClientID(spotifyID),
		spotifyauth.WithClientSecret(spotifySecret),
	)
	client := spotify.New(auth.Client(context.Background(), tok))

	// Use the client to make calls that require authorization
	user, err := client.CurrentUser(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("You are logged in as:", user.ID)

	// Example: List the first page of my playlists
	pp, err := client.CurrentUsersPlaylists(context.Background())
	if err != nil {
		log.Fatalf("Error getting user's playlists: %v\n", err)
	}
	for _, playlist := range pp.Playlists {
		fmt.Printf("playlist: %v, id: %v\n", playlist.Name, playlist.ID)
	}
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
