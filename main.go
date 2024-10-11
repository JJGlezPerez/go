package main

import (
    "fmt"
    "net/http"
    "encoding/json"
)

// Estructura para definir un recurso (por ejemplo, un Usuario)
type Usuario struct {
    ID   int    `json:"id"`
    Nombre string `json:"nombre"`
}

// Función para manejar la solicitud en la ruta "/usuarios"
func obtenerUsuarios(w http.ResponseWriter, r *http.Request) {
    usuarios := []Usuario{
        {ID: 1, Nombre: "Juan"},
        {ID: 2, Nombre: "Ana"},
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(usuarios)
}

// Función principal
func main() {
    http.HandleFunc("/usuarios", obtenerUsuarios) // Ruta "/usuarios"

    fmt.Println("Servidor escuchando en el puerto 8080")
    http.ListenAndServe(":8080", nil)
}