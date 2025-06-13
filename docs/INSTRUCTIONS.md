# Instructions - devops-challenge

---

## 🔁 Cloning the Repository

```bash
git clone https://github.com/PatentTest/devops-challenge.git
cd devops-challenge
```

Make sure you have:

* Docker + Docker Compose installed
* Go (1.24.x if using distroless build)

---

## 🐳 Running the App Locally

### Step 1: Prepare Environment Variables

Create a `.env` file at the project root:

```env
DYNAMODB_TABLE=devops-challenge
CONTAINER_URL=https://hub.docker.com/r/PatentTest/devops-challenge
PROJECT_URL=https://github.com/PatentTest/devops-challenge
AWS_REGION=us-east-1
AWS_ACCESS_KEY_ID=your-access-key     # optional if using IAM role
AWS_SECRET_ACCESS_KEY=your-secret-key # optional if using IAM role
```

> Alternatively, export them in your shell.

### Step 2: Start the Service

```bash
docker-compose up --build
```

The service should be available at:
`http://localhost:5000/health` and `http://localhost:5000/secret`

---

## 🧪 Testing

Run tests using Go (from project root):

```bash
go test -v ./app
```

> Uses [testify](https://github.com/stretchr/testify) for assertions.

---

## 🔍 Verify Locally

Use the provided verification script to check endpoints:

```bash
./verification.sh
```

This will test `/health` and `/secret` and report success/failure.

---

## 🔁 Continuous Integration

CI is managed via [Travis CI](https://travis-ci.com):

* Runs tests on each push
* Builds Docker image from `Dockerfile`
* Pushes image to Docker Hub when `main` is updated

---

## 🧼 Clean Up

To remove containers:

```bash
docker-compose down
```
