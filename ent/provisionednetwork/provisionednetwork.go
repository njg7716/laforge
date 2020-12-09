// Code generated by entc, DO NOT EDIT.

package provisionednetwork

const (
	// Label holds the string label denoting the provisionednetwork type in the database.
	Label = "provisioned_network"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCidr holds the string denoting the cidr field in the database.
	FieldCidr = "cidr"
	// FieldVars holds the string denoting the vars field in the database.
	FieldVars = "vars"

	// EdgeTag holds the string denoting the tag edge name in mutations.
	EdgeTag = "tag"
	// EdgeProvisionedHosts holds the string denoting the provisioned_hosts edge name in mutations.
	EdgeProvisionedHosts = "provisioned_hosts"
	// EdgeStatus holds the string denoting the status edge name in mutations.
	EdgeStatus = "status"
	// EdgeNetwork holds the string denoting the network edge name in mutations.
	EdgeNetwork = "network"
	// EdgeBuild holds the string denoting the build edge name in mutations.
	EdgeBuild = "build"

	// Table holds the table name of the provisionednetwork in the database.
	Table = "provisioned_networks"
	// TagTable is the table the holds the tag relation/edge.
	TagTable = "tags"
	// TagInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagInverseTable = "tags"
	// TagColumn is the table column denoting the tag relation/edge.
	TagColumn = "provisioned_network_tag"
	// ProvisionedHostsTable is the table the holds the provisioned_hosts relation/edge.
	ProvisionedHostsTable = "provisioned_hosts"
	// ProvisionedHostsInverseTable is the table name for the ProvisionedHost entity.
	// It exists in this package in order to avoid circular dependency with the "provisionedhost" package.
	ProvisionedHostsInverseTable = "provisioned_hosts"
	// ProvisionedHostsColumn is the table column denoting the provisioned_hosts relation/edge.
	ProvisionedHostsColumn = "provisioned_network_provisioned_hosts"
	// StatusTable is the table the holds the status relation/edge.
	StatusTable = "status"
	// StatusInverseTable is the table name for the Status entity.
	// It exists in this package in order to avoid circular dependency with the "status" package.
	StatusInverseTable = "status"
	// StatusColumn is the table column denoting the status relation/edge.
	StatusColumn = "provisioned_network_status"
	// NetworkTable is the table the holds the network relation/edge.
	NetworkTable = "networks"
	// NetworkInverseTable is the table name for the Network entity.
	// It exists in this package in order to avoid circular dependency with the "network" package.
	NetworkInverseTable = "networks"
	// NetworkColumn is the table column denoting the network relation/edge.
	NetworkColumn = "provisioned_network_network"
	// BuildTable is the table the holds the build relation/edge.
	BuildTable = "builds"
	// BuildInverseTable is the table name for the Build entity.
	// It exists in this package in order to avoid circular dependency with the "build" package.
	BuildInverseTable = "builds"
	// BuildColumn is the table column denoting the build relation/edge.
	BuildColumn = "provisioned_network_build"
)

// Columns holds all SQL columns for provisionednetwork fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCidr,
	FieldVars,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the ProvisionedNetwork type.
var ForeignKeys = []string{
	"provisioned_host_provisioned_network",
	"team_provisioned_networks",
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
