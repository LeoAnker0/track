package main

import (
    "track-server/internal/routes"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "mime"
    "fmt"
    "os"
)

func main() {
    // Initialize MIME types
    mime.AddExtensionType(".js", "application/javascript")

    // Create a new Gin router
    r := gin.Default()
    // Increase maximum file size to 2GB
    r.MaxMultipartMemory = 2 << 30 // 2GB

    // Use CORS middleware
    config := cors.DefaultConfig()
    config.AllowOrigins = []string{"http://localhost:*", "https://track.la0.uk"} // Change this to the specific origins you want to allow
    config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
    config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
    config.AllowCredentials = true
    r.Use(cors.New(config))

    // Create Gin routes
    routes.SetupRoutes(r)

    // Run the application on port 6000
     hostPort := os.Getenv("ADMIN_PORT")
    if hostPort == "" {
        hostPort = "6050" // Default port if not set
    }

    r.Run(fmt.Sprintf(":%s", hostPort))
}