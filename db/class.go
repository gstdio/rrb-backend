package db
import (
	"database/sql"
	"github.com/gwtony/gapi/errors"
	"github.com/gstdio/rrb-backend/structs"
)

func (mc *MysqlContext) ClassGetAll() ([]structs.Class, error) {
	result, err := mc.QueryRead(parseClass, CLASS_GETALL_SQL)
	if err != nil {
		return nil, err
	}
	if res, ok := result.([]structs.Class); ok {
		return res, nil
	} else {
		return nil, errors.InternalServerError
	}
}

func parseClass(rows *sql.Rows) (interface{}, error) {
	classArr := []structs.Class{}
	for rows.Next() {
		c := structs.Class{}
		err := rows.Scan(&c.Id, &c.Name, &c.Desc)
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

func (mc *MysqlContext) ClassInsert(c *structs.Class) error {
	return mc.QueryWrite(CLASS_INSERT_SQL, c.Name, c.Desc)
}

func (mc *MysqlContext) ClassUpdate(c *structs.Class) error {
	return mc.QueryWrite(CLASS_UPDATE_SQL, c.Name, c.Desc, c.Id)
}
