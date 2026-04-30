# Retail Pulse Image Processing Service

A high-performance Go-based microservice designed for asynchronous image processing of retail store visits. This service validates store data against a master record and performs image analysis (perimeter calculation) with simulated processing delays.

## 🚀 Features

- **Bulk Job Submission**: Process multiple store visits and their associated images in a single request.
- **Asynchronous Processing**: Jobs run in the background using Go routines, allowing for immediate response with a `job_id`.
- **Store Validation**: Validates `store_id` against a master CSV database (`StoreMasterAssignment.csv`).
- **Image Analysis**: Downloads images, decodes JPEG data, and calculates image dimensions/perimeter.
- **Job Status Tracking**: Real-time status updates (`pending`, `ongoing`, `completed`, `failed`) with detailed error reporting.
- **Dockerized**: Easy deployment using Docker.

## 🛠️ Tech Stack

- **Language**: Go (Golang) 1.23+
- **Web Framework**: [Gorilla Mux](https://github.com/gorilla/mux)
- **Containerization**: Docker
- **Data Format**: JSON (API), CSV (Master Data)

## 📡 API Reference

### 1. Submit Image Processing Job
Submits a new job for processing.

- **URL**: `/api/submit/`
- **Method**: `POST`
- **Payload**:
  ```json
  {
    "count": 1,
    "visits": [
      {
        "store_id": "RP001",
        "image_url": ["https://example.com/image1.jpg"],
        "visit_time": "2023-10-27T10:00:00Z"
      }
    ]
  }
  ```
- **Response** (201 Created):
  ```json
  {
    "job_id": 1
  }
  ```

### 2. Get Job Status
Retrieves the current status of a submitted job.

- **URL**: `/api/status`
- **Method**: `GET`
- **Query Parameters**: `jobid` (integer)
- **Response** (Success):
  ```json
  {
    "status": "completed",
    "job_id": 1
  }
  ```
- **Response** (Failure):
  ```json
  {
    "status": "failed",
    "job_id": 1,
    "error": [
      {
        "store_id": "RP999",
        "error": "Store does not exist:"
      }
    ]
  }
  ```

## 🏃 Running the Project

### Using Go Locally
1. Ensure Go is installed.
2. Run the server:
   ```bash
   go run main.go
   ```
3. The server will start on `http://localhost:8080`.

### Using Docker
1. Build the image:
   ```bash
   docker build -t image-process-service .
   ```
2. Run the container:
   ```bash
   docker run -p 8080:8080 image-process-service
   ```

## 📂 Project Structure

- `main.go`: Application entry point and router configuration.
- `handlers/`: Contains the logic for HTTP requests and image processing.
  - `job_handler.go`: Handles job submission and status retrieval.
  - `image_process.go`: Logic for downloading and analyzing images.
- `models/`: Data structures and in-memory storage.
  - `job.go`: Job and Visit definitions.
  - `store.go`: Store validation logic and CSV parsing.
- `StoreMasterAssignment.csv`: Master dataset for store validation.

---
*Created as part of a Retail Pulse assessment.*
