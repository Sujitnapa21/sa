// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/Sujitnapa21/app/ent/status"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Status is the model entity for the Status schema.
type Status struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "Name" field.
	Name string `json:"Name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StatusQuery when eager-loading is set.
	Edges StatusEdges `json:"edges"`
}

// StatusEdges holds the relations/edges for other nodes in the graph.
type StatusEdges struct {
	// Patient holds the value of the patient edge.
	Patient []*Patient
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PatientOrErr returns the Patient value or an error if the edge
// was not loaded in eager-loading.
func (e StatusEdges) PatientOrErr() ([]*Patient, error) {
	if e.loadedTypes[0] {
		return e.Patient, nil
	}
	return nil, &NotLoadedError{edge: "patient"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Status) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Name
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Status fields.
func (s *Status) assignValues(values ...interface{}) error {
	if m, n := len(values), len(status.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	s.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Name", values[0])
	} else if value.Valid {
		s.Name = value.String
	}
	return nil
}

// QueryPatient queries the patient edge of the Status.
func (s *Status) QueryPatient() *PatientQuery {
	return (&StatusClient{config: s.config}).QueryPatient(s)
}

// Update returns a builder for updating this Status.
// Note that, you need to call Status.Unwrap() before calling this method, if this Status
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Status) Update() *StatusUpdateOne {
	return (&StatusClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (s *Status) Unwrap() *Status {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Status is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Status) String() string {
	var builder strings.Builder
	builder.WriteString("Status(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", Name=")
	builder.WriteString(s.Name)
	builder.WriteByte(')')
	return builder.String()
}

// StatusSlice is a parsable slice of Status.
type StatusSlice []*Status

func (s StatusSlice) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
