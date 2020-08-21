package config

import (
	"encoding/json"
	"github.com/kelseyhightower/envconfig"
	"io/ioutil"
	"os"
	"strconv"
)

var IsDebugEnvName = "DEBUG"

type Loader struct {
	Config
}

func CreateLoader(config *Config) *Loader {
	if config == nil {
		config = createDefaultConfig("")
		var isDebug, _ = GetEnvValue(IsDebugEnvName)
		config.IsDebug = isDebug == strconv.FormatBool(true)
	}
	return &Loader{Config:*config}
}

func GetEnvValue(key string) (string, bool) {
	return os.LookupEnv(key)
}

func (loader *Loader)LoadEnv(config interface{}) error {
	return envconfig.Process("", config)
}
func (loader *Loader)LoadJson(config interface{}, isDebug bool) error {
	var filePath = loader.getJsonFilePath(isDebug)
	dat, err := loader.getFile(filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(dat, config)
}

func (loader *Loader)getFile(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}

func (loader *Loader)getJsonFilePath(isDebug bool) string {
	if isDebug {
		return loader.JsonDebugFilePath
	}
	return loader.JsonReleaseFilePath
}