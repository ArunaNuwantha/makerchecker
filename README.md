# Maker Checker


## Installation

### Running Locally

- to run this application locally

   ```bash
   git clone https://github.com/arunanuwantha/makerchecker.git
   cd makerchecker
   go mod tidy
   go run main.go
   ```

- and also you can compile the code into binary file, for that use below command 

   ```bash
   go build
   ./makerchecker
   ```
   - then server will start on http://localhost:8080

- using docker

   ```bash
   docker build -t makerchecker:latest .
   docker run -d -p 8080:8080 --name makerchecker_app makerchecker:latest
   ```

## Usage
- to create the message, you have to call POST request to the /messages like below 

```bash
curl -X POST http://localhost:8080/messages \
-H "Content-Type: application/json" \
-d '{
  "recipient": "recipient@example.com",
  "content": "Hello, this is a test message."
}'
```

- to see the messages
```bash
curl http://localhost:8080/messages
```

- to approve message
```bash
curl -X POST http://localhost:8080/messages/{message_id}/approve
```

- to reject the message

```bash 
curl -X POST http://localhost:8080/messages/{id}/reject
```


