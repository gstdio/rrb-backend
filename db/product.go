package db
import (
	"database/sql"
	"github.com/gwtony/gapi/errors"
	"github.com/gstdio/rrb-backend/structs"
)

func (mc *MysqlContext) ProductGetBySubClassId(id int) ([]structs.Product, error) {
	result, err := mc.QueryRead(parseProductPrice, PRODUCT_GETBY_SUBCLASSID_SQL, id)
	if err != nil {
		return nil, err
	}
	if res, ok := result.([]structs.Product); ok {
		return res, nil
	} else {
		return nil, errors.InternalServerError
	}
}

func (mc *MysqlContext) ProductGetByShopId(id int) ([]structs.Product, error) {
	result, err := mc.QueryRead(parseProductPrice, PRODUCT_GETBY_SHOPID_SQL, id)
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
		err := rows.Scan(&c.Id, &c.SubClassId, &c.Name, &c.Desc, &c.Url)
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

func parseProductPrice(rows *sql.Rows) (interface{}, error) {
	arr := []structs.Product{}
	for rows.Next() {
		c := structs.Product{}
		err := rows.Scan(&c.Id, &c.SubClassId, &c.Name, &c.Desc, &c.Url, &c.Price)
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

func (mc *MysqlContext) ProductInsert(c *structs.Product) error {
	return mc.QueryWrite(PRODUCT_INSERT_SQL, c.SubClassId, c.Name, c.Desc, c.Url)
}

func (mc *MysqlContext) ProductUpdate(c *structs.Product) error {
	return mc.QueryWrite(PRODUCT_UPDATE_SQL, c.SubClassId, c.Name, c.Desc, c.Url, c.Id)
}
