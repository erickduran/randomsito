/*
Copyright © 2020 Erick Durán me@erickduran.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	homedir "github.com/mitchellh/go-homedir"
)

var cfgFile string
var Language string
var MongoHost string
var MongoPort string
var MongoDatabase string
var ConnectionString string
var BLUE = "\033[0;34m"
var NC = "\033[0m"

var rootCmd = &cobra.Command{
	Use:   "randomsito",
	Short: "Randomsito CLI.",
	Long: `--- Randomsito ---
A CLI for interactive classes. 

Erick Durán. Copyright © 2020. 
https://www.erickduran.com`,
	Run: startCmd.Run,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.randomsito.yaml)")
	rootCmd.PersistentFlags().StringVarP(&Language, "language", "l", "en", "select a language (en, es)")
	rootCmd.PersistentFlags().StringVarP(&MongoHost, "host", "m", "localhost", "mongodb host")
	rootCmd.PersistentFlags().StringVarP(&MongoPort, "port", "p", "27017", "mongodb port")
	rootCmd.PersistentFlags().StringVarP(&MongoDatabase, "database", "b", "randomsito", "mongodb database name")
	rootCmd.Flags().BoolP("dry-run", "d", false, "don't save changes")
}

func initConfig() {
	fmt.Println(`   ┏━┓┏━┓┏┓╻╺┳┓┏━┓┏┳┓┏━┓╻╺┳╸┏━┓   
╺━╸┣┳┛┣━┫┃┗┫ ┃┃┃ ┃┃┃┃┗━┓┃ ┃ ┃ ┃╺━╸
   ╹┗╸╹ ╹╹ ╹╺┻┛┗━┛╹ ╹┗━┛╹ ╹ ┗━┛   `)
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".randomsito")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("config:", viper.ConfigFileUsed())
		language := viper.Get("language")
		if language != nil {
			Language = language.(string)
		}

		mongoHost := viper.Get("mongodb_host")
		if mongoHost != nil {
			MongoHost = mongoHost.(string)
		}

		mongoPort := viper.Get("mongodb_port")
		if mongoPort != nil {
			MongoPort = mongoPort.(string)
		}

		mongoDatabase := viper.Get("mongodb_name")
		if mongoDatabase != nil {
			MongoDatabase = mongoDatabase.(string)
		}
	}

	ConnectionString = fmt.Sprintf("mongodb://%s:%s", MongoHost, MongoPort)
}
