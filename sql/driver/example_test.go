package driver_test

import (
	"context"
	"database/sql"
	sdriver "database/sql/driver"
	"fmt"
	"log"

	"github.com/genjidb/genji"
	"github.com/genjidb/genji/engine/memoryengine"
	"github.com/genjidb/genji/sql/driver"
)

type User struct {
	ID   int64
	Name string
	Age  uint32
}

func Example() {
	ctx := context.Background()
	mem := memoryengine.NewEngine()
	gdb2, err := genji.New(ctx, mem)
	if err != nil {
		log.Fatal(err)
	}
	drv := driver.NewDriver(gdb2)
	oc := drv.(interface {
		OpenConnector(name string) (sdriver.Connector, error)
	})
	conn, err := oc.OpenConnector("")
	if err != nil {
		log.Fatal(err)
	}

	// Create a database instance, here we'll store everything in memory
	db := sql.OpenDB(conn)
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS user")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE INDEX IF NOT EXISTS idx_user_name ON user (name)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT INTO user (id, name, age) VALUES (?, ?, ?)", 10, "foo", 15)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT INTO user VALUES ?, ?", &User{ID: 1, Name: "bar", Age: 100}, &User{ID: 2, Name: "baz"})
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM user WHERE name = ?", "bar")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var u User
		err = rows.Scan(driver.Scanner(&u))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(u)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// Output: {1 bar 100}
}
