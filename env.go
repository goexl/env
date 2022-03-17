package env

import (
	`os`

	`github.com/google/uuid`
	`github.com/rs/xid`
)

var _ = Get

// Get 获取环境变量
func Get(key string) (env string) {
	switch key {
	case envHostname:
		env = hostname()
	case envUuid:
		env = uuid.NewString()
	case envXid:
		env = xid.New().String()
	default:
		env = os.Getenv(key)
	}

	return
}
