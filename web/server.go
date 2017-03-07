package web

import (
	"fmt"
	"path/filepath"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mdouchement/lss/config"
	"github.com/mdouchement/lss/controllers"
	"github.com/mdouchement/lss/web/middlewares"
	"gopkg.in/urfave/cli.v2"
)

var (
	// ServerCommand defines the server command (CLI).
	ServerCommand = &cli.Command{
		Name:    "server",
		Aliases: []string{"s"},
		Usage:   "start server",
		Action:  serverAction,
		Flags:   serverFlags,
	}

	serverFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  "p, port",
			Usage: "Specify the port to listen to.",
		},
		&cli.StringFlag{
			Name:  "b, binding",
			Usage: "Binds server to the specified IP.",
		},
		&cli.StringFlag{
			Name:  "w, workspace",
			Usage: "Specify the files' workspace (overrides LSS_WORKSPACE)",
		},
		&cli.StringFlag{
			Name:  "s, size",
			Usage: "Specify the upload size limit (default: 8G - overrides LSS_UPLOAD_SIZE_LIMIT)",
		},
		&cli.StringFlag{
			Name:  "n, namespace",
			Usage: "Specify the router namespace (overrides LSS_ROUTER_NAMESPACE)",
		},
	}
)

func serverAction(context *cli.Context) error {
	if w := context.String("w"); w != "" {
		config.Cfg.Workspace = filepath.Clean(w)
	}
	if n := context.String("n"); n != "" {
		config.Cfg.RouterNamespace = n
	}
	if size := context.String("s"); size != "" {
		config.Cfg.UploadSizeLimit = size
	}
	engine := EchoEngine()
	printRoutes(engine)

	listen := fmt.Sprintf("%s:%s", context.String("b"), context.String("p"))
	config.Log.Infof("Upload limit: %s", config.Cfg.UploadSizeLimit)
	config.Log.Infof("Server listening on %s", listen)
	engine.Start(listen)

	return nil
}

// EchoEngine instanciates the LSS server.
func EchoEngine() *echo.Echo {
	engine := echo.New()
	engine.Use(middleware.Recover())
	engine.Use(middleware.BodyLimit(config.Cfg.UploadSizeLimit))
	if config.Env() == config.Production {
		engine.Use(middlewares.ProductionEchorus())
	} else {
		engine.Use(middlewares.DefaultEchorus())
	}
	engine.HTTPErrorHandler = middlewares.HTTPErrorHandler

	router := engine.Group(config.Cfg.RouterNamespace)

	router.GET("/version", controllers.Version)
	router.HEAD("/*", controllers.Exist)
	router.GET("/*", func(c echo.Context) error {
		if _, ok := c.QueryParams()["list"]; ok {
			return controllers.ListFiles(c)
		} else if _, ok := c.QueryParams()["metadata"]; ok {
			return controllers.Metadata(c)
		} else {
			return controllers.Download(c)
		}
	})
	router.POST("/*", controllers.Upload)
	router.DELETE("/*", controllers.Delete)

	return engine
}

func printRoutes(e *echo.Echo) {
	fmt.Println("Routes:")
	for _, route := range e.Routes() {
		fmt.Printf("%6s %s\n", route.Method, route.Path)
	}
}
