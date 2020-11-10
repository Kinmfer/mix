package gin

import "github.com/gin-gonic/gin"

// New
func New(routeDefinitionCallbacks ...func(router *gin.Engine)) *gin.Engine {
	engine := gin.New()
	for _, callback := range routeDefinitionCallbacks {
		callback(engine)
	}
	return engine
}

const (
	// DebugMode indicates gin mode is debug.
	DebugMode = "debug"
	// ReleaseMode indicates gin mode is release.
	ReleaseMode = "release"
	// TestMode indicates gin mode is test.
	TestMode = "test"
)

// SetMode
func SetMode(value string) {
	gin.SetMode(value)
}
