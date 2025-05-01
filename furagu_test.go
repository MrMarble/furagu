package furagu_test

import (
	"flag"
	"fmt"
	"os"

	"github.com/mrmarble/furagu"
)

var (
	debugVar = furagu.Bool("debug", false, "Enable debug mode")
	fooVar   = furagu.String("foo", "default", "Foo variable")
	barVar   = furagu.Int("bar", 0, "Bar variable")
)

func Example() {

	// Go test command injects some arguments into os.Args.
	// We need to remove them to avoid parsing errors.
	os.Args = os.Args[:1]
	flag.CommandLine = flag.NewFlagSet("--", flag.ContinueOnError)
	flag.CommandLine.SetOutput(os.Stdout)

	// Provide the prefix for environment variables.
	furagu.Parse("EXAMPLE_USE")

	// Use the flags.
	if *debugVar {
		fmt.Println(*fooVar, *barVar)
	}

	// Flag automatic usage generation.
	flag.PrintDefaults()

	// Output:
	//  -bar int
	//     	Bar variable (env: "EXAMPLE_USE_BAR")
	//   -debug
	//     	Enable debug mode (env: "EXAMPLE_USE_DEBUG")
	//   -foo string
	//     	Foo variable (env: "EXAMPLE_USE_FOO") (default "default")
}
