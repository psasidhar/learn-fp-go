// Generated by: gen
// TypeWriter: slice
// Directive: +gen on Int8

package num

// Int8Slice is a slice of type Int8. Use it where you would use []Int8.
type Int8Slice []Int8

// SumInt8 sums Int8 over elements in Int8Slice. See: http://clipperhouse.github.io/gen/#Sum
func (rcv Int8Slice) SumInt8(fn func(Int8) Int8) (result Int8) {
	for _, v := range rcv {
		result += fn(v)
	}
	return
}
