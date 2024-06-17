# Go LRU Cache with React Frontend

This project is a demonstration of an LRU (Least Recently Used) Cache implemented in Go with a React frontend. The backend provides APIs for setting, getting, and deleting cache entries, and the frontend interacts with these APIs. Additionally, WebSocket is used to dynamically update the frontend with the current cache state.

## Folder Structure

### Go Backend

/go-lru-cache
/cmd
/main
main.go
/internal
/cache
lru.go
/handlers
handlers.go
/pkg
/models
response.go
go.mod
go.sum

shell
Copy code

### React Frontend

/react-lru-cache
/public
/src
/components
CacheDisplay.js
CacheForm.js
/utils
api.js
App.js
index.js
package.json
.env

markdown
Copy code

## Setup Instructions

### Backend (Go)

1. **Install Go**: Ensure that Go is installed on your machine. You can download it from [golang.org](https://golang.org/dl/).

2. **Clone the Repository**: Clone this repository to your local machine.
    ```bash
    git clone https://github.com/yourusername/go-lru-cache.git
    cd go-lru-cache
    ```

3. **Initialize the Project**:
    ```bash
    go mod tidy
    ```

4. **Run the Backend Server**:
    ```bash
    go run cmd/main/main.go
    ```

The backend server will start on `http://localhost:8080`.

### Frontend (React)

1. **Install Node.js**: Ensure that Node.js and npm are installed on your machine. You can download them from [nodejs.org](https://nodejs.org/).

2. **Navigate to the Frontend Directory**:
    ```bash
    cd react-lru-cache
    ```

3. **Install Dependencies**:
    ```bash
    npm install
    ```

4. **Start the Frontend Server**:
    ```bash
    npm start
    ```

The frontend server will start on `http://localhost:3000`.

## API Endpoints

### Set Cache Entry

- **URL**: `/set`
- **Method**: `POST`
- **Request Body**:
    ```json
    {
        "key": "exampleKey",
        "value": "exampleValue",
        "expiration": 5
    }
    ```
- **Response**: `200 OK`

### Get Cache Entry

- **URL**: `/get?key=exampleKey`
- **Method**: `GET`
- **Response**:
    ```json
    {
        "key": "exampleKey",
        "value": "exampleValue"
    }
    ```

### Delete Cache Entry

- **URL**: `/delete?key=exampleKey`
- **Method**: `DELETE`
- **Response**: `200 OK`

### WebSocket

- **URL**: `/ws`
- **Description**: Connects to the WebSocket to get real-time updates of the cache state.

## Additional Information

### Concurrency

The LRU cache is implemented with thread-safety in mind, using read-write mutexes to handle concurrent access to the cache.

### Expiration

Cache entries support expiration times, after which they are automatically evicted from the cache.

### WebSocket Integration

The project includes WebSocket integration to reflect current key-value pairs dynamically in the React frontend.

### License

This project is licensed under the MIT License. See the `LICENSE` file for more details.

---
