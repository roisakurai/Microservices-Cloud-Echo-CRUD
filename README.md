# Cloud Echo CRUD Application

This project is a microservices-based CRUD application built with Go, Echo, MongoDB, Docker, Google Cloud Run, and Google Cloud Scheduler.

## Tech Stack

- Go
- Echo Framework
- MongoDB Atlas
- Docker
- Google Cloud Run
- Google Cloud Scheduler

## Microservices

This project consists of three independent services:

1. User Service
2. Product Service
3. Order Service

Each service has its own folder, Dockerfile, route, handler, usecase, repository, model, and configuration.

## Project Structure

```txt
.
├── user-service/
│   ├── config/
│   ├── handlers/
│   ├── models/
│   ├── repositories/
│   ├── routes/
│   ├── usecases/
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   ├── .env
│   └── main.go
│
├── product-service/
│   ├── config/
│   ├── handlers/
│   ├── models/
│   ├── repositories/
│   ├── routes/
│   ├── usecases/
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   ├── .env
│   └── main.go
│
└── order-service/
    ├── config/
    ├── handlers/
    ├── models/
    ├── repositories/
    ├── routes/
    ├── usecases/
    ├── Dockerfile
    ├── go.mod
    ├── go.sum
    ├── .env
    └── main.go

```

## Database

This project uses MongoDB Atlas.

Database name:

YOUR_DB_NAME

Collections:

users
products
orders
Environment Variables

Each service uses the following environment variables:

PORT=8080
MONGO_URI=mongodb+srv://paniroisakurai_db_user:Jb2yfCOP112zrmPN@cluster0.4rornyq.mongodb.net/?retryWrites=true&w=majority
DB_NAME=YOUR_DB_NAME

For local Docker testing, use:

MONGO_URI=mongodb://host.docker.internal:27017

For Cloud Run, use MongoDB Atlas URI.

## Deployment URLs
User Service
Base URL:
https://user-service-976985026390.asia-southeast2.run.app

Product Service
Base URL:
https://product-service-976985026390.asia-southeast2.run.app

Order Service
Base URL:
https://order-service-976985026390.asia-southeast2.run.app

## API Documentation
## User Service
Create User

Endpoint:

POST /users

Request body:

```txt
{
  "name": "Roi",
  "email": "roi@example.com"
}
```

Success response:

```txt
{
  "id": "665fabc1234567890abc1111",
  "name": "Roi",
  "email": "roi@example.com",
  "created_at": "21 May 2026 10:00:00"
}
```

Get User by ID

Endpoint:

GET /users/:id

Success response:

```txt
{
  "id": "665fabc1234567890abc1111",
  "name": "Roi",
  "email": "roi@example.com",
  "created_at": "21 May 2026 10:00:00"
}
```

Update User

Endpoint:

PUT /users/:id

Request body:

```txt
{
  "name": "Roi Updated",
  "email": "roi.updated@example.com"
}
```

Success response:

```txt
{
  "message": "user updated successfully"
}
```

Delete User

Endpoint:

DELETE /users/:id

Success response:

```txt
{
  "message": "user deleted successfully"
}
```

## Product Service
Create Product

Endpoint:

POST /products

Request body:

```txt
{
  "name": "Keyboard",
  "price": 250000,
  "stock": 10
}
```

Success response:

```txt
{
  "id": "665fdef1234567890abc2222",
  "name": "Keyboard",
  "price": 250000,
  "stock": 10,
  "created_at": "21 May 2026 10:05:00"
}
```

Get Product by ID

Endpoint:

GET /products/:id

Success response:

```txt
{
  "id": "665fdef1234567890abc2222",
  "name": "Keyboard",
  "price": 250000,
  "stock": 10,
  "created_at": "21 May 2026 10:05:00"
}
```

Update Product

Endpoint:

PUT /products/:id

Request body:

```txt
{
  "name": "Keyboard Mechanical",
  "price": 350000,
  "stock": 15
}
```

Success response:

```txt
{
  "message": "product updated successfully"
}
```

Delete Product

Endpoint:

DELETE /products/:id

Success response:

```txt
{
  "message": "product deleted successfully"
}
```

## Order Service
Create Order

Endpoint:

POST /orders

Request body:

```txt
{
  "user_id": "665fabc1234567890abc1111",
  "product_id": "665fdef1234567890abc2222",
  "quantity": 2
}
```

Notes:

The total price is calculated automatically by the server using product price * quantity.
The product stock is reduced automatically when an order is created.

Success response:

```txt
{
  "id": "665faaa1234567890abc3333",
  "user_id": "665fabc1234567890abc1111",
  "product_id": "665fdef1234567890abc2222",
  "quantity": 2,
  "total": 500000,
  "status": "pending",
  "created_at": "21 May 2026 10:10:00"
}
```

