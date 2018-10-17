package rrbend

import (
	"github.com/gwtony/gapi/log"
	"github.com/gwtony/gapi/api"
	"github.com/gwtony/gapi/config"

	"github.com/gstdio/rrb-backend/db"
	"github.com/gstdio/rrb-backend/handler"
)


var AdminToken string

// InitContext inits dcron context
func InitContext(conf *config.Config, log log.Log) error {
	cf := &RrbendConfig{}
	err := cf.ParseConfig(conf)
	if err != nil {
		log.Error("Rrbend parse config failed")
		return err
	}
	AdminToken = cf.adminToken

	mc, err := db.InitMysqlContext(cf.maddr, cf.dbname, cf.dbuser, cf.dbpwd, log)
	if err != nil {
		log.Error("Rrbend init mysql context failed")
		return err
	}

	apiLoc := cf.apiLoc

	api.AddHttpHandler(apiLoc + CLASS_GETALL_LOC, &handler.ClassGetAllHandler{Mc: mc, Log: log})
	api.AddHttpHandler(apiLoc + CLASS_INSERT_LOC, &handler.ClassInsertHandler{Mc: mc, Log: log})
	api.AddHttpHandler(apiLoc + CLASS_UPDATE_LOC, &handler.ClassUpdateHandler{Mc: mc, Log: log})

	api.AddHttpHandler(apiLoc + SUBCLASS_GETALL_LOC, &handler.SubClassGetAllHandler{Mc: mc, Log: log})
	api.AddHttpHandler(apiLoc + SUBCLASS_GETBY_CLASSID_LOC, &handler.SubClassGetByClassIdHandler{Mc: mc, Log: log})
	api.AddHttpHandler(apiLoc + SUBCLASS_INSERT_LOC, &handler.SubClassInsertHandler{Mc: mc, Log: log})
	api.AddHttpHandler(apiLoc + SUBCLASS_UPDATE_LOC, &handler.SubClassUpdateHandler{Mc: mc, Log: log})

	api.AddHttpHandler(apiLoc + PRODUCT_GETBY_SUBCLASSID_LOC, &handler.ProductGetBySubClassIdHandler{Mc: mc, Log: log})
	api.AddHttpHandler(apiLoc + PRODUCT_GETBY_SHOPID_LOC, &handler.ProductGetByShopIdHandler{Mc: mc, Log: log})
	api.AddHttpHandler(apiLoc + PRODUCT_INSERT_LOC, &handler.ProductInsertHandler{Mc: mc, Log: log})
	api.AddHttpHandler(apiLoc + PRODUCT_UPDATE_LOC, &handler.ProductUpdateHandler{Mc: mc, Log: log})

	return nil
}
