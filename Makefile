check: 
	go test -v

release:
	git tag -a $(version) -m "Generated release "$(version)
	git push origin $(version)
