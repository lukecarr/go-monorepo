go_work_file := "go.work"

# Runs a service located in `./services/`
run SERVICE:
  (cd ./services/{{SERVICE}} && go mod tidy)
  go run ./services/{{SERVICE}}/


# Lists all Go modules within this monorepo
modules TYPE="":
  grep -o '\./{{TYPE}}[^)]*' "{{go_work_file}}" | awk '{print $1}'

[private]
run-in-modules COMMAND:
  for module in `just modules`; do (cd $module && {{COMMAND}}) & done; wait

# Performs `go mod tidy` against all modules
mod-tidy: (run-in-modules "go mod tidy && echo \"Tidied $module\"")

# Performs `go fmt` against all modules
fmt: (run-in-modules "go fmt && echo \"Formatted $module\"")

# Performs `go test ./...` against all modules
test: (run-in-modules "go test ./...")

# Creates a new "library" module
new-lib LIB: (new "libs" LIB)

# Creates a new "service" module
new-service SERVICE: (new "services" SERVICE)

[private]
new TYPE NAME:
  mkdir -p ./{{TYPE}}/{{NAME}}
  (cd ./{{TYPE}}/{{NAME}} && go mod init github.com/lukecarr/go-monorepo/{{TYPE}}/{{NAME}})
  go work use ./{{TYPE}}/{{NAME}}

# Removes an existing "library" module
[confirm]
rm-lib LIB: (rm "libs" LIB)

# Removes an existing "service" module
[confirm]
rm-service SERVICE: (rm "services" SERVICE)

[private]
rm TYPE NAME:
  sed -i '/\.\/{{TYPE}}\/{{NAME}}/d' {{go_work_file}}
  rm -rf ./{{TYPE}}/{{NAME}}