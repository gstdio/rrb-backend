package db
import (
	"database/sql"
	"github.com/gwtony/gapi/errors"
	"github.com/gstdio/rrb-backend/structs"
)

func (mc *MysqlContext) GetAllClass() ([]structs.Class, error) {
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
