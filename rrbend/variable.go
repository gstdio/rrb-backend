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

	SUBCLASS_GETALL_LOC          = "/subclass/getall"
	SUBCLASS_GETBY_CLASSID_LOC   = "/subclass"
	SUBCLASS_INSERT_LOC          = "/subclass/insert"
	SUBCLASS_UPDATE_LOC          = "/subclass/update"

	PRODUCT_GETBY_SUBCLASSID_LOC = "/product/subclass"
	PRODUCT_GETBY_SHOPID_LOC     = "/product/shop"
	PRODUCT_INSERT_LOC           = "/product/insert"
	PRODUCT_UPDATE_LOC           = "/product/update"

	DEFAULT_RETENTION_TIME       = "1d"
	RRBEND_LOC                   = "/rrbend"
)
