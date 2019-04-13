package main
///////////////////////////////////////  IMPORTANDO LIBRERIAS   ///////////////////////////

//importando librerias necesarias
import (
	//json: Sirve para codificar y decodificar los mensajes que se reciben y envian en formato JSON
	"encoding/json"
	//log: para visualizar los errores del servidor
	"log"
	//http: Para escribir toda la funcionalidad del servidor
	"net/http"
	//mux: Para las rutas de API
	"github.com/gorilla/mux"
)


/////////////////////////////////////   CREANDO LA ESTRUCTURA (MODELOS)  ////////////////////////

//estructura de como el cliente enviara los datos y como el servidor los reenviara
// crearemos un tipo de dato persona, para almacenar las personas
type Person struct{
	// los datos a recibir seran en formato json
	// el ID lo enviara con el nombre de id (minuscula) y que no este vacio, para poder almacenarlo
	ID string `json:"id,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	LastName string `json:"lastname,omitempty"`
	tlf string `json:"tlf,omitempty"`
	Direccion *Direccion `json:"direccion,omitempty"`
}
//Difinicion de la direccion para las persona
type Direccion struct{
	Ciudad string `json:"ciudad,omitempty"`
	Estado string `json:"estado,omitempty"`
}

//////////////////////////////////    DEFINIENDO ARRAY (BASE DE DATOS)    /////////////////////////////

//Arreglo de persona para almacenar, hacemos referencia a la estructura Person
var ArrPersonasBD []Person


////////////////////////////////   DEFINIENDO FUNCIONES (CRUD)    /////////////////////////////////////
//Funcion para obtener los datos de todas las personas en la BD
//recibe dos parametros w: para enviar respuestas y req: para recibir solicitures
func GetPersonsEndPoint(w http.ResponseWriter, req *http.Request) {
	
	//codificamos los datos en formato json y los mostramos
	//Encode: para que el cliente entienda lo que el servidor quiere enviar
	json.NewEncoder(w).Encode(ArrPersonasBD)

}


//Funcion para obtener los datos de una sola persona
func GetPersonEndPoint(w http.ResponseWriter, req *http.Request) {
	
	//guardamos el dato que el cliente quiere, en este caso el id que se pasa por el request
	params := mux.Vars(req)
	//for para iterar en el array y buscar el id
	for _, item := range ArrPersonasBD {
		//si encuentra el id retorna los datos de la persona
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	//Si no encuentra la persona, retornara una persona en blanco
	json.NewEncoder(w).Encode(&Person{})


}

//funcion para crear una persona
func CreatePersonEndPoint(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	//crearemos una persona de tipo Person
	var person Person

	//Decode: para que el servidor entienda lo que el cliente envia
	//almacenamos los datos en variable person
	_ = json.NewDecoder(req.Body).Decode(&person)
	// agregamos el id a la persona recivido por request
	person.ID = params ["id"]
	// lo agregamos al array
	ArrPersonasBD = append(ArrPersonasBD, person)
	//lo mostramos en formato json
	json.NewEncoder(w).Encode(ArrPersonasBD)


}

//funcion para borrar una persona
func DeletePersonEndPoint(w http.ResponseWriter, req *http.Request) {
	
	// guardamos el request en la variable params
	params := mux.Vars(req)
	//iteramos en el arreglo para buscar y eliminar la persona
	for index, item := range ArrPersonasBD{
		if item.ID == params["id"] {
			//eliminamos a la persona en la posicion del arreglo index
			 ArrPersonasBD = append (ArrPersonasBD[:index],  ArrPersonasBD[index +1:]...)
			 break
		}
	}
	// mostramos todas las personas
	json.NewEncoder(w).Encode(ArrPersonasBD)

}

////////////////////////////////////////  PRINCIPAL   /////////////////////////////////////////////
func main() {
	//Definicion de rutas de la API
	router := mux.NewRouter()

	//llenamos el arreglo con append: recibe el arreglo y los datos
	ArrPersonasBD = append(ArrPersonasBD, Person{ID:"1", FirstName: "Darwin", LastName: "Garcia", tlf: "02372832398", Direccion: &Direccion{Ciudad:"Guayana", Estado:"Bolivar"}})
	ArrPersonasBD = append(ArrPersonasBD, Person{ID:"2", FirstName: "Juan", LastName: "Garcia", tlf: "02372832398"})
	ArrPersonasBD = append(ArrPersonasBD, Person{ID:"3", FirstName: "Manuel", LastName: "Perez", tlf: "03456452398"})

	//endpoints (rutas del API)
	//Ruta para obtener todas las personas recibe por parametro el nombre de la ruta y la funcion a ejecutar, se implementa el metodo GET
	router.HandleFunc("/persons", GetPersonsEndPoint).Methods("GET")
	//Ruta para obtener los datos de una sola persona
	router.HandleFunc("/person/{id}", GetPersonEndPoint).Methods("GET")
	//Ruta para crear una persona
	router.HandleFunc("/person/{id}", CreatePersonEndPoint).Methods("POST")
	//Ruta pata eliminar una persona
	router.HandleFunc("/person/{id}", DeletePersonEndPoint).Methods("DELETE")

	// las rutas pueden tener el mismo nombre pero con diferentes metodos lo que hace que no interfiera una con la otra


	//ejecutamos nuestro servidor, le pasamos el puerto de escucha y las routas
	log.Fatal(http.ListenAndServe(":3000", router))
}