# Go Monorepo

> Go monorepo PoC containing shared libraries, a full-stack web application, a
> REST API, and a management CLI.

## Commands

This monorepo uses the [`just`][just] command runner for interaction and general
commands.

### `just new-lib <LIB>`

Creates a new library within the monorepo by creating a new directory in
`libs/`, initialising the Go module with the appropriate module path, and
registering the new module with the `go.work` file.

### `just new-service <SERVICE>`

Functions the same as [`just new-lib <LIB>`](#just-new-lib-lib), but creates a
new service (instead of a library).

### `just rm-lib <LIB>`

Removes an existing library within the monorepo by deleting the library's entire
directory and removing the library's `go.work` entry.

### `just rm-service <SERVICE>`

Functions the same as [`just rm-lib <LIB>`](#just-rm-lib-lib), but removes an
existing service (instead of a library).

### `just modules [TYPE]`

Lists all Go modules in the monorepo. If `[TYPE]` is provided (e.g. `libs` or
`services`), the resultant list will be filtered to modules matching that type.

### `just mod-tidy`

Performs `go mod tidy` within each module in the monorepo.

### `just fmt`

Performs `go fmt` within each module in the monorepo.

### `just test`

Performs `go test ./...` within each module in the monorepo.

### `just run <SERVICE>`

Runs a service in the monorepo (from the provided `<SERVICE>` name).

> [!NOTE]
> Before running the service, `go mod tidy` is invoked on the service's Go
> module.

[just]: https://github.com/casey/just
