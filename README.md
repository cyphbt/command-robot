# Command Robot

Auto run command by webhook call.

## Quick Start

close repository
```sh
git clone git@github.com:cypj/command-robot.git
```

set your config.
```json
{
  "path": "/root/abc/",    // project directory
  "port": "3000",          // listen webhook port
  "cmd": "git",            // command
  "args": [                // arguments
    "pull",
    "origin"
  ],
  "event": "push",           // trigger event, refer https://docs.github.com/en/webhooks/webhook-events-and-payloads#delivery-headers
  "secret": "your webhook secret"
}
```

run server
```sh
go run .
```

If you don't have a go runtime, you can download `command-robot` binary and `config.json` to your folder, then `./command-robot` run it.

Otherwise, you can use `daemon service` / `tmux` / `nohup` to run it background.

