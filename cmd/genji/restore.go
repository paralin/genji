package main

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"os"

	"github.com/dgraph-io/badger/v2"
	"github.com/genjidb/genji"
	"github.com/genjidb/genji/engine"
	"github.com/genjidb/genji/engine/badgerengine"
	"github.com/genjidb/genji/engine/boltengine"
)

func runRestoreCommand(ctx context.Context, rd io.Reader, ngine, table, dbPath string) error {
	// If the database file doesn't exist we take into account
	// the engine flag to create the new database. Otherwise we only
	// check the type of the database file to determine the engine to use.
	fileinfo, err := os.Stat(dbPath)
	switch {
	case os.IsNotExist(err): // we use the engine specified in the flag
	case fileinfo.IsDir():
		ngine = "badger"
	default:
		ngine = "bolt"
	}

	var ng engine.Engine
	switch ngine {
	case "bolt":
		ng, err = boltengine.NewEngine(dbPath, 0660, nil)
	case "badger":
		ng, err = badgerengine.NewEngine(badger.DefaultOptions(dbPath).WithLogger(nil))
	}
	if err != nil {
		return err
	}

	db, err := genji.New(ctx, ng)
	if err != nil {
		return err
	}
	defer db.Close()

	tx, err := db.Begin(true)
	if err != nil {
		return err
	}

	r := bufio.NewReader(rd)
	for {
		q, err := r.ReadBytes(';')
		if err != nil {
			if err == io.EOF {
				break
			}
			tx.Rollback()
			return err
		}

		// Ignore "BEGIN TRANSACTION" and "COMMIT" statements.
		qu := bytes.ToUpper(q)
		if bytes.Contains(qu, []byte("TRANSACTION")) || bytes.Contains(qu, []byte("COMMIT")) {
			continue
		}

		err = tx.Exec(string(q))
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return nil
}
