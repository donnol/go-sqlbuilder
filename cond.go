// Copyright 2018 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package sqlbuilder

// Cond provides several helper methods to build conditions.
type Cond struct {
	Args *Args
}

// NewCond returns a new Cond.
func NewCond() *Cond {
	return &Cond{
		Args: &Args{},
	}
}

type Condsult interface {
	Init(s StringBuilder)
	Value() interface{}
	WriteBefore()
	WriteString(index int, ph string)
	WriteAfter()
	Cond() string
}

var _ Condsult = (*condsult)(nil)

type condsult struct {
	s     StringBuilder
	field string
	op    string
	value interface{}
}

func (cs *condsult) Init(s StringBuilder) {
	cs.s = s
}

func (cs *condsult) Value() interface{} {
	return cs.value
}

func (cs *condsult) WriteBefore() {
	cs.s.WriteString(cs.field)
	cs.s.WriteString(cs.op)
}

func (cs *condsult) WriteString(index int, ph string) {
	cs.s.WriteString(ph)
}

func (cs *condsult) WriteAfter() {
}

func (cs *condsult) Cond() string {
	return cs.s.String()
}

// Equal represents "field = value".
func Equal(field string, value interface{}) Condsult {
	return &condsult{
		field: Escape(field),
		op:    " = ",
		value: value,
	}
}

// Equal represents "field = value".
func (c *Cond) Equal(field string, value interface{}) string {
	buf := newStringBuilder()
	buf.WriteString(Escape(field))
	buf.WriteString(" = ")
	buf.WriteString(c.Args.Add(value))
	return buf.String()
}

// E is an alias of Equal.
func (c *Cond) E(field string, value interface{}) string {
	return c.Equal(field, value)
}

// EQ is an alias of Equal.
func (c *Cond) EQ(field string, value interface{}) string {
	return c.Equal(field, value)
}

// NotEqual represents "field <> value".
func NotEqual(field string, value interface{}) Condsult {
	return &condsult{
		field: Escape(field),
		op:    " <> ",
		value: value,
	}
}

// NotEqual represents "field <> value".
func (c *Cond) NotEqual(field string, value interface{}) string {
	buf := newStringBuilder()
	buf.WriteString(Escape(field))
	buf.WriteString(" <> ")
	buf.WriteString(c.Args.Add(value))
	return buf.String()
}

// NE is an alias of NotEqual.
func (c *Cond) NE(field string, value interface{}) string {
	return c.NotEqual(field, value)
}

// NEQ is an alias of NotEqual.
func (c *Cond) NEQ(field string, value interface{}) string {
	return c.NotEqual(field, value)
}

// GreaterThan represents "field > value".
func GreaterThan(field string, value interface{}) Condsult {
	return &condsult{
		field: Escape(field),
		op:    " > ",
		value: value,
	}
}

// GreaterThan represents "field > value".
func (c *Cond) GreaterThan(field string, value interface{}) string {
	buf := newStringBuilder()
	buf.WriteString(Escape(field))
	buf.WriteString(" > ")
	buf.WriteString(c.Args.Add(value))
	return buf.String()
}

// G is an alias of GreaterThan.
func (c *Cond) G(field string, value interface{}) string {
	return c.GreaterThan(field, value)
}

// GT is an alias of GreaterThan.
func (c *Cond) GT(field string, value interface{}) string {
	return c.GreaterThan(field, value)
}

// GreaterEqualThan represents "field >= value".
func GreaterEqualThan(field string, value interface{}) Condsult {
	return &condsult{
		field: Escape(field),
		op:    " >= ",
		value: value,
	}
}

// GreaterEqualThan represents "field >= value".
func (c *Cond) GreaterEqualThan(field string, value interface{}) string {
	buf := newStringBuilder()
	buf.WriteString(Escape(field))
	buf.WriteString(" >= ")
	buf.WriteString(c.Args.Add(value))
	return buf.String()
}

// GE is an alias of GreaterEqualThan.
func (c *Cond) GE(field string, value interface{}) string {
	return c.GreaterEqualThan(field, value)
}

// GTE is an alias of GreaterEqualThan.
func (c *Cond) GTE(field string, value interface{}) string {
	return c.GreaterEqualThan(field, value)
}

