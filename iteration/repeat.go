package iteration

func Repeat(char string, numrepeats int) string {
	var name string
	for i := 0; i < numrepeats; i++ {
		name += char
	}
	return name
}
