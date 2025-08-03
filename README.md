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


GET
curl http://localhost:8080/book
curl http://localhost:8080/book/1


PUT
curl -X PUT http://localhost:8080/books/1 \
-H "Content-Type: application/json" \
-d '{
  "title": "Updated Title",
  "author": "Updated Author",
  "published_year": 2023,
  "price": 400
}'