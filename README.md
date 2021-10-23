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

**NOTE:** Cannot build on "dirty" git working dir


## Publishing 

* Dry run
```
goreleaser release --skip-publish --rm-dist

```

* 

**NOTE:** Need to set the Access token as env var
```
goreleaser release  --rm-dist
```