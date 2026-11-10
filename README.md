# Lorekeeper

Lorekeeper is a web application concept for helping Dungeons & Dragons players keep track of campaign information and quickly rediscover details from their notes.

The project is being developed as a technical challenge. The current codebase establishes the application foundation: a Vue frontend, a Go REST API, MySQL persistence, authentication, database migrations and automated tests. Campaign-note ingestion, search and AI-assisted question answering are the intended product features and are not yet implemented in the current scaffold.

## Goals

The longer-term goal is to let a player:

- Organise information by campaign.
- Upload or add session notes and other campaign documents.
- Search those notes for characters, locations, events, quests and other details.
- Ask natural-language questions and receive answers grounded in the campaign material.
- Inspect the source notes behind an answer, so information can be checked rather than blindly trusted.

The immediate engineering goals are to build a maintainable full-stack application, keep the frontend and backend responsibilities separate, persist application data, and provide a repeatable local development and testing environment.


## Prerequisites

Install the following before starting:

- Git
- Docker Desktop (or Docker Engine) with the Docker Compose plugin
- Go version compatible with `backend/go.mod`
- Node.js and npm (Node 18 is used by the current CI workflow)

The backend and database run in Docker Compose. The Vue development server is currently run separately using the Vue CLI.

## Local setup

### 1. Clone the repository

```bash
git clone <repository-url>
cd lorekeeper
```

### 2. Configure local hostnames

Add the following line to your operating system's hosts file:

```text
127.0.0.1 jb-lorekeeper.com api.jb-lorekeeper.com
```

This makes the local frontend and API hostnames resolve to your own machine.

Typical hosts-file locations:

- Windows: `C:\Windows\System32\drivers\etc\hosts` (edit as Administrator)
- macOS / Linux: `/etc/hosts` (edit with `sudo`)

### 3. Configure environment variables

Copy the example environment file:

```bash
cp .env.example .env
```

Review the values in `.env` before starting the application. In particular:

- Set `JWT_SECRET` to a unique, sufficiently long local secret.
- Set `ADMIN_USERNAME` and `ADMIN_PASSWORD` to local development credentials.
- Keep `DB_HOST=lorekeeper_mysql` so the API connects to the database using the Compose service name.
- **Set `DB_PORT=3306`**: this is the port MySQL listens on inside the Compose network. `EXPOSE_DB_PORT` controls the host-side published port and is separate.

Do not commit `.env` or real credentials. The repository's `.env.example` is a template, not a production configuration.

### 4. Start the API and database

From the repository root:

```bash
docker compose up --build
```

To run in the background:

```bash
docker compose up --build -d
```

The Compose configuration starts the Go API and MySQL. The API container waits for the database before starting. Database data is stored in the named `database_mysql` volume and is retained when containers are recreated.

Useful commands:

```bash
# Follow service logs
docker compose logs -f

# Follow only API logs
docker compose logs -f lorekeeper_api

# Stop the services (preserve the database volume)
docker compose down

# Stop services and delete the database volume
docker compose down -v
```

Deleting the volume with `down -v` permanently removes the local MySQL data.

### 5. Start the Vue frontend

In a second terminal:

```bash
cd frontend
cp .env.example .env
npm install
npm run serve
```

Open the frontend at:

```text
http://jb-lorekeeper.com:8080
```

The frontend's Axios client uses `VUE_APP_API_BASE_URL`. The example value points to the locally exposed API:

```text
http://api.jb-lorekeeper.com:7788
```

If you change the API host port in the root `.env` (`EXPOSE_PORT`), update `VUE_APP_API_BASE_URL` to match and restart the Vue development server.

### 6. Build the frontend

```bash
cd frontend
npm run build
```

The production build is written to `frontend/dist/`.

## Technology stack

| Area | Technology |
|---|---|
| Frontend | Vue 2.7, Vue Router, Vuex |
| HTTP client | Axios |
| Styling | Tailwind CSS |
| Frontend tooling | Vue CLI, Babel, ESLint |
| Backend | Go, Echo |
| Database | MySQL 8 |
| Database access | GORM (legacy `github.com/jinzhu/gorm`) and sqlx |
| Database migrations | go-migrations |
| Authentication | JWT access tokens and HttpOnly refresh-token cookie |
| Validation | go-playground/validator |
| API documentation generation | Swaggo |
| Local environment | Docker, Docker Compose |
| CI | GitHub Actions |

## API

The API is served by the Go Echo application. With the default example configuration, its base URL is:

```text
http://api.jb-lorekeeper.com:7788
```

All current routes are under `/auth`. JSON request and response bodies are used unless otherwise noted.

### Authentication endpoints

