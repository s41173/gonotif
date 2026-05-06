# Go API - Social E-Commerce

A modular **Go API** for community social e-commerce platform.

---

## 🚀 Features
- Social E-Commerce API using **Go**, **Gin** & **GORM**
- **MySQL** for persistent data
- **Redis** for caching, session management, rate limiting & pub/sub notification
- Product list & detail cached on first load
- Event list cached for fast retrieval
- Session stored in Redis after login
- **JWT** for authentication
- **Google OAuth** login support
- **Swagger** API documentation
- **WhatsApp Notification** for OTP, payment & event updates (async via Redis Pub/Sub)
- **Voucher & Discount** system with target audience segmentation and usage limit control
- **Payment Gateway** integration via Flip (including webhook callback)
- **Shipping Integration** via RajaOngkir (shipping rate & AWB tracking)
- **Event Management** (participant, public & merchant registration)
- **Product Management** with full CRUD, search autocomplete, city filter, and front page display (latest & best seller)
- **Wallet & Points** system for member rewards
- **Wishlist** management
- **Cron Job** for auto-delete unpaid orders & expired vouchers
- **Rate Limiter** via Redis (100 requests/minute per IP)
- **Database Connection Pooling** for optimized DB performance
- **Structured Logging** via Uber Zap
- **Parallel cache warmup** on startup via goroutine & WaitGroup
- **Database indexing** strategy for optimized query performance
- Environment-based configuration via `.env`
- Integration Testing available

---

## 🏗 Architecture / Tech Stack

Client (Browser / Postman)
│
▼
Go REST API (Gin + GORM)
│
├──────────────────────────┐
▼                          ▼
┌───────────────┐   ┌───────────────┐
│   MySQL DB    │   │  Redis Cache  │ │
└───────────────┘   └───────────────┘


---------------

| Layer          | Technology              |
|----------------|-------------------------|
| Language       | Go                      |
| Framework      | Gin                     |
| ORM            | GORM                    |
| Database       | MySQL                   |
| Cache          | Redis                   |
| Auth           | JWT + Google OAuth      |
| Payment        | Flip                    |
| Shipping       | RajaOngkir              |
| Notification   | Fontee API              |
| Logging        | Uber Zap                |
| Docs           | Swagger                 |
| Rate Limiter   | Redis + ulule/limiter   |
| Deployment     | VPS + CI/CD Github      |


---

## 🏛 Clean Architecture

This project follows a modular clean architecture pattern inspired by HMVC, where each domain is self-contained.

internal/
├── auth/
│   ├── dto/
│   ├── handler/
│   ├── model/
│   ├── repository/
│   ├── service/
│   └── routes/
├── user/
├── event/
├── product/
├── cart/
├── order/
├── chapter/
├── slider/
└── ...

Each module follows a strict layered pattern:

| Layer          | Responsibility                                      |
|----------------|-----------------------------------------------------|
| `handler`      | Handle HTTP request & response                      |
| `service`      | Business logic                                      |
| `repository`   | Database queries via GORM                           |
| `model`        | Database struct & table mapping                     |
| `dto`          | Request & response data transfer objects            |
| `routes`       | Route registration per module                       |

-----------------

## 🔄 CI/CD

Automated pipeline via **GitHub Actions**
- ✅ Build check on every push
- ✅ Auto deploy to VPS on merge to `master`

-----------------

## 📦 Endpoints

### General
| Method | Endpoint  | Description       |
|--------|-----------|-------------------|
| GET    | /city     | Get all cities    |

### Auth
| Method | Endpoint  | Description           |
|--------|-----------|-----------------------|
| POST   | /login    | Login                 |
| POST   | /oauth    | Login with Google     |
| GET    | /decode   | Decode user session   |
| GET    | /logout   | Logout                |
| POST   | /forgot   | Reset password        |
| POST   | /otp      | Request OTP           |
| POST   | /verify   | Verify user account   |

### User
| Method | Endpoint         | Description              |
|--------|------------------|--------------------------|
| POST   | /register        | Register new user        |
| GET    | /user            | Get current user         |
| PUT    | /update          | Update user profile      |
| PUT    | /password        | Change password          |
| POST   | /image           | Upload profile image     |
| POST   | /notif           | Get notifications        |
| GET    | /notif_detail/:id| Get notification detail  |
| POST   | /wallet          | Get wallet & points      |

