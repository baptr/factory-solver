proto:
	protoc --go_out=/home/bpu/go/src config.proto

dyson:
	go run github.com/baptr/factory-solver games/dyson-sphere-program.textproto "Information matrix" 30
