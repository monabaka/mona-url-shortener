# URL Shortener

A modern (and one of the many already existing) URL shortener service built with Go (Fiber) backend and planned Next.js frontend. This project allows users to create shortened versions of long URLs and track click statistics.

## Features

Current:
- URL shortening with unique 8-character IDs
- URL redirection
- Click tracking for shortened URLs
- RESTful API
- PostgreSQL database for persistent storage

Planned:
- Modern Next.js frontend with Tailwind CSS
- User-friendly URL submission form
- Click analytics visualization
- and more...?

## Tech Stack

Backend:
- Go 1.24
- Fiber (web framework)
- GORM (ORM)
- PostgreSQL
- Google UUID

Planned Frontend:
- Next.js
- Tailwind CSS

## API Endpoints

### Shorten URL
```http
POST /shorten
Content-Type: application/json

{
    "url": "https://example.com/very-long-url"
}
```

### Access Shortened URL
```http
GET /:shortID
```

## Environment Variables

Create a `.env` file in the backend directory with the following variables:
```env
DB_URL=postgresql://user:password@localhost:5432/dbname
DOMAIN=http://localhost
PORT=8080
```

## Getting Started

### Backend Setup

1. Clone the repository
```bash
git clone https://github.com/monabaka/mona-url-shortener.git
cd url-shortener
```

2. Install dependencies
```bash
cd backend
go mod download
```

3. Set up your PostgreSQL database and configure the environment variables

4. Run the server
```bash
go run .
```

### Frontend Setup (Coming Soon)

The frontend implementation using Next.js and Tailwind CSS is currently under development.

## Contributing

This is a personal project that probably will be seen by no one, but suggestions and feedback are welcome! Feel free to open issues or submit pull requests.

## Author

Oleg 'starman' A

---
*Note: This project is currently under development. Frontend features will be added soon.*
