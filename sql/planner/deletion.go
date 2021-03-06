package planner

import (
	"errors"
	"fmt"

	"github.com/genjidb/genji/database"
	"github.com/genjidb/genji/document"
	"github.com/genjidb/genji/sql/query/expr"
)

// deleteBufferSize is the size of the buffer used to delete documents.
const deleteBufferSize = 100

type deletionNode struct {
	node

	tableName string
	table     *database.Table
}

var _ operationNode = (*deletionNode)(nil)

// NewDeletionNode creates a node that delete every document of a stream
// from their respective table.
func NewDeletionNode(n Node, tableName string) Node {
	return &deletionNode{
		node: node{
			op:   Deletion,
			left: n,
		},
		tableName: tableName,
	}
}

func (n *deletionNode) Bind(tx *database.Transaction, params []expr.Param) (err error) {
	n.table, err = tx.GetTable(n.tableName)
	return
}

// toResult deletes matching documents by batches of deleteBufferSize documents.
// Some engines can't iterate while deleting keys (https://github.com/etcd-io/bbolt/issues/146)
// and some can't create more than one iterator per read-write transaction (https://github.com/dgraph-io/badger/issues/1093).
// To deal with these limitations, Run will iterate on a limited number of documents, copy the keys
// to a buffer and delete them after the iteration is complete, and it will do that until there is no document
// left to delete.
// Increasing deleteBufferSize will occasionate less key searches (O(log n) for most engines) but will take more memory.
func (n *deletionNode) toStream(st document.Stream) (document.Stream, error) {
	st = st.Limit(deleteBufferSize)

	keys := make([][]byte, deleteBufferSize)

	for {
		var i int

		err := st.Iterate(func(d document.Document) error {
			k, ok := d.(document.Keyer)
			if !ok {
				return errors.New("attempt to delete document without key")
			}
			// copy the key and reuse the buffer
			keys[i] = append(keys[i][0:0], k.Key()...)
			i++
			return nil
		})
		if err != nil {
			return document.Stream{}, err
		}

		keys = keys[:i]

		for _, key := range keys {
			err = n.table.Delete(key)
			if err != nil {
				return document.Stream{}, err
			}
		}

		if i < deleteBufferSize {
			break
		}
	}

	return document.Stream{}, nil
}

func (n *deletionNode) String() string {
	return fmt.Sprintf("Delete(%s)", n.tableName)
}
