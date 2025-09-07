// 代码生成时间: 2025-09-08 03:15:08
package main

import (
    "fmt"
    "log"
    "os"
    "time"

    "github.com/gobuffalo/buffalo"
    "github.com/gobuffalo/pop"
    "github.com/gobuffalo/buffalo/generators"
    "github.com/gobuffalo/buffalo/generators/assets/esbuild"
)

// TestDatum represents the data structure for generated test data
type TestDatum struct {
    ID        uint      `db:"id"`
    Name      string    `db:"name"`
    CreatedAt time.Time `db:"created_at"`
    UpdatedAt time.Time `db:"updated_at"`
}

// TestDataGenerator is the main struct for the test data generator
type TestDataGenerator struct {
    DB *pop.Connection
}

// NewTestDatum generates a new TestDatum instance with random data
func NewTestDatum() TestDatum {
    return TestDatum{
        Name:      fmt.Sprintf("Test %d", time.Now().UnixNano()),
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }
}

// Generate generates a specified number of test data entries
func (g *TestDataGenerator) Generate(count int) error {
    for i := 0; i < count; i++ {
        td := NewTestDatum()
        err := g.DB.Create(&td)
        if err != nil {
            return err
        }
    }
    return nil
}

// main is the entry point for the test data generator application
func main() {
    // Initialize the Buffalo application
    app := buffalo.Automatic()
    app.Serve()

    // Create a new TestDataGenerator instance
    generator := TestDataGenerator{DB: buffalo.DB().Connection("default")}

    // Define the number of test data entries to generate
    testDataCount := 100 // You can adjust this value as needed

    // Generate the test data entries
    err := generator.Generate(testDataCount)
    if err != nil {
        log.Fatalf("Error generating test data: %s", err)
    } else {
        fmt.Printf("Successfully generated %d test data entries.
", testDataCount)
    }
}
