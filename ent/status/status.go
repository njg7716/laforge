// Code generated by entc, DO NOT EDIT.

package status

import (
	"fmt"
)

const (
	// Label holds the string label denoting the status type in the database.
	Label = "status"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldStartedAt holds the string denoting the started_at field in the database.
	FieldStartedAt = "started_at"
	// FieldEndedAt holds the string denoting the ended_at field in the database.
	FieldEndedAt = "ended_at"
	// FieldFailed holds the string denoting the failed field in the database.
	FieldFailed = "failed"
	// FieldCompleted holds the string denoting the completed field in the database.
	FieldCompleted = "completed"
	// FieldError holds the string denoting the error field in the database.
	FieldError = "error"

	// EdgeTag holds the string denoting the tag edge name in mutations.
	EdgeTag = "tag"

	// Table holds the table name of the status in the database.
	Table = "status"
	// TagTable is the table the holds the tag relation/edge.
	TagTable = "tags"
	// TagInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagInverseTable = "tags"
	// TagColumn is the table column denoting the tag relation/edge.
	TagColumn = "status_tag"
)

// Columns holds all SQL columns for status fields.
var Columns = []string{
	FieldID,
	FieldState,
	FieldStartedAt,
	FieldEndedAt,
	FieldFailed,
	FieldCompleted,
	FieldError,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Status type.
var ForeignKeys = []string{
	"provisioned_host_status",
	"provisioned_network_status",
	"provisioning_step_status",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// State defines the type for the state enum field.
type State string

// State values.
const (
	StateStaged    State = "staged"
	StateBuilding  State = "building"
	StateFailed    State = "failed"
	StateSucceeded State = "succeeded"
)

func (s State) String() string {
	return string(s)
}

// StateValidator is a validator for the "state" field enum values. It is called by the builders before save.
func StateValidator(s State) error {
	switch s {
	case StateStaged, StateBuilding, StateFailed, StateSucceeded:
		return nil
	default:
		return fmt.Errorf("status: invalid enum value for state field: %q", s)
	}
}
