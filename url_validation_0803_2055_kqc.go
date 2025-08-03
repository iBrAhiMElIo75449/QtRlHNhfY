// 代码生成时间: 2025-08-03 20:55:48
* Usage:
*   - Install BUFFALO: $ buffalo -w
*   - Start the server: $ buffalo dev
*
* Endpoint: POST /validate-url
* Payload:
*   {"url": "http://example.com"}
*
* Response:
*   {"isValid": true, "message": "URL is valid"} or {"isValid": false, "message": "URL is not valid"}
*
*/

package main

import (
	"buffalo"
	"github.com/markbates/validate"
	"net/http"
	"net/url"
)

// ValidateURLRequest defines the request payload structure.
type ValidateURLRequest struct {
	URL string `json:"url" validate:"required,url"`
}

// ValidateURLResponse defines the response payload structure.
type ValidateURLResponse struct {
	IsValid bool   `json:"isValid"`
	Message string `json:"message"`
}

// validateURLHandler handles the validation of a URL.
func validateURLHandler(c buffalo.Context) error {
	// Parse the request payload.
	payload := &ValidateURLRequest{}
	if err := c.Bind(payload); err != nil {
		return err
	}

	// Validate the URL using the validator package.
	validateErr := validate.New().Struct(payload)
	if validateErr != nil {
		return c.Render(http.StatusBadRequest, r.JSON(ValidateURLResponse{
			IsValid: false,
			Message: validateErr.Error(),
		}))
	}

	// If validation passes, return a success message.
	return c.Render(http.StatusOK, r.JSON(ValidateURLResponse{
		IsValid: true,
		Message: "URL is valid",
	}))
}

// main is the entry point of the application.
func main() {
	app := buffalo.Automatic()

	// Add the handler for the URL validation endpoint.
	app.POST("/validate-url", validateURLHandler)

	// Start the server.
	app.Serve()
}
