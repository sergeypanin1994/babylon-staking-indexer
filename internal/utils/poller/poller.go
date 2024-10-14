package poller

import (
	"context"
	"go/types"
	"time"

	"github.com/rs/zerolog/log"
)

type Poller struct {
	interval   time.Duration
	quit       chan struct{}
	pollMethod func(ctx context.Context) *types.Error
}

func NewPoller(interval time.Duration, pollMethod func(ctx context.Context) *types.Error) *Poller {
	return &Poller{
		interval:   interval,
		quit:       make(chan struct{}),
		pollMethod: pollMethod,
	}
}

func (p *Poller) Start(ctx context.Context) {
	ticker := time.NewTicker(p.interval)

	for {
		select {
		case <-ticker.C:
			if err := p.pollMethod(ctx); err != nil {
				log.Error().Err(err).Msg("Error polling")
			}
		case <-ctx.Done():
			// Handle context cancellation.
			log.Info().Msg("Poller stopped due to context cancellation")
			return
		case <-p.quit:
			ticker.Stop() // Stop the ticker
			return
		}
	}
}

func (p *Poller) Stop() {
	close(p.quit)
}