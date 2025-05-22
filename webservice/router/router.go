package route

import (
	"fmt"
	"net/http"

	"github.com/MayuraJam/StitchAlienAPIGO/webservice/controller"
)

func corsMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-control-allow-origin", "*")
		w.Header().Add("Access-control-allow-method", "POST , GET , OPTIONS , PUT , DELETE")
		w.Header().Add("Access-control-allow-headers", "Accept , Content-Type , Content-Lenght , Authorization")
		handler.ServeHTTP(w, r)
	})
}
func SetupRoutes(apiBasePath string, path string) {
	creaturesHandler := http.HandlerFunc(controller.HandlerCreatures)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, path), corsMiddleware(creaturesHandler))
	creatureHandler := http.HandlerFunc(controller.HandlerCreature)
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, path), corsMiddleware(creatureHandler))
}
