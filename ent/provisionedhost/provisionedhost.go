// Code generated by entc, DO NOT EDIT.

package provisionedhost

const (
	// Label holds the string label denoting the provisionedhost type in the database.
	Label = "provisioned_host"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSubnetIP holds the string denoting the subnet_ip field in the database.
	FieldSubnetIP = "subnet_ip"

	// EdgeStatus holds the string denoting the status edge name in mutations.
	EdgeStatus = "status"
	// EdgeProvisioningSteps holds the string denoting the provisioning_steps edge name in mutations.
	EdgeProvisioningSteps = "provisioning_steps"
	// EdgeProvisionedNetwork holds the string denoting the provisioned_network edge name in mutations.
	EdgeProvisionedNetwork = "provisioned_network"
	// EdgeHost holds the string denoting the host edge name in mutations.
	EdgeHost = "host"
	// EdgeTag holds the string denoting the tag edge name in mutations.
	EdgeTag = "tag"

	// Table holds the table name of the provisionedhost in the database.
	Table = "provisioned_hosts"
	// StatusTable is the table the holds the status relation/edge.
	StatusTable = "status"
	// StatusInverseTable is the table name for the Status entity.
	// It exists in this package in order to avoid circular dependency with the "status" package.
	StatusInverseTable = "status"
	// StatusColumn is the table column denoting the status relation/edge.
	StatusColumn = "provisioned_host_status"
	// ProvisioningStepsTable is the table the holds the provisioning_steps relation/edge.
	ProvisioningStepsTable = "provisioning_steps"
	// ProvisioningStepsInverseTable is the table name for the ProvisioningStep entity.
	// It exists in this package in order to avoid circular dependency with the "provisioningstep" package.
	ProvisioningStepsInverseTable = "provisioning_steps"
	// ProvisioningStepsColumn is the table column denoting the provisioning_steps relation/edge.
	ProvisioningStepsColumn = "provisioned_host_provisioning_steps"
	// ProvisionedNetworkTable is the table the holds the provisioned_network relation/edge.
	ProvisionedNetworkTable = "provisioned_networks"
	// ProvisionedNetworkInverseTable is the table name for the ProvisionedNetwork entity.
	// It exists in this package in order to avoid circular dependency with the "provisionednetwork" package.
	ProvisionedNetworkInverseTable = "provisioned_networks"
	// ProvisionedNetworkColumn is the table column denoting the provisioned_network relation/edge.
	ProvisionedNetworkColumn = "provisioned_host_provisioned_network"
	// HostTable is the table the holds the host relation/edge.
	HostTable = "hosts"
	// HostInverseTable is the table name for the Host entity.
	// It exists in this package in order to avoid circular dependency with the "host" package.
	HostInverseTable = "hosts"
	// HostColumn is the table column denoting the host relation/edge.
	HostColumn = "provisioned_host_host"
	// TagTable is the table the holds the tag relation/edge.
	TagTable = "tags"
	// TagInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagInverseTable = "tags"
	// TagColumn is the table column denoting the tag relation/edge.
	TagColumn = "provisioned_host_tag"
)

// Columns holds all SQL columns for provisionedhost fields.
var Columns = []string{
	FieldID,
	FieldSubnetIP,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the ProvisionedHost type.
var ForeignKeys = []string{
	"provisioned_network_provisioned_hosts",
	"provisioning_step_provisioned_host",
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
