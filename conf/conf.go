package conf

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

//环境变量
var (
	configFile = ""
)

type Config struct {
	Server struct {
		Name              string `yaml:"name"`
		HttpListen        string `yaml:"httpListen"`
		LogPath           string `yaml:"logPath"`
		Debug             bool   `yaml:"debug"`
		PrivateHttpListen string `yaml:"privateHttpListen"`
		JobMaster         bool   `yaml:"JobMaster"`
	} `yaml:"server,omitempty"`

	Mysql struct {
		Host     string `yaml:"host,omitempty"`
		Port     string `yaml:"port,omitempty"`
		Password string `yaml:"password,omitempty"`
		User     string `yaml:"user,omitempty"`
		DbName   string `yaml:"dbName,omitempty"`
	} `yaml:"mysql"`

	Redis struct {
		Host     string `yaml:"host,omitempty"`
		Port     string `yaml:"port,omitempty"`
		Password string `yaml:"password,omitempty"`
	} `yaml:"redis"`
}

var AppConfig *Config

func Init() {
	configFile = os.Getenv("config")
	if len(configFile) == 0 {
		flag.StringVar(&configFile, "c", "conf/config.yaml", "config file")
		flag.Parse()
	}
	bytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(bytes, &AppConfig)
	if err != nil {
		panic(err)
	}

}
