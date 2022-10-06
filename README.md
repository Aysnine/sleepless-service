# sleepless-service

## Development

Prepare:

  - Redis run at `localhost:6379`

Start:

```bash
go run ./main.go --dev --port 5001
```

## Release

```bash
next_version=v0.0.1 # update

git commit -m "chore(release): $next_version" --allow-empty
git tag $next_version
git push origin --all && git push origin --tags
```
