package routes

import(
	"github.com/gorilla/mux"
	"github.com/Akanom/Golang-Auto-Shop-Management-APIs/pkg/controllers"
)

var RegisterAutoShopRoutes=func(router *mux.Router){
	router.HandleFunc("/auto/", controllers.CreateNewAuto).Methods("POST")
	router.HandleFunc("/auto/", controllers.GetAuto).Methods("GET")
	router.HandleFunc("/auto/{autoid}", controllers.GetAutoById).Methods("GET")
	router.HandleFunc("/auto/{autoid}", controllers.UpdateAuto).Methods("PUT")
	router.HandleFunc("/auto/{autoid}",controllers.DeleteAuto).Methods("DELETE")
}