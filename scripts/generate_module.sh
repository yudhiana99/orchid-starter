#!/bin/bash

# Ambil nama module dari argumen atau gunakan default
MODULE_NAME="${1:-default}"
MODULES_ROOT="./modules/${MODULE_NAME}"

# Function to determine package name from file path
get_package_name() {
  local file_path="$1"
  local dir_name=$(dirname "$file_path")

  case "$dir_name" in
    "delivery/api/rest/v2") echo "v2" ;;
    "delivery/api/rest") echo "rest" ;;
    "delivery/event") echo "event" ;;
    "repository") echo "repository" ;;
    "usecase") echo "usecase" ;;
    "model/db") echo "modelDB" ;;
    "model/request") echo "modelRequest" ;;
    "model/response") echo "modelResponse" ;;
    "model/repository") echo "modelRepository" ;;
    "model/usecase") echo "usecase" ;;
    *) echo "main" ;;
  esac
}

# Struktur paths relatif dari MODULES_ROOT
paths=(
  "delivery/api/rest/handler.go"
  "delivery/api/rest/v2/handler_v2.go"
  "delivery/event/event.go"
  "repository/${MODULE_NAME}_repository.go"
  "repository/repository_interface.go"
  "usecase/${MODULE_NAME}_usecase.go"
  "usecase/usecase_interface.go"
  "model/db/${MODULE_NAME}.go"
  "model/request/request_${MODULE_NAME}.go"
  "model/response/response_${MODULE_NAME}.go"
  "model/repository/${MODULE_NAME}_repo_input.go"
  "model/usecase/${MODULE_NAME}_usecase_input.go"
)

# Loop buat direktori dan file
for path in "${paths[@]}"; do
  dir_path="${MODULES_ROOT}/$(dirname "$path")"
  file_path="${MODULES_ROOT}/${path}"
  package_name=$(get_package_name "$path")

  mkdir -p "$dir_path"

  if [ ! -f "$file_path" ]; then
    echo "package $package_name" > "$file_path"
    echo "Created: $file_path (package: $package_name)"
  else
    echo "Already exists: $file_path"
  fi
done

echo "✅ Module '${MODULE_NAME}' structure created at: $MODULES_ROOT"
