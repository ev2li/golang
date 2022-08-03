package main

import "time"

type TimedData interface {
	GetCacheTime() int64
}

//整理缓存,删除加入最早的
func UpdateCache(cacheMap *map[string]TimedData) (delkey string) {
	earliestTime := time.Now().UnixNano()
	for key, value := range *cacheMap {
		if value.GetCacheTime() < earliestTime {
			earliestTime = value.GetCacheTime()
			delkey = key
		}
	}
	delete(*cacheMap, delkey)
	return delkey
}
