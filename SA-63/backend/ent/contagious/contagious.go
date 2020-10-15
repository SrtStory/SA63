// Code generated by entc, DO NOT EDIT.

package contagious

const (
	// Label holds the string label denoting the contagious type in the database.
	Label = "contagious"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"

	// EdgeStatistic holds the string denoting the statistic edge name in mutations.
	EdgeStatistic = "statistic"

	// Table holds the table name of the contagious in the database.
	Table = "contagious"
	// StatisticTable is the table the holds the statistic relation/edge.
	StatisticTable = "statistics"
	// StatisticInverseTable is the table name for the Statistic entity.
	// It exists in this package in order to avoid circular dependency with the "statistic" package.
	StatisticInverseTable = "statistics"
	// StatisticColumn is the table column denoting the statistic relation/edge.
	StatisticColumn = "contagious_statistic"
)

// Columns holds all SQL columns for contagious fields.
var Columns = []string{
	FieldID,
	FieldName,
}