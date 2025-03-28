package firmware

import (
	"iotsec/config"
	"iotsec/crypt"
)

type Firmware struct {
	Version   string `json:"version"`
	Data      string `json:"data"`
	Signature string `json:"signature"`
}

func VerifyFirmware(fw Firmware) bool {
	return crypt.VerifyHMAC(fw.Data, config.HMACKey, fw.Signature)
}
