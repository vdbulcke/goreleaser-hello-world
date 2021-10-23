# Testing goreleaser



## Dev Build 

```bash
goreleaser build  --snapshot --rm-dist
```

## Tagged Release

```
git tag -a -s v0.1
git push --tags
```

```
goreleaser build  --rm-dist
```

