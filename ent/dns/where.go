// Code generated by entc, DO NOT EDIT.

package dns

import (
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/gen0cide/laforge/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
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
func IDGT(id int) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// RootDomain applies equality check predicate on the "root_domain" field. It's identical to RootDomainEQ.
func RootDomain(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRootDomain), v))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.DNS {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DNS(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.DNS {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DNS(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldType), v))
	})
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldType), v))
	})
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldType), v))
	})
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldType), v))
	})
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldType), v))
	})
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldType), v))
	})
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldType), v))
	})
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldType), v))
	})
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldType), v))
	})
}

// RootDomainEQ applies the EQ predicate on the "root_domain" field.
func RootDomainEQ(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRootDomain), v))
	})
}

// RootDomainNEQ applies the NEQ predicate on the "root_domain" field.
func RootDomainNEQ(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRootDomain), v))
	})
}

// RootDomainIn applies the In predicate on the "root_domain" field.
func RootDomainIn(vs ...string) predicate.DNS {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DNS(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRootDomain), v...))
	})
}

// RootDomainNotIn applies the NotIn predicate on the "root_domain" field.
func RootDomainNotIn(vs ...string) predicate.DNS {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DNS(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRootDomain), v...))
	})
}

// RootDomainGT applies the GT predicate on the "root_domain" field.
func RootDomainGT(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRootDomain), v))
	})
}

// RootDomainGTE applies the GTE predicate on the "root_domain" field.
func RootDomainGTE(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRootDomain), v))
	})
}

// RootDomainLT applies the LT predicate on the "root_domain" field.
func RootDomainLT(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRootDomain), v))
	})
}

// RootDomainLTE applies the LTE predicate on the "root_domain" field.
func RootDomainLTE(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRootDomain), v))
	})
}

// RootDomainContains applies the Contains predicate on the "root_domain" field.
func RootDomainContains(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRootDomain), v))
	})
}

// RootDomainHasPrefix applies the HasPrefix predicate on the "root_domain" field.
func RootDomainHasPrefix(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRootDomain), v))
	})
}

// RootDomainHasSuffix applies the HasSuffix predicate on the "root_domain" field.
func RootDomainHasSuffix(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRootDomain), v))
	})
}

// RootDomainEqualFold applies the EqualFold predicate on the "root_domain" field.
func RootDomainEqualFold(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRootDomain), v))
	})
}

// RootDomainContainsFold applies the ContainsFold predicate on the "root_domain" field.
func RootDomainContainsFold(v string) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRootDomain), v))
	})
}

// HasTag applies the HasEdge predicate on the "tag" edge.
func HasTag() predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TagTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TagTable, TagColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTagWith applies the HasEdge predicate on the "tag" edge with a given conditions (other predicates).
func HasTagWith(preds ...predicate.Tag) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
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

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.DNS) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.DNS) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
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
func Not(p predicate.DNS) predicate.DNS {
	return predicate.DNS(func(s *sql.Selector) {
		p(s.Not())
	})
}
