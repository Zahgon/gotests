package process

import (
	"io"
	"os"
	"regexp"

	"github.com/cweill/gotests"
)

const newFilePerm os.FileMode = 0644

const (
	specifyFlagMessage = "Please specify either the -only, -excl, -exported, or -all flag"
	specifyFileMessage = "Please specify a file or directory containing the source"
)

type Options struct {
	OnlyFuncs          string
	ExclFuncs          string
	ExportedFuncs      bool
	AllFuncs           bool
	PrintInputs        bool
	Subtests           bool
	Parallel           bool
	Named              bool
	WriteOutput        bool
	Template           string
	TemplateDir        string
	TemplateParamsPath string
	TemplateParams     string
	TemplateData       [][]byte
	UseGoCmp           bool
	UseAI              bool
	AIModel            string
	AIEndpoint         string
	AIMinCases         int
	AIMaxCases         int
}

func Run(out io.Writer, args []string, opts *Options) { _ = "STUB: not implemented"; return }

func parseOptions(out io.Writer, opt *Options) *gotests.Options {
	_ = "STUB: not implemented"
	return nil
}

func parseRegexp(s string) (*regexp.Regexp, error) { _ = "STUB: not implemented"; return nil, nil }

func generateTests(out io.Writer, path string, writeOutput bool, opt *gotests.Options) {
	_ = "STUB: not implemented"
	return
}

func outputTest(out io.Writer, t *gotests.GeneratedTest, writeOutput bool) {
	_ = "STUB: not implemented"
	return
}
