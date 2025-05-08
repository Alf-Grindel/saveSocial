package conf

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

type server struct {
	Version string
	Name    string
}

type mySQL struct {
	Addr     string
	Database string
	UserName string
	Password string
}

type config struct {
	Server server
	MySQL  mySQL
}

var (
	Server       *server
	MySQL        *mySQL
	runtimeViper = viper.New()
)

func Init() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("[FATAL] conf.Init: %s", err)
	}
	dir := getDir(path)
	runtimeViper.SetConfigName("config")
	runtimeViper.SetConfigType("yml")
	runtimeViper.AddConfigPath(dir)

	if err := runtimeViper.ReadInConfig(); err != nil {
		log.Fatalf("[FATAL] conf.Init: %s", err)
	}

}

func getDir(path string) string {
	dir := path
	for {
		if _, err := os.Stat(filepath.Join(dir, "config.yaml")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			log.Panicln("[ERROR] conf.getDir: can not find the config file dir")
		}
		dir = parent
	}
}

func configMapping() {
	c := &config{}
	if err := runtimeViper.Unmarshal(&c); err != nil {
		log.Panicln("[ERROR] conf.configMapping: can not unmarshal config")
	}
	Server = &c.Server
	MySQL = &c.MySQL
}
