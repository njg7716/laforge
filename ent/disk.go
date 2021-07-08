// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/gen0cide/laforge/ent/disk"
	"github.com/gen0cide/laforge/ent/host"
	"github.com/google/uuid"
)

// Disk is the model entity for the Disk schema.
type Disk struct {
	config ` json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Size holds the value of the "size" field.
	Size int `json:"size,omitempty" hcl:"size,attr"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DiskQuery when eager-loading is set.
	Edges DiskEdges `json:"edges"`

	// Edges put into the main struct to be loaded via hcl
	// DiskToHost holds the value of the DiskToHost edge.
	HCLDiskToHost *Host `json:"DiskToHost,omitempty"`
	//
	host_host_to_disk *uuid.UUID
}

// DiskEdges holds the relations/edges for other nodes in the graph.
type DiskEdges struct {
	// DiskToHost holds the value of the DiskToHost edge.
	DiskToHost *Host `json:"DiskToHost,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// DiskToHostOrErr returns the DiskToHost value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DiskEdges) DiskToHostOrErr() (*Host, error) {
	if e.loadedTypes[0] {
		if e.DiskToHost == nil {
			// The edge DiskToHost was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: host.Label}
		}
		return e.DiskToHost, nil
	}
	return nil, &NotLoadedError{edge: "DiskToHost"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Disk) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case disk.FieldSize:
			values[i] = new(sql.NullInt64)
		case disk.FieldID:
			values[i] = new(uuid.UUID)
		case disk.ForeignKeys[0]: // host_host_to_disk
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Disk", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Disk fields.
func (d *Disk) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case disk.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				d.ID = *value
			}
		case disk.FieldSize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field size", values[i])
			} else if value.Valid {
				d.Size = int(value.Int64)
			}
		case disk.ForeignKeys[0]:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field host_host_to_disk", values[i])
			} else if value != nil {
				d.host_host_to_disk = value
			}
		}
	}
	return nil
}

// QueryDiskToHost queries the "DiskToHost" edge of the Disk entity.
func (d *Disk) QueryDiskToHost() *HostQuery {
	return (&DiskClient{config: d.config}).QueryDiskToHost(d)
}

// Update returns a builder for updating this Disk.
// Note that you need to call Disk.Unwrap() before calling this method if this Disk
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Disk) Update() *DiskUpdateOne {
	return (&DiskClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the Disk entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Disk) Unwrap() *Disk {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Disk is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Disk) String() string {
	var builder strings.Builder
	builder.WriteString("Disk(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteString(", size=")
	builder.WriteString(fmt.Sprintf("%v", d.Size))
	builder.WriteByte(')')
	return builder.String()
}

// Disks is a parsable slice of Disk.
type Disks []*Disk

func (d Disks) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
