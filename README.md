# aacevski.com

minimal portfolio built with go templates. deployed on cloudflare pages.

## development

run the dev server:
```bash
make dev
# or
air
```

visit: http://localhost:8080

## build static site

generate static files:
```bash
make build
# or
go run cmd/build/main.go
```

output will be in `./dist`

### books section (koreader integration)

the build process can fetch reading statistics from your KOReader database synced to Koofr.

**setup:**

1. copy the example env file:
```bash
cp env.example .env
```

2. edit `.env` with your Koofr credentials:
```bash
KOOFR_EMAIL=your-email@example.com
KOOFR_PASSWORD=your-password
KOREADER_DB_PATH=/KOReader/statistics.sqlite3  # adjust if needed
```

3. build as usual:
```bash
make build
```

if credentials aren't provided, the build will continue without book statistics.

**note:** make sure your KOReader is configured to sync to Koofr and the database path matches your setup.

## deploy to cloudflare pages

### option 1: one command deploy

```bash
# install wrangler first (if you haven't)
npm install -g wrangler

# login to cloudflare
wrangler login

# build and deploy
make deploy
# or
go run cmd/deploy/main.go
```

## available commands

```bash
make dev      # run development server
make build    # build static site
make deploy   # build and deploy to cloudflare pages
make clean    # remove build artifacts
```

## features

- **home page** - portfolio with work history and projects
- **blog** - markdown-based blog with syntax highlighting
- **books** - reading statistics synced from koreader on koofr
- **cmd+k search** - quick navigation modal
- **rss feed** - stay updated with new posts

## stack

- go templates for templating
- no frontend framework nonsense
- static site generation
- cloudflare pages for hosting
- koreader + koofr integration for book tracking
- fast. minimal. hacker vibes. 💻

---

breaking code, building tools.

