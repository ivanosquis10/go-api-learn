# Go API Learn

Este proyecto es una API RESTful desarrollada en Go utilizando el framework Gin, con autenticación JWT, conexión a base de datos PostgreSQL y arquitectura modular para controladores y middlewares.

## Características principales

- CRUD de usuarios y álbumes
- Autenticación y autorización con JWT
- Hash de contraseñas con bcrypt
- Organización modular (controllers, middlewares, models, utils)
- Variables de entorno para configuración
- Docker Compose para base de datos PostgreSQL
- Ejemplo de manejo de sesiones y cookies seguras

## Estructura del proyecto

```
├── controllers/      # Lógica de negocio y endpoints
├── initializers/     # Conexión y sincronización de la base de datos
├── middlwares/       # Middlewares personalizados (auth, etc)
├── models/           # Definición de modelos de datos
├── responses/        # Estructuras de respuesta para la API
├── utils/            # Utilidades y helpers
├── main.go           # Punto de entrada de la aplicación
├── go.mod            # Dependencias del proyecto
├── docker-compose.yml# Configuración de servicios (Postgres)
```

## Instalación y ejecución

1. Clona el repositorio:
   ```sh
   git clone https://github.com/ivanosquis10/go-api-learn.git
   cd go-api-learn
   ```
2. Copia el archivo de variables de entorno:
   ```sh
   cp .env.example .env
   # Edita .env según tu configuración
   ```
3. Levanta la base de datos con Docker Compose:
   ```sh
   docker-compose up -d
   ```
4. Instala las dependencias y ejecuta la API:
   ```sh
   go mod tidy
   go run .
   # o usa air para hot reload si está instalado
   air
   ```

## Variables de entorno

- `PORT`           → Puerto donde corre la API
- `DB_URL`         → URL de conexión a PostgreSQL
- `JWT_SECRET`     → Secreto para firmar los tokens JWT

## Requisitos

- Go 1.20+
- Docker y Docker Compose
- PostgreSQL (si no usas Docker)

## Licencia

MIT

---

> Desarrollado por [ivanosquis10](https://github.com/ivanosquis10)
