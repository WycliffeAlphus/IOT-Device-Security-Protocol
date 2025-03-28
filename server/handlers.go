package server

import (
	"encoding/json"
	"net/http"

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

func authenticateDevice(w http.ResponseWriter, r *http.Request) {}

func updateFirmware(w http.ResponseWriter, r *http.Request) {}
