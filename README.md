# Lorekeeper

Lorekeeper is a web application that helps Dungeons & Dragons players organise campaign information and rediscover 
details from their session notes. Players can browse campaigns and characters, upload PDF session notes, and ask 
natural-language questions about a campaign. Lorekeeper retrieves relevant passages from the notes and uses a locally 
hosted language model to generate an answer.

The project was built for the Enterprise Automation Services Trainee and Junior DevOps Engineer technical challenge. It 
demonstrates a full-stack application, persistent data, authenticated API routes, document processing, information 
retrieval, local AI integration, Docker-based services and automated backend/frontend checks.

## Goals

Lorekeeper is intended to help a player:

- Organise information by campaign.
- Browse campaign and character information.
- Upload session notes and retain their source documents.
- Extract text from PDF notes while preserving page numbers.
- Search campaign material for characters, locations, events, quests and other details.
- Ask natural-language questions and receive answers based on retrieved campaign excerpts.
- Retain source information with retrieved passages so answers can be checked against the original notes.

The implementation deliberately separates document extraction, chunking, retrieval and answer generation. This makes the 
retrieval stage testable independently of the language model and allows the model integration to be replaced without 
rewriting the rest of the pipeline.

## Prerequisites

Install the following before starting:

- Git
- Docker Desktop (or Docker Engine) with the Docker Compose plugin
- Go 1.26.5 (see `backend/go.mod`)
- Node.js 18 and npm
- Enough disk space for Docker images, the MySQL database and the Ollama model

The root Docker Compose configuration runs the Go API, MySQL 8 and Ollama. A one-shot Compose service downloads the 
configured Ollama model before the API starts. The Vue development server is run separately.

## Local setup

### 1. Clone the repository

```bash
git clone <repository-url>
cd lorekeeper
```

### 2. Configure local hostnames

The default frontend and CORS configuration use `jb-lorekeeper.com` for the frontend and `api.jb-lorekeeper.com` for the
API. Add this line to your operating system's hosts file:

```text
127.0.0.1 jb-lorekeeper.com api.jb-lorekeeper.com
```

Typical hosts-file locations:

- Windows: `C:\Windows\System32\drivers\etc\hosts` (edit as Administrator)
- macOS / Linux: `/etc/hosts` (edit with `sudo`)

If you prefer different hostnames or ports, update the frontend API URL and the backend `CORS_ORIGINS` setting to match 
your local setup.

### 3. Configure environment variables

Copy the root example environment file:

```bash
cp .env.example .env
```

Review the values before starting. The main settings are:

| Variable                            | Purpose                                       | Example value       |
|-------------------------------------|-----------------------------------------------|---------------------|
| `ENVIRONMENT`                       | Runtime environment                           | `development`       |
| `PORT`                              | Port the API listens on inside its container  | `7788`              |
| `DB_DRIVER`                         | Database driver                               | `mysql`             |
| `DB_HOST`                           | Database hostname on the Compose network      | `lorekeeper_mysql`  |
| `DB_PORT`                           | MySQL port on the Compose network             | `3306`              |
| `DB_NAME`                           | Database name                                 | `lorekeeper`        |
| `DB_USER` / `DB_PASSWORD`           | Database credentials                          | Set local values    |
| `EXPOSE_PORT`                       | Host port published for the API               | `7788`              |
| `EXPOSE_DB_PORT`                    | Host port published for MySQL                 | `56741`             |
| `JWT_SECRET`                        | Secret used to sign JWTs                      | Replace the example |
| `ADMIN_USERNAME` / `ADMIN_PASSWORD` | Credentials used for the seeded admin account | Set local values    |
| `OLLAMA_PORT`                       | Ollama port inside the Compose network        | `11434`             |
| `OLLAMA_EXPOSE_PORT`                | Host port published for Ollama                | `11434`             |
| `OLLAMA_MODEL`                      | Model downloaded and used by Ollama           | `qwen2.5:3b`        |

Important configuration notes:

- Set `JWT_SECRET` to a unique, sufficiently long secret.
- Change the example database and admin credentials for your local environment.
- Keep `DB_HOST=lorekeeper_mysql` and `DB_PORT=3306` when using the supplied Compose network. `EXPOSE_DB_PORT` is the 
separate host-side port.
- The API receives its Ollama URL from Compose as `http://lorekeeper_ollama:${OLLAMA_PORT}`. The API does not need to 
use the host-published Ollama port.
- The root `.env` is used by Docker Compose and is loaded by the Go application when running outside Docker. Do not 
commit real credentials or secrets. `.env.example` is a template, not a production configuration.

