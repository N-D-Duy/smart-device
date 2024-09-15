# Entry point script
#!/bin/sh
# Wait for PostgreSQL to be available
until pg_isready -h postgresql -p 5432 -U postgres; do
  echo "Waiting for PostgreSQL..."
  sleep 2
done

# Start application
exec "$@"
