package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/logto-io/go/v2/client"
)

var (
	ContentTypeHtml = "text/html; charset=utf-8"
)

func main() {
	logtoConfig := &client.LogtoConfig{
		// see .env.example for more details and examples
		Endpoint:  os.Getenv("ENDPOINT"),
		AppId:     os.Getenv("APP_ID"),
		AppSecret: os.Getenv("APP_SECRET"),
		Resources: []string{},
	}
	resources, present := os.LookupEnv("RESOURCES")
	if present {
		logtoConfig.Resources = strings.Split(resources, ",")
	}

	store := memstore.NewStore([]byte("secret"))

	router := gin.Default()

	router.Use(sessions.Sessions("logto-session", store))

	router.GET("/", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		logtoClient := client.NewLogtoClient(logtoConfig, &SessionStorage{session: session})

		authState := "You are not logged in to this website. :("

		if logtoClient.IsAuthenticated() {
			authState = "You are logged in to this website! :)"
		}

		homePage := `<h1>Hello Logto</h1>` +
			"<div>" + authState + "</div>" +
			`<div><a href="/sign-in">Sign In</a></div>` +
			`<div><a href="/sign-out">Sign Out</a></div>` +
			`<div><a href="/user-id-token-claims">ID Token Claims</a></div>` +
			`<div><a href="/protected">Protected Resource</a></div>` +
			`<div><a href="/user-info">User Info</a></div>`

		ctx.Data(http.StatusOK, ContentTypeHtml, []byte(homePage))
	})

	router.GET("/sign-in", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		logtoClient := client.NewLogtoClient(logtoConfig, &SessionStorage{session: session})
		signInUri, err := logtoClient.SignInWithRedirectUri(os.Getenv("REDIRECT_URI"))
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.Redirect(http.StatusTemporaryRedirect, signInUri)
	})

	router.GET("/sign-in-callback", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		logtoClient := client.NewLogtoClient(logtoConfig, &SessionStorage{session: session})

		err := logtoClient.HandleSignInCallback(ctx.Request)
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Redirect(http.StatusTemporaryRedirect, "/")
	})

	router.GET("/user-id-token-claims", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		logtoClient := client.NewLogtoClient(logtoConfig, &SessionStorage{session: session})

		idTokenClaims, err := logtoClient.GetIdTokenClaims()

		if err != nil {
			ctx.String(http.StatusOK, err.Error())
		}

		ctx.JSON(http.StatusOK, idTokenClaims)
	})

	router.GET("/sign-out", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		logtoClient := client.NewLogtoClient(logtoConfig, &SessionStorage{session: session})

		signOutUri, signOutErr := logtoClient.SignOut(os.Getenv("POST_SIGN_OUT_REDIRECT_URI"))

		if signOutErr != nil {
			ctx.String(http.StatusOK, signOutErr.Error())
			return
		}

		ctx.Redirect(http.StatusTemporaryRedirect, signOutUri)
	})

	router.GET("/protected", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		logtoClient := client.NewLogtoClient(logtoConfig, &SessionStorage{session: session})

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

	router.GET("/user-info", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		logtoClient := client.NewLogtoClient(logtoConfig, &SessionStorage{session: session})

		if logtoClient.IsAuthenticated() {
			userInfoResponse, fetchUserInfoErr := logtoClient.FetchUserInfo()

			if fetchUserInfoErr != nil {
				ctx.String(http.StatusOK, fetchUserInfoErr.Error())
				return
			}

			ctx.JSON(http.StatusOK, userInfoResponse)
			return
		}

		unauthorizedPage := `
		<h1>Unauthorized</h1>
		<div><a href="/">Home</a></div>
		`
		ctx.Data(http.StatusOK, ContentTypeHtml, []byte(unauthorizedPage))
	})

	router.Run("0.0.0.0:8080")
}