// LessThan represents "field < value".
func LessThan(field string, value interface{}) Condsult {
	return &condsult{
		field: Escape(field),
		op:    " < ",
		value: value,
	}
}

// LessThan represents "field < value".
func (c *Cond) LessThan(field string, value interface{}) string {
	buf := newStringBuilder()
	buf.WriteString(Escape(field))
	buf.WriteString(" < ")
	buf.WriteString(c.Args.Add(value))
	return buf.String()
}

// L is an alias of LessThan.
func (c *Cond) L(field string, value interface{}) string {
	return c.LessThan(field, value)
}

// LT is an alias of LessThan.
func (c *Cond) LT(field string, value interface{}) string {
	return c.LessThan(field, value)
}

// LessEqualThan represents "field <= value".
func LessEqualThan(field string, value interface{}) Condsult {
	return &condsult{
		field: Escape(field),
		op:    " <= ",
		value: value,
	}
}

// LessEqualThan represents "field <= value".
func (c *Cond) LessEqualThan(field string, value interface{}) string {
	buf := newStringBuilder()
	buf.WriteString(Escape(field))
	buf.WriteString(" <= ")
	buf.WriteString(c.Args.Add(value))
	return buf.String()
}

// LE is an alias of LessEqualThan.
func (c *Cond) LE(field string, value interface{}) string {
	return c.LessEqualThan(field, value)
}

// LTE is an alias of LessEqualThan.
func (c *Cond) LTE(field string, value interface{}) string {
	return c.LessEqualThan(field, value)
}

type inCondsult struct {
	sep string

	condsult
}

func (cs *inCondsult) Init(s StringBuilder) {
	cs.s = s
}

func (cs *inCondsult) Value() interface{} {
	return cs.value
}

func (cs *inCondsult) WriteBefore() {
	cs.s.WriteString(cs.field)
	cs.s.WriteString(cs.op)
	cs.s.WriteString("(")
}

func (cs *inCondsult) WriteString(index int, ph string) {
	if index != 0 {
		cs.s.WriteString(cs.sep)
	}
	cs.s.WriteString(ph)
}

func (cs *inCondsult) WriteAfter() {
	cs.s.WriteString(")")
}

func (cs *inCondsult) Cond() string {
	return cs.s.String()
}

// In represents "field IN (value...)".
func In(field string, value ...interface{}) Condsult {
	return &inCondsult{
		sep: ", ",

		condsult: condsult{
			field: Escape(field),
			op:    " IN ",
			value: value,
		},
	}
}

// In represents "field IN (value...)".
func (c *Cond) In(field string, value ...interface{}) string {
	vs := make([]string, 0, len(value))

	for _, v := range value {
		vs = append(vs, c.Args.Add(v))
	}

	buf := newStringBuilder()
	buf.WriteString(Escape(field))
	buf.WriteString(" IN (")
	buf.WriteStrings(vs, ", ")
	buf.WriteString(")")
	return buf.String()
}

// NotIn represents "field NOT IN (value...)".
func NotIn(field string, value ...interface{}) Condsult {
	return &inCondsult{
		sep: ", ",

		condsult: condsult{
			field: Escape(field),
			op:    " NOT IN ",
			value: value,
		},
	}
}

// NotIn represents "field NOT IN (value...)".
func (c *Cond) NotIn(field string, value ...interface{}) string {
	vs := make([]string, 0, len(value))

	for _, v := range value {
		vs = append(vs, c.Args.Add(v))
	}

	buf := newStringBuilder()
	buf.WriteString(Escape(field))
	buf.WriteString(" NOT IN (")
	buf.WriteStrings(vs, ", ")
	buf.WriteString(")")
	return buf.String()
}

// Like represents "field LIKE value".
func Like(field string, value interface{}) Condsult {
	return &condsult{
		field: Escape(field),
		op:    " LIKE ",
		value: value,
	}
}

// Like represents "field LIKE value".
func (c *Cond) Like(field string, value interface{}) string {
	buf := newStringBuilder()
	buf.WriteString(Escape(field))
	buf.WriteString(" LIKE ")
	buf.WriteString(c.Args.Add(value))
	return buf.String()
}

