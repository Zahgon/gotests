package gotests

import (
	"go/types"
	"regexp"

	"github.com/cweill/gotests/internal/goparser"
	"github.com/cweill/gotests/internal/models"
)

type Options struct {
	Only           *regexp.Regexp
	Exclude        *regexp.Regexp
	Exported       bool
	PrintInputs    bool
	Subtests       bool
	Parallel       bool
	Named          bool
	Importer       func() types.Importer
	Template       string
	TemplateDir    string
	TemplateParams map[string]interface{}
	TemplateData   [][]byte
	UseGoCmp       bool
	UseAI          bool
	AIModel        string
	AIEndpoint     string
	AIMinCases     int
	AIMaxCases     int
}

type GeneratedTest struct {
	Path      string
	Functions []*models.Function
	Output    []byte
}

func GenerateTests(srcPath string, opt *Options) ([]*GeneratedTest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type result struct {
	gt  *GeneratedTest
	err error
}

func parallelize(srcFiles, files []models.Path, opt *Options) ([]*GeneratedTest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readResults(rs <-chan *result) ([]*GeneratedTest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func generateTest(src models.Path, files []models.Path, opt *Options) (*GeneratedTest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseTestFile(p *goparser.Parser, testPath string, h *models.Header) (*models.Header, []string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func testableFuncs(funcs []*models.Function, only, excl *regexp.Regexp, exp bool, testFuncs []string) []*models.Function {
	_ = "STUB: not implemented"
	return nil
}

func isInvalid(f *models.Function) bool { _ = "STUB: not implemented"; return false }

func isTestFunction(f *models.Function, testFuncs []string) bool {
	_ = "STUB: not implemented"
	return false
}

func isExcluded(f *models.Function, excl *regexp.Regexp) bool {
	_ = "STUB: not implemented"
	return false
}

func isUnexported(f *models.Function, exp bool) bool { _ = "STUB: not implemented"; return false }

func isIncluded(f *models.Function, only *regexp.Regexp) bool {
	_ = "STUB: not implemented"
	return false
}

func contains(ss []string, s string) bool { _ = "STUB: not implemented"; return false }
