# Ensure a command executes successfully

ensure repeatly runs a command until it succeeds.

## Example
```shell
# ensure the success of a command when having a network broken
ensure curl http://example.com

# ensure a file is created
ensure test -f file-to-create

# ensure a service being restarted when having a crash
ensure redis-server
```
