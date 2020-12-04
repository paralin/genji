package dbutil

import (
	"context"

	"github.com/genjidb/genji"
	"github.com/genjidb/genji/engine"
	"github.com/genjidb/genji/engine/memoryengine"
	"github.com/genjidb/genji/stringutil"
)

// OpenDB opens a database at the given path, using the selected engine.
func OpenDB(ctx context.Context, dbPath, engineName string) (*genji.DB, error) {
	var (
		ng  engine.Engine
		err error
	)

	switch engineName {
	case "memory":
		ng = memoryengine.NewEngine()
	default:
		return nil, stringutil.Errorf(`engine should be "memory", got %q`, engineName)
	}
	if err != nil {
		return nil, err
	}

	return genji.New(ctx, ng)
}
