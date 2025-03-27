package configuration

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var enableStdOut bool = false

func init() {
	viper.AddConfigPath("assets/application/configuration/")
	viper.SetConfigName("configuration") // Register config file name (no extension)
	viper.SetConfigType("yaml")          // Look for specific type

	readConfigPerEnvironment()

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("Cannot read application configuration")
	}
	initLogger()

	viper.WatchConfig()
	viper.OnConfigChange(onConfigChange)
	enableStdOut = viper.GetBool("log.configuration.stdout")
}

func readConfigPerEnvironment() {
	env := os.Getenv("env")
	if env != "" {
		viper.AddConfigPath("assets/application/configuration/")
		viper.SetConfigName(fmt.Sprintf("%s-%s", "configuration", strings.ToLower(env))) // Register config file name (no extension)
		viper.SetConfigType("yaml")
		err := viper.MergeInConfig()
		if err != nil {
			log.Errorln("Cannot read application configuration")
		}
	}

}

func onConfigChange(e fsnotify.Event) {
	if enableStdOut {
		log.Debugln(fmt.Sprintf("Config file changed: %s", e.Name))
	}
}

func GetItem(key string) any {
	if log != nil && enableStdOut {
		log.Debugln(fmt.Sprintf("Returning item %s", key))
	}
	return viper.Get(key)
}

func GetStringSlice(key string) []string {
	if log != nil && enableStdOut {
		log.Debugln(fmt.Sprintf("Returning slice for %s", key))
	}
	return viper.GetStringSlice(key)
}

func GetItemOrDefault(key string, val any) any {
	if log != nil && enableStdOut {
		log.Debugln(fmt.Sprintf("Returning item %s", key))
	}
	if viper.Get(key) == nil {
		return val
	}
	return viper.Get(key)
}

func GetItemToStruct(key string, a any) {
	err := viper.UnmarshalKey(key, a)
	if err != nil {
		log.Errorln(fmt.Sprintf("Cannot unmarshall %s", key))
	}
}

func GetString(key string) string {
	if log != nil && enableStdOut {
		log.Debugln(fmt.Sprintf("Returning item %s", key))
	}
	return viper.GetString(key)
}

func GetStringOrDefault(key string, val string) string {
	if log != nil && enableStdOut {
		log.Debugln(fmt.Sprintf("Returning item %s", key))
	}
	if viper.Get(key) == nil {
		return val
	}
	return viper.GetString(key)
}

func GetInt(key string) int {
	if log != nil && enableStdOut {
		log.Debugln(fmt.Sprintf("Returning item %s", key))
	}
	return viper.GetInt(key)
}

func GetIntOrDefault(key string, val int) int {
	if log != nil && enableStdOut {
		log.Debugln(fmt.Sprintf("Returning item %s", key))
	}
	if viper.Get(key) == nil {
		return val
	}
	return viper.GetInt(key)
}

func GetBool(key string) bool {
	if log != nil && enableStdOut {
		log.Debugln(fmt.Sprintf("Returning item %s", key))
	}
	return viper.GetBool(key)
}

func GetBoolOrDefault(key string, val bool) bool {
	if log != nil && enableStdOut {
		log.Debugln(fmt.Sprintf("Returning item %s", key))
	}
	if viper.Get(key) == nil {
		return val
	}
	return viper.GetBool(key)
}
