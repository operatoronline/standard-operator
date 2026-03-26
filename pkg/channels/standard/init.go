package standard

import (
	"github.com/standardws/operator/pkg/bus"
	"github.com/standardws/operator/pkg/channels"
	"github.com/standardws/operator/pkg/config"
)

func init() {
	channels.RegisterFactory("standard", func(cfg *config.Config, b *bus.MessageBus) (channels.Channel, error) {
		return NewStandardChannel(cfg.Channels.Standard, b)
	})
}
