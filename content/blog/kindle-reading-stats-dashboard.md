---
title: "How I Built a Kindle Reading Stats Dashboard That Actually Works"
date: "2025-11-08"
excerpt: "forget goodreads. i turned my kindle into a reading stats machine that syncs to my portfolio website automatically. here's how."
---

## the problem

i love reading, but tracking what i read was a mess. kindle's built-in stats are basic, goodreads feels bloated, and i wanted something that just... worked. something minimal. something i controlled.

## enter koreader

[koreader](https://koreader.rocks/) is an open-source ebook reader that runs on kindle (and other e-readers). it's fast, customizable, and most importantly - it tracks everything in a sqlite database.

the killer feature? it can sync that database to cloud storage. including koofr.

## why koofr?

koofr offers webdav support out of the box. no complicated apis, no rate limits for personal use. just mount it like a network drive and you're done.

plus, they respect privacy and offer generous free storage.

## the setup

### 1. install koreader on kindle

- download koreader from their [github releases](https://github.com/koreader/koreader/releases)
- follow the installation instructions for your device
- launch koreader

### 2. configure koofr sync

in koreader:
- go to settings → network → cloud storage
- add koofr webdav
- enter credentials:
  - server: `https://app.koofr.net/dav/Koofr`
  - username: your koofr email
  - password: app-specific password from koofr settings

enable automatic sync and you're done. every time you read, koreader syncs `statistics.sqlite3` to koofr.

### 3. build the integration

now the fun part. i wanted to display my reading stats on my portfolio website.

the stack:
- **webdav client** to fetch the sqlite database from koofr
- **sqlite** parsing to extract book data
- **web framework** of your choice for rendering

here's the basic flow:

```typescript
// example using typescript
async function fetchKOReaderStats() {
    // 1. connect to koofr via webdav
    const client = createWebDAVClient(webdavURL, {
        username: email,
        password: password
    });
    
    // 2. download statistics.sqlite3
    const data = await client.getFileContents(dbPath);
    
    // 3. parse the database
    const db = await openSQLite(data);
    
    // 4. extract books, pages, reading time
    return parseKOReaderDatabase(db);
}
```

the database schema:
- `book` table: titles, authors, pages, last_open
- `page_stat_data` table: reading sessions with timestamps and duration

### 4. parsing the database

here's how i extract the data:

```typescript
async function parseKOReaderDatabase(db: SQLiteDB) {
    const books = await db.all(`
        SELECT title, authors, pages, last_open
        FROM book
        ORDER BY last_open DESC
    `);
    
    for (const book of books) {
        const bookId = await db.get(
            'SELECT id FROM book WHERE title = ?',
            [book.title]
        );
        
        // get latest page from most recent session
        const currentPage = await db.get(`
            SELECT page FROM page_stat_data 
            WHERE id_book = ? 
            ORDER BY start_time DESC LIMIT 1
        `, [bookId.id]);
        
        // sum total reading time
        const timeData = await db.get(`
            SELECT SUM(duration) as total
            FROM page_stat_data
            WHERE id_book = ?
        `, [bookId.id]);
        
        // calculate progress and format
        const progress = (currentPage?.page / book.pages) * 100;
        const readingTime = Math.floor(timeData.total / 60);
    }
}
```

### 5. accurate progress tracking

for accurate progress tracking, i fetch the latest page from the most recent reading session:

```sql
SELECT page 
FROM page_stat_data 
WHERE id_book = ? 
ORDER BY start_time DESC 
LIMIT 1
```

not just `MAX(page)` - because if you flip back to check something, that shouldn't be your "current" page.

## the result

now my [books page](/books) shows:
- **currently reading**: with real-time progress bars
- **finished books**: sorted by completion date
- **reading time**: actual minutes spent, not estimates
- **pages read**: accurate counts from database

all updated automatically whenever i sync my kindle.

## environment variables

set these in `.env` for local dev:

```bash
KOOFR_WEBDAV_URL=https://app.koofr.net/dav/Koofr
KOOFR_EMAIL=your-email@example.com
KOOFR_PASSWORD=your-app-password
KOREADER_DB_PATH=/KOReader/statistics.sqlite3
```

for production, add them as environment variables in your hosting dashboard.

## my setup

- amazon kindle (16 gb) – black
- koreader 2024.11
- koofr free tier
- auto-sync on close

total cost? $0/month.

## what's next?

planning to add:
- reading streaks (consecutive days)
- monthly reading goals
- reading speed analytics
- year-in-review stats

but for now? it works. it's fast. it's mine.

## source code

this website is fully open source. check out the complete implementation:

[github.com/aacevski/aacevski.com](https://github.com/aacevski/aacevski.com)

the koreader integration lives in `pkg/koreader/` - feel free to use it, fork it, make it better.

