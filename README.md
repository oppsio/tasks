# tasks

## Monitoring

- `docker exec -it f3911deb2996 sh` Start a shell in the running Redis container
- `/ $ redis-cli` enter the redis-cli
- type `MONITOR` to monitor queue activity

## Useful

Find project external dependencies:   
`go list -f '{{.Deps}}' | tr "[" " " | tr "]" " " | xargs go list -f '{{if not .Standard}}{{.ImportPath}}{{end}}'`
