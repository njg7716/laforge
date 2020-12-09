// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/gen0cide/laforge/ent/competition"
)

// Competition is the model entity for the Competition schema.
type Competition struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// RootPassword holds the value of the "root_password" field.
	RootPassword string `json:"root_password,omitempty"`
	// Config holds the value of the "config" field.
	Config []string `json:"config,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CompetitionQuery when eager-loading is set.
	Edges                   CompetitionEdges `json:"edges"`
	environment_competition *int
}

// CompetitionEdges holds the relations/edges for other nodes in the graph.
type CompetitionEdges struct {
	// DNS holds the value of the dns edge.
	DNS []*DNS
	// Tag holds the value of the tag edge.
	Tag []*Tag
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// DNSOrErr returns the DNS value or an error if the edge
// was not loaded in eager-loading.
func (e CompetitionEdges) DNSOrErr() ([]*DNS, error) {
	if e.loadedTypes[0] {
		return e.DNS, nil
	}
	return nil, &NotLoadedError{edge: "dns"}
}

// TagOrErr returns the Tag value or an error if the edge
// was not loaded in eager-loading.
func (e CompetitionEdges) TagOrErr() ([]*Tag, error) {
	if e.loadedTypes[1] {
		return e.Tag, nil
	}
	return nil, &NotLoadedError{edge: "tag"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Competition) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // root_password
		&[]byte{},         // config
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Competition) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // environment_competition
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Competition fields.
func (c *Competition) assignValues(values ...interface{}) error {
	if m, n := len(values), len(competition.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	c.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field root_password", values[0])
	} else if value.Valid {
		c.RootPassword = value.String
	}

	if value, ok := values[1].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field config", values[1])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &c.Config); err != nil {
			return fmt.Errorf("unmarshal field config: %v", err)
		}
	}
	values = values[2:]
	if len(values) == len(competition.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field environment_competition", value)
		} else if value.Valid {
			c.environment_competition = new(int)
			*c.environment_competition = int(value.Int64)
		}
	}
	return nil
}

// QueryDNS queries the dns edge of the Competition.
func (c *Competition) QueryDNS() *DNSQuery {
	return (&CompetitionClient{config: c.config}).QueryDNS(c)
}

// QueryTag queries the tag edge of the Competition.
func (c *Competition) QueryTag() *TagQuery {
	return (&CompetitionClient{config: c.config}).QueryTag(c)
}

// Update returns a builder for updating this Competition.
// Note that, you need to call Competition.Unwrap() before calling this method, if this Competition
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Competition) Update() *CompetitionUpdateOne {
	return (&CompetitionClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *Competition) Unwrap() *Competition {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Competition is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Competition) String() string {
	var builder strings.Builder
	builder.WriteString("Competition(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", root_password=")
	builder.WriteString(c.RootPassword)
	builder.WriteString(", config=")
	builder.WriteString(fmt.Sprintf("%v", c.Config))
	builder.WriteByte(')')
	return builder.String()
}

// Competitions is a parsable slice of Competition.
type Competitions []*Competition

func (c Competitions) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
