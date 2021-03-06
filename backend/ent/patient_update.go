// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/Sujitnapa21/app/ent/bloodtype"
	"github.com/Sujitnapa21/app/ent/employee"
	"github.com/Sujitnapa21/app/ent/gender"
	"github.com/Sujitnapa21/app/ent/nametitle"
	"github.com/Sujitnapa21/app/ent/patient"
	"github.com/Sujitnapa21/app/ent/predicate"
	"github.com/Sujitnapa21/app/ent/status"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// PatientUpdate is the builder for updating Patient entities.
type PatientUpdate struct {
	config
	hooks      []Hook
	mutation   *PatientMutation
	predicates []predicate.Patient
}

// Where adds a new predicate for the builder.
func (pu *PatientUpdate) Where(ps ...predicate.Patient) *PatientUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetIdcard sets the Idcard field.
func (pu *PatientUpdate) SetIdcard(s string) *PatientUpdate {
	pu.mutation.SetIdcard(s)
	return pu
}

// SetName sets the Name field.
func (pu *PatientUpdate) SetName(s string) *PatientUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetAddress sets the Address field.
func (pu *PatientUpdate) SetAddress(s string) *PatientUpdate {
	pu.mutation.SetAddress(s)
	return pu
}

// SetCongenital sets the Congenital field.
func (pu *PatientUpdate) SetCongenital(s string) *PatientUpdate {
	pu.mutation.SetCongenital(s)
	return pu
}

// SetAllergic sets the Allergic field.
func (pu *PatientUpdate) SetAllergic(s string) *PatientUpdate {
	pu.mutation.SetAllergic(s)
	return pu
}

// SetEmployeeID sets the employee edge to Employee by id.
func (pu *PatientUpdate) SetEmployeeID(id int) *PatientUpdate {
	pu.mutation.SetEmployeeID(id)
	return pu
}

// SetNillableEmployeeID sets the employee edge to Employee by id if the given value is not nil.
func (pu *PatientUpdate) SetNillableEmployeeID(id *int) *PatientUpdate {
	if id != nil {
		pu = pu.SetEmployeeID(*id)
	}
	return pu
}

// SetEmployee sets the employee edge to Employee.
func (pu *PatientUpdate) SetEmployee(e *Employee) *PatientUpdate {
	return pu.SetEmployeeID(e.ID)
}

// SetStatusID sets the status edge to Status by id.
func (pu *PatientUpdate) SetStatusID(id int) *PatientUpdate {
	pu.mutation.SetStatusID(id)
	return pu
}

// SetNillableStatusID sets the status edge to Status by id if the given value is not nil.
func (pu *PatientUpdate) SetNillableStatusID(id *int) *PatientUpdate {
	if id != nil {
		pu = pu.SetStatusID(*id)
	}
	return pu
}

// SetStatus sets the status edge to Status.
func (pu *PatientUpdate) SetStatus(s *Status) *PatientUpdate {
	return pu.SetStatusID(s.ID)
}

// SetBloodtypeID sets the bloodtype edge to Bloodtype by id.
func (pu *PatientUpdate) SetBloodtypeID(id int) *PatientUpdate {
	pu.mutation.SetBloodtypeID(id)
	return pu
}

// SetNillableBloodtypeID sets the bloodtype edge to Bloodtype by id if the given value is not nil.
func (pu *PatientUpdate) SetNillableBloodtypeID(id *int) *PatientUpdate {
	if id != nil {
		pu = pu.SetBloodtypeID(*id)
	}
	return pu
}

// SetBloodtype sets the bloodtype edge to Bloodtype.
func (pu *PatientUpdate) SetBloodtype(b *Bloodtype) *PatientUpdate {
	return pu.SetBloodtypeID(b.ID)
}

// SetGenderID sets the gender edge to Gender by id.
func (pu *PatientUpdate) SetGenderID(id int) *PatientUpdate {
	pu.mutation.SetGenderID(id)
	return pu
}

// SetNillableGenderID sets the gender edge to Gender by id if the given value is not nil.
func (pu *PatientUpdate) SetNillableGenderID(id *int) *PatientUpdate {
	if id != nil {
		pu = pu.SetGenderID(*id)
	}
	return pu
}

// SetGender sets the gender edge to Gender.
func (pu *PatientUpdate) SetGender(g *Gender) *PatientUpdate {
	return pu.SetGenderID(g.ID)
}

// SetNametitleID sets the nametitle edge to NameTitle by id.
func (pu *PatientUpdate) SetNametitleID(id int) *PatientUpdate {
	pu.mutation.SetNametitleID(id)
	return pu
}

// SetNillableNametitleID sets the nametitle edge to NameTitle by id if the given value is not nil.
func (pu *PatientUpdate) SetNillableNametitleID(id *int) *PatientUpdate {
	if id != nil {
		pu = pu.SetNametitleID(*id)
	}
	return pu
}

// SetNametitle sets the nametitle edge to NameTitle.
func (pu *PatientUpdate) SetNametitle(n *NameTitle) *PatientUpdate {
	return pu.SetNametitleID(n.ID)
}

