package tools

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestInitLog(t *testing.T) {
	InitLog("./log/", "t", "gbk")
	log.WithFields(log.Fields{"animal": "walrus"}).Info("A walrus appears")
	log.Info("测试中文")
	log.Warning("this is test")
}