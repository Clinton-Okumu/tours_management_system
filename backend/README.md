# Tours Management System

This is a comprehensive Tours Management System designed to streamline the process of managing tours, bookings, reviews, and user interactions. The system is built with a robust backend and is designed to be scalable and maintainable.

## Features

### Backend
*   **User Management**: Register, login, and manage user accounts with different roles (e.g., user, admin).
*   **Tour Management**: Create, read, update, and delete tour information.
*   **Booking Management**: Handle tour bookings.
*   **Review System**: Allow users to post and manage reviews for tours.
*   **Location Management**: Manage tour locations.
*   **Authentication**: Secure user authentication using tokens.

## Technologies Used

### Backend
*   **Go**: Primary programming language.
*   **GORM**: ORM for database interactions.
*   **PostgreSQL**: Relational database.
*   **Chi**: HTTP router and URL matcher.
*   **Bcrypt**: For secure password hashing.
*   **JWT (JSON Web Tokens)**: For authentication (to be implemented/integrated).

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

*   Go (version 1.20 or higher)
*   PostgreSQL
*   Docker (optional, for running PostgreSQL in a container)

### Installation

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/your-username/tours_management_system.git
    cd tours_management_system/backend
    ```

2.  **Set up environment variables:**
    Create a `.env` file in the `backend` directory with your database connection details.
    ```
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=your_user
    DB_PASSWORD=your_password
    DB_NAME=your_database_name
    DB_SSLMODE=disable
    ```

3.  **Run database migrations:**
    The application will automatically run migrations on startup. Ensure your database is running and accessible.

4.  **Run the backend application:**
    ```bash
    go run main.go
    ```
    The backend server will start on the port specified in your configuration (default usually 8080).

## API Endpoints

The backend exposes RESTful API endpoints for managing users, tours, bookings, reviews, and locations. Detailed API documentation will be provided separately (e.g., Postman collection, Swagger/OpenAPI spec).

**Example Endpoints:**
*   `POST /api/v1/users/register` - Register a new user
*   `POST /api/v1/users/login` - Login a user
*   `GET /api/v1/tours` - Get all tours

## Frontend

This section describes the frontend application.

*   **Next.js**: React framework for building user interfaces.
*   **Tailwind CSS**: A utility-first CSS framework for rapid UI development.

