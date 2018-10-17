package db

import (
	"fmt"
	//"time"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gwtony/gapi/log"
	"github.com/gwtony/gapi/errors"
)

type MysqlContext struct {
	addr    string
	dbname  string
	dbuser  string
	dbpwd   string

	db      *sql.DB
	login   string

	log     log.Log
}

func InitMysqlContext(addr, dbname, dbuser, dbpwd string, log log.Log) (*MysqlContext, error) {
	mc := &MysqlContext{}

	mc.log      = log
	mc.addr     = addr
	mc.dbname   = dbname
	mc.dbuser   = dbuser
	mc.dbpwd    = dbpwd
	mc.login    = fmt.Sprintf("%s:%s@tcp(%s)/%s", dbuser, dbpwd, addr, dbname)
	db, err     := sql.Open("mysql", mc.login)
	if err != nil {
		mc.log.Error("Open mysql context failed: %s", err)
		return nil, err
	}
	mc.db = db

	return mc, nil
}

func (mc *MysqlContext) Close() error{
	return mc.db.Close()
}

func (mc *MysqlContext) QueryRead(callBack func(*sql.Rows)(interface{}, error), querySql string, args ...interface{}) (interface{}, error) {
	var rows *sql.Rows
	var err error

	if len(args) == 0 {
		fmt.Println("Query read no args")
		rows, err = mc.db.Query(querySql)
	} else {
		fmt.Println("Query read args is", args)
		rows, err = mc.db.Query(querySql, args...)
	}
	if err != nil {
		if err == sql.ErrNoRows {
			mc.log.Error("Scan no answer")
			return "", errors.NoContentError
		}
		mc.log.Error("Execute get result for sql (%s) failed: %s", querySql, err)
		return "", errors.BadGatewayError
	}
	defer rows.Close()

	return callBack(rows)
}

func (mc *MysqlContext) QueryWrite(query string, args ...interface{}) error {
	res, err := mc.db.Exec(query, args...)

	if err != nil {
		mc.log.Error("Execute write sql: ", query, args, " failed: ", err)
		return errors.BadGatewayError
	}
	affected, err := res.RowsAffected()
	if err != nil {
		mc.log.Error("Write get rows affected failed: %s", err)
		return errors.InternalServerError
	}
	if int(affected) <= 0 {
		return errors.BadGatewayError
	}

	return nil
}

func (mc *MysqlContext) QueryUpdate(query string, args ...interface{}) error {
	res, err := mc.db.Exec(query, args...)

	if err != nil {
		mc.log.Error("Execute update sql: ", query, args, " failed: ", err)
		return errors.BadGatewayError
	}

	affected, err := res.RowsAffected()
	if err != nil {
		mc.log.Error("Update get rows affected failed: %s", err)
		return errors.InternalServerError
	}

	if int(affected) > 0 {
		return nil
	}

	return errors.NoContentError
}
