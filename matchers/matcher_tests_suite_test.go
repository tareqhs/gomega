package matchers_test

import (
	"fmt"
	"io"
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type myStringer struct {
	a string
}

func (s *myStringer) String() string {
	return s.a
}

type StringAlias string

type myCustomType struct {
	s   string
	n   int
	f   float32
	arr []string
}

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gomega Matchers")
}

func readFileContents(filePath string) []byte {
	f := openFile(filePath)
	b, err := io.ReadAll(f)
	if err != nil {
		panic(fmt.Errorf("failed to read file contents: %v", err))
	}
	return b
}

func openFile(filePath string) *os.File {
	f, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Errorf("failed to open file: %v", err))
	}
	return f
}
