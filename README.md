CRUD ke-10

POST
curl -X POST http://localhost:8080/books \
-H "Content-Type: application/json" \
-d '{
  "title": "Go Programming",
  "author": "John Doe",
  "published_year": 2022,
  "price": 350.50
}'