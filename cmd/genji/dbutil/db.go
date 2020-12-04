package dbutil

import (
	"context"
	"fmt"

	"github.com/genjidb/genji"
	"github.com/genjidb/genji/engine"
	"github.com/genjidb/genji/engine/memoryengine"
)

type DBOptions struct {
	EncryptionKey string
}

// OpenDB opens a database at the given path, using the selected engine.
func OpenDB(ctx context.Context, dbPath, engineName string, opts DBOptions) (*genji.DB, error) {
	var (
		ng  engine.Engine
		err error
	)

	switch engineName {
	case "memory":
		ng = memoryengine.NewEngine()
	default:
		return nil, fmt.Errorf(`engine unknown, got %q`, engineName)
	}
	if err != nil {
		return nil, err
	}

	return genji.New(ctx, ng)
}
