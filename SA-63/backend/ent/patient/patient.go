// Code generated by entc, DO NOT EDIT.

package patient

const (
	// Label holds the string label denoting the patient type in the database.
	Label = "patient"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"

	// EdgeStatistic holds the string denoting the statistic edge name in mutations.
	EdgeStatistic = "statistic"

	// Table holds the table name of the patient in the database.
	Table = "patients"
	// StatisticTable is the table the holds the statistic relation/edge.
	StatisticTable = "statistics"
	// StatisticInverseTable is the table name for the Statistic entity.
	// It exists in this package in order to avoid circular dependency with the "statistic" package.
	StatisticInverseTable = "statistics"
	// StatisticColumn is the table column denoting the statistic relation/edge.
	StatisticColumn = "patient_statistic"
)

// Columns holds all SQL columns for patient fields.
var Columns = []string{
	FieldID,
	FieldGender,
}