// Code generated by genji.
// DO NOT EDIT!

package testdata

import (
	"errors"

	"github.com/asdine/genji"
	"github.com/asdine/genji/field"
	"github.com/asdine/genji/query"
	"github.com/asdine/genji/record"
	"github.com/asdine/genji/store"
	"github.com/asdine/genji/table"
)

// Field implements the field method of the record.Record interface.
func (b *Basic) Field(name string) (field.Field, error) {
	switch name {
	case "A":
		return field.Field{
			Name: "A",
			Type: field.String,
			Data: field.EncodeString(b.A),
		}, nil
	case "B":
		return field.Field{
			Name: "B",
			Type: field.Int,
			Data: field.EncodeInt(b.B),
		}, nil
	case "C":
		return field.Field{
			Name: "C",
			Type: field.Int32,
			Data: field.EncodeInt32(b.C),
		}, nil
	case "D":
		return field.Field{
			Name: "D",
			Type: field.Int32,
			Data: field.EncodeInt32(b.D),
		}, nil
	}

	return field.Field{}, errors.New("unknown field")
}

// Iterate through all the fields one by one and pass each of them to the given function.
// It the given function returns an error, the iteration is interrupted.
func (b *Basic) Iterate(fn func(field.Field) error) error {
	var err error
	var f field.Field

	f, _ = b.Field("A")
	err = fn(f)
	if err != nil {
		return err
	}

	f, _ = b.Field("B")
	err = fn(f)
	if err != nil {
		return err
	}

	f, _ = b.Field("C")
	err = fn(f)
	if err != nil {
		return err
	}

	f, _ = b.Field("D")
	err = fn(f)
	if err != nil {
		return err
	}

	return nil
}

// ScanRecord extracts fields from record and assigns them to the struct fields.
// It implements the record.Scanner interface.
func (b *Basic) ScanRecord(rec record.Record) error {
	return rec.Iterate(func(f field.Field) error {
		var err error

		switch f.Name {
		case "A":
			b.A, err = field.DecodeString(f.Data)
		case "B":
			b.B, err = field.DecodeInt(f.Data)
		case "C":
			b.C, err = field.DecodeInt32(f.Data)
		case "D":
			b.D, err = field.DecodeInt32(f.Data)
		}
		return err
	})
}

// BasicStore manages the table. It provides several typed helpers
// that simplify common operations.
type BasicStore struct {
	*store.Store
}

// NewBasicStore creates a BasicStore.
func NewBasicStore(db *genji.DB) *BasicStore {
	var schema *record.Schema

	var indexes []string

	return &BasicStore{Store: store.New(db, "Basic", schema, indexes)}
}

// NewBasicStoreWithTx creates a BasicStore valid for the lifetime of the given transaction.
func NewBasicStoreWithTx(tx *genji.Tx) *BasicStore {
	var schema *record.Schema

	var indexes []string

	return &BasicStore{Store: store.NewWithTx(tx, "Basic", schema, indexes)}
}

// Insert a record in the table and return the primary key.
func (b *BasicStore) Insert(record *Basic) (recordID []byte, err error) {
	return b.Store.Insert(record)
}

// Get a record using its primary key.
// If the record doesn't exist, returns table.ErrRecordNotFound.
func (b *BasicStore) Get(recordID []byte) (*Basic, error) {
	rec, err := b.Store.Get(recordID)
	if err != nil {
		return nil, err
	}

	if v, ok := rec.(*Basic); ok {
		return v, nil
	}

	var record Basic

	err = record.ScanRecord(rec)
	if err != nil {
		return nil, err
	}

	return &record, nil
}

// Delete a record using its primary key.
// If the record doesn't exist, returns table.ErrRecordNotFound.
func (b *BasicStore) Delete(recordID []byte) error {
	return b.Store.Delete(recordID)
}

