package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

// Must be an exact match of one of the redirect URIs registered for the app.
// These can be found (and changed, if needed) on the Spotify for Developers
// Web Dashboard.
const redirectURI = "http://localhost:8080/callback"

var (
	auth  *spotifyauth.Authenticator
	ch    = make(chan *spotify.Client)
	state = "awio43n10348"
)

// TODO: Update long description
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with Spotify",
	Long:  `Authenticate with Spotify so that this app can access your Spotify data`,
	Run:   runAuthCmd,
}

func runAuthCmd(cmd *cobra.Command, args []string) {
	// Get the Spotify client ID & secret
	if !(viper.IsSet("spotify.id") && viper.IsSet("spotify.secret")) {
		fmt.Println("`spotify.id` and `spotify.secret` must be set in the config file")
		return
	}
	spotifyID := viper.GetString("spotify.id")
	spotifySecret := viper.GetString("spotify.secret")

	// Initialize auth
	auth = spotifyauth.New(
		spotifyauth.WithRedirectURL(redirectURI),
		spotifyauth.WithScopes(spotifyauth.ScopePlaylistReadPrivate),
		spotifyauth.WithClientID(spotifyID),
		spotifyauth.WithClientSecret(spotifySecret),
	)

	// Start an HTTP server to receive the callback
	http.HandleFunc("/callback", completeAuth)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	url := auth.AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	// Wait for auth to complete
	client := <-ch

	// Use the client to make calls that require authorization
	user, err := client.CurrentUser(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("You are logged in as:", user.ID)
}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	// Check state
	st := r.FormValue("state")
	if st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}

	// Exchange code for a token
	tok, err := auth.Token(r.Context(), state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}

	// Use the token to get an authenticated client
	client := spotify.New(auth.Client(r.Context(), tok))
	fmt.Fprintf(w, "Login completed!")
	ch <- client
}

func init() {
	rootCmd.AddCommand(authCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// authCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// authCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
