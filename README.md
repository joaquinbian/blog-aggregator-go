# Blog Aggregator - Go

Un agregador de feeds RSS desarrollado en Go que permite a los usuarios seguir y recolectar contenido de múltiples blogs y sitios web mediante sus feeds RSS.

## 📋 Descripción

Blog Aggregator es una herramienta de línea de comandos (CLI) que permite gestionar feeds RSS de manera eficiente. Los usuarios pueden registrarse, agregar feeds, seguirlos, y el sistema automáticamente recolecta y muestra el contenido actualizado de los feeds configurados.

### Características principales

- ✅ Gestión de usuarios (registro, login, listado)
- ✅ Administración de feeds RSS (agregar, listar, seguir, dejar de seguir)
- ✅ Recolección automática de contenido de feeds en intervalos configurables
- ✅ Almacenamiento persistente en PostgreSQL
- ✅ Sistema de autenticación basado en archivo de configuración
- ✅ Middleware para proteger comandos que requieren autenticación

## 🛠️ Tecnologías utilizadas

- **Go 1.25.1**: Lenguaje de programación principal
- **PostgreSQL**: Base de datos relacional
- **sqlc**: Generación de código type-safe para SQL
- **goose**: Herramienta de migraciones de base de datos
- **github.com/lib/pq**: Driver de PostgreSQL para Go
- **github.com/google/uuid**: Generación de UUIDs

## 📦 Dependencias

Las dependencias del proyecto están definidas en `go.mod`:

```go
github.com/google/uuid v1.6.0
github.com/lib/pq v1.10.9
```

## 🚀 Instalación

### Prerrequisitos

1. **Go 1.25.1 o superior**
   ```bash
   go version
   ```

2. **PostgreSQL** instalado y ejecutándose
   ```bash
   psql --version
   ```

3. **Goose** para migraciones (opcional pero recomendado)
   ```bash
   go install github.com/pressly/goose/v3/cmd/goose@latest
   ```

4. **sqlc** para generar código SQL (solo si modificas queries)
   ```bash
   go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
   ```

### Pasos de instalación

1. **Clonar el repositorio**
   ```bash
   git clone <url-del-repositorio>
   cd blog-aggregator-go
   ```

2. **Instalar dependencias**
   ```bash
   go mod download
   ```

3. **Configurar la base de datos**
   
   Crear una base de datos en PostgreSQL:
   ```bash
   createdb blog_aggregator
   ```

4. **Ejecutar migraciones**
   ```bash
   cd sql/schema
   goose postgres "user=tu_usuario dbname=blog_aggregator sslmode=disable" up
   cd ../..
   ```

5. **Crear archivo de configuración**
   
   Crear el archivo `.gatorconfig.json` en tu directorio home (`~/.gatorconfig.json`):
   ```json
   {
     "db_url": "postgres://usuario:contraseña@localhost:5432/blog_aggregator?sslmode=disable",
     "current_user_name": ""
   }
   ```

6. **Compilar el proyecto**
   ```bash
   go build -o gator
   ```

## 📖 Uso

### Comandos disponibles

#### Gestión de usuarios

**Registrar un nuevo usuario**
```bash
./gator register <nombre_usuario>
```

**Iniciar sesión**
```bash
./gator login <nombre_usuario>
```

**Listar todos los usuarios**
```bash
./gator users
```

**Resetear base de datos (elimina todos los usuarios)**
```bash
./gator reset
```

#### Gestión de feeds

**Agregar un nuevo feed**
```bash
./gator addfeed <nombre_feed> <url_feed>
```
Ejemplo:
```bash
./gator addfeed "Blog Golang" https://go.dev/blog/feed.atom
```

**Listar todos los feeds**
```bash
./gator feeds
```

**Seguir un feed existente**
```bash
./gator follow <url_feed>
```

**Ver feeds que estás siguiendo**
```bash
./gator following
```

