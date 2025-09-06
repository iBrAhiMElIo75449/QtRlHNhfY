// 代码生成时间: 2025-09-06 09:05:14
package main

import (
    "buffalo"
    "crypto/sha256"
    "encoding/hex"
    "net/http"
    "strings"
    "log"
)

// HashCalculatorController is the controller for calculating hash values
type HashCalculatorController struct {
    // Actions stores the actions for the controller
    Actions *buffalo.Actions
}

// CalcSHA256Action calculates the SHA-256 hash of a given input
func (c *HashCalculatorController) CalcSHA256Action(w http.ResponseWriter, r *http.Request) error {
    // Get the input from the request
    input := r.FormValue("input")
    if input == "" {
        // Return an error if the input is empty
        return buffalo.NewError(http.StatusBadRequest, "Input cannot be empty")
    }

    // Calculate the SHA-256 hash of the input
    hash := sha256.Sum256([]byte(input))
    hexHash := hex.EncodeToString(hash[:])

    // Return the calculated hash as a JSON response
    return Render(w, r, buffalo.R{
        "hash": hexHash,
    })
}

// Render is a helper function to render JSON responses
func Render(w http.ResponseWriter, r *http.Request, data buffalo.R) error {
    w.Header().Set("Content-Type", "application/json")
    return buffalo.Writer(w, data)
}

// main is the entry point of the application
func main() {
    // Create a new Buffalo application
    app := buffalo.New(buffalo.Options{})

    // Create a new instance of HashCalculatorController
    controller := &HashCalculatorController{Actions: buffalo.NewActions(controller)}

    // Register the CalcSHA256Action with the app
    app.GET("/hash", controller.CalcSHA256Action)

    // Run the application
    app.Serve()
}