Get Order by ID

Endpoint:

GET /orders/:id

Success response:

```txt
{
  "id": "665faaa1234567890abc3333",
  "user_id": "665fabc1234567890abc1111",
  "username": "Roi",
  "product_id": "665fdef1234567890abc2222",
  "product_name": "Keyboard",
  "product_price": 250000,
  "quantity": 2,
  "total": 500000,
  "status": "pending",
  "created_at": "21 May 2026 10:10:00"
}
```

Update Order

Endpoint:

PUT /orders/:id

Request body:

```txt
{
  "user_id": "665fabc1234567890abc1111",
  "product_id": "665fdef1234567890abc2222",
  "quantity": 3
}
```

Notes:

If quantity is increased, product stock is reduced.
If quantity is decreased, product stock is returned.
The total price is recalculated automatically.

Success response:

```txt
{
  "message": "order updated successfully"
}
```

Delete Order

Endpoint:

DELETE /orders/:id

Notes:

When an order is deleted, the product stock is returned based on the order quantity.

Success response:

```txt
{
  "message": "order deleted successfully"
}
```

Update Order Status Cron Endpoint

Endpoint:

GET /orders/update-status

Description:

This endpoint is used by Google Cloud Scheduler to update pending orders to completed.

Success response:
```txt
{
  "message": "order status updated successfully",
  "matched_count": 2,
  "modified_count": 2
}
```

## Docker
Build Docker Images
docker build -t user-service ./user-service
docker build -t product-service ./product-service
docker build -t order-service ./order-service
Run User Service
```txt
docker run --rm \
  -p 8081:8080 \
  -e PORT=8080 \
  -e MONGO_URI="mongodb+srv://paniroisakurai_db_user:Jb2yfCOP112zrmPN@cluster0.4rornyq.mongodb.net/?retryWrites=true&w=majority \
  -e DB_NAME=YOUR_DB_NAME \
  user-service
```
Run Product Service
```txt
docker run --rm \
  -p 8082:8080 \
  -e PORT=8080 \
  -e MONGO_URI="mongodb+srv://paniroisakurai_db_user:Jb2yfCOP112zrmPN@cluster0.4rornyq.mongodb.net/?retryWrites=true&w=majority \
  -e DB_NAME=YOUR_DB_NAME \
  product-service
```
Run Order Service
```txt
docker run --rm \
  -p 8083:8080 \
  -e PORT=8080 \
  -e MONGO_URI="mongodb+srv://paniroisakurai_db_user:Jb2yfCOP112zrmPN@cluster0.4rornyq.mongodb.net/?retryWrites=true&w=majority \
  -e DB_NAME=YOUR_DB_NAME \
  order-service
  ```

## Cloud Run Deployment

Each microservice is deployed independently to Google Cloud Run.

User Service
```txt
gcloud run deploy user-service \
  --image asia-southeast2-docker.pkg.dev/PROJECT_ID/cloud-echo-repo/user-service:v1 \
  --region asia-southeast2 \
  --allow-unauthenticated \
  --set-env-vars MONGO_URI="MONGO_URI_ATLAS",DB_NAME="YOUR_DB_NAME"
  ```
Product Service
```txt
gcloud run deploy product-service \
  --image asia-southeast2-docker.pkg.dev/PROJECT_ID/cloud-echo-repo/product-service:v1 \
  --region asia-southeast2 \
  --allow-unauthenticated \
  --set-env-vars MONGO_URI="MONGO_URI_ATLAS",DB_NAME="YOUR_DB_NAME"
  ```
Order Service
```txt
gcloud run deploy order-service \
  --image asia-southeast2-docker.pkg.dev/PROJECT_ID/cloud-echo-repo/order-service:v1 \
  --region asia-southeast2 \
  --allow-unauthenticated \
  --set-env-vars MONGO_URI="MONGO_URI_ATLAS",DB_NAME="YOUR_DB_NAME"
  ```

## Cloud Scheduler

Cloud Scheduler is used to call the Order Service cron endpoint daily.

Job name:

update-order-status-job

Target endpoint:

GET https://order-service-xxxxx.a.run.app/orders/update-status

Schedule:

0 0 * * *

Timezone:

Asia/Jakarta

Command:
```txt
gcloud scheduler jobs create http update-order-status-job \
  --location=asia-southeast2 \
  --schedule="0 0 * * *" \
  --time-zone="Asia/Jakarta" \
  --uri="https://order-service-xxxxx.a.run.app/orders/update-status" \
  --http-method=GET
  ```