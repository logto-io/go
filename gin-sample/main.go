package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/logto-io/go/client"
)

var (
	ContentTypeHtml = "text/html; charset=utf-8"
)

func main() {
	logtoConfig := &client.LogtoConfig{
		Endpoint:           "http://localhost:3001",
		AppId:              "wkMqPInVTZL0Ys4K3POl7",
		AppSecret:          "71kbRPQ1NAVA08IBStipz",
		Resources:          []string{"https://api.logto.io"},
		PersistAccessToken: true,
	}

	store := memstore.NewStore([]byte("secret"))

	router := gin.Default()

	router.Use(sessions.Sessions("logto-session", store))

	router.GET("/", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		logtoClient := client.NewLogtoClient(logtoConfig, &Storage{session: session})

		authState := "You are not logged in to this website. :("

		if logtoClient.IsAuthenticated() {
			authState = "You are logged in to this website! :)"
		}

		homePage := `<h1>Hello Logto</h1>` +
			"<div>" + authState + "</div>" +
			`<div><a href="/sign-in">Sign In</a></div>` +
			`<div><a href="/sign-out">Sign Out</a></div>` +
			`<div><a href="/user">Profile</a></div>` +
			`<div><a href="/protected">Protected Resource</a></div>`

		ctx.Data(http.StatusOK, ContentTypeHtml, []byte(homePage))
	})

	router.GET("/sign-in", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		logtoClient := client.NewLogtoClient(logtoConfig, &Storage{session: session})
		signInUri, err := logtoClient.SignIn("http://localhost:8080/sign-in-callback")
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.Redirect(http.StatusTemporaryRedirect, signInUri)
	})

	router.GET("/sign-in-callback", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		logtoClient := client.NewLogtoClient(logtoConfig, &Storage{session: session})

		err := logtoClient.HandleSignInCallback(ctx.Request)
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Redirect(http.StatusTemporaryRedirect, "/")
	})

	router.GET("/user", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		logtoClient := client.NewLogtoClient(logtoConfig, &Storage{session: session})

		idTokenClaims, err := logtoClient.GetIdTokenClaims()

		if err != nil {
			ctx.String(http.StatusOK, err.Error())
		}

		ctx.JSON(http.StatusOK, idTokenClaims)
	})

	router.GET("/sign-out", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		logtoClient := client.NewLogtoClient(logtoConfig, &Storage{session: session})

		signOutUri, signOutErr := logtoClient.SignOut("http://localhost:8080")

		if signOutErr != nil {
			ctx.String(http.StatusOK, signOutErr.Error())
			return
		}

		ctx.Redirect(http.StatusTemporaryRedirect, signOutUri)
	})

	router.GET("/protected", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		logtoClient := client.NewLogtoClient(logtoConfig, &Storage{session: session})

		if logtoClient.IsAuthenticated() {
			protectedPage := `
			<h1>Authenticated</h1>
			<div>Protected content</div>
			<div><a href="/">Home</a></div>
			`
			ctx.Data(http.StatusOK, ContentTypeHtml, []byte(protectedPage))
			return
		}

		unauthorizedPage := `
		<h1>Unauthorized</h1>
		<div>You cannot visit the protected content</div>
		<div><a href="/">Home</a></div>
		`
		ctx.Data(http.StatusOK, ContentTypeHtml, []byte(unauthorizedPage))
	})

	router.Run("0.0.0:8080")
}
