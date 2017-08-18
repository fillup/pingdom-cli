builddist:
	gox -output="dist/{{.OS}}/{{.Arch}}/pingdom"

test:
	go test -cover
