package route

import (
	"fmt"
	"net/http"

	"github.com/MayuraJam/StitchAlienAPIGO/webservice/controller"
)

// cors ย่อมาจาก Cross-Origin Resource Sharing ช่วยให้ทาง client สามารถเรียกใช้ API ได้ข้าม server ได้อย่างสะดวก
func corsMiddleware(handler http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-control-allow-origin", "*")
		w.Header().Add("Access-control-allow-method", "POST , GET , OPTIONS , PUT , DELETE")
		w.Header().Add("Access-control-allow-headers", "Accept , Content-Type , Content-Lenght , Authorization")
		handler.ServeHTTP(w, r)
	})
}
func SetupRoutes(apiBasePath string, path string) {
	//controller สำหรับดึงข้อมูลที่ไม่ต้องใช้ Id
	creaturesHandler := http.HandlerFunc(controller.HandlerCreatures)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, path), corsMiddleware(creaturesHandler))

	//controller สำหรับดึงข้อมูลที่ต้องใช้ Id
	creatureHandler := http.HandlerFunc(controller.HandlerCreature)
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, path), corsMiddleware(creatureHandler))
}
