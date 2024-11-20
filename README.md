# Product Management System

The Product Management System is a backend service built with Go, PostgreSQL, and Redis for managing product data efficiently. It supports CRUD operations, caching, and image processing simulation.

## Features

- **Create Product**: Add a new product with details.
- **Retrieve Product**: Get product details using its ID.
- **Query Products**: Fetch all products with filters (e.g., user ID, price range).
- **Caching**: Redis is used for performance optimization.
- **Image Processing**: Simulates tasks like resizing images.

## System Architecture

Refer to the System Architecture section above for a detailed explanation.

---

## Setup Instructions

### 1. Prerequisites

Ensure the following are installed:
- **Go (1.18+)**: [Install Go](https://golang.org/dl/)
- **PostgreSQL**: [Install PostgreSQL](https://www.postgresql.org/download/)
- **Redis**: [Install Redis](https://redis.io/download/)
- **Docker (optional)**: Use Docker for PostgreSQL and Redis.

### 2. Clone the Repository

```bash
git clone https://github.com/yourusername/product-management.git
cd product-management
