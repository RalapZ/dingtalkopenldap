package command

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var (
	ENV          string
	WorkDir      = "../"
	AuthConf     Authentication
	Listenconfig Listen
	Authconfig   Authentication
	Ldapconfig   LdapConfig
	DBconfig     DBConfig
)

type User struct {
	Username string
	Password string
}

type Authentication struct {
	AppKey    string `yaml:"appkey"`
	AppSecret int    `yaml:"appsecret"`
}

type Listen struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type LdapConfig struct {
	Url      string `yaml:"url"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type DBConfig struct {
	DBType   string `yaml:"dbtype"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type Config struct {
	Listen     `yaml:"listen"`
	LdapConfig `yaml:"ldap"`
	DBConfig   `yaml:"DB"`
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
		fmt.Println(config)
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
