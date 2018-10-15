package db
import (
	"database/sql"
	"github.com/gwtony/gapi/errors"
	"github.com/gstdio/rrb-backend/structs"
)

func (mc *MysqlContext) GetProductBySubClassId(id int) ([]structs.Product, error) {
	result, err := mc.QueryRead(parseProduct, PRODUCT_GETBY_SUBCLASSID_SQL, id)
	if err != nil {
		return nil, err
	}
	if res, ok := result.([]structs.Product); ok {
		return res, nil
	} else {
		return nil, errors.InternalServerError
	}
}

func (mc *MysqlContext) GetProductByShopId(id int) ([]structs.Product, error) {
	result, err := mc.QueryRead(parseProduct, PRODUCT_GETBY_SHOPID_SQL, id)
	if err != nil {
		return nil, err
	}
	if res, ok := result.([]structs.Product); ok {
		return res, nil
	} else {
		return nil, errors.InternalServerError
	}
}

func parseProduct(rows *sql.Rows) (interface{}, error) {
	arr := []structs.Product{}
	for rows.Next() {
		c := structs.Product{}
		err := rows.Scan(&c.Id, &c.Subclass_id, &c.Name, &c.Desc)
		if err == sql.ErrNoRows {
			return "", errors.NoContentError
		}
		if err != nil {
			return "", errors.InternalServerError
		}
		arr = append(arr, c)
	}
	return arr, nil
}
