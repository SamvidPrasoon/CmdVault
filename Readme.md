# 🗄️ cmdvault

> A blazing-fast CLI tool for developers to **save, organize, search, and execute** frequently used shell and infra commands — your personal command vault.

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go)
![Platform](https://img.shields.io/badge/platform-Linux%20%7C%20macOS%20%7C%20Windows-lightgrey?style=flat-square)
![Build](https://img.shields.io/badge/build-passing-brightgreen?style=flat-square)

---

## 🧠 Why cmdvault?

Every developer has that one `kubectl` command they can never remember. Or that `docker` cleanup one-liner buried in some Notion note. Or 10 aliases that only work on one machine.

**cmdvault** solves this by giving you a local, searchable, taggable command store — right in your terminal.

```bash
# Save it once
cmdvault add deploy-prod "kubectl apply -f ./k8s/" --desc "Deploy to production" --tags k8s,infra

# Run it anywhere
cmdvault run deploy-prod

# Find it fast
cmdvault search k8s
```

No cloud. No login. No config. Just a fast local binary.

---

## ✨ Features

| Feature              | Description                                                |
| -------------------- | ---------------------------------------------------------- |
| 💾 **Save commands** | Store any shell command with a name, description, and tags |
| ▶️ **Run commands**  | Execute saved commands instantly by name                   |
| 🔍 **Smart search**  | Search by name, description, or tag                        |
| 📋 **List all**      | View all saved commands in a clean table with run counts   |
| 🗑️ **Delete**        | Remove commands you no longer need                         |
| 📦 **Export**        | Export your entire vault to a portable JSON file           |
| 📥 **Import**        | Import commands from a JSON file (team sharing)            |
| 📊 **Run tracking**  | Tracks how many times each command has been executed       |

---

## 🚀 Installation

### Option 1: Build from source

```bash
git clone https://github.com/yourusername/cmdvault.git
cd cmdvault
go build -o cmdvault .

# Move to PATH (Linux/macOS)
mv cmdvault /usr/local/bin/
```

### Option 2: Go install

```bash
go install github.com/yourusername/cmdvault@latest
```

### Requirements

- Go 1.21+
- Works on Linux, macOS, Windows

---

## 📖 Usage

### Add a command

```bash
cmdvault add <name> "<command>" [flags]

# Examples
cmdvault add deploy-prod "kubectl apply -f ./k8s/"
cmdvault add deploy-prod "kubectl apply -f ./k8s/" --desc "Deploy all services to production"
cmdvault add docker-clean "docker system prune -af" --desc "Nuke all docker cache" --tags docker,cleanup
cmdvault add port-forward "kubectl port-forward svc/myapp 8080:80 -n staging" --tags k8s,debug
```

**Flags:**

```
--desc    Short description of what the command does
--tags    Comma-separated tags for grouping (e.g. k8s,infra,docker)
```

---

### List all commands

```bash
cmdvault list
```

```
NAME                 COMMAND                             TAGS                  RUNS
─────────────────────────────────────────────────────────────────────────────────────
deploy-prod          kubectl apply -f ./k8s/             k8s, infra            12
docker-clean         docker system prune -af             docker, cleanup       4
port-forward         kubectl port-forward svc/myapp...   k8s, debug            7
pg-dump              pg_dump -U postgres mydb > ...      postgres, backup      2
```

---

### Run a command

```bash
cmdvault run <name>

# Example
cmdvault run deploy-prod
# ▶ Running: kubectl apply -f ./k8s/
# deployment.apps/myapp configured
# service/myapp unchanged
```

---

### Search commands

```bash
cmdvault search <query>

# Search by tag
cmdvault search k8s

# Search by keyword in name or description
cmdvault search deploy
cmdvault search cleanup
```

---

### Delete a command

```bash
cmdvault delete <name>

# Example
cmdvault delete port-forward
# 🗑️  Deleted 'port-forward'
```

---

### Export & Import

```bash
# Export all commands to JSON (great for sharing with your team)
cmdvault export my-commands.json

# Import commands from JSON
cmdvault import my-commands.json
```

**Example export format:**

```json
[
  {
    "id": "1718123456789",
    "name": "deploy-prod",
    "cmd": "kubectl apply -f ./k8s/",
    "description": "Deploy all services to production",
    "tags": ["k8s", "infra"],
    "created_at": "2024-06-15T10:30:00Z",
    "run_count": 12
  }
]
```

---

## 🏗️ Architecture

```
cmdvault/
├── main.go           # Entry point
├── cmd/              # All CLI commands (Cobra)
│   ├── root.go       # Root command + DB init
│   ├── add.go        # Save a command
│   ├── run.go        # Execute a command
│   ├── list.go       # List all commands
│   ├── search.go     # Search commands
│   ├── delete.go     # Delete a command
│   └── export.go     # Export/Import JSON
└── store/
    └── store.go      # bbolt DB layer (CRUD)
```

### Tech Stack

| Library                                   | Purpose                                   |
| ----------------------------------------- | ----------------------------------------- |
| [Cobra](https://github.com/spf13/cobra)   | CLI command structure and flag parsing    |
| [bbolt](https://github.com/etcd-io/bbolt) | Embedded key-value store (no external DB) |

### Why bbolt?

bbolt is an embedded, zero-dependency key-value database written in pure Go. It stores data in a single local file (`~/.cmdvault.db`), meaning:

- ✅ No database server to run
- ✅ No network dependency
- ✅ Works completely offline
- ✅ Single binary distribution

---

## 🔧 How It Works

1. **Storage**: All commands are stored in `~/.cmdvault.db` (bbolt key-value file in your home directory)
2. **Key**: Command `name` is used as the key
3. **Value**: Full `Command` struct serialized as JSON
4. **Execution**: Commands are run via `os/exec` with `sh -c`, piping stdout/stderr directly to your terminal
5. **Search**: Linear scan over all stored commands matching against name, description, and tags

---

## 💡 Real-World Use Cases

```bash
# DevOps / Infra
cmdvault add k8s-nodes "kubectl get nodes -o wide" --tags k8s
cmdvault add restart-deploy "kubectl rollout restart deployment/myapp" --tags k8s
cmdvault add tf-plan "terraform plan -var-file=prod.tfvars" --tags terraform,infra

# Docker
cmdvault add docker-ps "docker ps --format 'table {{.Names}}\t{{.Status}}\t{{.Ports}}'" --tags docker
cmdvault add docker-logs "docker logs -f --tail=100 myapp" --tags docker,debug

# Database
cmdvault add pg-connect "psql -U postgres -h localhost -d mydb" --tags postgres
cmdvault add mongo-dump "mongodump --uri='mongodb://localhost:27017' --out=./backup" --tags mongo,backup

# Git
cmdvault add git-cleanup "git branch | grep -v 'main' | xargs git branch -D" --tags git,cleanup
cmdvault add git-undo "git reset --soft HEAD~1" --tags git
```

---

## 🛣️ Roadmap

- [ ] `cmdvault import` — Import commands from JSON
- [ ] `--dry-run` flag on `run` — Preview command before executing
- [ ] HTTP API mode — Expose vault as local REST API for team sharing
- [ ] Shell autocomplete — Tab completion for saved command names
- [ ] `cmdvault edit` — Edit a saved command in-place
- [ ] Encrypted vault — Optional AES encryption for sensitive commands

---

## 🤝 Contributing

Contributions are welcome!

```bash


Built by **[SAMVID]** — a Go CLI tool born out of frustration with forgetting commands at 2am.

> ⭐ If this saved you from Googling that one `docker` command again, give it a star!
```
