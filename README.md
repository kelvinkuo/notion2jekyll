# notion2jekyll
Convert markdown file exported by notion into jekyll project dir.

## Usage

```bash
notion2jekyll [flags]

Flags:
-a, --author string       author of the post
-c, --categories string   categories split by ','
-h, --help                help for notion2jekyll
-j, --jekylldir string    the dir of created jekyll post (default "./")
-n, --notiondir string    the dir of exported notion post (default "./")
-t, --tags string         tags split by ',' and should be lower-case
-m, --time string         created time format "2006-01-02 15:04:05" (default "2023-11-03 13:12:45")

```
### Examples
```bash
notion2jekyll -m "2023-01-25 15:04:05" -c 'device,nexus7' -t 'nexus7'
```