// 代码生成时间: 2025-09-12 19:37:25
package main

import (
    "os"
    testing "testing"

    "github.com/gobuffalo/buffalo/generators"
    "github.com/gobuffalo/buffalo/generators/assets/generators/test"
)

// runTests sets up and executes the automated testing suite for the Buffalo application.
func runTests() {
    os.Args = []string{os.Args[0], "test"}
    testApp := test.New(&generators.App{CmdFlags: generators.CmdFlags{NoDB: true}})
    app := testApp.App

    // Run the test suite.
    if err := app.Run(os.Args); err != nil {
        app.Stop(err)
    }
}

func main() {
    // Run the automated test suite.
    runTests()
}
