package store

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
}

type cfg struct {
	Runmode string `yaml:"runmode"`
	Http    struct {
		Port int `yaml:"port"`
	} `yaml:"http"`
	Etcd struct {
		Name   string `yaml:"name"`
		Addr   string `yaml:"addr"`
		Ttl    int64  `yaml:"ttl"`
		Key    string `yaml:"key"`
		Schema string `yaml:"schema"`
	} `yaml:"etcd"`
}

func (s *Config) Get() (conf cfg) {
	configOnce.Do(func() {
		path := ""
		flag.StringVar(&path, "conf", "./configs/config.yml", "help")
		flag.Parse()
		bytes, err := ioutil.ReadFile(path)
		if nil != err {
			panic(err)
		}
		err = yaml.Unmarshal(bytes, &conf)
		if nil != err {
			panic(err)
		}
		config = conf
	})
	conf = config
	return
}
