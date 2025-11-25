---
title: "What They Don't Tell You About Maintaining an Open Source Project"
date: "2025-11-25"
excerpt: "i built kaneo.app - an open source, self-hosted kanban board. turns out shipping code is the easy part. here's what maintaining it actually looks like."
---

## the beginning

building [kaneo](https://kaneo.app) was fun. a clean, minimal kanban board. self-hosted. open source. no tracking, no subscriptions, no bullshit.

i shipped v1, posted it on reddit, got some stars on github. people actually used it. that feeling when someone tells you they're using something you built? incredible.

then i learned that **shipping is just the beginning.**

## the documentation challenge

i spent hours writing documentation. setup guides, configuration examples, troubleshooting sections. tried to make it clear and comprehensive.

but here's the thing: **people come from different backgrounds.** what's obvious to me after building the thing isn't obvious to someone installing it for the first time.

someone opens an issue: "how do i install this?"

my first reaction was frustration. it's in the readme! but then i realized - maybe the readme assumes too much. maybe they're new to docker. maybe they're coming from windows and linux is foreign.

so i improved the docs:
- added more examples
- created a troubleshooting guide
- made a video walkthrough
- added a "common issues" section

it's a constant process. documentation is never "done."

## support is product development

maintaining kaneo means helping people debug their setups. and honestly? it's taught me more than i expected.

people run kaneo on setups i never imagined:
- behind corporate proxies
- on raspberry pi clusters
- in kubernetes with custom networking
- on nas devices with limited resources

each support request reveals an assumption i made. each "it doesn't work" issue (even the ones without details) points to a failure mode i didn't consider.

the challenge is balancing time. i want to help everyone. but i also have a day job. and new features to build. and bugs to fix.

i'm still learning how to set boundaries while being helpful.

## feature requests are humbling

people want kaneo to do more. and that's amazing! it means they're actually using it. they care enough to imagine what it could be.

but **every feature request is a decision:**
- does this fit the vision?
- can i maintain this long-term?
- will this complicate the codebase?
- what else won't get built if i build this?

saying no is hard. especially when the request is thoughtful and well-reasoned. especially when someone offers to help implement it.

i've learned to be transparent: "i love this idea, but it's outside kaneo's scope. here's why..."

most people understand. some don't. that's okay.

## migrations are terrifying

the database schema needed a refactor. the current design was limiting. the new design would enable features people wanted.

but 200+ people were using kaneo in production. **their actual work data. their team's workflows.**

if i broke the migration, they'd lose trust. maybe lose data. definitely lose sleep.

so i:
1. wrote the migration script
2. tested it on every version going back to v1
3. wrote detailed upgrade notes
4. tested edge cases
5. tested the edge cases of edge cases
6. added validation checks
7. added dry-run mode

released it. held my breath.

most migrations went smoothly. a few didn't. not because people didn't read the notes - but because they had setups i couldn't have predicted:
- modified databases
- custom patches
- environments i'd never seen

we debugged together. they were patient. i was grateful.

every migration taught me something new about defensive programming.




## contributors are a gift

when someone submits a pr, it's incredible. **someone cared enough to spend their time improving kaneo.**

but integrating contributions is harder than i expected:
- different coding styles
- different assumptions about architecture
- different ideas about what kaneo should be

sometimes a pr is perfect. sometimes it needs work. sometimes it's solving a problem in a way that'll create more problems later.

i've learned to:
- appreciate the effort, always
- explain my reasoning when requesting changes
- be okay with saying "this doesn't fit, but i appreciate you"
- sometimes just fix it myself if it's close

the contributors who stick around? they're amazing. they've made kaneo better than i could alone.


## the diversity of environments

self-hosting means people run kaneo **everywhere:**

```bash
# docker on their laptops
docker compose up -d

# kubernetes clusters at work
kubectl apply -f kaneo.yaml

# raspberry pi in their home lab
# (with 1GB of RAM and dreams)

# bare metal on old servers
# (that have been running since 2015)

# nas devices with arm processors
# (that i've never even heard of)
```

each environment teaches me something. each "it doesn't work on my setup" issue reveals an assumption i made about how systems work.

i can't test every environment. but i can make kaneo more resilient:
- better error messages
- clearer logs
- more graceful failures

the people running kaneo on weird setups? they're often the most helpful. they understand their environment. they provide detailed logs. they test fixes.

we figure it out together.

## keeping documentation alive

documentation is never finished. every feature needs docs. every bug fix might need docs. every question reveals a gap in docs.

i've learned to:
- update docs in the same pr as code changes
- treat "docs are wrong" issues as high priority
- appreciate when people submit doc fixes
- accept that docs will never be perfect

the goal isn't perfect documentation. it's documentation that helps most people most of the time.

and when it doesn't? that's feedback. that's how it gets better.

## the comparison question

> "why not just use trello/notion/linear?"

it's a fair question. those tools are great. they have teams of engineers, designers, product managers. they're polished. they're fast. they're feature-rich.

kaneo is different:

| them | kaneo |
|------|-------|
| cloud-hosted | self-hosted (your data, your server) |
| closed source | open source (you can read every line) |
| feature-rich | minimal (does one thing well) |
| subscription | free (as in freedom and beer) |

it's not better. it's **different.** for some people, that difference matters.

and honestly? building kaneo taught me more than using those tools ever could.

## the emotional reality

maintaining open source is a rollercoaster:

someone stars your repo → feels good

someone opens a detailed bug report with logs and reproduction steps → feels great

someone says "kaneo saved our team" → feels incredible

someone opens an issue titled "this is trash" → hurts more than it should

you spend a weekend implementing a requested feature → crickets

you fix a small bug → three people thank you

you realize you haven't worked on your own roadmap in months → exhausting

someone submits a thoughtful pr → you're not alone

the highs are high. the lows are low. but the people who use kaneo, who contribute, who care? they make it worth it.

## what i learned

### 1. scope is everything

kaneo does **one thing:** kanban boards. not project management. not time tracking. not team chat.

every feature you add is a feature you maintain forever.

being clear about scope isn't limiting - it's liberating. it lets you focus. it lets you say no without guilt.

### 2. automate everything you can

```yaml
# .github/workflows/ci.yml
name: CI
on: [push, pull_request]
jobs:
  test:
    - run: npm test
  security:
    - run: npm audit
  release:
    - run: semantic-release
```

automation isn't lazy. it's sustainable:
- automated tests catch bugs before users do
- automated releases mean less manual work
- automated security scans give peace of mind
- automated dependency updates keep things current

it frees you to focus on what matters.

### 3. good issue templates help everyone

github issue templates help people provide:
- system info
- error logs
- steps to reproduce

it's not about gatekeeping. it's about making debugging possible. most people want to help you help them. templates make that easier.

### 4. saying no is an act of respect

you can't build everything. saying yes to everything means doing nothing well.

being honest about what you can and can't do respects everyone's time. including yours.

### 5. users are collaborators

the people using kaneo aren't just users. they're:
- beta testers finding bugs
- documentation editors spotting gaps
- feature designers sharing ideas
- community builders helping each other

they're not demanding. they're engaged. **that's a gift.**

when someone opens an issue, they're investing time in making kaneo better. even if the issue is unclear, the intent is good.

patience and kindness aren't just nice. they're necessary.

## the honest truth

maintaining an open source, self-hosted project is:
- more work than building it
- different fun than building it
- more rewarding than you'd expect
- harder than you'd expect
- worth it

you learn:
- technical skills (migrations, security, scalability)
- people skills (communication, patience, boundaries)
- product skills (prioritization, scope, vision)
- how to appreciate every contribution
- how to build something people actually want

## my setup (the real one)

```
┌─────────────────────────────────────┐
│  kaneo infrastructure               │
├─────────────────────────────────────┤
│  github                             │
│  ├─ code + issues                   │
│  ├─ actions (ci/cd)                 │
│  └─ container registry              │
│                                     │
│  hetzner ($7/month)                 │
│  └─ cloud instance                  │
│                                     │
│  cloudflare (free)                  │
│  └─ dns + ddos protection           │
│                                     │
│  plausible                          │
│  └─ privacy-friendly analytics      │
│                                     │
│  coffee (priceless)                 │
│  └─ way too much                    │
└─────────────────────────────────────┘
```

## what i'd tell past me

**1. invest in documentation early** - good docs reduce support burden and help people succeed. it's time well spent.

**2. automate from day one** - tests, releases, security scans. automation scales. you don't.

**3. be clear about scope** - say what your project is AND what it isn't. it helps everyone.

**4. migrations are worth the extra effort** - test thoroughly. add rollbacks. write clear upgrade notes. your users trust you with their data.

**5. it's okay to be slow** - you're not a company. you're a person. set expectations. take breaks. protect your energy.

**6. celebrate your users** - every person using kaneo is amazing. they chose to trust something you built. that's incredible.

**7. the community is the product** - the code matters, but the people matter more. invest in both.

## the conclusion

would i do it again? 

**absolutely.**

kaneo exists because i wanted a simple kanban board that i controlled. but it became something more: a community of people who value privacy, simplicity, and owning their tools.

the maintenance is real work. the migrations are stressful. the support takes time.

but people are using kaneo to:
- run their businesses
- manage their side projects
- organize their teams
- learn about self-hosting

they send thank you messages. they submit thoughtful bug reports. they contribute code. they help each other in discussions.

that's not just cool. that's why i do this.

---

*kaneo is open source and free forever. check it out: [github.com/usekaneo/kaneo](https://github.com/usekaneo/kaneo)*

*if you're using it, thank you. if you're contributing, you're amazing. if you're thinking about it, the docs are pretty good.*

*and if you find a bug? i'll fix it. probably at 11pm. but i'll fix it.*

