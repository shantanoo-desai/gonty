package main

import (
	"net/http"

	"shantanoo-desai/gonty/composer"

	"shantanoo-desai/gonty/expectation"

	"shantanoo-desai/gonty/joke"

	"shantanoo-desai/gonty/nobleman"

	"github.com/gin-gonic/gin"
)

// Health Check API
func getHealth(ctx *gin.Context) {
	ctx.String(http.StatusOK, "OK")
}

// Get a Baroque Composer
func getComposer(ctx *gin.Context) {
	genre := ctx.Params.ByName("genre")
	switch genre {
	case "baroque":
		ctx.JSON(http.StatusOK, composer.GreatestBaroqueComposer)
	default:
		ctx.String(http.StatusNotFound, "This is all getting rather silly!")

	}
}

// Get Jokes that may be deadly!
func getJokes(ctx *gin.Context) {

	check_funniest := ctx.Query("funniest")
	switch check_funniest {
	case "true": // So you want to hear the funniest Joke! Beware!
		ctx.JSON(http.StatusOK, joke.Jokes[0])
	case "false": // Safe Bet!
		ctx.JSON(http.StatusOK, joke.Jokes[1])
	default: // send all jokes if no parameters
		ctx.JSON(http.StatusOK, joke.Jokes)
	}
}

// Were you expecting something or someone?
func getExpectation(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, expectation.NeverExpected)
}

// Do you know about this Roman Nobleman?
func getNobleman(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nobleman.Biggus)
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", getHealth)
	router.GET("/composer/:genre", getComposer)
	router.GET("/jokes", getJokes)
	router.GET("/expect", getExpectation)
	router.GET("/history/nobleman", getNobleman)

	return router
}

func main() {
	router := setupRouter()
	router.Run()
}
