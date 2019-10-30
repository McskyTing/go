package main

import (
	"flag"
	"sync"

	"github.com/larspensjo/config"
)

var Conf_Main_Topic = "DEFAULT"

var (
	//config.ini为配置文件格式为 key=value
	configFile = flag.String("configfile", "config.ini", "General configuration file")
)
var commonConf = make(map[string]string)
var lock sync.RWMutex

func LoadCommonConfiguration() {
	lock.Lock()
	defer lock.Unlock()
	cfg, err := config.ReadDefault(*configFile)
	if err != nil {
		//....
	}
	commonConf = make(map[string]string)
	if cfg.HasSection(Conf_Main_Topic) {
		section, err := cfg.SectionOptions(Conf_Main_Topic)
		if err != nil {
			//....
		}
		for _, v := range section {
			options, err := cfg.String(Conf_Main_Topic, v)
			if err != nil {
				//....
			}
			commonConf[v] = options
		}
	}
}

//通过GetConf方法将key传入获取value值
func GetConf(key string) string {
	lock.RLock()
	defer lock.RUnlock()
	return commonConf[key]
}
