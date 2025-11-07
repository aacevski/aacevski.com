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

### option 2: via cloudflare dashboard

1. push your code to github
2. go to cloudflare dashboard > pages
3. connect your github repo
4. set build command: `go run cmd/build/main.go`
5. set build output directory: `dist`
6. deploy!

cloudflare pages will automatically rebuild on every push to main.

## available commands

```bash
make dev      # run development server
make build    # build static site
make deploy   # build and deploy to cloudflare pages
make clean    # remove build artifacts
```

## stack

- go templates for templating
- no frontend framework nonsense
- static site generation
- cloudflare pages for hosting
- fast. minimal. hacker vibes. 💻

---

breaking code, building tools.

