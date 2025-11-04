#!/bin/bash

# Script to run backend natively with environment variables from .env

set -e

# Get the directory where the script is located
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"

# Load environment variables from .env file
if [ -f "$PROJECT_ROOT/.env" ]; then
    echo "📝 Loading environment variables from .env..."
    export $(cat "$PROJECT_ROOT/.env" | grep -v '^#' | xargs)
    echo "✅ Environment variables loaded"
else
    echo "⚠️  Warning: .env file not found. Using defaults."
fi

# Set defaults if not set
export DB_HOST=${DB_HOST:-localhost}
export DB_PORT=${DB_PORT:-5432}
export DB_USER=${DB_USER:-postgres}
export DB_PASSWORD=${DB_PASSWORD:-password}
export DB_NAME=${DB_NAME:-sumb}
export DB_SSLMODE=${DB_SSLMODE:-disable}
export SERVER_PORT=${SERVER_PORT:-8080}
export DEBUG_PORT=${DEBUG_PORT:-40000}
export GO_ENV=${GO_ENV:-development}
export JWT_SECRET=${JWT_SECRET:-dev-secret}
export JWT_EXPIRE_HOURS=${JWT_EXPIRE_HOURS:-24}
export CORS_ALLOWED_ORIGINS=${CORS_ALLOWED_ORIGINS:-http://localhost:5173}

echo ""
echo "🚀 Starting backend application..."
echo "📍 Database: $DB_USER@$DB_HOST:$DB_PORT/$DB_NAME"
echo "🌐 Server port: $SERVER_PORT"
echo "🐛 Debug port: $DEBUG_PORT"
echo ""

cd "$PROJECT_ROOT/back"
go run ./cmd/server

