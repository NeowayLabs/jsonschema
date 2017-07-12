COVERFILE=coverage.txt
COVERFILE_HTML=cover.html

check:
	go test -coverprofile cover.out -race ./...

cover: check
	go tool cover -html=$(COVERFILE) -o=$(COVERFILE_HTML)
	xdg-open $(COVERFILE_HTML)

analyze:
	go vet .
	staticcheck .
	gosimple .
	unused .

devdeps:
	go get -u honnef.co/go/tools/...
