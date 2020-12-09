// Code generated by entc, DO NOT EDIT.

package filedelete

const (
	// Label holds the string label denoting the filedelete type in the database.
	Label = "file_delete"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"

	// EdgeTag holds the string denoting the tag edge name in mutations.
	EdgeTag = "tag"

	// Table holds the table name of the filedelete in the database.
	Table = "file_deletes"
	// TagTable is the table the holds the tag relation/edge.
	TagTable = "tags"
	// TagInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagInverseTable = "tags"
	// TagColumn is the table column denoting the tag relation/edge.
	TagColumn = "file_delete_tag"
)

// Columns holds all SQL columns for filedelete fields.
var Columns = []string{
	FieldID,
	FieldPath,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the FileDelete type.
var ForeignKeys = []string{
	"provisioning_step_file_delete",
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
