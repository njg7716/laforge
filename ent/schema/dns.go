package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DNS holds the schema definition for the DNS entity.
type DNS struct {
	ent.Schema
}

// Fields of the DNS.
func (DNS) Fields() []ent.Field {
	return []ent.Field{
		field.String("hcl_id").
			StructTag(`hcl:"id,label"`),
		field.String("type").
			StructTag(`hcl:"type,attr"`),
		field.String("root_domain").
			StructTag(`hcl:"root_domain,attr" `),
		field.JSON("dns_servers", []string{}).
			StructTag(`hcl:"dns_servers,attr"`),
		field.JSON("ntp_servers", []string{}).
			StructTag(`hcl:"ntp_servers,optional"`),
		field.JSON("config", map[string]string{}).
			StructTag(`hcl:"config,optional"`),
	}
}

// Edges of the DNS.
func (DNS) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("DNSToTag", Tag.Type),
		edge.From("DNSToEnvironment", Environment.Type).Ref("EnvironmentToDNS"),
		edge.From("DNSToCompetition", Competition.Type).Ref("CompetitionToDNS"),
	}
}
