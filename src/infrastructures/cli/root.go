package cli

import (
	"github.com/fatih/color"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile     string
	userLicense string

	RootCmd = &cobra.Command{
		Use:   "hexagonal",
		Short: "A hexagonal architect base Applications",
		Long:  `Hexagonal is a hexagonal architect library for Go that empowers applications. This application is a tool to generate the hexagonal architect structure`,
	}
)

func init() {
	cli := CLI{}
	cli.banner()
	cobra.OnInitialize(cli.initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	RootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	RootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	RootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	_ = viper.BindPFlag("author", RootCmd.PersistentFlags().Lookup("author"))
	_ = viper.BindPFlag("useViper", RootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	viper.SetDefault("license", "apache")
}

func (c CLI) Execute(addCommand func(*cobra.Command)) error {
	addCommand(RootCmd)
	return RootCmd.Execute()
}

func (c CLI) err(msg interface{}) {
	color.HiRed("Error: %s", msg)
	os.Exit(1)
}

func (c CLI) initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			c.err(err)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		color.HiRed("Using config file:" + viper.ConfigFileUsed())
	}
}

func (c CLI) banner() {
	color.Cyan("\n\n██╗  ██╗███████╗██╗  ██╗ █████╗  ██████╗  ██████╗ ███╗   ██╗ █████╗ ██╗     \n██║  ██║██╔════╝╚██╗██╔╝██╔══██╗██╔════╝ ██╔═══██╗████╗  ██║██╔══██╗██║     \n███████║█████╗   ╚███╔╝ ███████║██║  ███╗██║   ██║██╔██╗ ██║███████║██║     \n██╔══██║██╔══╝   ██╔██╗ ██╔══██║██║   ██║██║   ██║██║╚██╗██║██╔══██║██║     \n██║  ██║███████╗██╔╝ ██╗██║  ██║╚██████╔╝╚██████╔╝██║ ╚████║██║  ██║███████╗\n╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═╝  ╚═══╝╚═╝  ╚═╝╚══════╝ . " + os.Getenv("PKG"))
	color.HiGreen("Power By GIN\n\n")
}
