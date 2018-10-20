package db
import (
	"database/sql"
	"github.com/gwtony/gapi/errors"
	"github.com/gstdio/rrb-backend/structs"
)

func (mc *MysqlContext) SaleGetAll() ([]structs.Sale, error) {
	result, err := mc.QueryRead(parseSale, SALE_GETALL_SQL)
	if err != nil {
		return nil, err
	}
	if res, ok := result.([]structs.Sale); ok {
		return res, nil
	} else {
		return nil, errors.InternalServerError
	}
}

func parseSale(rows *sql.Rows) (interface{}, error) {
	classArr := []structs.Sale{}
	for rows.Next() {
		c := structs.Sale{}
		err := rows.Scan(&c.Id, &c.ProductId, &c.ShopId, &c.Price, &c.Title, &c.Desc, &c.CreateTime)
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

func (mc *MysqlContext) SaleInsert(c *structs.Sale) error {
	return mc.QueryWrite(SALE_INSERT_SQL, &c.ProductId, &c.ShopId, &c.Price, &c.Title, &c.Desc, &c.CreateTime)
}

func (mc *MysqlContext) SaleUpdate(c *structs.Sale) error {
	return mc.QueryWrite(SALE_UPDATE_SQL, &c.ProductId, &c.ShopId, &c.Price, &c.Title, &c.Desc, &c.CreateTime, &c.Id)
}

func (mc *MysqlContext) SaleDelete(c *structs.Sale) error {
	return mc.QueryWrite(SALE_DELETE_SQL, c.Id)
}
