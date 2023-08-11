package helpers

import "database/sql"

func ConvertUintToPointer(num uint) *uint {
	return &num
}

func CreateSqlNullInt64(num uint) sql.NullInt64 {
	if num != 0 {
		return sql.NullInt64{
			Int64: int64(num),
			Valid: true,
		}
	} else {
		return sql.NullInt64{
			Valid: false,
		}
	}
}
