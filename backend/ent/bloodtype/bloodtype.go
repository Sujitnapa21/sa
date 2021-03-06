// Code generated by entc, DO NOT EDIT.

package bloodtype

const (
	// Label holds the string label denoting the bloodtype type in the database.
	Label = "bloodtype"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"

	// EdgePatient holds the string denoting the patient edge name in mutations.
	EdgePatient = "patient"

	// Table holds the table name of the bloodtype in the database.
	Table = "bloodtypes"
	// PatientTable is the table the holds the patient relation/edge.
	PatientTable = "patients"
	// PatientInverseTable is the table name for the Patient entity.
	// It exists in this package in order to avoid circular dependency with the "patient" package.
	PatientInverseTable = "patients"
	// PatientColumn is the table column denoting the patient relation/edge.
	PatientColumn = "bloodtype_patient"
)

// Columns holds all SQL columns for bloodtype fields.
var Columns = []string{
	FieldID,
	FieldName,
}
