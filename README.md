# gonty
Some interesting RESTful API in Go / Gin Framework

## Run

### Debug Mode

    go run main.go

### Production Mode

    GIN_MODE=release go run main.go


## Testing

    CGO_ENABLED=0 GIN_MODE=release go test


## Usage

### Health Check

    curl http://localhost:8080/health

### A Famous German Baroque Composer

    curl http://localhost:8080/composer/:genre

try `/composer/baroque` to find out one of the greatest to ever exist!

try something else too e.g. `/composer/classical`

### Funniest Jokes in the World (proceed with Caution!)

1. Get all Jokes

        curl http://localhost:8080/jokes

2. Get the Funniest (Most possibly Dangerous too!)

        curl http://localhost:8080/jokes?funniest=true

### Check Your Expectation Levels

    curl http://localhost:8080/expect

And be Amazed!!

### Find out about a Nobleman in History

    curl http://localhost:8080/history/nobleman

## Pending Tasks

- [ ] Use a GitHub Actions Workflow to Containerize it
- [ ] Deploy it locally using Kubernetes