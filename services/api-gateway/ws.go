package main

import (
	"log"
	"net/http"
	"ride-sharing/shared/contracts"
	"ride-sharing/shared/util"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleRidersWebSocket(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Websocket upgrade Fail %s", err)
		return
	}

	defer conn.Close()

	userID := r.URL.Query().Get("userID")
	if userID == "" {
		log.Printf("userID not provided")
		return
	}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Unable to read Message %s", err)
			break
		}

		log.Printf("Received Message :%s", message)
	}
}

func handleDriversWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Websocket upgrade Fail %s", err)
		return
	}

	defer conn.Close()

	userID := r.URL.Query().Get("userID")
	if userID == "" {
		log.Printf("userID not provided")
		return
	}

	packageSlug := r.URL.Query().Get("packageSlug")
	if packageSlug == "" {
		log.Printf("No packageSlug provided")
		return
	}

	type Driver struct {
		Id             string //`json:="id"`
		Name           string //`json:="name"`
		ProfilePicture string //`json:="profilePicture"`
		CarPlate       string //`json:="carPlate"`
		PackageSlug    string //`json:="packageSlug"`
	}

	msg := contracts.WSMessage{
		Type: "driver.cmd.register",
		Data: Driver{
			Id:             userID,
			Name:           "Tiago",
			ProfilePicture: util.GetRandomAvatar(1),
			CarPlate:       "ABC123",
			PackageSlug:    packageSlug,
		},
	}

	if err := conn.WriteJSON(msg); err != nil {
		log.Printf("Error sending message: %v", err)
		return
	}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Unable to read Message %s", err)
			break
		}

		log.Printf("Received Message :%s", message)
	}
}
