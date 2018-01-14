package conf

import (
	"flag"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

var (
	// Configuration ..
	Configuration *Conf
)

// Conf : App Configuration Variables
type Conf struct {
	Name   string `yaml:"name"`
	Env    string `yaml:"env"`
	Server struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	}
	StaticPath string `yaml:"static-path"`
	AuthPath   string `yaml:"auth-path"`
}

// NewConf : NewConf
func NewConf() {
	var path string
	flag.StringVar(&path, "config", "conf/dev.yaml", "Parse Configuration File")
	flag.Parse()

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("parseconf error: %v", err)
	}

	err = yaml.Unmarshal(data, &Configuration)
	if err != nil {
		log.Fatalf("parseconf error: %v", err)
	}
}