**Dejar de seguir un feed**
```bash
./gator unfollow <url_feed>
```

#### Recolección de feeds

**Iniciar agregador automático**
```bash
./gator agg <intervalo>
```
Ejemplo:
```bash
./gator agg 1m    # Recolecta cada 1 minuto
./gator agg 30s   # Recolecta cada 30 segundos
./gator agg 1h    # Recolecta cada 1 hora
```

## 🗄️ Estructura de la base de datos

### Tablas

**users**
- `id`: UUID (PK)
- `created_at`: TIMESTAMP
- `updated_at`: TIMESTAMP
- `name`: VARCHAR(50) UNIQUE

**feeds**
- `id`: UUID (PK)
- `created_at`: TIMESTAMP
- `updated_at`: TIMESTAMP
- `name`: VARCHAR(50)
- `url`: VARCHAR(250) UNIQUE
- `user_id`: UUID (FK → users)
- `last_fetched_at`: TIMESTAMP

**feed_follows**
- `id`: UUID (PK)
- `created_at`: TIMESTAMP
- `updated_at`: TIMESTAMP
- `user_id`: UUID (FK → users)
- `feed_id`: UUID (FK → feeds)
- UNIQUE(user_id, feed_id)

## 📁 Estructura del proyecto

```
blog-aggregator-go/
├── internal/
│   ├── config/           # Manejo de archivo de configuración
│   │   └── config.go
│   ├── database/         # Código generado por sqlc
│   │   ├── db.go
│   │   ├── models.go
│   │   ├── users.sql.go
│   │   ├── feeds.sql.go
│   │   └── feeds_follows.sql.go
│   └── utils/            # Utilidades generales
│       └── utils.go
├── sql/
│   ├── schema/           # Migraciones de base de datos
│   │   ├── 0001_users.sql
│   │   ├── 0002_feeds.sql
│   │   ├── 0003_feeds_follows.sql
│   │   └── 0004_feeds.sql
│   └── queries/          # Queries SQL para sqlc
│       ├── users.sql
│       ├── feeds.sql
│       └── feeds_follows.sql
├── commands.go           # Sistema de comandos
├── middleware.go         # Middleware de autenticación
├── main.go              # Punto de entrada
├── rss_feed.go          # Manejo de feeds RSS
├── handle_*.go          # Handlers de comandos
├── go.mod
├── go.sum
├── sqlc.yaml            # Configuración de sqlc
└── README.md
```

## 🔧 Desarrollo

### Modificar queries SQL

Si necesitas modificar las queries de base de datos:

1. Edita los archivos en `sql/queries/`
2. Regenera el código con sqlc:
   ```bash
   sqlc generate
   ```

### Crear nuevas migraciones

```bash
cd sql/schema
goose create nombre_migracion sql
```

## 🤝 Contribuciones

Las contribuciones son bienvenidas. Por favor:

1. Haz fork del proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## 📝 Notas adicionales

- El archivo `.gatorconfig.json` almacena la configuración de conexión a la base de datos y el usuario actual logueado
- Los comandos que requieren autenticación (`addfeed`, `follow`, `following`, `unfollow`) verifican que haya un usuario logueado mediante middleware
- El sistema de agregación funciona en un loop infinito, recolectando el feed menos recientemente actualizado en cada iteración

## 🐛 Solución de problemas

**Error de conexión a la base de datos**
- Verifica que PostgreSQL esté ejecutándose
- Revisa las credenciales en `.gatorconfig.json`
- Asegúrate de que la base de datos existe

**Error "command does not exist"**
- Verifica que el comando esté correctamente escrito
- Consulta la lista de comandos disponibles arriba

**Error "no user logged in"**
- Debes ejecutar `./gator login <nombre>` antes de usar comandos protegidos

## 📄 Licencia

Este proyecto es de código abierto y está disponible bajo la licencia MIT (o la licencia que elijas especificar).

---

Desarrollado con ❤️ usando Go