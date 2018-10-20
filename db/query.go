package db

const (
	CLASS_GETALL_SQL             = "select * from classes"
	CLASS_INSERT_SQL             = "insert into classes (name, description) values (?, ?)"
	CLASS_UPDATE_SQL             = "update classes set name = ?, description = ? where id = ?"
	CLASS_DELETE_SQL             = "delete from classes where id = ?"

	SUBCLASS_GETALL_SQL          = "select * from subclasses"
	SUBCLASS_GETBY_CLASSID_SQL   = "select * from subclasses where class_id = ?"
	SUBCLASS_INSERT_SQL          = "insert into subclasses (class_id, name, description) values (?, ?, ?)"
	SUBCLASS_UPDATE_SQL          = "update subclasses set class_id = ?, name = ?, description = ? where id = ?"
	SUBCLASS_DELETE_SQL          = "delete from subclasses where id = ?"

	PRODUCT_GETBY_SUBCLASSID_SQL = "select products.*, min(price) from products inner join sales on products.id = sales.product_id and products.subclass_id = ? group by products.id"
	PRODUCT_GETBY_SHOPID_SQL     = "select products.*, sales.price from products inner join sales on products.id = sales.product_id and sales.shop_id = ?"
	PRODUCT_INSERT_SQL           = "insert into products (subclass_id, name, description, url) values (?, ?, ?, ?)"
	PRODUCT_UPDATE_SQL           = "update products set subclass_id = ?, name = ?, description = ?, url = ? where id = ?"
	PRODUCT_DELETE_SQL           = "delete from products where id = ?"

	SHOP_GETALL_SQL              = "select * from shops"
	SHOP_INSERT_SQL              = "insert into shops"
	SHOP_UPDATE_SQL              = "update shops set"
	SHOP_DELETE_SQL              = "delete from shops where id = ?"

	SALE_GETALL_SQL              = "select * from sales"
	SALE_INSERT_SQL              = "insert into saless (product_id, shop_id, price, title, description, create_time) values (?, ?, ?, ?, ?, ?)"
	SALE_UPDATE_SQL              = "update sales set price = ?, title = ?, description = ? where id = ?"
	SALE_DELETE_SQL              = "delete from sales where id = ?"
)
