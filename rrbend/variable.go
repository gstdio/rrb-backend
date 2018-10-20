package rrbend

const (
	// VERSION version
	VERSION                      = "0.1 alpha"

	API_CONTENT_HEADER           = "application/json;charset=utf-8"

	DEFAULT_ADMIN_TOKEN          = "CRON_TOKEN"
	ADMIN_TOKEN_HEADER           = "Admin-Token"

	//ETCD_CONTENT_HEADER           = "application/x-www-form-urlencoded"

	//http location for class
	CLASS_GETALL_LOC             = "/class/getall"
	CLASS_INSERT_LOC             = "/class/insert"
	CLASS_UPDATE_LOC             = "/class/update"
	CLASS_DELETE_LOC             = "/class/delete"

	SUBCLASS_GETALL_LOC          = "/subclass/getall"
	SUBCLASS_GETBY_CLASSID_LOC   = "/subclass"
	SUBCLASS_INSERT_LOC          = "/subclass/insert"
	SUBCLASS_UPDATE_LOC          = "/subclass/update"
	SUBCLASS_DELETE_LOC          = "/subclass/delete"

	PRODUCT_GETBY_SUBCLASSID_LOC = "/product/subclass"
	PRODUCT_GETBY_SHOPID_LOC     = "/product/shop"
	PRODUCT_INSERT_LOC           = "/product/insert"
	PRODUCT_UPDATE_LOC           = "/product/update"
	PRODUCT_DELETE_LOC           = "/product/delete"

	SHOP_INSERT_LOC              = "/shop/insert"
	SHOP_UPDATE_LOC              = "/shop/update"
	SHOP_DELETE_LOC              = "/shop/delete"

	SALE_INSERT_LOC              = "/sale/insert"
	SALE_UPDATE_LOC              = "/sale/update"
	SALE_DELETE_LOC              = "/sale/delete"
	SALE_TEST_LOC              = "/sale/test"

	DEFAULT_RETENTION_TIME       = "1d"
	RRBEND_LOC                   = "/rrbend"
)
