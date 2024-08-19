# URL Shortener

### A simple URl Shortener service built with Go, MySQL and Redis. This project provides a minimal system for shortening long url into code and redirecting users to the original URLs.

# Features

- Generate short, unique codes for URLs
- Redirect to origin url using code

# Architecture
### The URL Shortener composed of following components:
- API Server: Handles HTTP requests for shortening URLs and redirects.
- Redis Cache: Caches the code as the key and the original URL as the value.
- MySQL Database: Stores the original URL and its associated code.

# System Flow
1. Shortening a URL:
   - A user submit long url to the API.
   - The service check long url exists in database or not.
   - If it exists, a service return short url to the user.
   - If not, the service generates a unique code for the long URL and returns it to the user.
2. Redirecting:
  - A user accesses a short url.
  - The API checks if the short URL exists in Redis.
  - if found, the service redirects the user to the original URL.
  - if not, the service checks in the database.

# Technologies Used
 - Go
 - MySQL
 - Redis
 - Docker
