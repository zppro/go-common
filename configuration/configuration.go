package configuration

import (
	"flag"
)

const (
	defaultConfig = "configs/mmgop.conf"
	usageConfig   = "Path to configuration file"
)

var (
	c configuration
	configPath string
	configMap = make(map[string]string)
)


type configuration struct {
	Es elasticsearch
}


func init () {

}

func Load () {
	flag.StringVar(&configPath, "config", defaultConfig, usageConfig)
	flag.StringVar(&configPath, "c", defaultConfig, usageConfig)
	flag.Parse()
	var p Parser = confParser{}
	err := p.Parse(configPath, configMap)
	if err != nil {
		panic(err)
	}
	c.readES()
}

func Get () configuration {
	return c
}

func (c *configuration) readES () {
	for k, v := range configMap {
		if k == "es-host" {
			c.Es.Host = v
		}
		if k == "es-port" {
			c.Es.Port = v
		}
	}
}
