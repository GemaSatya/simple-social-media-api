# Simple Social Media API

A RESTful API for a simple social media platform built with Go. This is my first attempt at creating a social media API, implementing core features like user authentication, post management, and social interactions.

## 🚀 Features

- **User Authentication**: Secure JWT-based authentication system
- **User Management**: Create, read, update user profiles
- **Post Management**: Create, read, update, and delete posts
- **Social Interactions**: Like, comment, and share functionality
- **Friend System**: Follow/unfollow users and manage connections
- **RESTful Architecture**: Clean and intuitive API endpoints
- **Logging**: Built-in logging system for monitoring and debugging

## 📋 Prerequisites

Before running this application, make sure you have the following installed:

- [Go](https://golang.org/dl/) (version 1.16 or higher)
- [Git](https://git-scm.com/downloads)
- A database system (MongoDB/PostgreSQL/MySQL - check `go.mod` for specifics)

## 🛠️ Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/GemaSatya/simple-social-media-api.git
   cd simple-social-media-api
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   
   Create a `.env` file in the `env` directory based on your configuration needs:
   ```bash
   cp env/.env.example env/.env
   ```
   
   Configure the following variables:
   - Database connection strings
   - JWT secret key
   - Server port
   - Other application-specific settings

4. **Run the application**
   ```bash
   go run main.go
   ```

   The server should start on the configured port (default: `http://localhost:8080`)

## 📁 Project Structure

```
simple-social-media-api/
├── auth/           # Authentication logic and JWT handling
├── env/            # Environment configuration files
├── models/         # Data models and database schemas
├── utils/          # Utility functions and helpers
├── main.go         # Application entry point
├── log.txt         # Application logs
├── go.mod          # Go module dependencies
├── go.sum          # Dependency checksums
└── .gitignore      # Git ignore rules
```

## 🔑 API Endpoints

### Authentication
- `POST /api/auth/register` - Register a new user
- `POST /api/auth/login` - Login and receive JWT token
- `POST /api/auth/logout` - Logout user

### Users
- `GET /api/users` - Get all users
- `GET /api/users/:id` - Get user by ID
- `PUT /api/users/:id` - Update user profile
- `DELETE /api/users/:id` - Delete user account

### Posts
- `GET /api/posts` - Get all posts
- `GET /api/posts/:id` - Get post by ID
- `POST /api/posts` - Create a new post
- `PUT /api/posts/:id` - Update post
- `DELETE /api/posts/:id` - Delete post

### Social Features
- `POST /api/posts/:id/like` - Like a post
- `DELETE /api/posts/:id/like` - Unlike a post
- `POST /api/posts/:id/comments` - Add comment to post
- `POST /api/users/:id/follow` - Follow a user
- `DELETE /api/users/:id/follow` - Unfollow a user

*Note: Some endpoints require authentication. Include the JWT token in the Authorization header:*
```
Authorization: Bearer <your-jwt-token>
```

## 🧪 Testing

You can test the API endpoints using tools like:
- [Postman](https://www.postman.com/)
- [Insomnia](https://insomnia.rest/)
- [cURL](https://curl.se/)

### Example Request
```bash
# Register a new user
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "johndoe",
    "email": "john@example.com",
    "password": "securepassword"
  }'

# Login
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "securepassword"
  }'
```

## 🔧 Configuration

The application uses environment variables for configuration. Key settings include:

- `PORT` - Server port (default: 8080)
- `DATABASE_URL` - Database connection string
- `JWT_SECRET` - Secret key for JWT token generation
- `JWT_EXPIRATION` - Token expiration time
- `LOG_LEVEL` - Logging level (debug, info, warn, error)

## 📝 Logging

The application includes a logging system that writes to `log.txt`. Logs include:
- API requests and responses
- Database operations
- Authentication events
- Error tracking

## 👤 Author

**GemaSatya**
- GitHub: [@GemaSatya](https://github.com/GemaSatya)

## 🙏 Acknowledgments

- Thanks to the Go community for excellent documentation
- Inspired by various social media platforms
- Built as a learning project to understand RESTful API development

## 📚 Resources

- [Go Documentation](https://golang.org/doc/)
- [REST API Best Practices](https://restfulapi.net/)
- [JWT Authentication](https://jwt.io/introduction)

## 🐛 Known Issues

This is a learning project and may have some limitations:
- Limited error handling in some scenarios
- Basic security implementations
- No rate limiting implemented yet
- No caching layer

Feel free to open an issue if you find any bugs or have suggestions for improvements!

## 🗺️ Roadmap

Future enhancements planned:
- [ ] Add real-time notifications using WebSockets
- [ ] Implement file upload for profile pictures and post media
- [ ] Add rate limiting and request throttling
- [ ] Implement caching layer (Redis)
- [ ] Add comprehensive unit and integration tests
- [ ] Create API documentation with Swagger
- [ ] Add email verification system
- [ ] Implement password reset functionality
- [ ] Add search functionality for users and posts

---

⭐ If you find this project helpful, please consider giving it a star!
