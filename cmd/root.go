package cmd

import (
	"fmt"
	"github.com/bukowa/kutils/pkg"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)
var AppFs = afero.NewOsFs()

func projectDir(path string) string {
	return fmt.Sprintf("%v/%v", projectName, path)
}

var (
	projectName string
	hosts []string

	rootCmd = &cobra.Command{
		Use:   "cobra",
		Run: func(cmd *cobra.Command, args []string) {
			// create project dir
			err := AppFs.Mkdir(projectName, 0644); er(err)
			// create ingress
			if len(hosts) == 0 {
				hosts = viper.GetStringSlice("hosts")
			}
			ingress := &pkg.Ingress{
				Name:           fmt.Sprintf("ingress-%v", projectName),
				Hosts:          hosts,
				Class:          "nginx",
				ServiceName:    fmt.Sprintf("service-%v", projectName),
				ServicePort:    8080,
				PathTypePrefix: "Prefix",
				RewriteWWW: true,
			}
			err = pkg.SaveYaml(ingress.KubernetesObject(), projectDir("ingress.yaml")); er(err)
			for _, k := range viper.AllKeys() {
				fmt.Println(k, viper.Get(k))
			}
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&projectName, "name", "n", "", "project name")
	rootCmd.Flags().StringSliceVarP(&hosts, "hosts", "x", []string{}, "hosts to use")

	if err := viper.BindPFlag("name", rootCmd.PersistentFlags().Lookup("name")); err != nil {
		panic(err)
	}
	if err := viper.BindPFlag("hosts", rootCmd.Flags().Lookup("hosts")); err != nil {
		panic(err)
	}
	rootCmd.MarkPersistentFlagRequired("name")
}

func initConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigFile("cobra.yaml")

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		panic(err)
	}
	if err := viper.WriteConfig(); err == nil {
		fmt.Println("Saving config file:", viper.ConfigFileUsed())
	} else {
		panic(err)
	}

	fmt.Println("Config keys:", viper.AllKeys())
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func er(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
