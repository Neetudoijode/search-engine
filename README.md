
# Fullstack Search Engine

## Overview
A fullstack search engine built using Golang for the backend and React for the frontend. It reads data from Parquet files and allows keyword-based search.

## Backend
- Built in Go
- Reads Parquet files using `parquet-go`
- Stores data in memory
- Exposes `/search?q=keyword` endpoint

## Frontend
- Built in React
- Search UI with keyword input
- Shows results and metadata (count, duration)

## Setup

### Backend
```bash
cd backend
go run main.go
```

### Frontend
```bash
cd frontend
npm install
npm start
```

## Benchmarking
Logged time taken for each query

## Stretch Goals
- Upload Parquet files via UI (not implemented)

---