package providers

import (
	"fmt"
	"os/exec"

	"github.com/dtran421/json-wizard/types"
	"github.com/golang-collections/collections/set"
)

var SourceLanguage *set.Set = set.New(
	types.JSON,
	types.YAML,
)

var TargetLanguage *set.Set = set.New(
	types.TS,
	types.GO,
	types.RS,
)

type FlagArg struct {
	Flag  string
	Value string
}

func (f FlagArg) String() string {
	return fmt.Sprintf("%s=%s", f.Flag, f.Value)
}

type QuicktypeWrapper struct {
	srcLang    types.OutputFormat
	targetLang types.OutputFormat

	flagArgs []FlagArg
}

func NewQuicktypeWrapper(srcLang types.OutputFormat, targetLang types.OutputFormat) QuicktypeWrapper {
	if !SourceLanguage.Has(srcLang) {
		panic("Invalid source language")
	}

	if !TargetLanguage.Has(targetLang) {
		panic("Invalid target language")
	}

	return QuicktypeWrapper{
		srcLang:    srcLang,
		targetLang: targetLang,

		flagArgs: make([]FlagArg, 0),
	}
}

func (q *QuicktypeWrapper) SetOutFile(outFile types.Filepath) {
	q.AddFlagArg("--out", outFile.String())
}

func (q *QuicktypeWrapper) AddFlagArg(flag string, value string) {
	q.flagArgs = append(q.flagArgs, FlagArg{flag, value})
}

func (q *QuicktypeWrapper) String() string {
	cmd := fmt.Sprintf("quicktype %s %s", q.srcLang, q.targetLang)

	for _, arg := range q.flagArgs {
		cmd += fmt.Sprintf(" %s", arg)
	}

	return cmd
}

func (q *QuicktypeWrapper) Execute() {
	cmd := exec.Command(q.String())

	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
