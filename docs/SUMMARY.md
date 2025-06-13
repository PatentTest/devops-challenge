# Project summary - devops-challenge

## 📦 Project Structure

- `app/` — Contains the Go application
- `Dockerfile` — Multi-stage build, final image is distroless
- `docker-compose.yml` — Local testing with support for AWS creds
- `.travis.yml` — CI pipeline: test, build, and deploy image to Docker Hub
- `verification.sh` — Sanity-check endpoints after deployment

---

## 🛠 Development Steps

### 1. App Design
- Lightweight Go HTTP API with `/secret` and `/health`
- Uses AWS SDK v2 to query DynamoDB
- Structurally minimal for challenge clarity

### 2. Testing
- Unit tests using `testify`
- `go test` support locally and via CI

### 3. Dockerization
- Alpine-based builder image
- Final container runs distroless static binary
- Docker Compose supports env + AWS credentials

### 4. CI/CD (Travis CI)
- Pulls Go modules and runs tests
- Builds Docker image on `main` branch
- Pushes to Docker Hub with secure credentials

---

## 🧪 Verification

```bash
docker-compose up -d
./verification.sh
```