package storage

type Storage interface {
	Get(tinyUrl string) (longUrl string, err error)
	Set(tinyUrl, longUrl string) (success bool, err error)
}
