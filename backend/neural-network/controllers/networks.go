package controllers

import (
	"fastfoodrestaurant/neuralnetwork/intermediates"
	"fastfoodrestaurant/neuralnetwork/models/responses"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

const BasePath = "/network"
const trainPath = "/train"
const predictPath = "/ask"
const savePath = "/save"
const loadPath = "/load"

type NetworksController struct {
	networkIntermediate *intermediates.NetworkIntermediate
}

func NewNetworksController(networkIntermediate *intermediates.NetworkIntermediate) *NetworksController {
	return &NetworksController{
		networkIntermediate: networkIntermediate,
	}
}

func (n *NetworksController) saveNetwork(rw http.ResponseWriter, r *http.Request) {
	// filename string
	n.networkIntermediate.Save()
	rw.Write([]byte("Pesos guardados, revisar carpeta de backups"))
}

func (n *NetworksController) loadNetwork(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	direction, _ := ioutil.ReadAll(r.Body)
	path := string(direction)
	err := n.networkIntermediate.Load(path, 0)

	if err != nil {
		rw.Write([]byte("Ocurrio un error"))
		return
	}

	rw.Write([]byte("Pesos cargados exitosamente"))
}

func (n *NetworksController) trainChatBot(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	direction, _ := ioutil.ReadAll(r.Body)
	path := string(direction)

	n.networkIntermediate.Train(path, 500)
	rw.Write([]byte("Entrenamiento concluido"))
}

func (n *NetworksController) predict(rw http.ResponseWriter, r *http.Request) {
	messageResponse := &responses.MessageResponse{}
	if err := render.Bind(r, messageResponse); err != nil {
		return
	}

	message := n.networkIntermediate.PredictMessage(messageResponse.Message.Text)

	response := &responses.MessageResponse{Message: message}
	render.Render(rw, r, response)
}

func (n *NetworksController) GetRoutes() *chi.Mux {
	networkRouter := chi.NewRouter()

	networkRouter.Post(predictPath, n.predict)
	networkRouter.Post(trainPath, n.trainChatBot)
	networkRouter.Post(savePath, n.saveNetwork)
	networkRouter.Post(loadPath, n.loadNetwork)

	return networkRouter
}
