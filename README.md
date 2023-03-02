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


## Cosign Verify 

```bash
$ export COSIGN_EXPERIMENTAL=1
$ cosign verify-blob --certificate https://github.com/vdbulcke/goreleaser-hello-world/releases/download/v0.6.4/goreleaser-hello-world_0.6.4_Linux_x86_64.tar.gz.pem \ 
    --signature https://github.com/vdbulcke/goreleaser-hello-world/releases/download/v0.6.4/goreleaser-hello-world_0.6.4_Linux_x86_64.tar.gz.sig \
    --certificate-oidc-issuer https://token.actions.githubusercontent.com \
    --certificate-identity https://github.com/vdbulcke/goreleaser-hello-world/.github/workflows/release.yml@refs/tags/v0.6.4  \
    goreleaser-hello-world_0.6.4_Linux_x86_64.tar.gz

Verified OK
```