package db

const (
	CLASS_GETALL_SQL    = "select * from classes"
	SUBCLASS_GETALL_SQL = "select * from subclasses"
	SUBCLASS_GETBY_CLASSID_SQL = "select * from subclasses where class_id = ?"
	PRODUCT_GETBY_SUBCLASSID_SQL = "select * from products where subclass_id = ?"
	PRODUCT_GETBY_SHOPID_SQL = "select products.*, sales.price from products inner join sales on products.id = sales.product_id and sales.shop_id = ?"
)
