package namespace

import (
	"context"

	"github.com/pingcap-incubator/weir/pkg/proxy/driver"
)

type Frontend interface {
	Auth(username string, passwdBytes []byte, salt []byte) bool
	IsDatabaseAllowed(db string) bool
	ListDatabases() []string
}

type Backend interface {
	Close()
	GetPooledConn(context.Context) (driver.PooledBackendConn, error)
}
