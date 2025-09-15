# Docker Build Examples

## Default Build
```bash
docker build -t frontend .
```

## Custom Environment Variables
```bash
# Custom backend URL and API prefix
docker build \
  --build-arg BACKEND=http://my-backend:8080 \
  --build-arg PREFIX=/api/v2 \
  -t frontend .
```

## With Docker Compose
The environment variables can also be overridden at runtime:
```yaml
services:
  frontend:
    build: .
    environment:
      - BACKEND=http://backend:3001
      - PREFIX=/api/v1
```

## Environment Variables Reference

| Variable | Default | Description |
|----------|---------|-------------|
| `BACKEND` | `http://backend:3001` | Backend service URL for API rewrites |
| `PREFIX` | `/api/v1` | API prefix path for routing |
