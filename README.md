# Receipt Processor

This project is a receipt processing service that fulfills the documented API requirements. It processes receipts and calculates points based on specific rules.

## Table of Contents
1. [Project Setup](#project-setup)
2. [Running the Service](#running-the-service)
3. [Endpoints](#endpoints)
4. [Examples](#examples)
5. [Testing the Service](#testing-the-service)
6. [Dependencies](#dependencies)

---

## Project Setup

### Prerequisites

Ensure that you have the following installed:
- Go (version 1.16 or later)

### Installation

1. Clone the repository:
    ```bash
    git clone <repository-url>
    cd receipt-processor
    ```

2. Install the required Go modules:
    ```bash
    go get github.com/google/uuid
    ```

3. Initialize the Go module (if not already done):
    ```bash
    go mod init receipt-processor
    ```

4. Build the project:
    ```bash
    go build
    ```

---

## Running the Service

1. Run the service locally:
    ```bash
    go run main.go
    ```

2. The service will start and listen on port `8080`. You should see logs indicating the server has started:
    ```
    Starting server on :8080
    ```

---

## Endpoints

The service exposes two main endpoints:

### 1. Process Receipt (POST `/receipts/process`)

- **Description**: Submits a receipt for processing and returns a unique ID.
- **Endpoint**: `/receipts/process`
- **Method**: `POST`
- **Payload**: Receipt JSON
- **Response**: JSON containing an ID for the receipt.

#### Example Request(Should be from the path of the simple-receipt.json or the respective example file):
```bash
curl -X POST http://localhost:8080/receipts/process -d @simple-receipt.json --header "Content-Type: application/json"
```

#### Example Response:
```json
{
  "id": "7fb1377b-b223-49d9-a31a-5a02701dd310"
}
```

### 2. Get Points (GET `/receipts/{id}/points`)

- **Description**: Returns the points awarded for the given receipt ID.
- **Endpoint**: `/receipts/{id}/points`
- **Method**: `GET`
- **Response**: JSON object containing the number of points.

#### Example Request:
```bash
curl http://localhost:8080/receipts/7fb1377b-b223-49d9-a31a-5a02701dd310/points
```

#### Example Response:
```json
{
  "points": 28
}
```

---

## Examples

### Example 1: Submitting a Receipt

```json
{
    "retailer": "Target",
    "purchaseDate": "2022-01-02",
    "purchaseTime": "13:13",
    "total": "1.25",
    "items": [
        {"shortDescription": "Pepsi - 12-oz", "price": "1.25"}
    ]
}
```

#### Example Response:
```json
{
  "id": "e9fefab9-84a9-4dfa-a754-641b190a9e3c"
}
```

### Example 2: Getting Points for a Receipt

#### Example Request:
```bash
curl http://localhost:8080/receipts/e9fefab9-84a9-4dfa-a754-641b190a9e3c/points
```

#### Example Response:
```json
{
  "points": 31
}
```

---

## Testing the Service

You can test the service with sample JSON receipts. Two example files are included in the project directory:
- `simple-receipt.json`
- `morning-receipt.json`

To process a receipt, run:
```bash
curl -X POST http://localhost:8080/receipts/process -d @simple-receipt.json --header "Content-Type: application/json"
```

To retrieve points for the processed receipt, use the returned ID and run:
```bash
curl http://localhost:8080/receipts/{id}/points
```

---

## Dependencies

The project relies on the following Go modules:

- `github.com/google/uuid`: Used to generate unique receipt IDs.
  ```bash
  go get github.com/google/uuid
  ```

Make sure to install the dependencies using the `go get` command before running the service.

---

That's it! You should now have the Receipt Processor service up and running.
