package device


type Device struct {
	ID string `json:"device_id"`
	PublicKey string `json:"public_key"`
}

var devices []Device

func SaveDevice(dev Device) {
	devices = append(devices, dev)
}