![image](https://github.com/user-attachments/assets/a0b58225-a29c-46e2-9136-ab19bff24f43)![image](https://github.com/user-attachments/assets/83716f9f-f15b-46d2-8dbf-978e35a4f020)Product Management System
The Product Management System is a backend service designed for managing products, including features like creating products, retrieving product details, and querying all products. This system integrates with PostgreSQL for data storage, Redis for caching, and simulates image processing workflows.

Features
Create Product: Allows users to create new products with details like name, description, price, and images.
Retrieve Product: Fetch detailed information about a product using its ID.
Query Products: Retrieve a list of products with optional filtering options (e.g., user ID, price range).
Caching with Redis: Reduces database load by caching frequently accessed data.
Image Processing: Simulates image processing (e.g., compression or uploading to a storage service).

Architecture:
![image](https://github.com/user-attachments/assets/adfee4f5-470e-4db9-88b3-ea80bd1d1089)







Setup Instructions
1. Prerequisites
Ensure the following are installed:

Go (1.18+): Install Go
PostgreSQL: Install PostgreSQL
Redis: Install Redis
Docker (optional): Use Docker for PostgreSQL and Redis.
2. Clone the Repository
bash
Copy code
git clone https://github.com/yourusername/product-management.git
cd product-management
3. Configure Environment Variables
Copy the .env.example to .env:
bash
Copy code
cp .env.example .env
Update .env with your configuration:
plaintext
Copy code
DB_HOST=localhost
DB_PORT=5432
DB_NAME=product_management
DB_USER=pm_user
DB_PASSWORD=yourpassword
REDIS_HOST=localhost
REDIS_PORT=6379
4. Setup Database
Access PostgreSQL and create the database:
sql
Copy code
CREATE DATABASE product_management;
CREATE USER pm_user WITH PASSWORD 'yourpassword';
GRANT ALL PRIVILEGES ON DATABASE product_management TO pm_user;
Apply migrations:
bash
Copy code
psql -U pm_user -d product_management -f db/migrations/001_create_users_table.sql
psql -U pm_user -d product_management -f db/migrations/002_create_products_table.sql
5. Install Dependencies
Run the following:

bash
Copy code
go mod tidy
6. Run the Application
Start the application:

bash
Copy code
go run main.go
API available at: http://localhost:8080

API Endpoints
1. Create Product
POST /products
Request:
json
Copy code
{
  "user_id": 1,
  "product_name": "Sample Product",
  "product_description": "This is a sample product.",
  "product_images": ["image1.jpg"],
  "product_price": 49.99
}
Response:
json
Copy code
{
  "id": 1,
  "user_id": 1,
  "product_name": "Sample Product",
  "product_description": "This is a sample product.",
  "product_images": ["image1.jpg"],
  "product_price": 49.99
}
2. Get Product
GET /products/{id}
Response:
json
Copy code
{
  "id": 1,
  "user_id": 1,
  "product_name": "Sample Product",
  "product_description": "This is a sample product.",
  "product_images": ["image1.jpg"],
  "product_price": 49.99
}
3. Query Products
GET /products
Filters:
user_id: Filter by user ID.
price_min: Minimum price.
price_max: Maximum price.
Testing
Run all tests:
bash
Copy code
go test ./...
Generate coverage:
bash
Copy code
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out

1. Redis Connection Issues
Error: redis: dial tcp: lookup localhost: no such host
Cause:
Redis server is not running.
Incorrect Redis host or port configuration in .env.
Solution:
Verify Redis is running:
bash
Copy code
redis-cli ping
If you see PONG, Redis is running.
Check .env file for REDIS_HOST and REDIS_PORT. Default values are:
plaintext
Copy code
REDIS_HOST=localhost
REDIS_PORT=6379
Restart the Redis server if needed:
bash
Copy code
sudo systemctl restart redis
If using Docker:
bash
Copy code
docker start <redis-container-id>
2. PostgreSQL Connection Issues
Error: pq: connection to server failed
Cause:
PostgreSQL is not running.
Incorrect database credentials in .env.
Solution:
Verify PostgreSQL is running:
bash
Copy code
sudo systemctl status postgresql
Test database credentials manually:
bash
Copy code
psql -U pm_user -d product_management -h localhost
Check .env configuration:
plaintext
Copy code
DB_HOST=localhost
DB_PORT=5432
DB_NAME=product_management
DB_USER=pm_user
DB_PASSWORD=yourpassword
Restart PostgreSQL if necessary:
bash
Copy code
sudo systemctl restart postgresql
If using Docker for PostgreSQL:
bash
Copy code
docker start <postgres-container-id>
3. Port Already in Use
Error: listen tcp :8080: bind: address already in use
Cause: Another application is running on the same port.
Solution:
Find the process using the port:
bash
Copy code
sudo lsof -i :8080
Kill the process:
bash
Copy code
kill <PID>
Alternatively, change the port in main.go:
go
Copy code
log.Fatal(http.ListenAndServe(":8081", router))
4. Missing or Incorrect Environment Variables
Error: panic: missing environment variable DB_HOST
Cause: .env file is missing or variables are not correctly defined.
Solution:
Ensure the .env file exists in the root directory.
Verify all required variables are set:
plaintext
Copy code
DB_HOST=localhost
DB_PORT=5432
DB_NAME=product_management
DB_USER=pm_user
DB_PASSWORD=yourpassword
REDIS_HOST=localhost
REDIS_PORT=6379
Load environment variables:
bash
Copy code
source .env
5. API Endpoint Not Found
Error: 404 page not found
Cause: Incorrect API route or HTTP method.
Solution:
Check the API routes in api/routes.go.
Ensure the correct HTTP method is used (GET, POST, etc.).
Test the endpoint using curl or Postman:
bash
Copy code
curl -X GET http://localhost:8080/products
