# using make test to explicitly call to test with '-v' option. fzf behavior not testable otherwise
test:
	go test . -v -count=1
