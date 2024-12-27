package transport

import (
	"four/modules"
)

type baseHandler struct {
	db *modules.DB
}

func NewBaseHandler(db *modules.DB) *baseHandler {
	return &baseHandler{
		db: db,
	}
}

func NewBaseHandlerWithTableCustomers(db *modules.DB) *baseHandler {
	db.CreateTableCustomers()
	return &baseHandler{
		db: db,
	}
}