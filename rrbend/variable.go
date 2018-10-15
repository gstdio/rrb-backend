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
	SUBCLASS_GETALL_LOC          = "/subclass/getall"
	SUBCLASS_GETBY_CLASSID_LOC   = "/subclass"
	PRODUCT_GETBY_SUBCLASSID_LOC = "/product/subclass"
	PRODUCT_GETBY_SHOPID_LOC     = "/product/shop"

	DEFAULT_RETENTION_TIME       = "1d"
	RRBEND_LOC                   = "/rrbend"
)
