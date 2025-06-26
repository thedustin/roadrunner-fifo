package fifo

import (
	"context"
	"github.com/roadrunner-server/errors"
	"github.com/thedustin/roadrunner-fifo/impl"
	"go.uber.org/zap"
)

const (
	pluginName = "fifo"
)

type Configurer interface {
	// UnmarshalKey takes a single key and unmarshal it into a Struct.
	UnmarshalKey(name string, out any) error
	// Has checks if a config section exists.
	Has(name string) bool
}

type Logger interface {
	NamedLogger(name string) *zap.Logger
}

type Plugin struct {
	config *Config

	log *zap.Logger
}

func (p *Plugin) Init(cfg Configurer, log Logger) error {
	const op = errors.Op("fifo_plugin_init")

	if !cfg.Has(pluginName) {
		return errors.E(errors.Disabled)
	}

	p.log = log.NamedLogger(pluginName)

	err := cfg.UnmarshalKey(pluginName, &p.config)
	if err != nil {
		return errors.E(op, err)
	}

	err = p.config.InitDefaults()
	if err != nil {
		return errors.E(op, err)
	}

	return nil
}

func (p *Plugin) Serve() chan error {
	errCh := make(chan error, 1)
	const op = errors.Op("fifo_plugin_serve")

	// Here you would typically start your plugin's main logic
	// For example, starting a server or a worker that processes tasks
	go func() {
		// Simulate some work
		errCh <- nil // or send an error if something goes wrong
	}()

	return errCh
}

func (p *Plugin) Weight() uint {
	return 10
}

func (p *Plugin) Stop(ctx context.Context) error {
	return nil
}

func (p *Plugin) Name() string {
	return pluginName
}

// RPC returns associated rpc service.
func (p *Plugin) RPC() any {
	return &rpc{
		fifo: impl.NewOtterImpl(
			p.config.Ttl,
			p.config.MaxCacheSize,
		),
	}
}