### Product
| Method | Endpoint               | Description                        |
|--------|------------------------|------------------------------------|
| POST   | /product               | Get list of products               |
| GET    | /product/:sku          | Get product by SKU                 |
| POST   | /product/search        | Search products by name            |
| GET    | /product_front/:type   | Get front page products            |
| GET    | /product_city          | Get cities for product filter      |
| POST   | /product/refresh-cache | Refresh product cache              |

### Cart
| Method | Endpoint            | Description              |
|--------|---------------------|--------------------------|
| GET    | /cart               | Get user cart            |
| POST   | /cart               | Add product to cart      |
| DELETE | /cart/:id           | Delete cart item         |
| DELETE | /cart/clean         | Clean all cart items     |
| PUT    | /cart/notes/:id     | Set cart item notes      |
| PUT    | /cart/pickup/:id    | Toggle pickup option     |
| PUT    | /cart/publish/:id   | Toggle publish option    |

### Order
| Method | Endpoint              | Description              |
|--------|-----------------------|--------------------------|
| POST   | /order                | Get list of orders       |
| GET    | /order/:id            | Get order detail         |
| GET    | /checkout             | Checkout order           |
| GET    | /order/tracking/:awb  | Track shipment by AWB    |

### Event
| Method | Endpoint               | Description                    |
|--------|------------------------|--------------------------------|
| POST   | /events                | Get list of events             |
| GET    | /events/:id            | Get event detail               |
| GET    | /event/register/:id    | Register as participant        |
| POST   | /event/public          | Register as public participant |
| POST   | /event/merchant        | Register as merchant           |

### Wishlist
| Method | Endpoint              | Description                   |
|--------|-----------------------|-------------------------------|
| POST   | /wishlist             | Get user wishlist             |
| GET    | /wishlist/:product_id | Toggle product wishlist       |
| GET    | /iswishlist/:product_id | Check if product wishlisted |

### Shipping
| Method | Endpoint                      | Description                   |
|--------|-------------------------------|-------------------------------|
| GET    | /province_shipping            | Get all provinces             |
| GET    | /city_shipping/:province_id   | Get cities by province        |
| GET    | /district_shipping/:city_id   | Get districts by city         |
| PUT    | /set_shipping                 | Set shipping address          |

### Article
| Method | Endpoint               | Description                   |
|--------|------------------------|-------------------------------|
| POST   | /article               | Get list of articles          |
| GET    | /article/:permalink    | Get article by permalink      |
| GET    | /article_category      | Get article categories        |

### Chapter
| Method | Endpoint          | Description              |
|--------|-------------------|--------------------------|
| POST   | /chapter          | Get all chapters         |
| GET    | /chapter/:id      | Get chapter by ID        |
| POST   | /chapter_front    | Get chapters paginated   |

### Slider
| Method | Endpoint        | Description              |
|--------|-----------------|--------------------------|
| POST   | /slider         | Get list of sliders      |
| GET    | /slider/refresh | Refresh slider cache     |

---

## ⚡ Performance & Reliability

### Rate Limiting
- **100 requests per minute** per IP address
- Powered by Redis via `ulule/limiter`
- Returns `429 Too Many Requests` when limit exceeded

### Database Connection Pooling
| Setting              | Value     |
|----------------------|-----------|
| Max Open Connections | 25        |
| Max Idle Connections | 10        |
| Max Conn Lifetime    | 5 minutes |

---

## 📖 API Documentation

## ⚡ Performance & Reliability

### Rate Limiting
- **100 requests per minute** per IP address
- Powered by Redis via `ulule/limiter`
- Returns `429 Too Many Requests` when limit exceeded

### Database Connection Pooling
| Setting              | Value     |
|----------------------|-----------|
| Max Open Connections | 25        |
| Max Idle Connections | 10        |
| Max Conn Lifetime    | 5 minutes |


## ⚙️ Usage

### Run locally
```bash
cp .env.example .env
go mod tidy
go run main.go
```

### Run tests
```bash
go test -v ./tests/...
```

---

## 🌐 Demo & 📖 API Documentation

URL : https://goapi.dswip.com/
Swagger : https://goapi.dswip.com/swagger/index.html
Authentication : Bearer Token

> ⚠️ Demo credentials are for testing purposes only