### 4. Start the API, database and local AI model

From the repository root:

```bash
docker compose up --build
```

To run in the background:

```bash
docker compose up --build -d
```

The stack starts:

- `lorekeeper_mysql`: MySQL 8 database, with persistent data in the `database_mysql` named volume.
- `lorekeeper_ollama`: local Ollama model server, with model data in the `ollama_data` named volume.
- `lorekeeper_ollama_init`: one-shot service that waits for Ollama and pulls the model configured by `OLLAMA_MODEL`.
- `lorekeeper_api`: Go API. Compose waits for MySQL to become healthy and for model initialisation to complete before 
starting the API.

The first start can take longer because Docker images and the configured language model need to be downloaded. 
Subsequent starts reuse the named volumes.

Useful commands:

```bash
# Follow all service logs
docker compose logs -f

# Follow only API logs
docker compose logs -f lorekeeper_api

# Follow Ollama logs
docker compose logs -f lorekeeper_ollama

# Stop the services while preserving database and model data
docker compose down

# Stop services and permanently delete database and model data
docker compose down -v
```

Deleting the volumes with `down -v` removes the local MySQL data and downloaded Ollama model.

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

The frontend uses `VUE_APP_API_BASE_URL` to locate the API. The example value is:

```text
http://api.jb-lorekeeper.com:7788
```

If you change the API host port using `EXPOSE_PORT`, update `VUE_APP_API_BASE_URL` in `frontend/.env` to match and 
restart the Vue development server.

### 6. Build the frontend

```bash
cd frontend
npm run build
```

The production build is written to `frontend/dist/`.

## Technology stack

| Area                         | Technology                                             |
|------------------------------|--------------------------------------------------------|
| Frontend                     | Vue 2.7, Vue Router 3, Vuex 3                          |
| HTTP client                  | Axios                                                  |
| Styling                      | Tailwind CSS                                           |
| Frontend tooling             | Vue CLI 5, Babel, ESLint                               |
| Backend                      | Go 1.26.5, Echo                                        |
| Database                     | MySQL 8                                                |
| Database access              | GORM (`github.com/jinzhu/gorm`) and sqlx               |
| Database migrations          | go-migrations (`github.com/ShkrutDenis/go-migrations`) |
| Authentication               | JWT access tokens and HttpOnly refresh-token cookie    |
| Password hashing             | `golang.org/x/crypto`                                  |
| Validation                   | `go-playground/validator`                              |
| PDF extraction               | `github.com/ledongthuc/pdf`                            |
| Text retrieval               | In-process BM25 ranking implementation                 |
| AI inference                 | Ollama, using the configured local model               |
| File storage                 | Local filesystem in `backend/assets/files`             |
| API documentation generation | Swaggo                                                 |
| Local environment            | Docker, Docker Compose                                 |
| CI                           | GitHub Actions                                         |

## How it works

The main session-note question-answering flow is:

1. A signed-in user opens a campaign and uploads a PDF session-notes file.
2. The API checks that the user can access the campaign and validates that the upload is a PDF no larger than 10 MB.
3. The original file is saved to local filesystem storage and a file record is associated with the campaign.
4. The PDF extractor reads the document page by page, retaining each page number.
5. The text is split into overlapping chunks. The current extractor uses a target chunk size of 50 words and an overlap 
of 10 words; empty pages are skipped.
6. The chunks are stored in MySQL and linked to the campaign file and its source page.
7. When the user asks a question, the API checks campaign access and ranks that campaign's chunks using BM25.
8. Up to 10 matching chunks are passed to Ollama as reference context. The prompt instructs the model to use only those 
excerpts, avoid inventing facts, and cite relevant excerpts by source number.
9. The API returns the generated answer together with the retrieved source results.

The retrieval step does not depend on Ollama. BM25 tokenises text using Unicode letters and numbers, scores terms using 
inverse document frequency and document-length normalisation, filters out low-scoring matches, and returns the 
highest-ranked results. The implementation currently uses `k1 = 1.5`, `b = 0.75` and a minimum score threshold of `0.5`.

Ollama is called through its `/api/chat` endpoint. Requests are non-streaming and use a temperature of `0.2`. The model
name and Ollama base URL are supplied through environment variables, so a different Ollama model can be configured 
without changing the question service.

## Application functionality

The current frontend provides:

- Registration and login screens.
- Authenticated home, campaigns and characters pages.
- A paginated campaign list with name filtering.
- Campaign details with Lorekeeper, Characters and Session Notes tabs.
- A PDF session-notes upload flow.
- A question input on the Lorekeeper tab and display of returned answers.
- A paginated character list with name filtering.
- Session storage for the access token and an HttpOnly cookie for the refresh token.

