// Code generated by entc, DO NOT EDIT.

package environment

const (
	// Label holds the string label denoting the environment type in the database.
	Label = "environment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCompetitionID holds the string denoting the competition_id field in the database.
	FieldCompetitionID = "competition_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldBuilder holds the string denoting the builder field in the database.
	FieldBuilder = "builder"
	// FieldTeamCount holds the string denoting the team_count field in the database.
	FieldTeamCount = "team_count"
	// FieldRevision holds the string denoting the revision field in the database.
	FieldRevision = "revision"
	// FieldAdminCidrs holds the string denoting the admin_cidrs field in the database.
	FieldAdminCidrs = "admin_cidrs"
	// FieldExposedVdiPorts holds the string denoting the exposed_vdi_ports field in the database.
	FieldExposedVdiPorts = "exposed_vdi_ports"
	// FieldConfig holds the string denoting the config field in the database.
	FieldConfig = "config"

	// EdgeTag holds the string denoting the tag edge name in mutations.
	EdgeTag = "tag"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeIncludedNetwork holds the string denoting the included_network edge name in mutations.
	EdgeIncludedNetwork = "included_network"
	// EdgeBuild holds the string denoting the build edge name in mutations.
	EdgeBuild = "build"
	// EdgeNetwork holds the string denoting the network edge name in mutations.
	EdgeNetwork = "network"
	// EdgeHost holds the string denoting the host edge name in mutations.
	EdgeHost = "host"
	// EdgeCompetition holds the string denoting the competition edge name in mutations.
	EdgeCompetition = "competition"

	// Table holds the table name of the environment in the database.
	Table = "environments"
	// TagTable is the table the holds the tag relation/edge.
	TagTable = "tags"
	// TagInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagInverseTable = "tags"
	// TagColumn is the table column denoting the tag relation/edge.
	TagColumn = "environment_tag"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "users"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "environment_user"
	// IncludedNetworkTable is the table the holds the included_network relation/edge.
	IncludedNetworkTable = "included_networks"
	// IncludedNetworkInverseTable is the table name for the IncludedNetwork entity.
	// It exists in this package in order to avoid circular dependency with the "includednetwork" package.
	IncludedNetworkInverseTable = "included_networks"
	// IncludedNetworkColumn is the table column denoting the included_network relation/edge.
	IncludedNetworkColumn = "environment_included_network"
	// BuildTable is the table the holds the build relation/edge.
	BuildTable = "builds"
	// BuildInverseTable is the table name for the Build entity.
	// It exists in this package in order to avoid circular dependency with the "build" package.
	BuildInverseTable = "builds"
	// BuildColumn is the table column denoting the build relation/edge.
	BuildColumn = "environment_build"
	// NetworkTable is the table the holds the network relation/edge.
	NetworkTable = "networks"
	// NetworkInverseTable is the table name for the Network entity.
	// It exists in this package in order to avoid circular dependency with the "network" package.
	NetworkInverseTable = "networks"
	// NetworkColumn is the table column denoting the network relation/edge.
	NetworkColumn = "environment_network"
	// HostTable is the table the holds the host relation/edge.
	HostTable = "hosts"
	// HostInverseTable is the table name for the Host entity.
	// It exists in this package in order to avoid circular dependency with the "host" package.
	HostInverseTable = "hosts"
	// HostColumn is the table column denoting the host relation/edge.
	HostColumn = "environment_host"
	// CompetitionTable is the table the holds the competition relation/edge.
	CompetitionTable = "competitions"
	// CompetitionInverseTable is the table name for the Competition entity.
	// It exists in this package in order to avoid circular dependency with the "competition" package.
	CompetitionInverseTable = "competitions"
	// CompetitionColumn is the table column denoting the competition relation/edge.
	CompetitionColumn = "environment_competition"
)

// Columns holds all SQL columns for environment fields.
var Columns = []string{
	FieldID,
	FieldCompetitionID,
	FieldName,
	FieldDescription,
	FieldBuilder,
	FieldTeamCount,
	FieldRevision,
	FieldAdminCidrs,
	FieldExposedVdiPorts,
	FieldConfig,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Environment type.
var ForeignKeys = []string{
	"team_environment",
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
