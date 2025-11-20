# Contributing Guide

## Development Setup

### Prerequisites

- go 1.21+
- node.js 20+ (for wrangler cli)
- air (for hot reload, optional)

### Getting Started

1. clone the repository
```bash
git clone https://github.com/aacevski/andrej.sh.git
cd andrej.sh
```

2. install dependencies
```bash
go mod download

# optional: install air for hot reload
go install github.com/air-verse/air@latest
```

3. run the dev server
```bash
make dev
# or
air
```

visit: http://localhost:8080

## Environment Configuration

### Create env file

```bash
cp env.example .env
```

### GitHub Contribution Graph

Displays your GitHub contribution activity for the last year.

1. create a personal access token at https://github.com/settings/tokens/new
   - required scope: `read:user`

2. add to `.env`:
```bash
GITHUB_USERNAME=your-github-username
GITHUB_TOKEN=ghp_your_personal_access_token_here
```

If not provided, the contribution graph will be skipped.

### KOReader Book Statistics

Fetches reading statistics from your KOReader database synced to Koofr.

1. configure KOReader to sync to Koofr

2. add to `.env`:
```bash
KOOFR_EMAIL=your-email@example.com
KOOFR_PASSWORD=your-password
KOREADER_DB_PATH=/KOReader/statistics.sqlite3
```

Adjust `KOREADER_DB_PATH` if your setup differs.

If not provided, the build continues without book statistics.

## Building

### Build Static Site

```bash
make build
# or
go run cmd/build/main.go
```

Output will be in `./dist`

The build process:
1. cleans the dist directory
2. fetches GitHub contributions (if configured)
3. fetches KOReader statistics (if configured)
4. generates HTML from go templates
5. parses markdown blog posts
6. generates RSS feed
7. copies static assets

### Clean Build Artifacts

```bash
make clean
```

## Deployment

### Automated Deployment (Recommended)

The site automatically rebuilds and redeploys every 24 hours via GitHub Actions.

**Setup:**

Configure GitHub secrets in repository settings:
- `GH_PAT` - GitHub personal access token
- `GH_USERNAME` - your GitHub username
- `KOOFR_EMAIL` - Koofr email
- `KOOFR_PASSWORD` - Koofr password
- `KOREADER_DB_PATH` - path to KOReader db
- `CLOUDFLARE_API_TOKEN` - Cloudflare API token
- `CLOUDFLARE_ACCOUNT_ID` - Cloudflare account ID

The workflow runs:
- daily at midnight UTC
- manually via GitHub Actions UI

See `.github/workflows/scheduled-rebuild.yml` for details.

### Manual Deployment

**Prerequisites:**
```bash
npm install -g wrangler
wrangler login
```

**Deploy:**
```bash
make deploy
# or
go run cmd/deploy/main.go
```

This will:
1. build the static site
2. deploy to Cloudflare Pages

### Cloudflare Pages Setup

1. create Cloudflare API token at https://dash.cloudflare.com/profile/api-tokens
   - use "Edit Cloudflare Workers" template
   - or create custom token with: Account > Cloudflare Pages > Edit

2. get your account ID from Cloudflare dashboard URL

3. configure in GitHub secrets or local env

## Writing Blog Posts

### Create a New Post

1. create a new markdown file in `content/blog/`:
```bash
touch content/blog/my-new-post.md
```

2. add frontmatter and content:
```markdown
---
title: "My New Post"
date: "2024-11-09"
excerpt: "A brief description of the post"
---

## heading

your content here...
```

3. rebuild to see changes:
```bash
make build
```

The slug will be generated from the filename (e.g., `my-new-post.md` → `/blog/my-new-post`)

### markdown features

- GitHub Flavored Markdown (GFM)
- syntax highlighting with monokai theme
- auto-generated heading IDs
- RSS feed auto-generation

## Customizing

### Templates

Templates are in `static/templates/`:
- `index.html` - home page
- `blog.html` - blog listing
- `blog-post.html` - individual post
- `books.html` - books/reading stats

### Styles

CSS files in `static/css/`:
- `main.css` - base styles
- `site.css` - component styles

### Static Assets

Place files in `static/`:
- `favicon.svg`
- `og-image.png`
- `robots.txt`
- `fonts/` - custom fonts

## Available Commands

```bash
make dev      # run development server with hot reload
make build    # build static site to ./dist
make deploy   # build and deploy to Cloudflare Pages
make clean    # remove dist/ and tmp/ directories
```

## Project Structure

```
.
├── cmd/
│   ├── build/       # static site builder
│   └── deploy/      # deployment script
├── content/
│   └── blog/        # markdown blog posts
├── handlers/        # http handlers for dev server
├── pkg/
│   ├── github/      # github api integration
│   ├── koreader/    # koreader stats fetching
│   └── utils/       # utility functions
├── static/
│   ├── css/         # stylesheets
│   ├── fonts/       # web fonts
│   └── templates/   # html templates
├── dist/            # build output (generated)
├── main.go          # dev server entry point
└── Makefile         # build commands
```

## Troubleshooting

### Build fails with GitHub API errors
- check your `GITHUB_TOKEN` is valid
- ensure token has `read:user` scope
- verify `GITHUB_USERNAME` is correct

### KOReader stats not showing
- confirm Koofr credentials are correct
- verify KOReader is syncing to Koofr
- check `KOREADER_DB_PATH` matches your setup

### Deployment fails
- verify wrangler is installed: `wrangler --version`
- check you're logged in: `wrangler whoami`
- ensure Cloudflare API token has correct permissions

### Hot reload not working
- install air: `go install github.com/air-verse/air@latest`
- check `.air.toml` configuration exists
- use `make dev` or `air` command

## Tips

- use `make clean` if builds are behaving strangely
- test builds locally before deploying
- keep API tokens secure, never commit `.env`
- optimize images before adding to `static/`
- rebuild after changing templates or content

## Questions?

Open an issue or reach out via the website.
