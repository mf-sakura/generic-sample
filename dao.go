package main

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

type Entity interface {
	// some method needed for all entity
}

type RDBEntity interface {
	Entity
	sql.Scanner
	driver.Valuer
}

type JsonEntity interface {
	json.Marshaler
	json.Unmarshaler
}

type DAO[E Entity, K comparable] interface {
	Save(e E) error
	Get(key K) (E, error)
	Delete(key K) error
	Update(key K, e E) error
}

type MySQLDAO[RE RDBEntity, K comparable] interface {
	DAO[RE, K]
}

type JsonDAO[JE JsonEntity, K comparable] interface {
	DAO[JE, K]
}
