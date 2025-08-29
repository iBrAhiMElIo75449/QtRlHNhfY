// 代码生成时间: 2025-08-30 05:12:12
package main

import (
    "buffalo"
    "github.com/markbates/inflect"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// User model represents a user in the system.
type User struct {
    gorm.Model
    Username string `gorm:"unique;not null"`
    Password string `gorm:"not null"`
}

// NewUser creates a new User instance.
func NewUser(username, password string) *User {
    return &User{
        Username: username,
        Password: password,
    }
}

// Login checks if the provided credentials are valid.
func Login(db *gorm.DB, username, password string) error {
    var user User
    // Attempt to find the user by username.
    if err := db.Where(&User{Username: username}).First(&user).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return fmt.Errorf("username or password is incorrect")
        }
        return err
    }
    // Check if the password is correct.
    if user.Password != password {
        return fmt.Errorf("username or password is incorrect")
    }
    return nil
}

func main() {
    // Initialize the SQLite database.
    db, err := gorm.Open(sqlite.Open("login.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Migrate the schema.
    db.AutoMigrate(&User{})

    app := buffalo.Automatic()

    // Define the login route.
    app.GET("/login", func(c buffalo.Context) error {
        // Render the login template.
        return c.Render(200, r.HTML("login.html"))
    })

    app.POST("/login", func(c buffalo.Context) error {
        // Extract credentials from the request.
        username := c.Request().FormValue("username")
        password := c.Request().FormValue("password")

        // Validate credentials.
        if err := Login(db, username, password); err != nil {
            return c.Render(401, r.String("Invalid credentials"))
        }

        // If credentials are valid, redirect to the home page.
        return c.Redirect(302, "/")
    })

    // Start the Buffalo application.
    if err := app.Serve(); err != nil {
        panic(err)
    }
}