The backend also seeds reference and demonstration data on startup, including roles, users, races, classes, subclasses, 
campaigns, characters and campaign session notes. Migrations are run during database initialisation.

## API

The API is served by the Go Echo application. With the default example configuration, its base URL is:

```text
http://api.jb-lorekeeper.com:7788
```

All routes under `/campaigns`, `/campaigns/:id/files`, `/campaigns/:id/search`, `/characters` and `/files` require a 
valid access token. The refresh endpoint uses the refresh-token cookie.

### Authentication endpoints

| Method | Endpoint         | Authentication       | Description                                           |
|--------|------------------|----------------------|-------------------------------------------------------|
| `POST` | `/auth/register` | Public               | Register a user                                       |
| `POST` | `/auth/login`    | Public               | Log in and obtain an access token                     |
| `POST` | `/auth/refresh`  | Refresh-token cookie | Exchange a valid refresh token for a new access token |
| `GET`  | `/auth/me`       | Access token         | Return the authenticated user's details               |
| `POST` | `/auth/logout`   | Access token         | Clear the refresh-token cookie                        |

### Character endpoints

| Method | Endpoint      | Query parameters              | Description                                                                        |
|--------|---------------|-------------------------------|------------------------------------------------------------------------------------|
| `GET`  | `/characters` | `page`, `page_size`, `search` | Return a paginated list of the current user's characters; `search` filters by name |

### Campaign endpoints

| Method | Endpoint                | Query parameters / request body | Description                                                                                     |
|--------|-------------------------|---------------------------------|-------------------------------------------------------------------------------------------------|
| `GET`  | `/campaigns`            | `page`, `page_size`, `search`   | Return campaigns the user owns or has a character in; `search` filters by campaign name         |
| `GET`  | `/campaigns/:id`        | Campaign ID in path             | Return a campaign the user can access, including its related campaign data                      |
| `POST` | `/campaigns/:id/files`  | Multipart form field `file`     | Upload a PDF session-notes file, extract and chunk its text, and associate it with the campaign |
| `POST` | `/campaigns/:id/search` | JSON body containing `question` | Retrieve relevant campaign-note chunks and generate an answer using Ollama                      |

### File endpoints

| Method | Endpoint           | Description                                        |
|--------|--------------------|----------------------------------------------------|
| `GET`  | `/files/:filepath` | Serve a stored file through an authenticated route |

### Authentication and API notes

- The access token is returned in the JSON response and stored by the frontend in `sessionStorage`.
- The refresh token is stored in an HttpOnly cookie, rather than being exposed to frontend JavaScript.
- The frontend Axios client uses `withCredentials: true` and attaches the access token as a Bearer token when one is 
present.
- The Axios response interceptor attempts to refresh an expired access token and retry the original request.
- Registration, login and refresh routes are rate-limited.
- Protected routes use JWT middleware.
- CORS allowed origins can be configured through `CORS_ORIGINS` as a comma-separated list. The backend has a development
fallback for `http://jb-lorekeeper.com:8080`.
- API request and response structures are defined in `backend/api/requests` and `backend/api/responses`.
- Swagger documentation is generated from the Go annotations by the CI workflow and Docker startup command into 
`backend/docs`.

## Database and storage

MySQL stores users, roles, characters, campaigns, files, campaign-file associations and document chunks. The migration 
list is maintained in `backend/db/migrations/list` and registered in `backend/db/migrations/process/process.go`. 
Migrations and seeders run when the backend initialises its database connection.

Uploaded documents are stored on the API container's local filesystem under `backend/assets/files` (mounted into the 
container through the backend directory bind mount). The database stores file metadata and the relationship between a 
file, campaign and its extracted chunks.

The Compose configuration persists MySQL data and Ollama model data in named volumes. Uploaded files are currently held 
in the backend directory rather than a dedicated object-storage service, so ensure that directory is included in any 
backup strategy.

## Testing

### Backend tests

Backend tests are located in `backend/tests/`. They include API-level tests using a test server and database, plus 
dependency-level tests for document processing, file storage and search.

Start the test MySQL container from the repository root:

```bash
docker compose -f ./docker-compose-test.yml up --build -d
```

The test database uses the same database credentials and database name from the root `.env`, and publishes MySQL on 
`TEST_EXPOSE_DB_PORT`.

From the `backend` directory, run:

```bash
go test ./tests/... -v -timeout 120s
```

To run all Go package tests:

```bash
go test ./...
```

Stop the test database when finished:

