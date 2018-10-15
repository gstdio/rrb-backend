package db

const (
	CLASS_GETALL_SQL    = "select * from classes"
	SUBCLASS_GETALL_SQL = "select * from subclasses"
	SUBCLASS_GETBY_CLASSID_SQL = "select * from subclasses where class_id = ?"
	PRODUCT_GETBY_SUBCLASSID_SQL = "select products.*, product_properties.display_url from products inner join product_properties on products.id = product_properties.product_id and products.subclass_id = ?"
	PRODUCT_GETBY_SHOPID_SQL = "select products.* from products inner join sales on products.id = sales.product_id and sales.shop_id = ?"
)
