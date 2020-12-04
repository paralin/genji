// +build !wasm

package genji

import (
	"context"
	"errors"

	"github.com/genjidb/genji/engine"
	"github.com/genjidb/genji/engine/memoryengine"
)

// Open creates a Genji database at the given path.
// If path is equal to ":memory:" it will open an in-memory database,
// otherwise it will create an on-disk database using the BoltDB engine.
func Open(path string) (*DB, error) {
	var ng engine.Engine
	var err error

	switch path {
	case ":memory:":
		ng = memoryengine.NewEngine()
	default:
		err = errors.New("unknown engine")
	}
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	return New(ctx, ng)
}
