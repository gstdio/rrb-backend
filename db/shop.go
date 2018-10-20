package db
import (
	"database/sql"
	"github.com/gwtony/gapi/errors"
	"github.com/gstdio/rrb-backend/structs"
)

func (mc *MysqlContext) ShopGetAll() ([]structs.Shop, error) {
	result, err := mc.QueryRead(parseShop, SHOP_GETALL_SQL)
	if err != nil {
		return nil, err
	}
	if res, ok := result.([]structs.Shop); ok {
		return res, nil
	} else {
		return nil, errors.InternalServerError
	}
}

func parseShop(rows *sql.Rows) (interface{}, error) {
	classArr := []structs.Shop{}
	for rows.Next() {
		c := structs.Shop{}
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

func (mc *MysqlContext) ShopInsert(c *structs.Shop) error {
	return mc.QueryWrite(SHOP_INSERT_SQL, c.RegionId, c.Name, c.Address, c.Desc, c.Url)
}

func (mc *MysqlContext) ShopUpdate(c *structs.Shop) error {
	return mc.QueryWrite(SHOP_UPDATE_SQL, c.RegionId, c.Name, c.Address, c.Desc, c.Url, c.Id)
}

func (mc *MysqlContext) ShopDelete(c *structs.Shop) error {
	return mc.QueryWrite(SHOP_DELETE_SQL, c.Id)
}
