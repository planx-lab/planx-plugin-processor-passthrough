package plugin

import (
	"context"
	"reflect"
	"testing"

	"github.com/planx-lab/planx-sdk-go/sdk"
)

func TestNewReturnsProcessorSPI(t *testing.T) {
	p := New()

	if p == nil {
		t.Fatal("New() returned nil, expected non-nil sdk.ProcessorSPI")
	}

	// Verify it implements the interface.
	var _ sdk.ProcessorSPI = p
}

func TestInitReturnsNil(t *testing.T) {
	p := New()

	err := p.Init(context.Background(), nil)
	if err != nil {
		t.Fatalf("Init(nil) returned error: %v", err)
	}
}

func TestProcessReturnsBatchUnchanged(t *testing.T) {
	p := New()
	_ = p.Init(context.Background(), nil)

	input := map[string]string{"key": "value", "foo": "bar"}
	output, err := p.Process(input)
	if err != nil {
		t.Fatalf("Process() returned error: %v", err)
	}

	// Passthrough must return the exact same batch.
	if !reflect.DeepEqual(output, input) {
		t.Errorf("Process() returned a different batch: got %v, want %v", output, input)
	}
}

func TestProcessPassthroughWithNilBatch(t *testing.T) {
	p := New()
	_ = p.Init(context.Background(), nil)

	output, err := p.Process(nil)
	if err != nil {
		t.Fatalf("Process(nil) returned error: %v", err)
	}
	if output != nil {
		t.Errorf("Process(nil) expected nil, got %v", output)
	}
}

func TestCloseReturnsNil(t *testing.T) {
	p := New()

	err := p.Close()
	if err != nil {
		t.Fatalf("Close() returned error: %v", err)
	}
}
