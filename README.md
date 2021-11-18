# Ensure a command executed successfully

ensure repeatly runs a command until it succeeds.

## Example
```shell
# ensure the success of a command when having a network broken
ensure curl http://example.com

# ensure a file is created
ensure test -f file-to-create

# ensure a service restarted when there is a crash
ensure redis-server
```