// NotLike represents "field NOT LIKE value".
func NotLike(field string, value interface{}) Condsult {
	return &condsult{
		field: Escape(field),
		op:    " NOT LIKE ",
		value: value,
	}
}

// NotLike represents "field NOT LIKE value".
func (c *Cond) NotLike(field string, value interface{}) string {
	buf := newStringBuilder()
	buf.WriteString(Escape(field))
	buf.WriteString(" NOT LIKE ")
	buf.WriteString(c.Args.Add(value))
	return buf.String()
}

type isCondsult struct {
	condsult
}

func (cs *isCondsult) Init(s StringBuilder) {
	cs.s = s
}

func (cs *isCondsult) Value() interface{} {
	return cs.value
}

func (cs *isCondsult) WriteBefore() {
	cs.s.WriteString(cs.field)
	cs.s.WriteString(cs.op)
}

func (cs *isCondsult) WriteString(index int, ph string) {
}

func (cs *isCondsult) WriteAfter() {
}

func (cs *isCondsult) Cond() string {
	return cs.s.String()
}

// IsNull represents "field IS NULL".
func IsNull(field string) Condsult {
	return &isCondsult{
		condsult: condsult{
			field: Escape(field),
			op:    " IS NULL",
			value: nil,
		},
	}
}

// IsNull represents "field IS NULL".
func (c *Cond) IsNull(field string) string {
	buf := newStringBuilder()
	buf.WriteString(Escape(field))
	buf.WriteString(" IS NULL")
	return buf.String()
}

// IsNotNull represents "field IS NOT NULL".
func IsNotNull(field string) Condsult {
	return &isCondsult{
		condsult: condsult{
			field: Escape(field),
			op:    " IS NOT NULL",
			value: nil,
		},
	}
}

// IsNotNull represents "field IS NOT NULL".
func (c *Cond) IsNotNull(field string) string {
	buf := newStringBuilder()
	buf.WriteString(Escape(field))
	buf.WriteString(" IS NOT NULL")
	return buf.String()
}

type betCondsult struct {
	condsult
}

func (cs *betCondsult) Init(s StringBuilder) {
	cs.s = s
}

func (cs *betCondsult) Value() interface{} {
	return cs.value
}

func (cs *betCondsult) WriteBefore() {
	cs.s.WriteString(cs.field)
	cs.s.WriteString(cs.op)
}

func (cs *betCondsult) WriteString(index int, ph string) {
	if index == 1 {
		cs.s.WriteString(" AND ")
	}
	cs.s.WriteString(ph)
}

func (cs *betCondsult) WriteAfter() {
}

func (cs *betCondsult) Cond() string {
	return cs.s.String()
}

// Between represents "field BETWEEN lower AND upper".
func Between(field string, lower, upper interface{}) Condsult {
	return &betCondsult{
		condsult: condsult{
			field: Escape(field),
			op:    " BETWEEN ",
			value: []interface{}{lower, upper},
		},
	}
}

// Between represents "field BETWEEN lower AND upper".
func (c *Cond) Between(field string, lower, upper interface{}) string {
	buf := newStringBuilder()
	buf.WriteString(Escape(field))
	buf.WriteString(" BETWEEN ")
	buf.WriteString(c.Args.Add(lower))
	buf.WriteString(" AND ")
	buf.WriteString(c.Args.Add(upper))
	return buf.String()
}

// NotBetween represents "field NOT BETWEEN lower AND upper".
func NotBetween(field string, lower, upper interface{}) Condsult {
	return &betCondsult{
		condsult: condsult{
			field: Escape(field),
			op:    " NOT BETWEEN ",
			value: []interface{}{lower, upper},
		},
	}
}

// NotBetween represents "field NOT BETWEEN lower AND upper".
func (c *Cond) NotBetween(field string, lower, upper interface{}) string {
	buf := newStringBuilder()
	buf.WriteString(Escape(field))
	buf.WriteString(" NOT BETWEEN ")
	buf.WriteString(c.Args.Add(lower))
	buf.WriteString(" AND ")
	buf.WriteString(c.Args.Add(upper))
	return buf.String()
}

type andOrCondsult struct {
	expr []string
	condsult
}

func (cs *andOrCondsult) Init(s StringBuilder) {
	cs.s = s
}

func (cs *andOrCondsult) Value() interface{} {
	return cs.value
}