// List records from the specified offset. If the limit is equal to -1, it returns all records after the selected offset.
func (b *BasicStore) List(offset, limit int) ([]Basic, error) {
	size := limit
	if size == -1 {
		size = 0
	}
	list := make([]Basic, 0, size)
	err := b.Store.List(offset, limit, func(recordID []byte, r record.Record) error {
		var record Basic
		err := record.ScanRecord(r)
		if err != nil {
			return err
		}
		list = append(list, record)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return list, nil
}

// Replace the selected record by the given one.
func (b *BasicStore) Replace(recordID []byte, record *Basic) error {
	return b.Store.Replace(recordID, record)
}

// BasicQuerySelector provides helpers for selecting fields from the Basic structure.
type BasicQuerySelector struct {
	A query.StringFieldSelector
	B query.IntFieldSelector
	C query.Int32FieldSelector
	D query.Int32FieldSelector
}

// NewBasicQuerySelector creates a BasicQuerySelector.
func NewBasicQuerySelector() BasicQuerySelector {
	return BasicQuerySelector{
		A: query.StringField("A"),
		B: query.IntField("B"),
		C: query.Int32Field("C"),
		D: query.Int32Field("D"),
	}
}

// Table returns a query.TableSelector for Basic.
func (*BasicQuerySelector) Table() query.TableSelector {
	return query.Table("Basic")
}

// All returns a list of all selectors for Basic.
func (s *BasicQuerySelector) All() []query.FieldSelector {
	return []query.FieldSelector{
		s.A,
		s.B,
		s.C,
		s.D,
	}
}

// BasicResult can be used to store the result of queries.
// Selected fields must map the Basic fields.
type BasicResult []Basic

// ScanTable iterates over table.Reader and stores all the records in the slice.
func (b *BasicResult) ScanTable(tr table.Reader) error {
	return tr.Iterate(func(_ []byte, r record.Record) error {
		var record Basic
		err := record.ScanRecord(r)
		if err != nil {
			return err
		}

		*b = append(*b, record)
		return nil
	})
}

// Field implements the field method of the record.Record interface.
func (b *basic) Field(name string) (field.Field, error) {
	switch name {
	case "A":
		return field.Field{
			Name: "A",
			Type: field.Bytes,
			Data: field.EncodeBytes(b.A),
		}, nil
	case "B":
		return field.Field{
			Name: "B",
			Type: field.Uint16,
			Data: field.EncodeUint16(b.B),
		}, nil
	case "C":
		return field.Field{
			Name: "C",
			Type: field.Float32,
			Data: field.EncodeFloat32(b.C),
		}, nil
	case "D":
		return field.Field{
			Name: "D",
			Type: field.Float32,
			Data: field.EncodeFloat32(b.D),
		}, nil
	}

	return field.Field{}, errors.New("unknown field")
}

// Iterate through all the fields one by one and pass each of them to the given function.
// It the given function returns an error, the iteration is interrupted.
func (b *basic) Iterate(fn func(field.Field) error) error {
	var err error
	var f field.Field

	f, _ = b.Field("A")
	err = fn(f)
	if err != nil {
		return err
	}

	f, _ = b.Field("B")
	err = fn(f)
	if err != nil {
		return err
	}

	f, _ = b.Field("C")
	err = fn(f)
	if err != nil {
		return err
	}

	f, _ = b.Field("D")
	err = fn(f)
	if err != nil {
		return err
	}

	return nil
}

// ScanRecord extracts fields from record and assigns them to the struct fields.
// It implements the record.Scanner interface.
func (b *basic) ScanRecord(rec record.Record) error {
	return rec.Iterate(func(f field.Field) error {
		var err error

		switch f.Name {
		case "A":
			b.A, err = field.DecodeBytes(f.Data)
		case "B":
			b.B, err = field.DecodeUint16(f.Data)
		case "C":
			b.C, err = field.DecodeFloat32(f.Data)
		case "D":
			b.D, err = field.DecodeFloat32(f.Data)
		}
		return err
	})
}

// basicStore manages the table. It provides several typed helpers
// that simplify common operations.
type basicStore struct {
	*store.Store
}

// newBasicStore creates a basicStore.
func newBasicStore(db *genji.DB) *basicStore {
	var schema *record.Schema

	var indexes []string

	return &basicStore{Store: store.New(db, "Basic", schema, indexes)}
}

// newBasicStoreWithTx creates a basicStore valid for the lifetime of the given transaction.
func newBasicStoreWithTx(tx *genji.Tx) *basicStore {
	var schema *record.Schema

	var indexes []string

	return &basicStore{Store: store.NewWithTx(tx, "Basic", schema, indexes)}
}

// Insert a record in the table and return the primary key.
func (b *basicStore) Insert(record *basic) (recordID []byte, err error) {
	return b.Store.Insert(record)
}

// Get a record using its primary key.
// If the record doesn't exist, returns table.ErrRecordNotFound.
func (b *basicStore) Get(recordID []byte) (*basic, error) {
	rec, err := b.Store.Get(recordID)
	if err != nil {
		return nil, err
	}

	if v, ok := rec.(*basic); ok {
		return v, nil
	}

	var record basic

	err = record.ScanRecord(rec)
	if err != nil {
		return nil, err
	}

	return &record, nil
}

// Delete a record using its primary key.
// If the record doesn't exist, returns table.ErrRecordNotFound.
func (b *basicStore) Delete(recordID []byte) error {
	return b.Store.Delete(recordID)
}

// List records from the specified offset. If the limit is equal to -1, it returns all records after the selected offset.
func (b *basicStore) List(offset, limit int) ([]basic, error) {
	size := limit
	if size == -1 {
		size = 0
	}
	list := make([]basic, 0, size)
	err := b.Store.List(offset, limit, func(recordID []byte, r record.Record) error {
		var record basic
		err := record.ScanRecord(r)
		if err != nil {
			return err
		}
		list = append(list, record)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return list, nil
}

// Replace the selected record by the given one.
func (b *basicStore) Replace(recordID []byte, record *basic) error {
	return b.Store.Replace(recordID, record)
}

// basicQuerySelector provides helpers for selecting fields from the basic structure.
type basicQuerySelector struct {
	A query.BytesFieldSelector
	B query.Uint16FieldSelector
	C query.Float32FieldSelector
	D query.Float32FieldSelector
}

// newbasicQuerySelector creates a basicQuerySelector.
func newBasicQuerySelector() basicQuerySelector {
	return basicQuerySelector{
		A: query.BytesField("A"),
		B: query.Uint16Field("B"),
		C: query.Float32Field("C"),
		D: query.Float32Field("D"),
	}
}

// Table returns a query.TableSelector for basic.
func (*basicQuerySelector) Table() query.TableSelector {
	return query.Table("basic")
}

// All returns a list of all selectors for basic.
func (s *basicQuerySelector) All() []query.FieldSelector {
	return []query.FieldSelector{
		s.A,
		s.B,
		s.C,
		s.D,
	}
}

// basicResult can be used to store the result of queries.
// Selected fields must map the basic fields.
type basicResult []basic

// ScanTable iterates over table.Reader and stores all the records in the slice.
func (b *basicResult) ScanTable(tr table.Reader) error {
	return tr.Iterate(func(_ []byte, r record.Record) error {
		var record basic
		err := record.ScanRecord(r)
		if err != nil {
			return err
		}

		*b = append(*b, record)
		return nil
	})
}

// Field implements the field method of the record.Record interface.
func (b *BasicSchemaful) Field(name string) (field.Field, error) {
	switch name {
	case "A":
		return field.Field{
			Name: "A",
			Type: field.String,
			Data: field.EncodeString(b.A),
		}, nil
	case "B":
		return field.Field{
			Name: "B",
			Type: field.Int,
			Data: field.EncodeInt(b.B),
		}, nil
	case "C":
		return field.Field{
			Name: "C",
			Type: field.Int32,
			Data: field.EncodeInt32(b.C),
		}, nil
	case "D":
		return field.Field{
			Name: "D",
			Type: field.Int32,
			Data: field.EncodeInt32(b.D),
		}, nil
	}

	return field.Field{}, errors.New("unknown field")
}

// Iterate through all the fields one by one and pass each of them to the given function.
// It the given function returns an error, the iteration is interrupted.
func (b *BasicSchemaful) Iterate(fn func(field.Field) error) error {
	var err error
	var f field.Field

	f, _ = b.Field("A")
	err = fn(f)
	if err != nil {
		return err
	}

	f, _ = b.Field("B")
	err = fn(f)
	if err != nil {
		return err
	}

	f, _ = b.Field("C")
	err = fn(f)
	if err != nil {
		return err
	}

	f, _ = b.Field("D")
	err = fn(f)
	if err != nil {
		return err
	}

	return nil
}

// ScanRecord extracts fields from record and assigns them to the struct fields.
// It implements the record.Scanner interface.
func (b *BasicSchemaful) ScanRecord(rec record.Record) error {
	return rec.Iterate(func(f field.Field) error {
		var err error

		switch f.Name {
		case "A":
			b.A, err = field.DecodeString(f.Data)
		case "B":
			b.B, err = field.DecodeInt(f.Data)
		case "C":
			b.C, err = field.DecodeInt32(f.Data)
		case "D":
			b.D, err = field.DecodeInt32(f.Data)
		}
		return err
	})
}

// BasicSchemafulStore manages the table. It provides several typed helpers
// that simplify common operations.
type BasicSchemafulStore struct {
	*store.Store
}

// NewBasicSchemafulStore creates a BasicSchemafulStore.
func NewBasicSchemafulStore(db *genji.DB) *BasicSchemafulStore {
	var schema *record.Schema
	schema = &record.Schema{
		Fields: []field.Field{
			{Name: "A", Type: field.String},
			{Name: "B", Type: field.Int},
			{Name: "C", Type: field.Int32},
			{Name: "D", Type: field.Int32},
		},
	}

	var indexes []string

	return &BasicSchemafulStore{Store: store.New(db, "BasicSchemaful", schema, indexes)}
}

// NewBasicSchemafulStoreWithTx creates a BasicSchemafulStore valid for the lifetime of the given transaction.
func NewBasicSchemafulStoreWithTx(tx *genji.Tx) *BasicSchemafulStore {
	var schema *record.Schema
	schema = &record.Schema{
		Fields: []field.Field{
			{Name: "A", Type: field.String},
			{Name: "B", Type: field.Int},
			{Name: "C", Type: field.Int32},
			{Name: "D", Type: field.Int32},
		},
	}

	var indexes []string

	return &BasicSchemafulStore{Store: store.NewWithTx(tx, "BasicSchemaful", schema, indexes)}
}

// Insert a record in the table and return the primary key.
func (b *BasicSchemafulStore) Insert(record *BasicSchemaful) (recordID []byte, err error) {
	return b.Store.Insert(record)
}

// Get a record using its primary key.
// If the record doesn't exist, returns table.ErrRecordNotFound.
func (b *BasicSchemafulStore) Get(recordID []byte) (*BasicSchemaful, error) {
	rec, err := b.Store.Get(recordID)
	if err != nil {
		return nil, err
	}

	if v, ok := rec.(*BasicSchemaful); ok {
		return v, nil
	}

	var record BasicSchemaful

	err = record.ScanRecord(rec)
	if err != nil {
		return nil, err
	}

	return &record, nil
}

// Delete a record using its primary key.
// If the record doesn't exist, returns table.ErrRecordNotFound.
func (b *BasicSchemafulStore) Delete(recordID []byte) error {
	return b.Store.Delete(recordID)
}

// List records from the specified offset. If the limit is equal to -1, it returns all records after the selected offset.
func (b *BasicSchemafulStore) List(offset, limit int) ([]BasicSchemaful, error) {
	size := limit
	if size == -1 {
		size = 0
	}
	list := make([]BasicSchemaful, 0, size)
	err := b.Store.List(offset, limit, func(recordID []byte, r record.Record) error {
		var record BasicSchemaful
		err := record.ScanRecord(r)
		if err != nil {
			return err
		}
		list = append(list, record)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return list, nil
}

// Replace the selected record by the given one.
func (b *BasicSchemafulStore) Replace(recordID []byte, record *BasicSchemaful) error {
	return b.Store.Replace(recordID, record)
}

// BasicSchemafulQuerySelector provides helpers for selecting fields from the BasicSchemaful structure.
type BasicSchemafulQuerySelector struct {
	A query.StringFieldSelector
	B query.IntFieldSelector
	C query.Int32FieldSelector
	D query.Int32FieldSelector
}

// NewBasicSchemafulQuerySelector creates a BasicSchemafulQuerySelector.
func NewBasicSchemafulQuerySelector() BasicSchemafulQuerySelector {
	return BasicSchemafulQuerySelector{
		A: query.StringField("A"),
		B: query.IntField("B"),
		C: query.Int32Field("C"),
		D: query.Int32Field("D"),
	}
}

// Table returns a query.TableSelector for BasicSchemaful.
func (*BasicSchemafulQuerySelector) Table() query.TableSelector {
	return query.Table("BasicSchemaful")
}

// All returns a list of all selectors for BasicSchemaful.
func (s *BasicSchemafulQuerySelector) All() []query.FieldSelector {
	return []query.FieldSelector{
		s.A,
		s.B,
		s.C,
		s.D,
	}
}

// BasicSchemafulResult can be used to store the result of queries.
// Selected fields must map the BasicSchemaful fields.
type BasicSchemafulResult []BasicSchemaful

// ScanTable iterates over table.Reader and stores all the records in the slice.
func (b *BasicSchemafulResult) ScanTable(tr table.Reader) error {
	return tr.Iterate(func(_ []byte, r record.Record) error {
		var record BasicSchemaful
		err := record.ScanRecord(r)
		if err != nil {
			return err
		}

		*b = append(*b, record)
		return nil
	})
}

// Field implements the field method of the record.Record interface.
func (p *Pk) Field(name string) (field.Field, error) {
	switch name {
	case "A":
		return field.Field{
			Name: "A",
			Type: field.String,
			Data: field.EncodeString(p.A),
		}, nil
	case "B":
		return field.Field{
			Name: "B",
			Type: field.Int64,
			Data: field.EncodeInt64(p.B),
		}, nil
	}

	return field.Field{}, errors.New("unknown field")
}

// Iterate through all the fields one by one and pass each of them to the given function.
// It the given function returns an error, the iteration is interrupted.
func (p *Pk) Iterate(fn func(field.Field) error) error {
	var err error
	var f field.Field

	f, _ = p.Field("A")
	err = fn(f)
	if err != nil {
		return err
	}

	f, _ = p.Field("B")
	err = fn(f)
	if err != nil {
		return err
	}

	return nil
}

// ScanRecord extracts fields from record and assigns them to the struct fields.
// It implements the record.Scanner interface.
func (p *Pk) ScanRecord(rec record.Record) error {
	return rec.Iterate(func(f field.Field) error {
		var err error

		switch f.Name {
		case "A":
			p.A, err = field.DecodeString(f.Data)
		case "B":
			p.B, err = field.DecodeInt64(f.Data)
		}
		return err
	})
}

// Pk returns the primary key. It implements the table.Pker interface.
func (p *Pk) Pk() ([]byte, error) {
	return field.EncodeInt64(p.B), nil
}

// PkStore manages the table. It provides several typed helpers
// that simplify common operations.
type PkStore struct {
	*store.Store
}

// NewPkStore creates a PkStore.
func NewPkStore(db *genji.DB) *PkStore {
	var schema *record.Schema

	var indexes []string

	return &PkStore{Store: store.New(db, "Pk", schema, indexes)}
}

// NewPkStoreWithTx creates a PkStore valid for the lifetime of the given transaction.
func NewPkStoreWithTx(tx *genji.Tx) *PkStore {
	var schema *record.Schema

	var indexes []string

	return &PkStore{Store: store.NewWithTx(tx, "Pk", schema, indexes)}
}

// Insert a record in the table and return the primary key.
func (p *PkStore) Insert(record *Pk) (err error) {
	_, err = p.Store.Insert(record)
	return err
}

// Get a record using its primary key.
// If the record doesn't exist, returns table.ErrRecordNotFound.
func (p *PkStore) Get(pk int64) (*Pk, error) {
	recordID := field.EncodeInt64(pk)
	rec, err := p.Store.Get(recordID)
	if err != nil {
		return nil, err
	}

	if v, ok := rec.(*Pk); ok {
		return v, nil
	}

	var record Pk

	err = record.ScanRecord(rec)
	if err != nil {
		return nil, err
	}

	return &record, nil
}

// Delete a record using its primary key.
// If the record doesn't exist, returns table.ErrRecordNotFound.
func (p *PkStore) Delete(pk int64) error {
	recordID := field.EncodeInt64(pk)
	return p.Store.Delete(recordID)
}

// List records from the specified offset. If the limit is equal to -1, it returns all records after the selected offset.
func (p *PkStore) List(offset, limit int) ([]Pk, error) {
	size := limit
	if size == -1 {
		size = 0
	}
	list := make([]Pk, 0, size)
	err := p.Store.List(offset, limit, func(recordID []byte, r record.Record) error {
		var record Pk
		err := record.ScanRecord(r)
		if err != nil {
			return err
		}
		list = append(list, record)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return list, nil
}

// Replace the selected record by the given one.
func (p *PkStore) Replace(pk int64, record *Pk) error {
	recordID := field.EncodeInt64(pk)
	if record.B != pk {
		record.B = pk
	}
	return p.Store.Replace(recordID, record)
}

// PkQuerySelector provides helpers for selecting fields from the Pk structure.
type PkQuerySelector struct {
	A query.StringFieldSelector
	B query.Int64FieldSelector
}

// NewPkQuerySelector creates a PkQuerySelector.
func NewPkQuerySelector() PkQuerySelector {
	return PkQuerySelector{
		A: query.StringField("A"),
		B: query.Int64Field("B"),
	}
}

// Table returns a query.TableSelector for Pk.
func (*PkQuerySelector) Table() query.TableSelector {
	return query.Table("Pk")
}

// All returns a list of all selectors for Pk.
func (s *PkQuerySelector) All() []query.FieldSelector {
	return []query.FieldSelector{
		s.A,
		s.B,
	}
}

// PkResult can be used to store the result of queries.
// Selected fields must map the Pk fields.
type PkResult []Pk

// ScanTable iterates over table.Reader and stores all the records in the slice.
func (p *PkResult) ScanTable(tr table.Reader) error {
	return tr.Iterate(func(_ []byte, r record.Record) error {
		var record Pk
		err := record.ScanRecord(r)
		if err != nil {
			return err
		}

		*p = append(*p, record)
		return nil
	})
}

// Field implements the field method of the record.Record interface.
func (i *Indexed) Field(name string) (field.Field, error) {
	switch name {
	case "A":
		return field.Field{
			Name: "A",
			Type: field.String,
			Data: field.EncodeString(i.A),
		}, nil
	case "B":
		return field.Field{
			Name: "B",
			Type: field.Int64,
			Data: field.EncodeInt64(i.B),
		}, nil
	}

	return field.Field{}, errors.New("unknown field")
}

// Iterate through all the fields one by one and pass each of them to the given function.
// It the given function returns an error, the iteration is interrupted.
func (i *Indexed) Iterate(fn func(field.Field) error) error {
	var err error
	var f field.Field

	f, _ = i.Field("A")
	err = fn(f)
	if err != nil {
		return err
	}

	f, _ = i.Field("B")
	err = fn(f)
	if err != nil {
		return err
	}

	return nil
}

// ScanRecord extracts fields from record and assigns them to the struct fields.
// It implements the record.Scanner interface.
func (i *Indexed) ScanRecord(rec record.Record) error {
	return rec.Iterate(func(f field.Field) error {
		var err error

		switch f.Name {
		case "A":
			i.A, err = field.DecodeString(f.Data)
		case "B":
			i.B, err = field.DecodeInt64(f.Data)
		}
		return err
	})
}

// IndexedStore manages the table. It provides several typed helpers
// that simplify common operations.
type IndexedStore struct {
	*store.Store
}

// NewIndexedStore creates a IndexedStore.
func NewIndexedStore(db *genji.DB) *IndexedStore {
	var schema *record.Schema

	var indexes []string
	indexes = append(indexes, "A")

	return &IndexedStore{Store: store.New(db, "Indexed", schema, indexes)}
}

// NewIndexedStoreWithTx creates a IndexedStore valid for the lifetime of the given transaction.
func NewIndexedStoreWithTx(tx *genji.Tx) *IndexedStore {
	var schema *record.Schema

	var indexes []string

	indexes = append(indexes, "A")

	return &IndexedStore{Store: store.NewWithTx(tx, "Indexed", schema, indexes)}
}

// Insert a record in the table and return the primary key.
func (i *IndexedStore) Insert(record *Indexed) (recordID []byte, err error) {
	return i.Store.Insert(record)
}

// Get a record using its primary key.
// If the record doesn't exist, returns table.ErrRecordNotFound.
func (i *IndexedStore) Get(recordID []byte) (*Indexed, error) {
	rec, err := i.Store.Get(recordID)
	if err != nil {
		return nil, err
	}

	if v, ok := rec.(*Indexed); ok {
		return v, nil
	}

	var record Indexed

	err = record.ScanRecord(rec)
	if err != nil {
		return nil, err
	}

	return &record, nil
}

// Delete a record using its primary key.
// If the record doesn't exist, returns table.ErrRecordNotFound.
func (i *IndexedStore) Delete(recordID []byte) error {
	return i.Store.Delete(recordID)
}

// List records from the specified offset. If the limit is equal to -1, it returns all records after the selected offset.
func (i *IndexedStore) List(offset, limit int) ([]Indexed, error) {
	size := limit
	if size == -1 {
		size = 0
	}
	list := make([]Indexed, 0, size)
	err := i.Store.List(offset, limit, func(recordID []byte, r record.Record) error {
		var record Indexed
		err := record.ScanRecord(r)
		if err != nil {
			return err
		}
		list = append(list, record)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return list, nil
}

// Replace the selected record by the given one.
func (i *IndexedStore) Replace(recordID []byte, record *Indexed) error {
	return i.Store.Replace(recordID, record)
}

// IndexedQuerySelector provides helpers for selecting fields from the Indexed structure.
type IndexedQuerySelector struct {
	A query.StringFieldSelector
	B query.Int64FieldSelector
}

// NewIndexedQuerySelector creates a IndexedQuerySelector.
func NewIndexedQuerySelector() IndexedQuerySelector {
	return IndexedQuerySelector{
		A: query.StringField("A"),
		B: query.Int64Field("B"),
	}
}

// Table returns a query.TableSelector for Indexed.
func (*IndexedQuerySelector) Table() query.TableSelector {
	return query.Table("Indexed")
}

// All returns a list of all selectors for Indexed.
func (s *IndexedQuerySelector) All() []query.FieldSelector {
	return []query.FieldSelector{
		s.A,
		s.B,
	}
}

// IndexedResult can be used to store the result of queries.
// Selected fields must map the Indexed fields.
type IndexedResult []Indexed

// ScanTable iterates over table.Reader and stores all the records in the slice.
func (i *IndexedResult) ScanTable(tr table.Reader) error {
	return tr.Iterate(func(_ []byte, r record.Record) error {
		var record Indexed
		err := record.ScanRecord(r)
		if err != nil {
			return err
		}

		*i = append(*i, record)
		return nil
	})
}

// Field implements the field method of the record.Record interface.
func (m *MultipleTags) Field(name string) (field.Field, error) {
	switch name {
	case "A":
		return field.Field{
			Name: "A",
			Type: field.String,
			Data: field.EncodeString(m.A),
		}, nil
	case "B":
		return field.Field{
			Name: "B",
			Type: field.Int64,
			Data: field.EncodeInt64(m.B),
		}, nil
	case "C":
		return field.Field{
			Name: "C",
			Type: field.Float32,
			Data: field.EncodeFloat32(m.C),
		}, nil
	case "D":
		return field.Field{
			Name: "D",
			Type: field.Bool,
			Data: field.EncodeBool(m.D),
		}, nil
	}

	return field.Field{}, errors.New("unknown field")
}

// Iterate through all the fields one by one and pass each of them to the given function.
// It the given function returns an error, the iteration is interrupted.
func (m *MultipleTags) Iterate(fn func(field.Field) error) error {
	var err error
	var f field.Field

	f, _ = m.Field("A")
	err = fn(f)
	if err != nil {
		return err
	}

	f, _ = m.Field("B")
	err = fn(f)
	if err != nil {
		return err
	}

	f, _ = m.Field("C")
	err = fn(f)
	if err != nil {
		return err
	}

	f, _ = m.Field("D")
	err = fn(f)
	if err != nil {
		return err
	}

	return nil
}

// ScanRecord extracts fields from record and assigns them to the struct fields.
// It implements the record.Scanner interface.
func (m *MultipleTags) ScanRecord(rec record.Record) error {
	return rec.Iterate(func(f field.Field) error {
		var err error

		switch f.Name {
		case "A":
			m.A, err = field.DecodeString(f.Data)
		case "B":
			m.B, err = field.DecodeInt64(f.Data)
		case "C":
			m.C, err = field.DecodeFloat32(f.Data)
		case "D":
			m.D, err = field.DecodeBool(f.Data)
		}
		return err
	})
}

// Pk returns the primary key. It implements the table.Pker interface.
func (m *MultipleTags) Pk() ([]byte, error) {
	return field.EncodeString(m.A), nil
}

// MultipleTagsStore manages the table. It provides several typed helpers
// that simplify common operations.
type MultipleTagsStore struct {
	*store.Store
}

// NewMultipleTagsStore creates a MultipleTagsStore.
func NewMultipleTagsStore(db *genji.DB) *MultipleTagsStore {
	var schema *record.Schema

	var indexes []string
	indexes = append(indexes, "D")

	return &MultipleTagsStore{Store: store.New(db, "MultipleTags", schema, indexes)}
}

// NewMultipleTagsStoreWithTx creates a MultipleTagsStore valid for the lifetime of the given transaction.
func NewMultipleTagsStoreWithTx(tx *genji.Tx) *MultipleTagsStore {
	var schema *record.Schema

	var indexes []string

	indexes = append(indexes, "D")

	return &MultipleTagsStore{Store: store.NewWithTx(tx, "MultipleTags", schema, indexes)}
}

// Insert a record in the table and return the primary key.
func (m *MultipleTagsStore) Insert(record *MultipleTags) (err error) {
	_, err = m.Store.Insert(record)
	return err
}

// Get a record using its primary key.
// If the record doesn't exist, returns table.ErrRecordNotFound.
func (m *MultipleTagsStore) Get(pk string) (*MultipleTags, error) {
	recordID := field.EncodeString(pk)
	rec, err := m.Store.Get(recordID)
	if err != nil {
		return nil, err
	}

	if v, ok := rec.(*MultipleTags); ok {
		return v, nil
	}

	var record MultipleTags

	err = record.ScanRecord(rec)
	if err != nil {
		return nil, err
	}

	return &record, nil
}

// Delete a record using its primary key.
// If the record doesn't exist, returns table.ErrRecordNotFound.
func (m *MultipleTagsStore) Delete(pk string) error {
	recordID := field.EncodeString(pk)
	return m.Store.Delete(recordID)
}

// List records from the specified offset. If the limit is equal to -1, it returns all records after the selected offset.
func (m *MultipleTagsStore) List(offset, limit int) ([]MultipleTags, error) {
	size := limit
	if size == -1 {
		size = 0
	}
	list := make([]MultipleTags, 0, size)
	err := m.Store.List(offset, limit, func(recordID []byte, r record.Record) error {
		var record MultipleTags
		err := record.ScanRecord(r)
		if err != nil {
			return err
		}
		list = append(list, record)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return list, nil
}

// Replace the selected record by the given one.
func (m *MultipleTagsStore) Replace(pk string, record *MultipleTags) error {
	recordID := field.EncodeString(pk)
	if record.A != pk {
		record.A = pk
	}
	return m.Store.Replace(recordID, record)
}

// MultipleTagsQuerySelector provides helpers for selecting fields from the MultipleTags structure.
type MultipleTagsQuerySelector struct {
	A query.StringFieldSelector
	B query.Int64FieldSelector
	C query.Float32FieldSelector
	D query.BoolFieldSelector
}

// NewMultipleTagsQuerySelector creates a MultipleTagsQuerySelector.
func NewMultipleTagsQuerySelector() MultipleTagsQuerySelector {
	return MultipleTagsQuerySelector{
		A: query.StringField("A"),
		B: query.Int64Field("B"),
		C: query.Float32Field("C"),
		D: query.BoolField("D"),
	}
}

// Table returns a query.TableSelector for MultipleTags.
func (*MultipleTagsQuerySelector) Table() query.TableSelector {
	return query.Table("MultipleTags")
}

// All returns a list of all selectors for MultipleTags.
func (s *MultipleTagsQuerySelector) All() []query.FieldSelector {
	return []query.FieldSelector{
		s.A,
		s.B,
		s.C,
		s.D,
	}
}

// MultipleTagsResult can be used to store the result of queries.
// Selected fields must map the MultipleTags fields.
type MultipleTagsResult []MultipleTags

// ScanTable iterates over table.Reader and stores all the records in the slice.
func (m *MultipleTagsResult) ScanTable(tr table.Reader) error {
	return tr.Iterate(func(_ []byte, r record.Record) error {
		var record MultipleTags
		err := record.ScanRecord(r)
		if err != nil {
			return err
		}

		*m = append(*m, record)
		return nil
	})
}

// ScanRecord extracts fields from record and assigns them to the struct fields.
// It implements the record.Scanner interface.
func (s *Sample) ScanRecord(rec record.Record) error {
	return rec.Iterate(func(f field.Field) error {
		var err error

		switch f.Name {
		case "A":
			s.A, err = field.DecodeString(f.Data)
		case "B":
			s.B, err = field.DecodeInt64(f.Data)
		}
		return err
	})
}

// SampleResult can be used to store the result of queries.
// Selected fields must map the Sample fields.
type SampleResult []Sample

// ScanTable iterates over table.Reader and stores all the records in the slice.
func (s *SampleResult) ScanTable(tr table.Reader) error {
	return tr.Iterate(func(_ []byte, r record.Record) error {
		var record Sample
		err := record.ScanRecord(r)
		if err != nil {
			return err
		}

		*s = append(*s, record)
		return nil
	})
}
