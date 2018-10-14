package db
import (
	"database/sql"
	"github.com/gwtony/gapi/errors"
	"github.com/gstdio/rrb-backend/structs"
)

func (mc *MysqlContext) GetAllSubClass() ([]structs.SubClass, error) {
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

func (mc *MysqlContext) GetSubClassByClassId(id int) ([]structs.SubClass, error) {
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