func (cs *andOrCondsult) WriteBefore() {
	cs.s.WriteString("(")
	cs.s.WriteStrings(cs.expr, cs.op)
}

func (cs *andOrCondsult) WriteString(index int, ph string) {
}

func (cs *andOrCondsult) WriteAfter() {
	cs.s.WriteString(")")
}

func (cs *andOrCondsult) Cond() string {
	return cs.s.String()
}

// Or represents OR logic like "expr1 OR expr2 OR expr3".
func Or(orExpr ...string) Condsult {
	return &andOrCondsult{
		expr: orExpr,
		condsult: condsult{
			op: " OR ",
		},
	}
}

// Or represents OR logic like "expr1 OR expr2 OR expr3".
func (c *Cond) Or(orExpr ...string) string {
	buf := newStringBuilder()
	buf.WriteString("(")
	buf.WriteStrings(orExpr, " OR ")
	buf.WriteString(")")
	return buf.String()
}

// And represents AND logic like "expr1 AND expr2 AND expr3".
func And(andExpr ...string) Condsult {
	return &andOrCondsult{
		expr: andExpr,
		condsult: condsult{
			op: " AND ",
		},
	}
}

// And represents AND logic like "expr1 AND expr2 AND expr3".
func (c *Cond) And(andExpr ...string) string {
	buf := newStringBuilder()
	buf.WriteString("(")
	buf.WriteStrings(andExpr, " AND ")
	buf.WriteString(")")
	return buf.String()
}

// Exists represents "EXISTS (subquery)".
func Exists(subquery interface{}) Condsult {
	return &inCondsult{
		sep: "",
		condsult: condsult{
			op:    "EXISTS ",
			value: subquery,
		},
	}
}

// Exists represents "EXISTS (subquery)".
func (c *Cond) Exists(subquery interface{}) string {
	buf := newStringBuilder()
	buf.WriteString("EXISTS (")
	buf.WriteString(c.Args.Add(subquery))
	buf.WriteString(")")
	return buf.String()
}

// Exists represents "EXISTS (subquery)".
func NotExists(subquery interface{}) Condsult {
	return &inCondsult{
		sep: "",
		condsult: condsult{
			op:    "NOT EXISTS ",
			value: subquery,
		},
	}
}

// NotExists represents "NOT EXISTS (subquery)".
func (c *Cond) NotExists(subquery interface{}) string {
	buf := newStringBuilder()
	buf.WriteString("NOT EXISTS (")
	buf.WriteString(c.Args.Add(subquery))
	buf.WriteString(")")
	return buf.String()
}

// Any represents "field op ANY (value...)".
func (c *Cond) Any(field, op string, value ...interface{}) string {
	vs := make([]string, 0, len(value))

	for _, v := range value {
		vs = append(vs, c.Args.Add(v))
	}

	buf := newStringBuilder()
	buf.WriteString(Escape(field))
	buf.WriteString(" ")
	buf.WriteString(op)
	buf.WriteString(" ANY (")
	buf.WriteStrings(vs, ", ")
	buf.WriteString(")")
	return buf.String()
}

// All represents "field op ALL (value...)".
func (c *Cond) All(field, op string, value ...interface{}) string {
	vs := make([]string, 0, len(value))

	for _, v := range value {
		vs = append(vs, c.Args.Add(v))
	}

	buf := newStringBuilder()
	buf.WriteString(Escape(field))
	buf.WriteString(" ")
	buf.WriteString(op)
	buf.WriteString(" ALL (")
	buf.WriteStrings(vs, ", ")
	buf.WriteString(")")
	return buf.String()
}

// Some represents "field op SOME (value...)".
func (c *Cond) Some(field, op string, value ...interface{}) string {
	vs := make([]string, 0, len(value))

	for _, v := range value {
		vs = append(vs, c.Args.Add(v))
	}

	buf := newStringBuilder()
	buf.WriteString(Escape(field))
	buf.WriteString(" ")
	buf.WriteString(op)
	buf.WriteString(" SOME (")
	buf.WriteStrings(vs, ", ")
	buf.WriteString(")")
	return buf.String()
}

// Var returns a placeholder for value.
func (c *Cond) Var(value interface{}) string {
	return c.Args.Add(value)
}
