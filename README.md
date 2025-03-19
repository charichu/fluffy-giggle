# fluffy-giggle
This is a simple, hardcoded demo program that implements a REST API using Go and the Gorilla mux router. The program is meant solely for demonstration purposes and does not interact with a real ERP system or database. All responses are hardcoded.

## Disclaimer
This is a hardcoded demo application. The responses are simulated and static, and the program is not intended for production use. No actual ERP or account validation is performed.

## Endpoints

### 1. GET `/product/{id}/stocklevel`
- **Description:**  
  Returns a hardcoded product JSON object with the following fields:
  - **produktname:** Name of the product (e.g., "Demo Produkt")
  - **produkt_id:** Product ID (extracted from the URL)
  - **preis:** Fixed price (e.g., 99.99)
  - **verfugbarkeit:** Availability (e.g., "Auf Lager")
  - **versanddauer:** Estimated shipping duration (e.g., "3-5 Tage")
  
- **Simulated Behavior:**  
  - Logs "Fetching stock level from ERP...." and waits for one second.
  - Logs "Estimate shipping time with ERP RPC...." before sending the response.

### 2. GET `/account/{id}/checkLimit?amount=<amount>`
- **Description:**  
  Checks the account limit for a given account and order amount. Although the HTTP status is 200, the JSON response contains an error message.
  
- **Response:**  
  A JSON object containing:
  - **error:** "Order exceeds account limit"
  - Optionally, the requested `amount` is also included.
  
- **Simulated Behavior:**  
  - Logs "Getting account information.."
  - Logs "Getting unpaid orders for account {id}" (with the provided account id)
  - Logs "Checking limit for new order..."

## Getting Started

### Prerequisites
- **Go:** Version 1.15 or later.
- **Gorilla mux:** Install via:
 ```
go get -u github.com/gorilla/mux
```

## Running the Server

1. **Clone the repository** or copy the source code to your local machine.
2. **Navigate to the project directory.**
3. **Start the server using:**
 ```
go run main.go
 ```

## Testing the Endpoints
cURL Commands
Linux / macOS
Open a terminal and use the following commands:

### Product Endpoint:
```
curl -X GET "http://localhost:8080/product/1/stocklevel"
```

### Account Endpoint:
```
curl -X GET "http://localhost:8080/account/1/checkLimit?amount=209,99"
```

## Windows
In a Windows terminal (e.g., PowerShell), the alias curl may refer to Invoke-WebRequest. To use the actual curl executable, use curl.exe:

### Product Endpoint:
```
curl.exe -X GET "http://localhost:8080/product/1/stocklevel"
```

### Account Endpoint:
```
curl.exe -X GET "http://localhost:8080/account/1/checkLimit?amount=209,99"
```

