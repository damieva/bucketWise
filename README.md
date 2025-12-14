# 🪣 BucketWise

**BucketWise** is an automated personal finance app that intelligently categorizes your banking transactions into budgets — so you can understand your spending without manual effort.

---

## 📘 Table of Contents

- [Overview](#-overview)
- [Architecture](#-architecture)
- [Local Development Setup](#-local-development-setup)
    - [Requirements](#requirements)
    - [Environment Variables](#environment-variables)
    - [Makefile Commands](#makefile-commands)

---

## 💡 Overview

BucketWise automatically classifies your online banking transactions into categories like **Food**, **Leisure**, **Bills**, or **Mortgage**.  
It learns from your habits and organizes your finances intelligently so you can focus on making better financial decisions — not spreadsheets.

---

## 🏗️ Architecture

project-root/
├── cmd/
│ └── api/ # Go API entry point
├── pkg/ # Core packages (logic, DTOs, etc.)
├── scripts/
│ └── mongo_init.sh # MongoDB initialization script
├── Makefile # Automation for local setup
├── go.mod
└── README.md


---

## 💻 Local Development Setup

### Requirements

- 🐳 **Docker** (for running MongoDB)
- 🐍 **Make** (for automation)
- 🦫 **Go 1.22+** (for the API)

---

### Environment Variables

BucketWise uses MongoDB locally.  
The environment variables are handled in the `Makefile`, but you can override them if needed.

| Variable | Default | Description |
|-----------|----------|-------------|
| `MONGO_CONTAINER_NAME` | `local-mongodb` | Docker container name |
| `MONGO_PORT` | `27017` | Local MongoDB port |
| `MONGO_VOLUME` | `mongo_data` | Docker volume for persistence |
| `MONGO_DATABASE` | `bucketWise` | Database name |
| `MONGO_URI` | `mongodb://localhost:27017/bucketWise` | Default connection string for local API |

---

### Makefile Commands

| Command | Description |
|----------|-------------|
| `make start-up-local-env` | Starts MongoDB in Docker and runs initialization script |
| `make stop-local-env` | Stops and removes the MongoDB container |
| `make clean-local-env` | Removes MongoDB container and volume (fresh start) |
| `make run` | Generates Swagger docs and starts the Go API |
| `make swagger` | Regenerates Swagger documentation |

---

## 🧭 Frontend Setup (Next.js)

### 1. Install nvm

```bash
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.7/install.sh | bash
echo 'export NVM_DIR="$HOME/.nvm"' >> ~/.zshrc
echo '[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"' >> ~/.zshrc
source ~/.zshrc
```

### 2. Install Node LTS

```bash
nvm install --lts
```

### 3. Create Next.js project

```bash
cd frontend
npx create-next-app@latest . --ts
```

Selected options:
- ESLint: Yes
- React Compiler: No
- Tailwind CSS: Yes
- `src/` directory: Yes
- App Router: Yes
- Custom import alias: No

### 4. Environment variables

Create `frontend/.env.local`:

```
NEXT_PUBLIC_API_URL=http://localhost:8001
```

### 5. Run frontend

```bash
npm run dev
```

Frontend: http://localhost:3000  
Backend: http://localhost:8001
