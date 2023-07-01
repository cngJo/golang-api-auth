<div align=center>
    <h1>Golang JWT auth with binary UUIDs</h1>
</div>

This example project shows and example implementation of JWT authentication in Golang using binary UUIDs as primary keys.

Under the hood, the project uses [GORM](https://gorm.io/) as ORM and [Gin](https://github.com/gin-gonic/gin).

This is heavily inspired by this blog post [JWT Authentication in Golang](https://codewithmukesh.com/blog/jwt-authentication-in-golang/) and an example of [how to use binary UUIDs with GORM](https://github.com/dipeshdulal/binary-uuid-gorm). 

**Missing Features**

- Refresh tokens

## Available Requests

### `POST /api/user/register`

**Body**

```json
{
    "name": "Test User",
    "email": "me@example.com",
    "username": "test.user",
    "password": "change!M3"
}
```

**Response**

```json
{
    "userId": "--uuid--",
    "email": "me@example.com",
    "username": "test.user"
}
```

**Example**

```bash
curl -X POST -H "Content-Type: application/json" \
-d '{"name": "Test User", "email": "me@example.com", "username": "test.user", "password": "change!M3"}' \
http://localhost:8080/api/user/register
```

### `POST /api/token`

**Body**

```json
{
    "email": "me@example.com",
    "password": "change!M3"
}
```

**Response**

```json
{
    "token": "--jwt--"
}
```

**Example**

```bash
curl -X POST -H "Content-Type: application/json" \
-d '{"email": "me@example.com", "password": "change!M3"}' \
http://localhost:8080/api/token
```

### `GET /api/secured/ping`

**Response**

```json
{
    "message": "pong"
}
```

**Example**

```bash
curl -X GET -H "Content-Type: application/json" \
-H "Authorization: Bearer --jwt--" \
http://localhost:8080/api/secured/ping
```