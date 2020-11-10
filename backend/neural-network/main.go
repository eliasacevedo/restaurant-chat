package main

import (
	"context"
	"fastfoodrestaurant/neuralnetwork/controllers"
	"fastfoodrestaurant/neuralnetwork/intermediates"
	"fastfoodrestaurant/neuralnetwork/services"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

const version = "/v1"

func getHandle() *chi.Mux {

	word2vecService := services.NewWordService()
	dataTrainService := &services.DataTrainService{}
	networkIntermediate := intermediates.NewTextClassificationNetwork(word2vecService, dataTrainService)
	networkController := controllers.NewNetworksController(networkIntermediate)

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

	mainRouter.Route(version, func(r chi.Router) {
		r.Mount(controllers.BasePath, networkController.GetRoutes())
	})

	return mainRouter
}

const port = ":9091"

func main() {
	handle := getHandle()

	l := log.New(os.Stdout, "FastFood - Neural network ", log.LstdFlags)

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
		l.Println("Starting server on port 9091")

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

// func main() {
// service := &services.DataTrainService{}
// data := service.GetData("https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/training/chats")

// separatorSentence := "#"
// separatorClassification := "("

// h, err := os.Open("backups/hweights.model")
// vec := mat.NewVecDense(3, nil)
// vec.Reset()
// vec.UnmarshalBinaryFrom(h)

// if err == nil {
// 	return
// }

// defer h.Close()

// activationFunction := &models.SygmoidFunction{}
// lossFunction := &models.BinaryCrossEntropy{}

// yHat := activationFunction.GetActivationValue(-0.350)
// y := 1.0
// fmt.Println(yHat)
// fmt.Println(lossFunction.ErrorBase(yHat, y))
// }
