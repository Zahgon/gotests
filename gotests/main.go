package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cweill/gotests/gotests/process"
)

var (
	onlyFuncs          = flag.String("only", "", `regexp. generate tests for functions and methods that match only. Takes precedence over -all`)
	exclFuncs          = flag.String("excl", "", `regexp. generate tests for functions and methods that don't match. Takes precedence over -only, -exported, and -all`)
	exportedFuncs      = flag.Bool("exported", false, `generate tests for exported functions and methods. Takes precedence over -only and -all`)
	allFuncs           = flag.Bool("all", false, "generate tests for all functions and methods")
	printInputs        = flag.Bool("i", false, "print test inputs in error messages")
	writeOutput        = flag.Bool("w", false, "write output to (test) files instead of stdout")
	templateDir        = flag.String("template_dir", "", `optional. Path to a directory containing custom test code templates. Takes precedence over -template. This can also be set via environment variable GOTESTS_TEMPLATE_DIR`)
	template           = flag.String("template", "", `optional. Specify custom test code templates, e.g. testify. This can also be set via environment variable GOTESTS_TEMPLATE`)
	templateParamsPath = flag.String("template_params_file", "", "read external parameters to template by json with file")
	templateParams     = flag.String("template_params", "", "read external parameters to template by json with stdin")
	useGoCmp           = flag.Bool("use_go_cmp", false, "use cmp.Equal (google/go-cmp) instead of reflect.DeepEqual")
	useAI              = flag.Bool("ai", false, "generate test cases using AI (requires Ollama)")
	aiModel            = flag.String("ai-model", "qwen2.5-coder:0.5b", "AI model to use for test generation")
	aiEndpoint         = flag.String("ai-endpoint", "http://localhost:11434", "Ollama API endpoint")
	aiMinCases         = flag.Int("ai-min-cases", 3, "minimum number of test cases to generate with AI")
	aiMaxCases         = flag.Int("ai-max-cases", 10, "maximum number of test cases to generate with AI")
	version            = flag.Bool("version", false, "print version information and exit")
)

var (
	nosubtests = true

	parallel bool

	named bool
)

func main() {
	flag.Parse()
	args := flag.Args()

	if *version {
		printVersion()
		return
	}

	if *useAI {

		fmt.Fprintf(os.Stderr, "⚠️  WARNING: Function source code will be sent to AI provider at %s\n", *aiEndpoint)
		fmt.Fprintf(os.Stderr, "   Ensure your code does not contain secrets or sensitive information.\n\n")

		if *aiModel == "" {
			fmt.Fprintf(os.Stderr, "Error: -ai-model cannot be empty when using -ai flag\n")
			os.Exit(1)
		}
		if *aiMinCases < 1 {
			fmt.Fprintf(os.Stderr, "Error: -ai-min-cases must be at least 1, got %d\n", *aiMinCases)
			os.Exit(1)
		}
		if *aiMaxCases > 100 {
			fmt.Fprintf(os.Stderr, "Error: -ai-max-cases must be at most 100, got %d\n", *aiMaxCases)
			os.Exit(1)
		}
		if *aiMinCases > *aiMaxCases {
			fmt.Fprintf(os.Stderr, "Error: -ai-min-cases (%d) cannot be greater than -ai-max-cases (%d)\n", *aiMinCases, *aiMaxCases)
			os.Exit(1)
		}
	}

	process.Run(os.Stdout, args, &process.Options{
		OnlyFuncs:          *onlyFuncs,
		ExclFuncs:          *exclFuncs,
		ExportedFuncs:      *exportedFuncs,
		AllFuncs:           *allFuncs,
		PrintInputs:        *printInputs,
		Subtests:           !nosubtests,
		Parallel:           parallel,
		Named:              named,
		WriteOutput:        *writeOutput,
		Template:           valOrGetenv(*template, "GOTESTS_TEMPLATE"),
		TemplateDir:        valOrGetenv(*templateDir, "GOTESTS_TEMPLATE_DIR"),
		TemplateParamsPath: *templateParamsPath,
		TemplateParams:     *templateParams,
		UseGoCmp:           *useGoCmp,
		UseAI:              *useAI,
		AIModel:            *aiModel,
		AIEndpoint:         *aiEndpoint,
		AIMinCases:         *aiMinCases,
		AIMaxCases:         *aiMaxCases,
	})
}

func printVersion() { _ = "STUB: not implemented"; return }

func valOrGetenv(val, key string) string { _ = "STUB: not implemented"; return "" }
