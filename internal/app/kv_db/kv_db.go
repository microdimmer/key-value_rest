package kv_db

import (
	"sync"
)

type DataMap struct {
	mu   sync.RWMutex
	data map[string]string
}

func Create() *DataMap {
	db := &DataMap{
		data: make(map[string]string),
	}
	return db
}

func (d *DataMap) Get(key string) (string, bool) {
	d.mu.RLock()
	val, ok := d.data[key]
	d.mu.Unlock()
	return val, ok
}

func (d *DataMap) Set(key string, value string) {
	d.mu.Lock()
	d.data[key] = value
	d.mu.Unlock()
}

func (d *DataMap) List(key string, value string) map[string]string {
	return d.data
}

func (d *DataMap) Delete(key string, value string) bool {
	d.mu.Lock()
	_, ok := d.data[key]
	if ok {
		delete(d.data, key)
	}
	d.mu.Unlock()
	return ok
}
