package model

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var (
	ENV          string
	WorkDir      = "."
	AuthConf     Authentication
	Listenconfig Listen
	Authconfig   Authentication
	Ldapconfig   LdapConfig
	DBconfig     DBConfig
	Defaultconfig DefaultConfig
)

type User struct {
	Username string
	Password string
}

type Authentication struct {
	AppKey    string `yaml:"appkey"`
	AppSecret string    `yaml:"appsecret"`
}

type Listen struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type LdapConfig struct {
	Addr      string `yaml:"addr"`
	//Port	 string `yaml:"port"`
	BindUserName     string `yaml:"bindUserName"`
	BindPassword string `yaml:"bindPassword"`
	SearchDN string `yaml:"searchdn"`
}

type DBConfig struct {
	DBType   string `yaml:"dbtype"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type  DefaultConfig struct{
	StatusList string `yaml:"status_list"`
	UserOffset string `yaml:"user_offset"`
	SchedulerTime int   `yaml:"schedulertime"`
}

type Config struct {
	DefaultConfig  `yaml:"default"`
	Listen         `yaml:"listen"`
	LdapConfig     `yaml:"ldap"`
	Authentication `yaml:"authentication"`
	DBConfig       `yaml:"DB"`
}

func InitConfig() {
	//fmt.Println(os.Getwd())
	setConf()
}

func setConf() {
	filePath := WorkDir + "/conf/conf.yml"
	if configurationContent, err := ioutil.ReadFile(filePath); err != nil {
		panic(fmt.Sprintf("fail to read configuration: %s", filePath))
	} else {
		// configuration := gjson.ParseBytes(configurationContent)
		config := Config{}
		err := yaml.Unmarshal(configurationContent, &config)
		if err != nil {
			log.Printf("conf: %s, error: %v", configurationContent, err)
		}
		//fmt.Println(config)
		//fmt.Println(config)
		//AuthConf=conf
		Defaultconfig=config.DefaultConfig
		Listenconfig =config.Listen
		Authconfig =config.Authentication
		Ldapconfig =config.LdapConfig
		DBconfig =config.DBConfig
		// listen
		//if config. != 0 {
		//	ServerPort = config.Conf.Listen.Port
		//}
		//
		//
		//initAuthentication(config.Authentication)
		//
		//initPlugins(config.Plugins)
	}
}
