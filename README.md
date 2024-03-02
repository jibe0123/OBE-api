# 📊 OBE API

## Description

🚀 Welcome to the OBE API! This project is a high-performance web service built with Go and the Gin framework. It's designed to handle CSV file uploads, process them according to predefined business logic, and perform various data collection and analysis operations. By adhering to Domain-Driven Design (DDD) principles, our codebase is neatly organized into distinct layers, enhancing both maintainability and scalability.

## Features

- 📁 **CSV File Uploads**: Accepts CSV files and processes them effectively.
- 🛠️ **DDD Architecture**: Our codebase is structured following the DDD principles to ensure clear separation of concerns.
- 📡 **RESTful Endpoints**: Offers simple and well-structured API endpoints for easy integration.
- 🐳 **Docker Support**: Facilitates easy containerization and deployment of the service.

## Getting Started

### Prerequisites

- Go version 1.15 or later.
- Docker (optional, for containerization).

### Installation

1. **Clone the repository:**

```bash
git clone https://github.com/yourusername/csv-processor-api.git
cd csv-processor-api
```

2. **Install dependencies:**

```bash
go mod tidy
```

### Running the Application

- **To run the server locally:**

```bash
go run cmd/main.go
```

- **To build and run with Docker:**

```bash
docker build -t csv-processor-api .
docker run -p 8080:8080 csv-processor-api
```

## Usage

To upload and process a CSV file, use the following `curl` command:

```bash
curl -X POST -F "file=@path/to/your/file.csv" http://localhost:8080/api/v1/upload
```

Replace `path/to/your/file.csv` with the actual file path you wish to upload.

## Contributing

🤝 Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the project.
2. Create your feature branch (`git checkout -b feature/AmazingFeature`).
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`).
4. Push to the branch (`git push origin feature/AmazingFeature`).
5. Open a pull request.

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

Your Name - [@your_twitter](https://twitter.com/your_twitter) - email@example.com

Project Link: [https://github.com/yourusername/csv-processor-api](https://github.com/yourusername/csv-processor-api)

## Acknowledgements

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [Go Programming Language](https://golang.org/)
```