package controller

import "fmt"

//MetricIntValue type adapter
type MetricIntValue struct {
	V int
}

func (m MetricIntValue) String() string {
	return fmt.Sprint(m.V)
}

//MetricStrValue type adapter
type MetricStrValue struct {
	V string
}

func (m MetricStrValue) String() string {
	return fmt.Sprint("\"", m.V, "\"")
}
