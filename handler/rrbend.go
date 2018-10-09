package handler

import (
	"sync"
	"github.com/gwtony/gapi/log"
	"github.com/gwtony/gapi/api"
	"github.com/gwtony/gapi/config"
)


var AdminToken string

// InitContext inits dcron context
func InitContext(conf *config.Config, log log.Log) error {
	var rh *RedisHandler
	cf := &DcronConfig{}
	err := cf.ParseConfig(conf)
	if err != nil {
		log.Error("Dcron parse config failed")
		return err
	}
	AdminToken = cf.adminToken

	mc, err := InitMysqlContext(cf.maddr, cf.dbname, cf.dbuser, cf.dbpwd, log)
	if err != nil {
		log.Error("Dcron init mysql context failed")
		return err
	}

	apiLoc := cf.apiLoc

	api.AddHttpHandler(apiLoc + DCRON_ADD_LOC, &AddHandler{eh: eh, ch: ch, cw: cw, log: log})

	ch.Run()

	return nil
}
