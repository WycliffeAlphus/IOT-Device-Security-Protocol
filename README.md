# 🔒 IoT Device Security Protocol

## 📜 Overview

The IoT Device Security Protocol is a framework designed to secure IoT devices by providing authentication, encryption, and integrity for device communication and management. This protocol helps protect sensitive data, prevent unauthorized access, and ensure secure device updates.


# Certificate Setup
Run these commands to generate certificates:

```bash
mkdir -p certs && cd certs

# Generate CA
openssl genrsa -out ca.key 2048
openssl req -x509 -new -nodes -key ca.key -sha256 -days 3650 -out ca.crt -subj "/CN=My IoT CA"

# Generate server cert
openssl genrsa -out server.key 2048
openssl req -new -key server.key -out server.csr -subj "/CN=localhost"
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 365 -sha256

# Generate client cert (for mTLS)
openssl genrsa -out client.key 2048
openssl req -new -key client.key -out client.csr -subj "/CN=iot-device-1"
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 365 -sha256

# Cleanup
rm *.csr *.srl