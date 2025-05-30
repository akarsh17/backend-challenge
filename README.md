# 🍔 Order Food Online API

A modular, production-ready RESTful API built with Golang and Gin for managing food products and placing orders. Based on OpenAPI 3.1.0.

---
### ⚠️ Note on `couponbase.gz` Files

> The `couponbase.gz` files are **large** and have not been included in this repository.  
> Please manually place them under the `./data/` directory.

#### Preprocessing Information

- The preprocessed output is stored at: `data/valid_coupons.json`.
- If you wish to **re-process** the `.gz` files:
  - Uncomment **lines 31 to 41** in `api/routes/api.go`.
  - ⚠️ Processing these files at startup can **significantly delay the API server startup**, which is why this code is commented out by default.

#### Alternative Approach

- You can implement and expose a **dedicated API endpoint** to perform the preprocessing.
- This endpoint can be triggered periodically by an **external scheduler** or **cron job** to reprocess coupon codes without impacting API startup time.


## 📁 Project Structure

```
backend-challenge/
├── api/
│   └── handlers
│       ├── order_handler.go
│       └── product_handler.go
│   └── middleware
│       ├── auth.go
│   └── routes
│       ├── api.go
├── config/
│       ├── config.go
├── internal
│   └── controllers/
│       ├── order_controller.go
│       └── product_controller.go
│   └── middleware/
│       └── auth.go
│   └── models/
│       ├── order.go
│       └── product.go
│   └── services/
│       ├── order_service.go
│       └── product_service.go
│       └── coupon_service.go
├── pkg/
│   └── errors
│       ├── api_errors.go
│   └── utils
│       ├──preprocessing.go
│       └── response.go
├── date/
│       ├── couponbase1.gz
│       ├── couponbase2.gz
│       ├── couponbase3.gz
├── go.mod
└── main.go
```

---

## 🚀 Getting Started

### Prerequisites

- Go 1.23+
- Git

### Installation

```bash
git clone https://github.com/akarsh17/backend-challenge.git
cd backend-challenge
go mod tidy
```

### Run the server

```bash
go run main.go
```

Server will start at: `http://localhost:8080`

---

## 🔐 Authentication

All `/order` endpoints require an API key:

```
Header: api_key: apitest
```

---

## 🧪 API Endpoints

### List Products

```http
GET /product
```

### Get Product by ID

```http
GET /product/{productId}
```

### Place Order

```http
POST /order
Headers:
  api_key: apitest
Body:
{
  "couponCode": "SAVE10",
  "items": [
    {
      "productId": "10",
      "quantity": 2
    }
  ]
}
```

---

## 🧱 Design Principles

- Follows Clean Architecture
- Layered (Controller → Service → Model)
- Error handling abstraction via `errors/api_errors.go`
- Optional response abstraction via `utils/response.go`
- Extensible middleware (e.g., API key auth)
- Preprocessing of couponbase files to generate a json with coupon occurancing >= 2 times across files

---

## 🛠️ To Do (Suggestions)

- Add persistent storage (e.g., PostgreSQL)
- Dockerize the application
