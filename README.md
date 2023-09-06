# Command Robot

auto run command by webhook call.

## Quick Start

set your config.
```json
{
  "path": ".",     // target directory
  "port": "3000",  // listen webhook port
  "cmd": "git",    // command
  "args": [        // arguments
    "pull"
  ]
}
```

run server
```sh
go run .
```
