package main

import (
	"github.com/planx-lab/planx-plugin-processor-passthrough/internal/plugin"
	"github.com/planx-lab/planx-sdk-go/sdk"
)

func main() {
	sdk.ServeProcessor(plugin.New)
}
