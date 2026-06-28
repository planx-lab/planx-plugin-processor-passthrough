package plugin

import (
	"context"

	"github.com/planx-lab/planx-sdk-go/sdk"
)

type Processor struct{}

func New() sdk.ProcessorSPI {
	return &Processor{}
}

func (p *Processor) Init(_ context.Context, _ []byte) error {
	return nil
}

func (p *Processor) Process(batch sdk.Batch) (sdk.Batch, error) {
	return batch, nil
}

func (p *Processor) Close() error {
	return nil
}
