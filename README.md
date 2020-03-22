# Crast

A command line, directory based, flexible todo list app.

## Installation

### Homebrew

```shell script
brew tap safe-k/tap
brew install crast
```

## Usage

```shell script
> crast --help
Usage:
  crast [command]

Available Commands:
  add         Adds task(s)
  clear       Removes all tasks from the list
  delete      Deletes the list
  do          Marks task(s) as done
  edit        Edits existing task(s)
  help        Help about any command
  init        Initialises a list
  move        Moves task(s) to a different list
  rm          Removes task(s)
  status      Prints the current list status
  undo        Marks task(s) as todo

Flags:
  -h, --help   help for crast
```

**Example:**
``` shell script
> pwd
/Users/seifkamal/src/project/internal/code
> crast status -a
---------------------------------------------------------
|     ID     |  ✔  | P |   TOPIC    |      SUMMARY      |
---------------------------------------------------------
| x91jHirZg  | [ ] | 2 | general    | read article      |
| aP4GDirWR  | [ ] | 3 | general    | install cool tool |
| tfzgvirWg  | [✔] | 4 | general    | update docs       |
| tBzRvi9Wgz | [ ] | 4 | feature    | setup cron        |
---------------------------------------------------------
Dir: /Users/seifkamal/src/project
```
