package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

var cfgFile string = ".gomusic.yaml"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gomusic",
	Short: "Displays basic usage information",
	Long: `gomusic is a sandbox-like project used to learn how to interface with various
APIs.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello!")

		// Get the config data
		conf := Conf{}
		str, err := yaml.Marshal(&conf)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		fmt.Println(string(str))
	},
}

type Conf struct {
	Spotify struct {
		ID     string `yaml:"id"`
		Secret string `yaml:"secret"`
		Code   string `yaml:"code"`
	} `yaml:"spotify"`
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
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gomusic.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in current directory with name ".gomusic" (without extension).
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".gomusic")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found
			// TODO Write a default
			fmt.Println("Config file not found!")
		} else {
			// Config file was found, but another error was produced
			log.Fatalf("error reading in config: %w", err)
		}
	}
	fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
}
