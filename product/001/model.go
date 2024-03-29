package main

type KfPerson struct {
	Id     int    `db:"id"`
	Name   string `db:"name"`
	Idcard string `db:"idcard"`
}

//缓存结果(存入时间和查询频次)
type QueryResult struct {
	Value     []KfPerson
	CacheTime int64
	Count     int
}

func (qr *QueryResult) GetCacheTime() int64 {
	return qr.CacheTime
}
