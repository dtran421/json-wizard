package convert

type ConvertStrategy interface {
	SetInput(input string)

	SetInputFile(inputFile string)

	/*
	 * Convert JSON to the specified output format.
	 */
	Convert() error
}
