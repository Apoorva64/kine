package redis

import (
	"context"

	"github.com/k3s-io/kine/pkg/server"
	"github.com/k3s-io/kine/pkg/tls"
)

type Config struct {
	//Redis server hostname
	serverHostname string
	//Redis server port
	serverPort string
	//Redis server password
	serverPassword string
	//Redis server database
	serverDatabase int64
}

func New(ctx context.Context, connection string, tlsInfo tls.Config) (server.Backend, error) {
	return newBackend(ctx, connection, tlsInfo, false)
}

// NewLegacy return an implementation of server.Backend using NATS + JetStream
// with legacy jetstream:// behavior, ignoring the embedded server.
func NewLegacy(ctx context.Context, connection string, tlsInfo tls.Config) (server.Backend, error) {
	return newBackend(ctx, connection, tlsInfo, true)
}

func newBackend(ctx context.Context, connection string, tlsInfo tls.Config, legacy bool) (server.Backend, error) {
}
