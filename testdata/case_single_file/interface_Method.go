// Code generated by struct2interface; DO NOT EDIT.

package case_single_file

// MethodInterface ...
// Method describes the code and documentation
// tied into a method
type MethodInterface interface {
	// Lines return a []string consisting of
	// the documentation and code appended
	// in chronological order
	Lines() []string
}
