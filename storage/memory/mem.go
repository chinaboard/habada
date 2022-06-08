package memory

import (
	"github.com/allegro/bigcache/v3"
	"time"
)

type MemoryStorage struct {
	c *bigcache.BigCache
}

func New() *MemoryStorage {
	cacheCfg := bigcache.DefaultConfig(24 * time.Hour)
	cacheCfg.MaxEntrySize = 10000
	cache, _ := bigcache.NewBigCache(cacheCfg)
	return &MemoryStorage{
		c: cache,
	}

}

func (m *MemoryStorage) Get(tinyUrl string) (longUrl string, err error) {
	b, err := m.c.Get(tinyUrl)
	if b == nil || err != nil {
		return "", err
	}
	return string(b), nil
}

func (m *MemoryStorage) Set(tinyUrl, longUrl string) (success bool, err error) {
	err = m.c.Set(tinyUrl, []byte(longUrl))
	return err == nil, err
}
