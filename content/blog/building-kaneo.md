---
title: "building kaneo: lessons from 2.4k stars"
date: "2025-01-20"
slug: "building-kaneo"
excerpt: "what i learned building an open source project management tool that people actually use."
---

# building kaneo: lessons from 2.4k stars

kaneo started as a side project to scratch my own itch. now it's a project with 2.4k+ stars on github and users actually depending on it. here's what i learned.

```yaml
services:
  postgres:
    image: postgres:16-alpine
    env_file:
      - .env
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
    restart: unless-stopped

  api:
    image: ghcr.io/usekaneo/api:latest
    ports:
      - "1337:1337"
    env_file:
      - .env
    restart: unless-stopped

  web:
    image: ghcr.io/usekaneo/web:latest
    ports:
      - "5173:5173"
    env_file:
      - .env
    restart: unless-stopped

volumes:
  postgres_data:
```

> a

## start with a real problem

i was frustrated with existing pm tools. they're either too complex (looking at you, jira) or too simple (hello, todo lists). i wanted something in between.

## keep it simple

the first version was intentionally minimal:
- create projects
- add tasks
- mark them done

that's it. no kanban boards, no gantt charts, no ai assistants. just the basics done well.

## listen to users (but not too much)

every feature request sounds reasonable in isolation. but if you implement them all, you end up with feature bloat.

i learned to ask "does this serve the core use case?" if not, it goes in the backlog (or gets closed).

## open source is marketing

making kaneo open source was the best decision. it built trust, attracted contributors, and created a community around the project.

## the tech stack

- **backend**: go for performance and simplicity
- **frontend**: typescript + react
- **database**: postgres because it just works
- **deployment**: docker for easy self-hosting

## what's next

we're working on better collaboration features, improved search, and making self-hosting even easier with drim (our deployment cli).

if you haven't checked it out yet: [github.com/usekaneo/kaneo](https://github.com/usekaneo/kaneo)

