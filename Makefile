COVERFILE=coverage.txt
COVERFILE_HTML=cover.html

check:
	go test -coverprofile $(COVERFILE) -race ./...

cover: check
	go tool cover -html=$(COVERFILE) -o=$(COVERFILE_HTML)
	xdg-open $(COVERFILE_HTML)

analyze:
	go vet .
	staticcheck .
	gosimple .
	unused .

docs:
	mdtoc -w ./docs/spec.md

depsdev:
	go get -u honnef.co/go/tools/...
	go get -u github.com/katcipis/mdtoc/cmd/mdtoc

.PHONY:docs depsdev analyze cover check
