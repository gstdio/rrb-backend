package db

const (
	CLASS_GETALL_SQL    = "select * from classes"
	SUBCLASS_GETALL_SQL = "select * from subclasses"
	SUBCLASS_GETBY_CLASSID_SQL = "select * from subclasses where class_id = ?"
)
