.PHONY: release-skip-publish
release-skip-publish: 
	goreleaser release --rm-dist --skip-publish 

.PHONY: release-snapshot
release-snapshot: 
	goreleaser release --rm-dist --skip-publish --snapshot

