Golang CRUD API with Gin

A simple RESTful API built with Go and Gin framework for managing products. This API supports basic CRUD operations: Create, Read, Update, and Delete.

Features

Get all products (GET /products)

Get a single product by ID (GET /products/:id)

Add a new product (POST /products)

Update a product partially (PATCH /products/:id)

Delete a product (DELETE /products/:id)

Technologies Used

Golang

Gin Web Framework

Installation

Clone the repository:

git clone https://github.com/Theo1039/golang-crud-api.git
cd golang-crud-api


Initialize Go modules (if not done already):

go mod init golang-crud-api


Install dependencies:

go get github.com/gin-gonic/gin


Run the application:

go run main.go


The API will run at http://localhost:8080.

API Endpoints
Get All Products
GET /products


Response Example:

[
  {"id":1,"name":"Car","price":90000},
  {"id":2,"name":"Tyre","price":560},
  {"id":3,"name":"Tipper motor","price":56700}
]

Get Product by ID
GET /products/:id


Response Example:

{"id":1,"name":"Car","price":90000}

Add New Product
POST /products


Request Body Example:

{
  "name": "Wheel",
  "price": 120
}


Response Example:

{"message": "product posted successfully"}

Update Product (Partial)
PATCH /products/:id


Request Body Example:

{
  "name": "Updated Car",
  "price": 95000
}


Response Example:

{"id":1,"name":"Updated Car","price":95000}

Delete Product
DELETE /products/:id


Response Example:

{"message": "product deleted successfully"}

Testing API with curl

You can test all endpoints directly from your terminal using curl.

1. Get all products

curl -X GET http://localhost:8080/products


2. Get a product by ID

curl -X GET http://localhost:8080/products/1


3. Add a new product

curl -X POST http://localhost:8080/products \
-H "Content-Type: application/json" \
-d '{"name":"Wheel","price":120}'


4. Update a product partially

curl -X PATCH http://localhost:8080/products/1 \
-H "Content-Type: application/json" \
-d '{"price":95000}'


5. Delete a product

curl -X DELETE http://localhost:8080/products/1

Notes

PATCH allows partial updates: you can update only name or only price.

Product IDs are automatically generated.

API returns proper HTTP status codes for success and errors.

License

This project is open-source and available under the MIT License.