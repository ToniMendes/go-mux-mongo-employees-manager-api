// Package enum defines domain-level status enumerations.
package enum

type Status int

const (
	StatusInactive Status = iota
	StatusActive
)

func (s Status) String() string {
	return [...]string{"Inactive", "Active"}[s]
}