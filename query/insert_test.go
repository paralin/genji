package query

import (
	"testing"

	"github.com/asdine/genji/record"
	"github.com/stretchr/testify/require"
)

func TestInsertStatement(t *testing.T) {
	t.Run("NoFields", func(t *testing.T) {
		tx, cleanup := createTable(t, 10, false)
		defer cleanup()

		res := Insert().Into(Table("test")).Values(IntValue(5), StringValue("hello"), IntValue(50), IntValue(5)).Exec(tx)
		require.Error(t, res.Err())
	})

	t.Run("WithFields", func(t *testing.T) {
		tx, cleanup := createTable(t, 10, false)
		defer cleanup()

		res := Insert().Into(Table("test")).Fields("a", "b").Values(IntValue(5), StringValue("hello")).Exec(tx)
		require.NoError(t, res.Err())

		tb, err := tx.GetTable("test")
		require.NoError(t, err)

		st := record.NewStream(tb)
		count, err := st.Count()
		require.NoError(t, err)
		require.Equal(t, 11, count)
	})
}
