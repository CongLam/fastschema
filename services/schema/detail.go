package schemaservice

import (
	"github.com/fastschema/fastschema/app"
	"github.com/fastschema/fastschema/schema"
)

func (ss *SchemaService) Detail(c app.Context, _ *any) (*schema.Schema, error) {
	s, err := ss.app.SchemaBuilder().Schema(c.Arg("name"))
	if err != nil {
		return nil, err
	}

	return s, nil
}