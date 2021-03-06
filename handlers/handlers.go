package handlers

import (
	"github.com/Ignis-Divine/api-nodemcu/middleW"
	"github.com/Ignis-Divine/api-nodemcu/routers"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

//Manejadores, configuro el puerto, el handler y se pone en escucha el servidor
func Manejadores() {
	router := mux.NewRouter()
	router.Use(
		middleware.Logger, //log every http request
		middleware.Recoverer, //recover if panic occurs
	)
	//-----------------------------RUTAS------------------------------------//
	//-----------------------------LOGIN------------------------------------//
	router.HandleFunc("/login", middleW.RevisarDB(routers.Login)).Methods("POST")
	//-----------------------------USUARIOS---------------------------------//
	router.HandleFunc("/usuarios/registro", middleW.RevisarDB(routers.RegistrarUsuarioAdmin)).Methods("POST")
	router.HandleFunc("/usuarios/ver", middleW.RevisarDB(middleW.ValidoJWT(routers.VerUsuario))).Methods("GET")
	router.HandleFunc("/usuarios/modificar", middleW.RevisarDB(middleW.ValidoJWT(routers.ModificarUsuario))).Methods("PUT")
	router.HandleFunc("/usuarios/eliminar", middleW.RevisarDB(middleW.ValidoJWT(routers.EliminarUsuario))).Methods("DELETE")
	router.HandleFunc("/usuarios/lista", middleW.RevisarDB(middleW.ValidoJWT(routers.ListarUsuarios))).Methods("GET")
	//-----------------------------NODEMCUS---------------------------------//
	router.HandleFunc("/nodemcu/registro", middleW.RevisarDB(routers.RegistrarNodemcu)).Methods("POST")
	router.HandleFunc("/nodemcu/ver", middleW.RevisarDB(middleW.ValidoJWT(routers.VerNodemcu))).Methods("GET")
	router.HandleFunc("/nodemcu/modificar", middleW.RevisarDB(middleW.ValidoJWT(routers.ModificarNodemcu))).Methods("PUT")
	router.HandleFunc("/nodemcu/eliminar", middleW.RevisarDB(middleW.ValidoJWT(routers.EliminarNodemcu))).Methods("DELETE")
	router.HandleFunc("/nodemcu/lista", middleW.RevisarDB(middleW.ValidoJWT(routers.ListarNodemcus))).Methods("GET")
	//-----------------------------DATOS SENSORES---------------------------//
	router.HandleFunc("/nodemcu", middleW.RevisarDB(routers.CrearRegistroDatos)).Methods("POST")
	router.HandleFunc("/nodemcu", middleW.RevisarDB(routers.ListarDatos)).Methods("GET")
	//----------------------------------------------------------------------//

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "5050"
	}
	//da permisos a todo mundo
	handler := cors.AllowAll().Handler(router)
	//pone a escuchar el puerto elegido y pasa el control al handler
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
