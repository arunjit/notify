-- Ensure the GOPATH is available
tup.export('GOPATH')

-- Build the protos
tup.foreach_rule(tup.glob('*.proto'), 'protoc --go_out=. %f', {'%B.pb.go'})

-- Build the binary
tup.rule(tup.glob('*.go'), 'go build -o %o %f', {'build/notify'})

-- Update gitignore
tup.creategitignore()
