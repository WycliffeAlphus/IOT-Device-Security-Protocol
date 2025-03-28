package server

import (
	"encoding/json"
	"net/http"

	"iotsec/auth"

	"iotsec/device"
)

func registerDevice(w http.ResponseWriter, r *http.Request) {
	// Register device
	var dev device.Device
	if err := json.NewDecoder(r.Body).Decode(&dev); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// save device to the database
	device.SaveDevice(dev)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "Device registered successfully"})
}

func authenticateDevice(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GenerateJWT("device123")
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func updateFirmware(w http.ResponseWriter, r *http.Request) {}
