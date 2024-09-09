package app

import (
	"cellify_backend/features/addusers/controller"
	mobilecontroller "cellify_backend/features/mobile_products/mobile_controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func AppRoutes() {
	userRoutes := controller.UserController{}
	mobileInfoRoute := mobilecontroller.MobileInfoController{}
	r := mux.NewRouter()
	r.HandleFunc("/save_user_info", userRoutes.SaveUserInfo).Methods("POST")
	r.HandleFunc("/save_mobile_info", mobileInfoRoute.SaveMobileInfo).Methods("POST")
	r.HandleFunc("/login_user",userRoutes.LoginUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}
