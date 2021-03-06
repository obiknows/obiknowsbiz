package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/buffalo/middleware/ssl"
	"github.com/gobuffalo/envy"
	"github.com/unrolled/secure"

	"github.com/gobuffalo/buffalo/middleware/csrf"
	"github.com/gobuffalo/buffalo/middleware/i18n"
	"github.com/gobuffalo/packr"
	"github.com/obiknows/obiknowsbiz/models"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App

// T is the international translator
var T *i18n.Translator

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "_obiknowsbiz_session",
		})
		// Automatically redirect to SSL
		app.Use(ssl.ForceSSL(secure.Options{
			SSLRedirect:     ENV == "production",
			SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
		}))

		if ENV == "development" {
			app.Use(middleware.ParameterLogger)
		}

		// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
		// Remove to disable this.
		app.Use(csrf.New)

		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.PopTransaction)
		// Remove to disable this.
		app.Use(middleware.PopTransaction(models.DB))

		// Setup and use translations:
		var err error
		if T, err = i18n.New(packr.NewBox("../locales"), "en-US"); err != nil {
			app.Stop(err)
		}
		app.Use(T.Middleware())

		// Testing
		app.ServeFiles("/assets", assetsBox)
		app.GET("/routes", RoutesHandler)

		// Api Domain ([domain.com]/api/v1/...)
		api := app.Group("/api/v1")
		api.GET("/code/getRepos", CodeToGetRepos)

		// TheSitePages
		app.GET("/kubed", KubeHandler)

		app.GET("/", IndexHandler)
		app.GET("/vue", IndexHandler)
		app.GET("/home", HomeHandler)
		app.GET("/card", CardIndex)
		app.GET("/art", ArtIndex)
		app.GET("/knxws", ResearchIndex)
		app.GET("/crypto", CryptoIndex)
		app.GET("/code", CodeIndex)
		app.GET("/food", FoodIndex)
		app.GET("/pocket", PocketIndex)
		app.GET("/pocket/edit", PocketEdit)

		app.GET("/hemp", HempIndex)
		app.GET("/sounds", BeatsIndex)
		app.GET("/ipfs/index", IpfsIndex)
	}

	return app
}
