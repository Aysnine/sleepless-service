# sleepless-service

## Development

Prepare:

  - Redis run at `localhost:6379`
  - Golang 1.19+
  - protoc & protoc-gen-go

Start:

```bash
go run ./main.go --wx-app-id=xxx --wx-app-secret=xxx --jwt-secret=xxx --dev --port 5001
```

Compile proto file:

```bash
protoc --go_out=internal ./internal/message/message.proto
```

## Release

```bash
next_version=v0.0.1 # update

git commit -m "chore(release): $next_version" --allow-empty
git tag $next_version
git push origin --all && git push origin --tags
```

## Maintain

*How to upgrade deps?*

```bash
go get -u ./...
```
