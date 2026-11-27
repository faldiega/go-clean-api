# golang-clean-api
Learn golang from the scratch, step by step, with Echo framework and Clean Architecture.

```
golang-simple-api
├─ cmd
│  └─ api
│     └─ main.go
├─ go.mod
├─ go.sum
├─ internal
│  ├─ config
│  │  ├─ configuration.go
│  │  ├─ constants.go
│  │  └─ dotenv.go
│  ├─ delivery
│  │  └─ http
│  │     ├─ dto
│  │     │  ├─ auth_dto.go
│  │     │  ├─ user_dto.go
│  │     │  └─ validator_dto.go
│  │     ├─ handler
│  │     │  ├─ auth_handler.go
│  │     │  └─ user_handler.go
│  │     ├─ middleware
│  │     │  ├─ check_url_unauthorized.go
│  │     │  ├─ jwt.go
│  │     │  ├─ middlewares.go
│  │     │  ├─ recovery.go
│  │     │  ├─ timeout.go
│  │     │  └─ trace.go
│  │     └─ routes
│  │        ├─ routes.go
│  │        └─ routes_module.go
│  ├─ domain
│  │  ├─ entity
│  │  │  └─ user.go
│  │  ├─ repository
│  │  │  └─ user_repository.go
│  │  └─ usecase
│  │     └─ user_usecase.go
│  ├─ infrastructure
│  │  ├─ database
│  │  │  └─ postgres.go
│  │  ├─ jwt.go
│  │  └─ logger
│  │     └─ zap_logger.go
│  ├─ initialize
│  │  └─ initialize.go
│  ├─ repository
│  │  └─ user_repository.go
│  ├─ server
│  │  └─ server.go
│  ├─ usecase
│  │  └─ user_usecase.go
│  └─ utils
│     └─ converter.go
├─ pkg
│  └─ common
│     ├─ constants
│     │  ├─ context.go
│     │  ├─ errors.go
│     │  └─ http.go
│     ├─ graceful
│     │  └─ graceful.go
│     ├─ logger
│     │  └─ logger.go
│     ├─ pagination
│     │  └─ pagination.go
│     ├─ response
│     │  ├─ http_response.go
│     │  └─ response.go
│     └─ validator
│        ├─ error.go
│        └─ validator.go
└─ README.md

```