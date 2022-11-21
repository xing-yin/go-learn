package test

//go:generate mockgen -source=db.go -destination mock/mock_db.go -package=mock
type DB interface {
	Get(key string) (int, error)
}

// 方式1：源码模式生成。mockgen -source test/db.go -destination mock/mock_db.go
// 方式2：通过注释使用。

func GetFromDB(db DB, key string) int {
	if value, err := db.Get(key); err == nil {
		return value
	}
	return -1
}
