package cmd

import (
	"log"
  "fmt"
  "os"
  "github.com/spf13/cobra"

  homedir "github.com/mitchellh/go-homedir"
  "github.com/spf13/viper"

)

var(
  cfgFile string
  dataFile string
)


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
  Use:   "todo-cli",
  Short: "todo-cli is a ToDo application",
  Long: `ToDo will help you get more done in less time :)`,
}


func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
  cobra.OnInitialize(initConfig)

  rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.todo-cli.yaml)")

  rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

  home, err := homedir.Dir()
  if err != nil {
    log.Println("Unable to detect Home directory, please set data using --datafile option")
  }

  rootCmd.PersistentFlags().StringVar(&dataFile, "datafile", home + string(os.PathSeparator) + ".todos.json", "datafile to store ToDos") 

}


// initConfig reads in config file and ENV variables if set.
func initConfig() {
  if cfgFile != "" {
    // Use config file from the flag.
    viper.SetConfigFile(cfgFile)
  } else {
    // Find home directory.
    home, err := homedir.Dir()
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    // Search config in home directory with name ".todo-cli" (without extension).
    viper.AddConfigPath(home)
    viper.SetConfigName(".todo-cli")
  }

  viper.AutomaticEnv() // read in environment variables that match

  // If a config file is found, read it in.
  if err := viper.ReadInConfig(); err == nil {
    fmt.Println("Using config file:", viper.ConfigFileUsed())
  }
}

