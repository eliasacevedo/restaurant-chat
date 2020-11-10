package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	botController "fastfoodrestaurant.com/api/controllers/bot"
	userController "fastfoodrestaurant.com/api/controllers/user"
	botsService "fastfoodrestaurant.com/api/services/bot"
	usersService "fastfoodrestaurant.com/api/services/users"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

const version = "/v1"

const port string = ":9090"

func getHandle() *chi.Mux {
	userService := usersService.NewUserService()
	botService := botsService.NewBotService()
	mainRouter := chi.NewRouter()

	mainRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	mainRouter.Use(render.SetContentType(render.ContentTypeJSON))
	botControllers := botController.NewBotController(botService)

	userControllers := userController.NewUserController(userService)

	mainRouter.Route(version, func(r chi.Router) {
		r.Mount(botController.BotBasePath, botControllers.GetRoutes())
		r.Mount(userController.UserBasePath, userControllers.GetRoutes())
	})

	return mainRouter
}

func main() {
	handle := getHandle()

	l := log.New(os.Stdout, "FastFood - ApiGateway ", log.LstdFlags)

	s := http.Server{
		Addr:         port,
		Handler:      handle,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// start the server
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
