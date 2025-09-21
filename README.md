# URL Shortener

This is a simple URL shortener application built with Go, using the Gin framework. It provides a basic API for creating and redirecting short URLs.

## Features

- **Gin Framework:** Utilizes the high-performance Gin framework for handling HTTP requests.
- **PostgreSQL Database:** Uses a PostgreSQL database to store and retrieve URL data.
- **Structured Logging:** Implements custom logging for better error handling and application monitoring.
- **Environment Variables:** Configures the application using environment variables, ensuring sensitive information is kept out of the codebase.

## Prerequisites

- Go (1.20 or newer)
- PostgreSQL
- Git

## Getting Started

Follow these steps to get your local environment set up and running.

### 1\. Clone the repository

git clone &lt;repository_url&gt;  
cd url-shortener  

### 2\. Configure Environment Variables

Create a file named .env in the root directory of the project and add your database credentials and application settings. The APP_PORT variable will be used to run the application on a specific port.

\# Database Credentials  
DB_HOST=localhost  
DB_PORT=5432  
DB_USER=mirza.pradana  
DB_PASSWORD=password  
DB_NAME=url_shortener  
<br/>\# Application Settings  
APP_PORT=:8080  

### 3\. Install Dependencies

Run the following command to download and install the required Go modules:

go mod tidy  

### 4\. Run the Application

Execute the application from the root directory:

go run main.go  

The application will start on the port defined in your .env file (e.g., <http://localhost:8080>).

## Project Structure
```
├── main.go # Main application entry point  
├── config/ # Database and application configuration  
│ └── config.go  
├── internal/  
│ └── handler/ # HTTP request handlers  
│ └── handler.go  
├── repository/ # Database interaction logic  
│ └── repository.go  
├── pkg/  
│ ├── errors/ # Custom error handling package  
│ │ └── errors.go  
│ └── logger/ # Centralized logging package  
│ └── logger.go  
├── go.mod # Go module file  
├── go.sum # Go module checksums  
├── .env # Environment variables file  
└── README.md # Project documentation
```
