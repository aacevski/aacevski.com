---
title: "markdown guide: code and images"
date: "2025-01-10"
slug: "markdown-guide"
excerpt: "a quick guide showing how code blocks and images work in this blog."
---

# markdown guide

this post shows how code blocks and images work in the blog.

## code blocks

here's a javascript example:

```javascript
const greet = (name) => {
  console.log(`hello, ${name}!`);
};

greet("world");
```

and here's some go:

```go
package main

import "fmt"

func main() {
    fmt.Println("hello from go")
}
```

hover over any code block to see the copy button!

## images

you can add images using standard markdown syntax:

```markdown
![alt text](image-url.jpg)
```

images will be automatically styled with borders and proper sizing.

### image tips

- use descriptive alt text for accessibility
- images will scale to fit the content width
- they get nice rounded corners and borders automatically
- add captions by using the `![alt](url "caption")` syntax

---

that's it! now you know how to use code blocks and images.

