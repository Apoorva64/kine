package redis

import (
	"context"
	"time"

	"github.com/k3s-io/kine/pkg/server"
	"github.com/k3s-io/kine/pkg/tls"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type Driver struct {
	red           *redis.Client
	slowThreshold time.Duration
}

func (d *Driver) logMethod(dur time.Duration, str string, args ...any) {
	if dur > d.slowThreshold {
		logrus.Warnf(str, args...)
	} else {
		logrus.Tracef(str, args...)
	}
}

func New(ctx context.Context, connection string, tlsInfo tls.Config) (server.Backend, error) {
	return newBackend(ctx, connection, tlsInfo, false)
}

func NewLegacy(ctx context.Context, connection string, tlsInfo tls.Config) (server.Backend, error) {
	return newBackend(ctx, connection, tlsInfo, true)
}

func newBackend(ctx context.Context, connection string, tlsInfo tls.Config, legacy bool) (server.Backend, error) {
	opts, err := redis.ParseURL(connection)
	if err != nil {
		return nil, err
	}
	rdb := redis.NewClient(opts)

	return &Driver{
		red:           rdb,
		slowThreshold: time.Duration(500) * time.Millisecond, //500ms
	}, nil
}

func (d *Driver) Start(ctx context.Context) error {
	if _, err := d.Create(ctx, "/registry/health", []byte(`{"health":"true"}`), 0); err != nil {
		if err != server.ErrKeyExists {
			logrus.Errorf("Failed to create health check key: %v", err)
		}
	}
	return nil
}

// Count implements server.Backend.
func (*Driver) Count(ctx context.Context, prefix string) (int64, int64, error) {
	panic("unimplemented")
}

// Create implements server.Backend.
func (*Driver) Create(ctx context.Context, key string, value []byte, lease int64) (int64, error) {
	panic("unimplemented")
}

// DbSize implements server.Backend.
func (*Driver) DbSize(ctx context.Context) (int64, error) {
	panic("unimplemented")
}

// Delete implements server.Backend.
func (*Driver) Delete(ctx context.Context, key string, revision int64) (int64, *server.KeyValue, bool, error) {
	panic("unimplemented")
}

// Get implements server.Backend.
func (*Driver) Get(ctx context.Context, key string, rangeEnd string, limit int64, revision int64) (int64, *server.KeyValue, error) {
	panic("unimplemented")
}

// List implements server.Backend.
func (*Driver) List(ctx context.Context, prefix string, startKey string, limit int64, revision int64) (int64, []*server.KeyValue, error) {
	panic("unimplemented")
}

// Update implements server.Backend.
func (*Driver) Update(ctx context.Context, key string, value []byte, revision int64, lease int64) (int64, *server.KeyValue, bool, error) {
	panic("unimplemented")
}

// Watch implements server.Backend.
func (*Driver) Watch(ctx context.Context, key string, revision int64) <-chan []*server.Event {
	panic("unimplemented")
}
