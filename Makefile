setup:
	go get golang.org/x/sys
	go get golang.org/x/oauth
	go get golang.org/x/text
	go get firebase.google.com/go
	go get github.com/gin-gonic/gin
	go get github.com/99designs/gqlgen
	go get github.com/onsi/ginkgo/ginkgo
	go get github.com/onsi/gomega/...
	go get github.com/golang/mock/gomock
	go install github.com/golang/mock/mockgen

mock-objects:
	mockgen -destination server/wrapper/mock/firestore.go  github.com/project-hermes/hermes-app/server/wrapper \
	DBClientInterface,CollectionRefInterface,DocRefInterface,DocumentInteratorInterface

graphql-schema:
	go run server/scripts/gqlgen.go