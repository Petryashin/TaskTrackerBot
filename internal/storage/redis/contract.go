package rediscache

type redisCacheInterface interface {
	Set(key string, json string) error
	Get(key string) (string, error)
}
