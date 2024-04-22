package utils

func GetIndent(level int) string {
	indent := ""
	for i := 0; i < level; i++ {
		indent += " "
	}

	return indent
}

func GetCustomIndent(level int, indentSize int) string {
	indent := ""
	for i := 0; i < level*indentSize; i++ {
		indent += " "
	}

	return indent
}
