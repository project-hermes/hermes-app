setup:
	go get github.com/99designs/gqlgen
	go get github.com/onsi/ginkgo/ginkgo
	go get github.com/onsi/gomega/...

graphql-schema:
	go run server/scripts/gqlgen.go