| Method | Endpoint | Authentication | Description |
|---|---|---|---|
| `POST` | `/auth/register` | Public | Register a user |
| `POST` | `/auth/login` | Public | Log in and obtain an access token |
| `POST` | `/auth/refresh` | Refresh-token cookie | Exchange a valid refresh token for a new access token |
| `GET` | `/auth/me` | Access token | Return the authenticated user's details |
| `POST` | `/auth/logout` | Access token | Clear the refresh-token cookie |

#### Register

`POST /auth/register`

```json
{
  "username": "adventurer",
  "email": "adventurer@example.com",
  "password": "change-me",
  "confirm_password": "change-me"
}
```

On success, the API returns an access token and user details. It also sets the `refresh_token` cookie.

#### Login

`POST /auth/login`

```json
{
  "username": "adventurer",
  "password": "change-me"
}
```

On success, the response contains an `access_token` and `user` object. The refresh token is set as an HttpOnly cookie.

#### Refresh access token

`POST /auth/refresh`

The refresh token is read from the `refresh_token` cookie. The request does not require a JSON body. The response contains a new access token and user details, and refreshes the cookie.

#### Get the current user

`GET /auth/me`

Send the access token in the request header:

```text
Authorization: Bearer <access_token>
```

#### Logout

`POST /auth/logout`

Send the access token in the request header. The API expires the refresh-token cookie.

### Authentication notes

- The access token is returned in the JSON response.
- The refresh token is stored in an HttpOnly cookie, rather than being exposed to frontend JavaScript.
- The frontend Axios instance is configured with `withCredentials: true` and attaches the access token when one is present in session storage.
- Login, registration and refresh routes are rate-limited.
- Protected routes use JWT middleware.
- CORS allowed origins are configured with `CORS_ORIGINS` (comma-separated). The backend has a development fallback if it is unset.

Campaign, document, search and question-answering endpoints will be documented here when those features are implemented.

## Testing

### Backend tests

Backend tests are located in `backend/tests/`.

The test server uses a MySQL database configured through the test environment variables. Start a suitable MySQL instance and provide the test database settings before running the suite.

From the `backend` directory:

```bash
go test ./tests/... -v -timeout 120s
```

To run all Go package tests:

```bash
go test ./...
```

### Frontend checks

The frontend currently defines build and lint scripts:

```bash
cd frontend
npm run lint
npm run build
```

There is not currently a dedicated frontend unit-test script or frontend test suite configured.

### Continuous integration

The GitHub Actions workflow at `.github/workflows/ci.yml` runs on pull requests targeting `main` or `development`. It:

1. Starts a MySQL 8 service for backend tests.
2. Sets up the Go toolchain and downloads Go modules.
3. Generates Swagger documentation and builds the backend.
4. Runs the backend test suite.
5. Sets up Node.js, installs frontend dependencies and builds the Vue application.

## Repository layout

```text
.
├── .github/
│   └── workflows/       # GitHub Actions CI
├── backend/
│   ├── api/              # Echo server, handlers, requests, responses and routes
│   ├── cmd/api/          # API application entry point
│   ├── config/           # Backend configuration
│   ├── db/
│   │   ├── migrations/   # Database migrations
│   │   ├── models/       # GORM models
│   │   ├── repositories/ # Database repositories
│   │   └── seeders/      # Initial/reference data
│   ├── services/         # Application/business logic
│   ├── tests/            # Backend integration tests and test helpers
│   └── docker/           # Backend Dockerfile
├── frontend/
│   └── src/
│       ├── @core/        # Shared frontend utilities and components
│       ├── components/   # Reusable Vue components
│       ├── router/       # Vue Router configuration
│       ├── services/     # Axios API client and frontend services
│       ├── store/        # Vuex store and modules
│       └── views/        # Page-level Vue components
├── .dockerignore
├── .env.example
├── .gitignore
├── AI.md
├── CHANGELOG.md
├── docker-compose.yml
├── docker-compose-test.yml
├── LICENSE.md
├── README.md
└── RESEARCH.md
```

## Current limitations and future work

The current repository is the foundation for Lorekeeper rather than a complete campaign assistant. Planned next steps include:

- Campaign creation and campaign-scoped data.
- Uploading and managing campaign notes (initially plain text and Markdown).
- Extracting and chunking document content for retrieval.
- Search across campaign material, with relevant source references.
- AI-assisted answers grounded in retrieved passages, including a clear response when the notes do not contain an answer.
- Tests for document processing, retrieval quality and new API endpoints.
- A frontend test suite and end-to-end coverage.

The priority is to implement these capabilities incrementally and keep the retrieval and source-attribution behaviour testable independently of any AI model.

## Development notes

- Keep secrets in local environment files; never commit real credentials.
- Use migrations for database schema changes.
- Keep HTTP handlers focused on request validation and response handling; put application logic in services and persistence logic in repositories.
- Add tests alongside new behaviour.
- Update this README, `RESEARCH.md` and `AI.md` as the implementation and technical decisions evolve.
