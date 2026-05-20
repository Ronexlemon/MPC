# MPC (Multi-Party Computation) Service
- [Description](#description)
- [Installation](#installation)
- [Features](#features)
- [Architecture](#architecture)
- [Installation](#installation)
- [Usage](#usage)
- [Api-Endpoints](#api-endpoints)
- [Project-Structure](#project-structure)
- [Contribution](#contribution)
- [License](#license)
- [Contact](#contact)

---
# Description
This project is a simple implementation of a Multi-Party Computation (MPC) system built with Go. The system enables secure sharing and computation of secrets(Private key in Evm Wallets infra) between multiple parties(different nodes) without revealing the original secret to any single participant.

The project demonstrates core MPC concepts such as:
- Secret sharing
- Distributed trust
- Secure computation 
- Cryptographic verification
- Isolated enclave communication

The architecture includes:
- **Gateway Service** — Handles external requests and communication.
- **Enclave Service** — Performs secure internal computation and secret processing.
- **Cryptographic Utilities** — Used for hashing, verification and secret handling.

This project can serve as a foundation for:
- Secure wallets
- Threshold signatures
- Distributed authentication systems
- Privacy-preserving applications
- Secure fintech infrastructure

---
# Features
- Secure secret provisioning
- MPC-inspired architecture
- REST API endpoints
- Sharmir secret sharing algorithm 
- JSON-based communication
- Dockerized services
- Lightweight Go implementation
---
# Architecture
```
Client
   │
   ▼
Gateway Service
   │
   ▼
Enclave Service
   │
   ▼
Secure Secret Processing
```
### Components

| component | Description |
| -------- | -------- |
| Gateway  | Receives client requests and forwards secure operations   |
| Enclave    | Handles protected computations internally   |
|Cryptography Layer | Manages hashing and secret verification   |
| Docker Environment |    Provides isolated deployment setup       |


# Installation
### Prerequisites
Make sure you have the following installed:
- Go 1.25+
- Docker
- Docker Compose

Clone the Repository

```
git clone https://github.com/ronexlemon/mpc.git
cd mpc


```
### Install dependecies
```
go mod download
```

### Run with Docker
```
docker compose up --build

```
---
# Usage
### Start the Application
```
docker compose up
```
---
# Api endpoints
### Provision Secret
```
curl -X POST http://localhost:8080/v1/vaults \      
  -H "X-API-Key: b2b_secret_auth_token" \         
  -H "Content-Type: application/json" \
  -d '{"workspace_id": "enterprise_corp_alpha"}'

```
### Example Response
```

```

### Signing a Transaction
```
curl -X POST http://localhost:8080/v1/transactions \
  -H "X-API-Key: b2b_secret_auth_token" \         
  -H "Content-Type: application/json" \
  -d '{                                         
    "workspace_id": "enterprise_corp_alpha",
    "to_address": "0x742d35Cc6634C0532925a3b844Bc454e4438f44e",
    "amount": "1000000000000000000",
    "nonce": 12
  }'

```
### Example Response
```

```
---
# Project Structure
```
mpc/
|---enclave
|     |---internal
|           |---handler
|     |---pkg
|          |---config
|          |---types
|     |---Dockerfile
|     |---main.go
|---gateway
|      |---cmd
|           |---api
|      |---internal
|              |---handler
|              |---helper
|      |---pkg
|           |---types
|      |---Dockerfile
|---docker-compose.yml
|----go.mod
|---go.sum
|--- README.md

```

---
# Contribution
Contributions are welcome.
### Steps
1. Fork the repository
2. Create a new branch
```
git checkout -b feature/my-feature
```
3. Commit your changes
```
git commit -m "Add new feature"
```
4. Push to your branch
```
git push origin feature/my-feature
```
5. Open a Pull Request
---
# License
This project is licensed under the MIT License.
---
# Contact
For questions, suggestions, or collaboration:
- GitHub:@ronexlemon
- Email: ronexlemon@gmail.com
