package webhandler

import (
	"go-forum-thingy/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html"
	"github.com/rs/zerolog/log"
	"net/http"
)

func NewApp() {
	engine := html.New("./webhandler/views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(filesystem.New(filesystem.Config{
		Root: http.Dir("./views/assets"),
		Browse: true,
		NotFoundFile: "Golang fucked up, I swear to god...",
	}))
	app.Static("/stylesheet", "./webhandler/views/assets", fiber.Static{
		Browse: true,
		Index: "main.css",
	})
	app.Use(favicon.New(favicon.Config{
		File: "./webhandler/views/assets/favicon.ico",
	}))

	routeSetup(app)
	err := app.Listen(":8080")
	if err != nil {
		log.Fatal().Err(err)
	}
}

func routeSetup(app *fiber.App) {
	app.Get("/register", registerpage)
	app.Post("/api/register", controllers.Register)
	app.Get("/login", loginpage)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/logout", controllers.Logout)

	app.Get("/api/user", controllers.User)

	app.Get("/", home)
}

func home(c *fiber.Ctx) error {
return c.Render("index", fiber.Map{
            "Title": "Hello, World!",
        })
}

func registerpage(c *fiber.Ctx) error {
	return c.Render("register", fiber.Map{
	})
}

func loginpage(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{

	})
}
