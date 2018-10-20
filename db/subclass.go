package db
import (
	"database/sql"
	"github.com/gwtony/gapi/errors"
	"github.com/gstdio/rrb-backend/structs"
)

func (mc *MysqlContext) SubClassGetAll() ([]structs.SubClass, error) {
	result, err := mc.QueryRead(parseSubClass, SUBCLASS_GETALL_SQL)
	if err != nil {
		return nil, err
	}
	if res, ok := result.([]structs.SubClass); ok {
		return res, nil
	} else {
		return nil, errors.InternalServerError
	}
}

func (mc *MysqlContext) SubClassGetByClassId(id int) ([]structs.SubClass, error) {
	result, err := mc.QueryRead(parseSubClass, SUBCLASS_GETBY_CLASSID_SQL, id)
	if err != nil {
		return nil, err
	}
	if res, ok := result.([]structs.SubClass); ok {
		return res, nil
	} else {
		return nil, errors.InternalServerError
	}
}

func parseSubClass(rows *sql.Rows) (interface{}, error) {
	classArr := []structs.SubClass{}
	for rows.Next() {
		c := structs.SubClass{}
		err := rows.Scan(&c.Id, &c.ClassId, &c.Name, &c.Desc)
		if err == sql.ErrNoRows {
			return "", errors.NoContentError
		}
		if err != nil {
			return "", errors.InternalServerError
		}
		classArr = append(classArr, c)
	}
	return classArr, nil
}

func (mc *MysqlContext) SubClassInsert(c *structs.SubClass) error {
	return mc.QueryWrite(SUBCLASS_INSERT_SQL, c.ClassId, c.Name, c.Desc)
}

func (mc *MysqlContext) SubClassUpdate(c *structs.SubClass) error {
	return mc.QueryWrite(SUBCLASS_UPDATE_SQL, c.ClassId, c.Name, c.Desc, c.Id)
}

func (mc *MysqlContext) SubClassDelete(c *structs.SubClass) error {
	return mc.QueryWrite(SUBCLASS_DELETE_SQL, c.Id)
}
