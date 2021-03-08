package tests

import "github.com/gcarrenolg/go-kallax"

type ResultSetFixture struct {
	kallax.Model `table:"resultset"`
	ID           kallax.ULID `pk:""`
	Foo          string
}

func newResultSetFixture(f string) *ResultSetFixture {
	return &ResultSetFixture{ID: kallax.NewULID(), Foo: f}
}
