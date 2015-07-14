package goconfig

import (
	"code.google.com/p/gcfg"
	"log"
	"os"
)

type ConfigLoader struct {
	DefaultFilename string
}

func NewConfigLoader(defaultFilename string) *ConfigLoader {
	cl := new(ConfigLoader)
	cl.DefaultFilename = defaultFilename
	return cl
}

func (cl *ConfigLoader) Load(cfg interface{}) {
	var file string
	if len(os.Args) > 1 {
		file = os.Args[1]
	} else {
		file = cl.DefaultFilename
	}
	log.Printf("Loading from config file: %v", file)
	err := gcfg.ReadFileInto(cfg, file)
	if err != nil {
		log.Fatal("Failed to read config: ", err)
	}
}
