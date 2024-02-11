package main

import (
	"app/database"
	"app/middleware"
	"app/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setupRoutes(app *fiber.App) {

	app.Post("/api/posts", middleware.DeserializeUser, routes.CreatePost)
	app.Put("/api/posts/:id", middleware.DeserializeUser, routes.UpdatePost)
	app.Delete("/api/posts/:id", middleware.DeserializeUser, routes.DeletePost)

	app.Post("/api/posts/:postId/comments", middleware.DeserializeUser, routes.CreateComment)
	app.Post("/api/comments/:commentId/replies", middleware.DeserializeUser, routes.CreateReply)
	app.Get("/api/posts/:postId/comments", middleware.DeserializeUser, routes.GetComments)

	app.Post("/api/posts/like", middleware.DeserializeUser, routes.LikePost)
	app.Delete("/api/posts/like/:id", middleware.DeserializeUser, routes.TakeBackPostLike)

	app.Post("/api/comment/like", middleware.DeserializeUser, routes.LikeComment)
	app.Delete("/api/comment/like/:id", middleware.DeserializeUser, routes.TakeBackCommentLike)

	app.Post("/api/users/friendship", middleware.DeserializeUser, routes.BeFriend)
	app.Delete("/api/users/friendship/:id", middleware.DeserializeUser, routes.RemoveFriendship)
}

func main() {
	database.ConnectDb()

	app := fiber.New()
	micro := fiber.New()
	app.Mount("/api", micro)
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST",
		AllowCredentials: true,
	}))

	micro.Route("/auth", func(router fiber.Router) {
		router.Post("/register", routes.SignUpUser)
		router.Post("/login", routes.SignInUser)
		router.Get("/logout", middleware.DeserializeUser, routes.LogoutUser)
	})

	micro.Get("/users/me", middleware.DeserializeUser, routes.GetMe)

	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))

}
