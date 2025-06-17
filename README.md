<h1 align="center">🧺 Laundry App REST API</h1>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.22-blue?style=flat-square" />
  <img src="https://img.shields.io/badge/Gin_Framework-lightblue?style=flat-square" />
  <img src="https://img.shields.io/badge/GORM-ORM-blueviolet?style=flat-square" />
</p>

<p align="center">
  A RESTful API for managing customers, employees, and orders in a laundry service application. Built using <strong>Go</strong>, <strong>Gin</strong>, and <strong>GORM</strong> with full JWT-based authentication and Swagger documentation.
</p>

---

## 📌 Features

- 🧍 Customer and employee account management
- 🔐 JWT-based login & role-based access
- 🗃️ Customer and employee search, update, view by ID or account ID
- 📦 Clean Swagger 2.0 API documentation
- 📁 Configurable via `.env` file

---

## 🛠️ Tech Stack

| Tool   | Role                       |
| ------ | -------------------------- |
| Go     | Backend language           |
| Gin    | Web framework              |
| GORM   | ORM for PostgreSQL         |
| JWT    | Token-based authentication |
| Swaggo | Swagger generation         |

---

## 🚀 Getting Started

```bash
# Clone the repo
git clone https://github.com/maulanadityaa/laundry-app-rest-api.git
cd laundry-app-rest-api

# Set up env vars
cp .env.example .env

# Install dependencies
go mod tidy

# Generate Swagger docs
go install github.com/swaggo/swag/cmd/swag@latest
swag init

# Run the server
go run main.go
```

---

## 🔐 Authentication

Most routes require a JWT token.  
Header format:

```http
Authorization: Bearer <your_jwt_token>
```

---

## 📦 Key API Endpoints

### 🔑 Auth

| Method | Endpoint                | Description                            |
| ------ | ----------------------- | -------------------------------------- |
| POST   | `/api/v1/auth/login`    | Login and get JWT token                |
| POST   | `/api/v1/auth/register` | Register a user (customer or employee) |

### 👤 Customers

| Method | Endpoint                                | Description                          |
| ------ | --------------------------------------- | ------------------------------------ |
| GET    | `/api/v1/customers`                     | List all customers (with pagination) |
| GET    | `/api/v1/customers/{id}`                | Get customer by ID                   |
| GET    | `/api/v1/customers/account/{accountID}` | Get customer by account ID           |
| PUT    | `/api/v1/customers`                     | Update customer info                 |

### 🧑‍💼 Employees

| Method | Endpoint                                | Description                          |
| ------ | --------------------------------------- | ------------------------------------ |
| GET    | `/api/v1/employees`                     | List all employees (with pagination) |
| GET    | `/api/v1/employees/{id}`                | Get employee by ID                   |
| GET    | `/api/v1/employees/account/{accountID}` | Get employee by account ID           |
| PUT    | `/api/v1/employees`                     | Update employee info                 |

---

## 🧪 Example Requests

### Register

```http
POST /api/v1/auth/register
Content-Type: application/json

{
  "name": "John Doe",
  "email": "johndoe@example.com",
  "password": "12345678",
  "phoneNumber": "081234567890",
  "address": "Jl. Address No. 1",
  "role": "ROLE_CUSTOMER"
}
```

**Response**

```json
{
  "address": "Jl. Address No. 1",
  "email": "johndoe@example.com",
  "name": "John Doe",
  "password": "12345678",
  "phoneNumber": "081234567890",
  "role": "ROLE_CUSTOMER"
}
```

---

### Login

```http
POST /api/v1/auth/login
Content-Type: application/json

{
  "email": "johndoe@example.com",
  "password": "12345678"
}
```

**Response:**

```json
{
  "token": "ValidJWTToken"
}
```

---

### Get Customers (with Pagination)

```http
GET /api/v1/customers?page=1&rowsPerPage=10&name=John
Authorization: Bearer <jwt_token>
```

**Response:**

```json
{
  "data": [
    {
      "address": "Jl. Address No. 1",
      "createdAt": "2021-01-01T00:00:00Z",
      "id": "ValidUUIDv4",
      "name": "John Doe",
      "phoneNumber": "081234567890",
      "updatedAt": "2021-01-01T00:00:00Z"
    }
  ],
  "paging": {
    "currentPage": 1,
    "rowsPerPage": 10,
    "totalRows": 1
  }
}
```

---

## 📘 Swagger Docs

Swagger UI is available at:
📄 [`/swagger/docs`](https://laundry-app-rest-api.vercel.app/swagger/docs/index.html)

To generate docs:

```bash
swag init
```

---

## 🤝 Contributing

```bash
git checkout -b feature/your-feature
git commit -m "Add feature"
git push origin feature/your-feature
```

Then open a PR 🚀

---

## 📧 Contact

**Author:** [maulanadityaa](https://github.com/maulanadityaa)  
**Email:** maulanadityaaa@gmail.com
