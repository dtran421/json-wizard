package convert

type OutputFormat string

const (
	YAML OutputFormat = "yaml"
	XML  OutputFormat = "xml"
	TS   OutputFormat = "ts"
	GO   OutputFormat = "go"
	RS   OutputFormat = "rs"
)

type Extension string

func (f OutputFormat) GetExtension() Extension {
	switch f {
	case YAML:
		return ".yaml"
	case XML:
		return ".xml"
	case TS:
		return ".ts"
	case GO:
		return ".go"
	case RS:
		return ".rs"
	default:
		return ""
	}
}
