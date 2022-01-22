# log-generator
A simple CLI tool that generates *near real* logs for testing

## Overview

```
Usage:
  log-generator [flags]

Flags:
      --folder string   please specify a folder path for storing logs (default "logs")
      --format string   please specify log format (text or json) (default "json")
  -h, --help            help for log-generator
      --interval int    please specify interval time (default 3)
      --max int         please specify a maximum limit for writing logs (default 30)
      --separation      please input true or false
```

## Usages

You can run the following commands for testing:

```
./log-generator --folder="logs" --format=json --separation=true --max=30 --interval=3
```

