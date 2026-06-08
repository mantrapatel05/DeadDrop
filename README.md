# DeadDrop

DeadDrop is a Go API for encrypted, one-time secrets.

You create a drop, choose when it can be opened, and after it is read once it gets burned.

## Demo

![DeadDrop demo](docs/image.png)

The demo shows a full DeadDrop flow: health check, create drop, check status, knock three times, read the secret, then try to read it again after it has burned.

## What It Does

- Stores secret content encrypted with AES-GCM
- Supports time-based unlocks with `reveal_at`
- Supports knock-based unlocks with `knock_target`
- Burns a drop after a successful read
- Uses PostgreSQL for storage

## Setup

Install dependencies:

```bash
go mod tidy
```

Create a `.env` file:

```bash
cp .env.example .env
```

On PowerShell:

```powershell
Copy-Item .env.example .env
```

Fill in:

```env
ENCRYPTION_KEY=<64-character-hex-key>
PORT=8080
DATABASE_URL=postgres://postgres:postgres@localhost:5432/deaddrop?sslmode=disable
```

Generate an encryption key:

```bash
openssl rand -hex 32
```

Make sure PostgreSQL is running and the database in `DATABASE_URL` exists.

Run the app:

```bash
go run .
```

The server starts at:

```text
http://localhost:8080
```

## API

| Method | Route | Description |
| --- | --- | --- |
| `GET` | `/health` | Check if the server is alive |
| `POST` | `/drops` | Create a new encrypted drop |
| `GET` | `/drops/{id}/status` | Check if a drop is ready or burned |
| `POST` | `/drops/{id}/knock` | Add one knock to a knock-based drop |
| `GET` | `/drops/{id}` | Read a ready drop and burn it |
| `PATCH` | `/drops/{id}` | Update unlock settings |
| `DELETE` | `/drops/{id}` | Burn a drop manually |

## Example

Create a knock-based drop:

```cmd
curl -X POST http://localhost:8080/drops -H "Content-Type: application/json" --data "{\"content\":\"my secret\",\"knock_target\":3}"
```

Copy the returned `id`.

Knock three times:

```cmd
curl -X POST http://localhost:8080/drops/<id>/knock
curl -X POST http://localhost:8080/drops/<id>/knock
curl -X POST http://localhost:8080/drops/<id>/knock
```

Read the drop:

```cmd
curl http://localhost:8080/drops/<id>
```

Read it again and it returns `410 Gone`, because the drop has already been burned.
