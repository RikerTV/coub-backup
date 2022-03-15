coub-saver is a simple tool to automate the downloading of coubs from [coub.com](https://coub.com).

This tool will parse a user profile for all of their coubs, and assemble directory structure by year and month for each coub.

It will also generate an info file for each coub, containing the following information:

```text
Title
Post Date
Views
Likes
Reposts
Original Source
Recoubs
Tags
```

It will generate json that looks like:

```json
{
  "Title": "Coub Title",
  "Post Date": "2018-01-01",
  "Views": 123,
  "Likes": 456,
  "Reposts": 789,
  "Original Source": "https://youtube.com/watch?v=12345",
  "Recoubs": [
    {
      "Title": "Coub Title",
      "Post Date": "2018-01-01",
      "Views": 123,
      "Likes": 456,
      "Reposts": 789
    }
  ],
  "tags": [
    "tag1",
    "tag2"
  ]
}
```

Usage:

`coub-saver -d <directory> <username>`
