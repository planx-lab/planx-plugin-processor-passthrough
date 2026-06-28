package main

import (
	"github.com/planx-lab/planx-plugin-processor-passthrough/internal/plugin"
	"github.com/planx-lab/planx-sdk-go/sdk"
)

func main() {
	sdk.Serve(sdk.Plugin{
		ID:          "processor-passthrough",
		Version:     "1.0.0",
		DisplayName: "Passthrough Processor",
		Description: "Passthrough processor for testing and demonstration",
		Summary:     "Returns each batch unchanged (1:1).",
		Components: []sdk.ComponentSpec{{
			ID:          "processor",
			Kind:        sdk.KindProcessor,
			DisplayName: "Passthrough",
			Processor:   plugin.New,
		}},
	})
}
