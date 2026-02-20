package main

import (
	"event-publisher/internal/http/handlers"
	"event-publisher/internal/rabbit"
	"log"
	"net/http"
)

func main() {

	cfg := rabbit.LoadConfig()

	conn := rabbit.NewConnection(cfg)

	rabbit.InitTopology(conn, cfg)

	publisher := rabbit.NewPublisher(conn, cfg)

	http.HandleFunc("/api/xml-recebido", handlers.XmlRecebidoHandler(publisher))

	log.Println("API publisher rodando :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
