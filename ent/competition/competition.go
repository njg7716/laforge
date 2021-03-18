// Code generated by entc, DO NOT EDIT.

package competition

const (
	// Label holds the string label denoting the competition type in the database.
	Label = "competition"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldHclID holds the string denoting the hcl_id field in the database.
	FieldHclID = "hcl_id"
	// FieldRootPassword holds the string denoting the root_password field in the database.
	FieldRootPassword = "root_password"
	// FieldConfig holds the string denoting the config field in the database.
	FieldConfig = "config"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"

	// EdgeCompetitionToTag holds the string denoting the competitiontotag edge name in mutations.
	EdgeCompetitionToTag = "CompetitionToTag"
	// EdgeCompetitionToDNS holds the string denoting the competitiontodns edge name in mutations.
	EdgeCompetitionToDNS = "CompetitionToDNS"
	// EdgeCompetitionToEnvironment holds the string denoting the competitiontoenvironment edge name in mutations.
	EdgeCompetitionToEnvironment = "CompetitionToEnvironment"

	// Table holds the table name of the competition in the database.
	Table = "competitions"
	// CompetitionToTagTable is the table the holds the CompetitionToTag relation/edge.
	CompetitionToTagTable = "tags"
	// CompetitionToTagInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	CompetitionToTagInverseTable = "tags"
	// CompetitionToTagColumn is the table column denoting the CompetitionToTag relation/edge.
	CompetitionToTagColumn = "competition_competition_to_tag"
	// CompetitionToDNSTable is the table the holds the CompetitionToDNS relation/edge. The primary key declared below.
	CompetitionToDNSTable = "competition_CompetitionToDNS"
	// CompetitionToDNSInverseTable is the table name for the DNS entity.
	// It exists in this package in order to avoid circular dependency with the "dns" package.
	CompetitionToDNSInverseTable = "dn_ss"
	// CompetitionToEnvironmentTable is the table the holds the CompetitionToEnvironment relation/edge. The primary key declared below.
	CompetitionToEnvironmentTable = "environment_EnvironmentToCompetition"
	// CompetitionToEnvironmentInverseTable is the table name for the Environment entity.
	// It exists in this package in order to avoid circular dependency with the "environment" package.
	CompetitionToEnvironmentInverseTable = "environments"
)

// Columns holds all SQL columns for competition fields.
var Columns = []string{
	FieldID,
	FieldHclID,
	FieldRootPassword,
	FieldConfig,
	FieldTags,
}

var (
	// CompetitionToDNSPrimaryKey and CompetitionToDNSColumn2 are the table columns denoting the
	// primary key for the CompetitionToDNS relation (M2M).
	CompetitionToDNSPrimaryKey = []string{"competition_id", "dns_id"}
	// CompetitionToEnvironmentPrimaryKey and CompetitionToEnvironmentColumn2 are the table columns denoting the
	// primary key for the CompetitionToEnvironment relation (M2M).
	CompetitionToEnvironmentPrimaryKey = []string{"environment_id", "competition_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
