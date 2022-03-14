package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
)

func main() {
	var transformations = getTransformations()
	var modeMap = getModeMap(transformations)

	var transformation func(string) string
	var mode string
	var verbose bool

	flag.StringVar(&mode, "m", "kebab", fmt.Sprintf("One transformation: %v", strings.Join(getModes(transformations), ",")))
	flag.BoolVar(&verbose, "v", false, "Verbosity")
	flag.Parse()

	if t, ok := modeMap[mode]; ok {
		if verbose {
			fmt.Println("[debug] Mode:", t.Label)
		}
		transformation = t.Fn
	} else {
		fmt.Println("Unknown mode: ", mode)
		flag.Usage()
		return
	}

	if flag.NArg() < 1 {
		flag.Usage()
		return
	}

	if isTerminal() {
		for _, arg := range flag.Args() {
			os.Stdout.WriteString(transformation(arg))
			os.Stdout.WriteString("\n")
		}
	} else {
		ts := os.Getenv("IFS")
		if ts == "" {
			ts = " "
		}

		for i, arg := range flag.Args() {
			os.Stdout.WriteString(transformation(arg))

			if i < flag.NFlag() {
				os.Stdout.WriteString(ts)
			}
		}
	}
}

func isTerminal() bool {
	fileInfo, _ := os.Stdout.Stat()

	return (fileInfo.Mode() & os.ModeCharDevice) != 0
}

func getModeMap(transformations []*Transformation) map[string]*Transformation {
	modeMap := map[string]*Transformation{}

	for _, t := range transformations {
		for _, m := range t.Modes {
			modeMap[m] = t
		}
	}

	return modeMap
}

func getModes(transformations []*Transformation) []string {
	modes := make([]string, len(transformations))

	for i, t := range transformations {
		modes[i] = fmt.Sprintf("%v", t.Modes)
	}

	return modes
}

type Transformation struct {
	Label string
	Demo  string
	Modes []string
	Fn    func(string) string
}

func getTransformations() []*Transformation {
	return []*Transformation{
		{
			Label: "Kebab",
			Demo:  "hello-world",
			Modes: []string{"k", "kebab"},
			Fn:    strcase.ToKebab,
		},
		{
			Label: "Snake",
			Demo:  "hello_world",
			Modes: []string{"s", "snake"},
			Fn:    strcase.ToSnake,
		},
		{
			Label: "Camel",
			Demo:  "HelloWorld",
			Modes: []string{"c", "camel"},
			Fn:    strcase.ToCamel,
		},
		{
			Label: "Lower Camel",
			Demo:  "helloWorld",
			Modes: []string{"lc", "lowercamel", "lowerCamel"},
			Fn:    strcase.ToLowerCamel,
		},
	}
}
