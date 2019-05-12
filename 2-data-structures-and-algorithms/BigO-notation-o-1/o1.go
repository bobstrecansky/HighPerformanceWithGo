package oone

func ThreeWords() string {
	threewords := [3]string{"foo", "bar", "baz"}
	return threewords[1]
}

func TenWords() string {
	tenwords := [10]string{"foo", "bar", "baz", "qux", "grault", "waldo", "plugh", "xyzzy", "thud", "spam"}
	return tenwords[6]
}
