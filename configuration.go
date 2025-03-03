package configuration

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func init() {
	viper.AddConfigPath("assets/application/configuration/")
	viper.SetConfigName("configuration") // Register config file name (no extension)
	viper.SetConfigType("yaml")          // Look for specific type

	readConfigPerEnvironment()

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatalln("Cannot read application configuration")
	}
	initLogger()

	viper.WatchConfig()
	viper.OnConfigChange(onConfigChange)
}

func readConfigPerEnvironment() {
	env := os.Getenv("env")
	if env != "" {
		viper.AddConfigPath("assets/application/configuration/")
		viper.SetConfigName(fmt.Sprintf("%s-%s", "configuration", strings.ToLower(env))) // Register config file name (no extension)
		viper.SetConfigType("yaml")
		err := viper.MergeInConfig()
		if err != nil {
			logger.Errorln("Cannot read application configuration")
		}
	}

}

func onConfigChange(e fsnotify.Event) {
	logger.Debugln(fmt.Sprintf("Config file changed: %s", e.Name))
}

func GetItem(key string) any {
	logger.Debugln(fmt.Sprintf("Returning item %s", key))
	return viper.Get(key)
}

func GetStringSlice(key string) []string {
	logger.Debugln(fmt.Sprintf("Returning slice for %s", key))
	return viper.GetStringSlice(key)
}

func GetItemOrDefault(key string, val any) any {
	logger.Debugln(fmt.Sprintf("Returning item %s", key))
	if viper.Get(key) == nil {
		return val
	}
	return viper.Get(key)
}

func GetItemToStruct(key string, a any) {
	err := viper.UnmarshalKey(key, a)
	if err != nil {
		logger.Errorln(fmt.Sprintf("Cannot unmarshall %s", key))
	}
}

func GetString(key string) string {
	logger.Debugln(fmt.Sprintf("Returning item %s", key))
	return viper.GetString(key)
}

func GetStringOrDefault(key string, val string) string {
	logger.Debugln(fmt.Sprintf("Returning item %s", key))
	if viper.Get(key) == nil {
		return val
	}
	return viper.GetString(key)
}

func GetInt(key string) int {
	logger.Debugln(fmt.Sprintf("Returning item %s", key))
	return viper.GetInt(key)
}

func GetIntOrDefault(key string, val int) int {
	logger.Debugln(fmt.Sprintf("Returning item %s", key))
	if viper.Get(key) == nil {
		return val
	}
	return viper.GetInt(key)
}

func GetBool(key string) bool {
	logger.Debugln(fmt.Sprintf("Returning item %s", key))
	return viper.GetBool(key)
}

func GetBoolOrDefault(key string, val bool) bool {
	logger.Debugln(fmt.Sprintf("Returning item %s", key))
	if viper.Get(key) == nil {
		return val
	}
	return viper.GetBool(key)
}
