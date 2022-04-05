package env

import (
	"os"
	"runtime"

	"github.com/google/uuid"
	"github.com/rs/xid"
)

var _ = Get

// Get 获取环境变量
func Get(key string) (env string) {
	switch key {
	case envHostname:
		env = hostname()
	case envOsType:
		env = runtime.GOOS
	case envUuid:
		env = uuid.NewString()
	case envXid:
		env = xid.New().String()
	default:
		env = os.Getenv(key)
	}

	return
}
