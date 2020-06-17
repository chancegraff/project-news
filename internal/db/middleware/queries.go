package middleware

import "github.com/jinzhu/gorm"

// WithLimit will return a store with a limit attached as a query
func WithLimit(a *gorm.DB, limit int) *gorm.DB {
	store := a
	if limit != 0 {
		store = store.Limit(limit)
	}
	return store
}

// WithOffset will return a store with an offset attached as a query
func WithOffset(a *gorm.DB, offset int) *gorm.DB {
	store := a
	if offset != 0 {
		store = store.Offset(offset)
	}
	return store
}
