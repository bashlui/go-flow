# Go Flow

> **⚠️ Development Status**: This project is currently under active development. Features and APIs are subject to change.

## Overview

Go Flow is a backend API built with Go (Golang) that provides comprehensive financial market data and insights. The application integrates with external financial APIs to deliver real-time stock quotes, historical data, and portfolio management capabilities through a clean REST API.

## Features

### 🚀 Current Features
- **Stock Data Integration**: Fetch real-time stock quotes and historical data from Alpha Vantage API
- **Database Storage**: PostgreSQL database with proper migrations for data persistence
- **RESTful API**: Clean HTTP endpoints for stock data retrieval and management
- **Docker Support**: Containerized PostgreSQL database setup
- **Repository Pattern**: Clean architecture with separation of concerns

### 🔄 In Development
- User authentication and authorization
- Portfolio management and tracking
- Stock watchlists and alerts
- Technical indicators and analysis
- News headlines integration
- WebSocket support for real-time updates

## Tech Stack

- **Backend**: Go (Golang) with Gin web framework
- **Database**: PostgreSQL with pgx driver
- **External APIs**: Alpha Vantage for financial data
- **Containerization**: Docker & Docker Compose
- **Migration**: Custom migration system

## Project Structure

```
go-flow/
├── cmd/
│   ├── server/          # Main server application
│   └── migrate/         # Database migration runner
├── internal/
│   ├── api/
│   │   └── handler/     # HTTP request handlers
│   ├── models/          # Data models and structs
│   ├── repository/      # Database layer
│   └── service/         # Business logic and external API clients
├── db/
│   └── migrations/      # SQL migration files
├── docker-compose.yaml  # Database container setup
├── Makefile            # Development commands
└── .env                # Environment variables
```

## Getting Started

### Prerequisites
- Go 1.19+ installed
- Docker and Docker Compose
- Alpha Vantage API key (free at [alphavantage.co](https://www.alphavantage.co/support/#api-key))

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/bashlui/go-flow.git
   cd go-flow
   ```

2. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

3. **Start the database**
   ```bash
   docker-compose up -d
   ```

4. **Run database migrations**
   ```bash
   make migrate
   ```

5. **Install dependencies**
   ```bash
   make deps
   ```

6. **Start the server**
   ```bash
   make run
   ```

The API will be available at `http://localhost:8080`

## API Endpoints

### Stock Operations

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/api/stocks` | Get all stocks from database |
| `GET` | `/api/stocks/:id` | Get specific stock by ID/Symbol |
| `POST` | `/api/stocks/fetch/:symbol` | Fetch stock data from Alpha Vantage and store |

### Example Usage

```bash
# Fetch Apple stock data from Alpha Vantage
curl -X POST http://localhost:8080/api/stocks/fetch/AAPL

# Get all stored stocks
curl http://localhost:8080/api/stocks

# Get specific stock
curl http://localhost:8080/api/stocks/AAPL
```

## Database Schema

The application uses PostgreSQL with the following main tables:

- **stocks**: Basic stock information (symbol, name, last_price)
- **stock_history**: Historical price data (OHLCV)
- **users**: User accounts (planned)
- **stock_watchlist**: User stock watchlists (planned)
- **stock_alerts**: Price alerts (planned)

## Development

### Available Make Commands

```bash
make migrate    # Run database migrations
make run       # Start the development server
make build     # Build the application binaries
make deps      # Install/update dependencies
make clean     # Clean build artifacts
make test      # Run tests
```

### Environment Variables

```env
DATABASE_URL=postgres://goflow_user:password@localhost:5432/goflow_db?sslmode=disable
ALPHA_VANTAGE_API_KEY=your_api_key_here
PORT=8080
```

## Contributing

This project is in early development. Contributions, suggestions, and feedback are welcome!

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Architecture Notes

### Design Patterns
- **Repository Pattern**: Database operations are abstracted through interfaces
- **Service Layer**: Business logic separated from HTTP handlers
- **Dependency Injection**: Services and repositories are injected into handlers

### Code Organization
- `handlers/`: HTTP request/response handling
- `services/`: External API clients and business logic
- `repository/`: Database operations and queries
- `models/`: Data structures and domain models

## Roadmap

- [ ] Complete user authentication system
- [ ] Implement portfolio tracking
- [ ] Add real-time WebSocket updates
- [ ] Integrate additional financial data sources
- [ ] Add comprehensive testing suite
- [ ] Implement caching layer (Redis)
- [ ] Add API rate limiting
- [ ] Create frontend dashboard
- [ ] Deploy to cloud platform

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

**Antonio Basilio** - [@bashlui](https://github.com/bashlui)

Project Link: [https://github.com/bashlui/go-flow](https://github.com/bashlui/go-flow)
