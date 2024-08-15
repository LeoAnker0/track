// static_file_routes.go

package routes

import (
    "github.com/gin-gonic/gin"
    "path/filepath"
    "net/http"
    "time"
    "fmt"
    "os"
)

// Middleware function to check TTL
func TTLMiddleware(ttl time.Duration) gin.HandlerFunc {
    return func(c *gin.Context) {
        startTime := time.Now()

        // Process the request
        c.Next()

        // Check if the elapsed time exceeds TTL
        elapsed := time.Since(startTime)
        if elapsed > ttl {
            // If TTL exceeded, return a 404 response
            fmt.Println("Request took too long. TTL exceeded.")
            c.AbortWithStatusJSON(404, gin.H{"error": "Request took too long"})
        }
    }
}

func SetupStaticFileRoutes(router *gin.Engine) {
    // Define TTL duration of 30 days in seconds
    ttlDuration := 30 * 24 * time.Hour

    // Apply TTLMiddleware to all routes
    router.Use(TTLMiddleware(ttlDuration))

    // For serving index.html (from /) and all web assets (css/js/svg...)
    router.NoRoute(handleCoreWebFiles)
}

func handleCoreWebFiles(c *gin.Context) {
    path := c.Request.URL.Path

    // If the path is "/", serve index.html
    if path == "/" {
        path = "/index.html"
    }

    // Construct the full file path
    filePath := "./global-resources/track-server/core-web" + path

    // Check if the file exists
    if fileInfo, err := os.Stat(filePath); os.IsNotExist(err) {
        // File does not exist
        // Check if the path lacks an extension
        if filepath.Ext(path) == "" {
            // check if the path is actually a page, and should therefore be processed as such?
            
            // Serve 404.html if the path doesn't have an extension
            filePath = "./global-resources/track-server/template-pages/404.html"

        } else {
            // Otherwise, return a 404 error
            c.String(http.StatusNotFound, "File not found")
            return
        }
    } else if fileInfo.IsDir() {
        // check if the path is actually a page, and should therefore be processed as such?
        
        // If the path points to a directory, serve index.html within that directory
        filePath ="./global-resources/admin-server/template-pages/404.html"

    }

    // Serve the file
    c.File(filePath)
}
