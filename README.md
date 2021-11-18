# Ensure a command to execute successfully

ensure repeatly runs a command until it success.

## Example
```shell
# ensure a command to success when has a temprary network broken
ensure curl http://example.com

# ensure a file is created
ensure test -f file-to-create

# ensure to restart when has a crash
ensure redis-server
```
