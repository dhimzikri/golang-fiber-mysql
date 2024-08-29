package main

import (
	"log"

	"github.com/dhimzikri/golang-fiber-mysql/initializers"
	"github.com/dhimzikri/golang-fiber-mysql/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	initializers.ConnectDB(&config)
}

func main() {
	app := fiber.New()

	// Adding CORS middleware with specific origin
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowMethods:     "GET,POST,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Content-Type,Authorization,Accept,Origin,Access-Control-Request-Method,Access-Control-Request-Headers,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Access-Control-Allow-Methods,Access-Control-Expose-Headers,Access-Control-Max-Age,Access-Control-Allow-Credentials",
		AllowCredentials: true,
	}))

	// Setup routes
	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":8000"))
}

// =====================old
// func main() {
// 	app := fiber.New()
// 	micro := fiber.New()

// 	app.Mount("/api", micro)
// 	app.Use(logger.New())
// 	app.Use(cors.New(cors.Config{
// 		AllowOrigins:     "http://localhost:3000",
// 		AllowMethods:     "GET,POST,PUT,DELETE,PATCH,OPTIONS",
// 		AllowHeaders:     "Content-Type,Authorization,Accept,Origin,Access-Control-Request-Method,Access-Control-Request-Headers,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Access-Control-Allow-Methods,Access-Control-Expose-Headers,Access-Control-Max-Age,Access-Control-Allow-Credentials",
// 		AllowCredentials: true,
// 	}))

// 	micro.Route("/notes", func(router fiber.Router) {
// 		router.Post("/", controllers.CreateNoteHandler)
// 		router.Get("", controllers.FindNotes)
// 	})
// 	micro.Route("/notes/:noteId", func(router fiber.Router) {
// 		router.Delete("", controllers.DeleteNote)
// 		router.Get("", controllers.FindNoteById)
// 		router.Patch("", controllers.UpdateNote)
// 	})
// 	micro.Get("/healthchecker", func(c *fiber.Ctx) error {
// 		return c.Status(200).JSON(fiber.Map{
// 			"status":  "success",
// 			"message": "Welcome to Golang, Fiber, MySQL, and GORM",
// 		})
// 	})

// 	log.Fatal(app.Listen(":8000"))
// }
