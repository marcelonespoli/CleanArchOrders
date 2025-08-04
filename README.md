# CleanArchOrders
Projeto implementado em GO para registrar e listar orders utilizando WebServer, gRPC e graphQL  

## 🚀 Getting Started

### 🐳 Run with Docker

```bash
docker compose up --build
````

This will:

- ✅ Build the Go app  
- ✅ Start PostgreSQL  
- ✅ Apply migrations (if configured)  
- ✅ Expose the following services:

| Service     | Port  |
|-------------|-------|
| REST API    | 8000  |
| gRPC        | 50051 |
| GraphQL     | 8080  |
