package handlers

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ptilotta/ecomgo/routers"
)

/*Manejadores seteo mi puerto, el Handler y pongo a escuchar al Servidor */
func Manejadores(event events.APIGatewayProxyRequest) (int, string) {

	switch event.Path {
	case "/signup":
		if event.HTTPMethod == http.MethodPost {
			return routers.Registro(event)
		}
	}

	return 200, "Todo OK"
	/*	PORT := os.Getenv("PORT")
		if PORT == "" {
			PORT = "8080"
		}
		handler := cors.AllowAll().Handler(router)
		log.Fatal(http.ListenAndServe(":"+PORT, handler)) */
}
