## V0.0.5 - Local File Storage

### [Added]

- Added local file storage package with new endpoint for getting files
- Added logos to class models

## V0.0.4 - Characters Search

### [Added]

- Added search bar to characters list view
- Added no characters card to characters list view

## V0.0.3 - Campaigns List

### [Added]

- Added campaigns table with migration and new model
- Added campaign list endpoint with new handler, service repo and tests
- Added campaigns seeder and updated seeded characters to include campaign ids
- Added character with no campaign id
- Added campaigns route link card to home page
- Added campaigns list page with debounced search bar and new service

### [Changed]

- Updated characters to include a campaign id
- Updated character response objects to include their user and campaign
- Changed page layout for all pages so they're all the full height of the screen, updated all to match

### [Fixed]

- Fixed character seeder trying to insert duplicate primary keys

## V0.0.2 - Characters List

### [Added]

- Added races table migration with new model and seeded values
- Added classes table migration with new model and seeded values
- Added subclasses table migration with new model and seeded values
- Added characters table migration with new model
- Added character_classes and character_subclasses link table migrations
- Added character list endpoint with new service and repo
- Added generic paginate method to db
- Added character, class, subclass and race API response objects
- Added character list endpoint tests with factories for required models
- Added characters route link card to home page view
- Added characters list view with new character service
- Added characters seeder with example characters
- Added character endpoints to readme

### [Changed]

- Updated test instructions in readme to include test docker container

### [Removed]

- Removed redundant endpoint descriptions from readme

## V0.0.1 - Project Initialisation

### [Added]

- Initialised GitHub repository with `README.md`, `CHANGELOG.md`, `LICENSE.md`, and `.gitignore`
- Added `RESEARCH.md` and `AI.md` in accordance with the project brief
- Added notes on how AI was used during for Research & Learning, Planning & Architecture and Documentation & Testing
- Added docker compose files and env example to project route
- Added GitHub CI workflow and PR template
- Created Go Echo server with GORM and go-migrations structure with initial auth routes
- Added tests for auth routes
- Created Vue2 frontend with axios api request service
- Created auth pages and route protection

> - **[Added]** — New features, files, endpoints, components, or dependencies introduced in this version
> - **[Changed]** — Modifications to existing behaviour, UI, configuration, or structure that are not bug fixes
> - **[Fixed]** — Bug fixes and corrections to unintended behaviour
> - **[Removed]** — Features, files, endpoints, or dependencies that have been deleted