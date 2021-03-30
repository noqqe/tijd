## tijd

is dutch for time. I wrote this as a small go exercise to deal with the
`time` module

```
Your time: 	15:09:54 +0200 (Local)
UTC time: 	13:09:54 +0000 (UTC)

London		14:09:54 +0100 (Europe/London)
Punin		18:39:54 +0530 (Asia/Kolkata)
Sofia		16:09:54 +0300 (Europe/Sofia)
```

## Install

    brew install noqqe/tap/tijd

## Config

In `~/.tijd.json` place something like this:

```json
{
  "Locations": {
    "London": "Europe/London",
    "Punin": "Asia/Kolkata",
    "Sofia": "Europe/Sofia"
  }
}
```