// Mutation returns the PatientMutation object of the builder.
func (pu *PatientUpdate) Mutation() *PatientMutation {
	return pu.mutation
}

// ClearEmployee clears the employee edge to Employee.
func (pu *PatientUpdate) ClearEmployee() *PatientUpdate {
	pu.mutation.ClearEmployee()
	return pu
}

// ClearStatus clears the status edge to Status.
func (pu *PatientUpdate) ClearStatus() *PatientUpdate {
	pu.mutation.ClearStatus()
	return pu
}

// ClearBloodtype clears the bloodtype edge to Bloodtype.
func (pu *PatientUpdate) ClearBloodtype() *PatientUpdate {
	pu.mutation.ClearBloodtype()
	return pu
}

// ClearGender clears the gender edge to Gender.
func (pu *PatientUpdate) ClearGender() *PatientUpdate {
	pu.mutation.ClearGender()
	return pu
}

// ClearNametitle clears the nametitle edge to NameTitle.
func (pu *PatientUpdate) ClearNametitle() *PatientUpdate {
	pu.mutation.ClearNametitle()
	return pu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *PatientUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PatientMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PatientUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PatientUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PatientUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PatientUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   patient.Table,
			Columns: patient.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patient.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Idcard(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldIdcard,
		})
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldName,
		})
	}
	if value, ok := pu.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldAddress,
		})
	}
	if value, ok := pu.mutation.Congenital(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldCongenital,
		})
	}
	if value, ok := pu.mutation.Allergic(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldAllergic,
		})
	}
	if pu.mutation.EmployeeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.EmployeeTable,
			Columns: []string{patient.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.EmployeeTable,
			Columns: []string{patient.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.StatusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.StatusTable,
			Columns: []string{patient.StatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: status.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.StatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.StatusTable,
			Columns: []string{patient.StatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: status.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.BloodtypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.BloodtypeTable,
			Columns: []string{patient.BloodtypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bloodtype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.BloodtypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.BloodtypeTable,
			Columns: []string{patient.BloodtypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bloodtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.GenderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.GenderTable,
			Columns: []string{patient.GenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gender.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.GenderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.GenderTable,
			Columns: []string{patient.GenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gender.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.NametitleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.NametitleTable,
			Columns: []string{patient.NametitleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nametitle.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.NametitleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.NametitleTable,
			Columns: []string{patient.NametitleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nametitle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{patient.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PatientUpdateOne is the builder for updating a single Patient entity.
type PatientUpdateOne struct {
	config
	hooks    []Hook
	mutation *PatientMutation
}

// SetIdcard sets the Idcard field.
func (puo *PatientUpdateOne) SetIdcard(s string) *PatientUpdateOne {
	puo.mutation.SetIdcard(s)
	return puo
}

// SetName sets the Name field.
func (puo *PatientUpdateOne) SetName(s string) *PatientUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetAddress sets the Address field.
func (puo *PatientUpdateOne) SetAddress(s string) *PatientUpdateOne {
	puo.mutation.SetAddress(s)
	return puo
}

// SetCongenital sets the Congenital field.
func (puo *PatientUpdateOne) SetCongenital(s string) *PatientUpdateOne {
	puo.mutation.SetCongenital(s)
	return puo
}

// SetAllergic sets the Allergic field.
func (puo *PatientUpdateOne) SetAllergic(s string) *PatientUpdateOne {
	puo.mutation.SetAllergic(s)
	return puo
}

// SetEmployeeID sets the employee edge to Employee by id.
func (puo *PatientUpdateOne) SetEmployeeID(id int) *PatientUpdateOne {
	puo.mutation.SetEmployeeID(id)
	return puo
}

// SetNillableEmployeeID sets the employee edge to Employee by id if the given value is not nil.
func (puo *PatientUpdateOne) SetNillableEmployeeID(id *int) *PatientUpdateOne {
	if id != nil {
		puo = puo.SetEmployeeID(*id)
	}
	return puo
}

// SetEmployee sets the employee edge to Employee.
func (puo *PatientUpdateOne) SetEmployee(e *Employee) *PatientUpdateOne {
	return puo.SetEmployeeID(e.ID)
}

// SetStatusID sets the status edge to Status by id.
func (puo *PatientUpdateOne) SetStatusID(id int) *PatientUpdateOne {
	puo.mutation.SetStatusID(id)
	return puo
}

// SetNillableStatusID sets the status edge to Status by id if the given value is not nil.
func (puo *PatientUpdateOne) SetNillableStatusID(id *int) *PatientUpdateOne {
	if id != nil {
		puo = puo.SetStatusID(*id)
	}
	return puo
}

// SetStatus sets the status edge to Status.
func (puo *PatientUpdateOne) SetStatus(s *Status) *PatientUpdateOne {
	return puo.SetStatusID(s.ID)
}

// SetBloodtypeID sets the bloodtype edge to Bloodtype by id.
func (puo *PatientUpdateOne) SetBloodtypeID(id int) *PatientUpdateOne {
	puo.mutation.SetBloodtypeID(id)
	return puo
}

// SetNillableBloodtypeID sets the bloodtype edge to Bloodtype by id if the given value is not nil.
func (puo *PatientUpdateOne) SetNillableBloodtypeID(id *int) *PatientUpdateOne {
	if id != nil {
		puo = puo.SetBloodtypeID(*id)
	}
	return puo
}

// SetBloodtype sets the bloodtype edge to Bloodtype.
func (puo *PatientUpdateOne) SetBloodtype(b *Bloodtype) *PatientUpdateOne {
	return puo.SetBloodtypeID(b.ID)
}

// SetGenderID sets the gender edge to Gender by id.
func (puo *PatientUpdateOne) SetGenderID(id int) *PatientUpdateOne {
	puo.mutation.SetGenderID(id)
	return puo
}

// SetNillableGenderID sets the gender edge to Gender by id if the given value is not nil.
func (puo *PatientUpdateOne) SetNillableGenderID(id *int) *PatientUpdateOne {
	if id != nil {
		puo = puo.SetGenderID(*id)
	}
	return puo
}

// SetGender sets the gender edge to Gender.
func (puo *PatientUpdateOne) SetGender(g *Gender) *PatientUpdateOne {
	return puo.SetGenderID(g.ID)
}

// SetNametitleID sets the nametitle edge to NameTitle by id.
func (puo *PatientUpdateOne) SetNametitleID(id int) *PatientUpdateOne {
	puo.mutation.SetNametitleID(id)
	return puo
}

// SetNillableNametitleID sets the nametitle edge to NameTitle by id if the given value is not nil.
func (puo *PatientUpdateOne) SetNillableNametitleID(id *int) *PatientUpdateOne {
	if id != nil {
		puo = puo.SetNametitleID(*id)
	}
	return puo
}

// SetNametitle sets the nametitle edge to NameTitle.
func (puo *PatientUpdateOne) SetNametitle(n *NameTitle) *PatientUpdateOne {
	return puo.SetNametitleID(n.ID)
}

// Mutation returns the PatientMutation object of the builder.
func (puo *PatientUpdateOne) Mutation() *PatientMutation {
	return puo.mutation
}

// ClearEmployee clears the employee edge to Employee.
func (puo *PatientUpdateOne) ClearEmployee() *PatientUpdateOne {
	puo.mutation.ClearEmployee()
	return puo
}

// ClearStatus clears the status edge to Status.
func (puo *PatientUpdateOne) ClearStatus() *PatientUpdateOne {
	puo.mutation.ClearStatus()
	return puo
}

// ClearBloodtype clears the bloodtype edge to Bloodtype.
func (puo *PatientUpdateOne) ClearBloodtype() *PatientUpdateOne {
	puo.mutation.ClearBloodtype()
	return puo
}

// ClearGender clears the gender edge to Gender.
func (puo *PatientUpdateOne) ClearGender() *PatientUpdateOne {
	puo.mutation.ClearGender()
	return puo
}

// ClearNametitle clears the nametitle edge to NameTitle.
func (puo *PatientUpdateOne) ClearNametitle() *PatientUpdateOne {
	puo.mutation.ClearNametitle()
	return puo
}

// Save executes the query and returns the updated entity.
func (puo *PatientUpdateOne) Save(ctx context.Context) (*Patient, error) {

	var (
		err  error
		node *Patient
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PatientMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PatientUpdateOne) SaveX(ctx context.Context) *Patient {
	pa, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pa
}

// Exec executes the query on the entity.
func (puo *PatientUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PatientUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PatientUpdateOne) sqlSave(ctx context.Context) (pa *Patient, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   patient.Table,
			Columns: patient.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patient.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Patient.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.Idcard(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldIdcard,
		})
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldName,
		})
	}
	if value, ok := puo.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldAddress,
		})
	}
	if value, ok := puo.mutation.Congenital(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldCongenital,
		})
	}
	if value, ok := puo.mutation.Allergic(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldAllergic,
		})
	}
	if puo.mutation.EmployeeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.EmployeeTable,
			Columns: []string{patient.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.EmployeeTable,
			Columns: []string{patient.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.StatusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.StatusTable,
			Columns: []string{patient.StatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: status.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.StatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.StatusTable,
			Columns: []string{patient.StatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: status.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.BloodtypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.BloodtypeTable,
			Columns: []string{patient.BloodtypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bloodtype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.BloodtypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.BloodtypeTable,
			Columns: []string{patient.BloodtypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bloodtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.GenderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.GenderTable,
			Columns: []string{patient.GenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gender.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.GenderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.GenderTable,
			Columns: []string{patient.GenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gender.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.NametitleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.NametitleTable,
			Columns: []string{patient.NametitleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nametitle.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.NametitleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   patient.NametitleTable,
			Columns: []string{patient.NametitleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nametitle.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	pa = &Patient{config: puo.config}
	_spec.Assign = pa.assignValues
	_spec.ScanValues = pa.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{patient.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pa, nil
}
