package conf

import (
	"Ortau/static"
	"gopkg.in/ini.v1"
	"log"
	"os"
)

func checkCfgFileIsExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if err != nil {
		return false
	} else {
		return true
	}
}

func makeCfg() bool {
	fileName := "config.ini"
	if checkCfgFileIsExist(fileName) == false {
		f, err := os.Create(fileName)
		if err != nil {
			log.Println("1", err)
		}

		defer f.Close()
		_, err = f.Write([]byte(static.Config))
		if err != nil {
			log.Println("2", err)
		}

		return false
	} else {
		return true
	}
}

func GetCfgSectionKey(section string, key string) string {
	var cfgValue string

	for{
		if makeCfg() == false {
			log.Println("[-] Info Not found cfg...Try to make config.ini")
		} else {
			cfg, err := ini.Load("config.ini")
			if err != nil {
				log.Println("Fail to read file: %v", err)
				os.Exit(1)
			}
			cfgValue = cfg.Section(section).Key(key).String()
			break
		}
	}

	return cfgValue
}
