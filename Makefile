.PHONY: release-skip-publish
release-skip-publish: 
	goreleaser release --clean --skip-publish 

.PHONY: release-snapshot
release-snapshot: 
	goreleaser release --clean --skip-publish --snapshot

