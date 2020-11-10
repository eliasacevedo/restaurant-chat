package bot

import (
	"net/http"

	"fastfoodrestaurant.com/api/models"
	services "fastfoodrestaurant.com/api/services/bot"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type BotController struct {
	botService services.IBotService
}

func NewBotController(s services.IBotService) *BotController {
	return &BotController{
		botService: s,
	}
}

const BotBasePath = "/bot"
const botQuestionPath = "/ask"

func (c *BotController) GetRoutes() *chi.Mux {
	botRouter := chi.NewRouter()
	botRouter.Post(botQuestionPath, c.askQuestion)
	return botRouter
}

func (c *BotController) askQuestion(rw http.ResponseWriter, r *http.Request) {
	message := &models.Message{}
	if err := render.Bind(r, message); err != nil {
		return
	}

	response, _ := c.botService.AskQuestion(message.Text)
	render.Render(rw, r, response)
	return
}
