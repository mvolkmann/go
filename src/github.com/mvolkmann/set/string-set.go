// See set-demo/main.go for an example of using this.
package set

type empty struct{} // memory-efficient value for sets

// Strings is a set of strings.
type Strings map[string]empty

var present = empty{}

// Add adds a given value to the set.
func (set Strings) Add(value string) {
	set[value] = present
}

// Contains determines if a given value is in the set.
func (set Strings) Contains(value string) bool {
	_, ok := set[value]
	return ok
}

// Remove removes a given value from the set.
func (set Strings) Remove(value string) {
	delete(set, value)
}
