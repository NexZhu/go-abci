package aceiclient

import (
	"fmt"

	"github.com/daotl/go-log/v2"
	gsync "github.com/daotl/guts/sync"

	"github.com/daotl/go-acei/types"
)

// Creator creates new ACEI clients.
type Creator func(log.StandardLogger) (Client, error)

// NewLocalCreator returns a Creator for the given app,
// which will be running locally.
func NewLocalCreator(logger log.StandardLogger, app types.Application) Creator {
	mtx := new(gsync.Mutex)

	return func(log.StandardLogger) (Client, error) {
		return NewLocalClient(logger, mtx, app)
	}
}

// NewRemoteCreator returns a Creator for the given address (e.g.
// "192.168.0.1") and transport (e.g. "tcp"). Set mustConnect to true if you
// want the client to connect before reporting success.
func NewRemoteCreator(logger log.StandardLogger, addr, transport string, mustConnect bool) Creator {
	return func(log.StandardLogger) (Client, error) {
		remoteApp, err := NewClient(logger, addr, transport, mustConnect)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to proxy: %w", err)
		}

		return remoteApp, nil
	}
}