```bash
docker compose -f ./docker-compose-test.yml down
```

The test database uses a named volume. To remove that data as well:

```bash
docker compose -f ./docker-compose-test.yml down -v
```

### Frontend checks

The frontend currently defines build and lint scripts:

```bash
cd frontend
npm run lint
npm run build
```

The lint script runs Vue CLI ESLint with `--fix`, so it may modify files. There is not currently a dedicated frontend 
unit-test script or frontend test suite configured.

### Continuous integration

The GitHub Actions workflow at `.github/workflows/ci.yml` runs on pull requests targeting `main` or `development`. It 
has separate backend and frontend jobs.

The backend job:

1. Starts a MySQL 8 service for tests.
2. Sets up the Go version from `backend/go.mod`.
3. Installs Swaggo and generates Swagger documentation.
4. Downloads Go modules and builds the API.
5. Creates a test `.env` file and runs `go test ./tests/...`.

The frontend job:

1. Sets up Node.js 18.
2. Installs frontend dependencies.
3. Builds the Vue application using a placeholder API URL.

## Repository layout

```text
.
├── .github/
│   └── workflows/          # GitHub Actions CI
├── backend/
│   ├── api/
│   │   ├── handlers/       # HTTP handlers
│   │   ├── requests/       # Request DTOs and validation
│   │   ├── responses/      # Response DTOs and wrappers
│   │   └── routes/         # Echo route registration
│   ├── assets/             # Seed/test assets and local file-storage location
│   ├── cmd/api/            # API application entry point
│   ├── bin/                # Build binary file destination
│   ├── config/             # Runtime configuration
│   ├── db/
│   │   ├── factories/      # Model factories
│   │   ├── migrations/     # Database migrations
│   │   ├── models/         # GORM models
│   │   ├── repositories/   # Database repositories
│   │   └── seeders/        # Reference and demonstration data
│   ├── docker/             # Backend Dockerfile
│   ├── docs/               # Generated Swagger documentation
│   ├── pkg/
│   │   ├── document/       # PDF page extraction and text chunking
│   │   ├── file_storage/   # Local filesystem storage abstraction
│   │   ├── middleware/     # JWT and logging middleware
│   │   ├── search/         # BM25 ranking
│   │   └── validation/     # Echo request validation
│   ├── services/            # Application and business logic
│   ├── tests/               # Backend tests, helpers and mocks
│   ├── go.mod
│   └── go.sum
├── frontend/
│   ├── public/
│   └── src/
│       ├── @core/           # Shared frontend utilities and components
│       ├── assets/          # Images and CSS
│       ├── components/      # Reusable Vue components
│       ├── layouts/         # Reusable Vue layouts
│       ├── libs/            # Frontend library configuration
│       ├── router/          # Vue Router configuration
│       ├── services/        # Axios API client and API services
│       ├── store/           # Vuex store and modules
│       └── views/           # Page-level Vue components
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

The current implementation provides a working end-to-end PDF ingestion and question-answering path, but there are areas 
that could be developed further:

- **Source display:** the API returns retrieved source chunks with an answer, but the current frontend primarily 
displays the answer text. A source panel linking excerpts to their document and page would make verification easier.
- **Document formats:** uploads currently accept PDF files only. Support for Markdown, plain text and other useful note 
formats could be added.
- **Retrieval quality:** BM25 provides a simple lexical baseline. Stop-word handling, stemming, configurable ranking 
parameters, evaluation datasets and hybrid/vector retrieval could be explored.
- **Chunking:** the current fixed word-count chunker could be improved to respect paragraph or section boundaries and 
configurable document characteristics.
- **Document lifecycle:** deleting/replacing uploads, retrying failed extraction, and cleaning up partially-created 
records and files need more complete workflows.
- **Storage:** local filesystem storage is suitable for local development, but object storage and a clearer 
backup/retention strategy would be preferable for deployment.
- **Security and operations:** add production-focused configuration validation, stronger upload hardening, request 
limits, observability, and a deployment-specific secrets strategy.
- **Frontend quality:** add component/unit tests and end-to-end tests, improve error and empty states, and make the 
source references available in the interface.
- **AI evaluation:** test answer faithfulness, citation correctness, retrieval recall and behaviour when the notes do 
not contain an answer across a representative set of questions and documents.
- **Deployment:** add a production deployment configuration and a CI/CD release workflow once the target hosting 
environment is chosen.

The application currently depends on the Ollama service and a downloaded local model for generated answers. If Ollama is
unavailable, document upload and BM25 retrieval remain separate backend capabilities, but question answering cannot 
complete.
