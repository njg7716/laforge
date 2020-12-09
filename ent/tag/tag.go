// Code generated by entc, DO NOT EDIT.

package tag

const (
	// Label holds the string label denoting the tag type in the database.
	Label = "tag"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"

	// EdgeTag holds the string denoting the tag edge name in mutations.
	EdgeTag = "tag"

	// Table holds the table name of the tag in the database.
	Table = "tags"
	// TagTable is the table the holds the tag relation/edge. The primary key declared below.
	TagTable = "tag_tag"
)

// Columns holds all SQL columns for tag fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldName,
	FieldDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Tag type.
var ForeignKeys = []string{
	"build_tag",
	"command_tag",
	"competition_tag",
	"dns_tag",
	"dns_record_tag",
	"disk_tag",
	"environment_tag",
	"file_delete_tag",
	"file_download_tag",
	"file_extract_tag",
	"finding_tag",
	"host_tag",
	"included_network_tag",
	"network_tag",
	"provisioned_host_tag",
	"provisioned_network_tag",
	"provisioning_step_tag",
	"script_tag",
	"status_tag",
	"team_tag",
	"user_tag",
}

var (
	// TagPrimaryKey and TagColumn2 are the table columns denoting the
	// primary key for the tag relation (M2M).
	TagPrimaryKey = []string{"tag_id", "tag_id"}
)

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
