// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/gen0cide/laforge/ent/fileextract"
)

// FileExtract is the model entity for the FileExtract schema.
type FileExtract struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Source holds the value of the "source" field.
	Source string `json:"source,omitempty"`
	// Destination holds the value of the "destination" field.
	Destination string `json:"destination,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FileExtractQuery when eager-loading is set.
	Edges                          FileExtractEdges `json:"edges"`
	provisioning_step_file_extract *int
}

// FileExtractEdges holds the relations/edges for other nodes in the graph.
type FileExtractEdges struct {
	// Tag holds the value of the tag edge.
	Tag []*Tag
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TagOrErr returns the Tag value or an error if the edge
// was not loaded in eager-loading.
func (e FileExtractEdges) TagOrErr() ([]*Tag, error) {
	if e.loadedTypes[0] {
		return e.Tag, nil
	}
	return nil, &NotLoadedError{edge: "tag"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FileExtract) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // source
		&sql.NullString{}, // destination
		&sql.NullString{}, // type
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*FileExtract) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // provisioning_step_file_extract
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FileExtract fields.
func (fe *FileExtract) assignValues(values ...interface{}) error {
	if m, n := len(values), len(fileextract.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	fe.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field source", values[0])
	} else if value.Valid {
		fe.Source = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field destination", values[1])
	} else if value.Valid {
		fe.Destination = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field type", values[2])
	} else if value.Valid {
		fe.Type = value.String
	}
	values = values[3:]
	if len(values) == len(fileextract.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field provisioning_step_file_extract", value)
		} else if value.Valid {
			fe.provisioning_step_file_extract = new(int)
			*fe.provisioning_step_file_extract = int(value.Int64)
		}
	}
	return nil
}

// QueryTag queries the tag edge of the FileExtract.
func (fe *FileExtract) QueryTag() *TagQuery {
	return (&FileExtractClient{config: fe.config}).QueryTag(fe)
}

// Update returns a builder for updating this FileExtract.
// Note that, you need to call FileExtract.Unwrap() before calling this method, if this FileExtract
// was returned from a transaction, and the transaction was committed or rolled back.
func (fe *FileExtract) Update() *FileExtractUpdateOne {
	return (&FileExtractClient{config: fe.config}).UpdateOne(fe)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (fe *FileExtract) Unwrap() *FileExtract {
	tx, ok := fe.config.driver.(*txDriver)
	if !ok {
		panic("ent: FileExtract is not a transactional entity")
	}
	fe.config.driver = tx.drv
	return fe
}

// String implements the fmt.Stringer.
func (fe *FileExtract) String() string {
	var builder strings.Builder
	builder.WriteString("FileExtract(")
	builder.WriteString(fmt.Sprintf("id=%v", fe.ID))
	builder.WriteString(", source=")
	builder.WriteString(fe.Source)
	builder.WriteString(", destination=")
	builder.WriteString(fe.Destination)
	builder.WriteString(", type=")
	builder.WriteString(fe.Type)
	builder.WriteByte(')')
	return builder.String()
}

// FileExtracts is a parsable slice of FileExtract.
type FileExtracts []*FileExtract

func (fe FileExtracts) config(cfg config) {
	for _i := range fe {
		fe[_i].config = cfg
	}
}
