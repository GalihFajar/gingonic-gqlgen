package Router

import (
	"github.com/gin-gonic/gin"

	"graphql_todos/graph"
	"graphql_todos/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"graphql_todos/middlewares"
)

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func SetupRouter() *gin.Engine{
	// Setting up Gin
	r := gin.Default()
	r.Use(middlewares.GinContextToContextMiddleware()) // The order is important! Put this first!
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	return r
}

