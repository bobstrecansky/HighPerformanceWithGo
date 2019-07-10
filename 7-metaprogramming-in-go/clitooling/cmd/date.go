package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var verbose bool
var DateCommand = &cobra.Command{
	Use:     "date",
	Aliases: []string{"time"},
	Short:   "Return the current date",
	Long:    "Returns the current date in a YYYY-MM-DD HH:MM:SS format",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Current Date :\t", time.Now().Format("2006.01.02 15:04:05"))
		if viper.GetBool("verbose") {
			fmt.Println("Author :\t", viper.GetString("author"))
			fmt.Println("Version :\t", viper.GetString("version"))
		}
	},
}

var LicenseCommand = &cobra.Command{
	Use:   "license",
	Short: "Print the License",
	Long:  "Print the License of this Command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("License: Apache-2.0")
	},
}

func init() {
	DateCommand.AddCommand(LicenseCommand)
	viper.SetDefault("Author", "bob")
	viper.SetDefault("Version", "0.0.1")
	viper.SetDefault("license", "Apache-2.0")
	DateCommand.PersistentFlags().BoolP("verbose", "v", false, "Date Command Verbose")
	DateCommand.PersistentFlags().StringP("author", "a", "bob", "Date Command Author")

	viper.BindPFlag("author", DateCommand.PersistentFlags().Lookup("author"))
	viper.BindPFlag("verbose", DateCommand.PersistentFlags().Lookup("verbose"))

}
