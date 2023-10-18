# go-crud
Aplikasi golang crud with middleware. buat latihan dasar

@host = localhost:8080

// Create Product
POST http://{{host}}/api/products HTTP/1.1
content-type: application/json
    
{
    "name": "test-product",
    "description": "random-description",
    "price": 100.00
}

###

// Get Product By ID
GET http://{{host}}/api/products/23 HTTP/1.1
content-type: application/json

###

// Get All Products
GET http://{{host}}/api/products/ HTTP/1.1
content-type: application/json

###

// Update Product
PUT http://{{host}}/api/products/23 HTTP/1.1
content-type: application/json

{
    "name": "updated-product",
    "description": "random-description-updated",
    "price": 100.00
}

###

// Delete Product
DELETE http://{{host}}/api/products/23 HTTP/1.1
content-type: application/json

###
// Register User
POST http://{{host}}/api/user/register HTTP/1.1
content-type: application/json
    
{
    "name": "Mukesh Murugan",
    "username": "mukesh.murugan",
    "email": "mukesh@go.com",
    "password": "123465789"
}
###

// Generate JWT
POST http://{{host}}/api/token HTTP/1.1
content-type: application/json
    
{
    "email": "mukesh@go.com",
    "password": "123465789"
}
###


// Access a Secured API Endpoint
GET http://{{host}}/api/secured/ping HTTP/1.1
content-type: application/json
authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im11a2VzaC5tdXJ1Z2FuIiwiZW1haWwiOiJtdWtlc2hAZ28uY29tIiwiZXhwIjoxNjUwNzQzMjA1fQ.7cAcWxvpqJ1DDZ-ZOM2kIKKedeCEWuUzl0Hj2VuMxYA
###
    
