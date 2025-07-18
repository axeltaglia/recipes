# Recipe Admin App

This is a simple Recipe Admin App that allows users to browse a list of recipes, view ingredients and steps and submit ratings.

## Project Structure

- `api/`: backend build in Golang 1.22
- `ui/`: frontend build in Vue 3

---

## How to Run the App

### Option 1: Using Docker Compose

Prerequisites: Make sure you have Docker and Docker Compose installed.

```bash
docker-compose up --build
```

- API will run at: [http://localhost:9123](http://localhost:9123)
- UI will run at: [http://localhost:9124](http://localhost:9124)

---

### Option 2: Running Manually

#### 1. Backend (Go + Fiber)

```bash
cd api
go mod tidy
go run main.go
```

API will be available at [http://localhost:9123](http://localhost:9123)

#### 2. Frontend (Vue 3)

```bash
cd ui
npm install
npm run dev
```

UI will be available at [http://localhost:5173](http://localhost:5173)

---

## Features

- List, search, and filter recipes
- View ingredients per portion
- View ordered preparation steps
- Submit ratings for recipes
- See top and worst rated recipes

---

## Notes

- No database is used. Data is stored in memory on the backend.
- Multiple ratings per recipe are allowed (no user system).
