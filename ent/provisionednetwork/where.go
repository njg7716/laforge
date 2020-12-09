// Code generated by entc, DO NOT EDIT.

package provisionednetwork

import (
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/gen0cide/laforge/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Cidr applies equality check predicate on the "cidr" field. It's identical to CidrEQ.
func Cidr(v string) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCidr), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.ProvisionedNetwork {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.ProvisionedNetwork {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// CidrEQ applies the EQ predicate on the "cidr" field.
func CidrEQ(v string) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCidr), v))
	})
}

// CidrNEQ applies the NEQ predicate on the "cidr" field.
func CidrNEQ(v string) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCidr), v))
	})
}

// CidrIn applies the In predicate on the "cidr" field.
func CidrIn(vs ...string) predicate.ProvisionedNetwork {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCidr), v...))
	})
}

// CidrNotIn applies the NotIn predicate on the "cidr" field.
func CidrNotIn(vs ...string) predicate.ProvisionedNetwork {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCidr), v...))
	})
}

// CidrGT applies the GT predicate on the "cidr" field.
func CidrGT(v string) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCidr), v))
	})
}

// CidrGTE applies the GTE predicate on the "cidr" field.
func CidrGTE(v string) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCidr), v))
	})
}

// CidrLT applies the LT predicate on the "cidr" field.
func CidrLT(v string) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCidr), v))
	})
}

// CidrLTE applies the LTE predicate on the "cidr" field.
func CidrLTE(v string) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCidr), v))
	})
}

// CidrContains applies the Contains predicate on the "cidr" field.
func CidrContains(v string) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCidr), v))
	})
}

// CidrHasPrefix applies the HasPrefix predicate on the "cidr" field.
func CidrHasPrefix(v string) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCidr), v))
	})
}

// CidrHasSuffix applies the HasSuffix predicate on the "cidr" field.
func CidrHasSuffix(v string) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCidr), v))
	})
}

// CidrEqualFold applies the EqualFold predicate on the "cidr" field.
func CidrEqualFold(v string) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCidr), v))
	})
}

// CidrContainsFold applies the ContainsFold predicate on the "cidr" field.
func CidrContainsFold(v string) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCidr), v))
	})
}

// HasTag applies the HasEdge predicate on the "tag" edge.
func HasTag() predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TagTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TagTable, TagColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTagWith applies the HasEdge predicate on the "tag" edge with a given conditions (other predicates).
func HasTagWith(preds ...predicate.Tag) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TagInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TagTable, TagColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProvisionedHosts applies the HasEdge predicate on the "provisioned_hosts" edge.
func HasProvisionedHosts() predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProvisionedHostsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProvisionedHostsTable, ProvisionedHostsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProvisionedHostsWith applies the HasEdge predicate on the "provisioned_hosts" edge with a given conditions (other predicates).
func HasProvisionedHostsWith(preds ...predicate.ProvisionedHost) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProvisionedHostsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProvisionedHostsTable, ProvisionedHostsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStatus applies the HasEdge predicate on the "status" edge.
func HasStatus() predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StatusTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StatusTable, StatusColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStatusWith applies the HasEdge predicate on the "status" edge with a given conditions (other predicates).
func HasStatusWith(preds ...predicate.Status) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StatusInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StatusTable, StatusColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNetwork applies the HasEdge predicate on the "network" edge.
func HasNetwork() predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NetworkTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, NetworkTable, NetworkColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNetworkWith applies the HasEdge predicate on the "network" edge with a given conditions (other predicates).
func HasNetworkWith(preds ...predicate.Network) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NetworkInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, NetworkTable, NetworkColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBuild applies the HasEdge predicate on the "build" edge.
func HasBuild() predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BuildTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BuildTable, BuildColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBuildWith applies the HasEdge predicate on the "build" edge with a given conditions (other predicates).
func HasBuildWith(preds ...predicate.Build) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BuildInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BuildTable, BuildColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.ProvisionedNetwork) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.ProvisionedNetwork) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ProvisionedNetwork) predicate.ProvisionedNetwork {
	return predicate.ProvisionedNetwork(func(s *sql.Selector) {
		p(s.Not())
	})
}
