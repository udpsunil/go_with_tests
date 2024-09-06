package iteration

// Repeat repeats a given character a specified number of times.
//
// ExampleRepeat:
//    Repeat("a", 3) returns "aaa"
//    Repeat("hello", 2) returns "hellohello"
func Repeat(character string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated = repeated + character
	}
	return repeated
}
