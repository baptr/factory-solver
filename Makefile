dyson: configpb/config.pb.go
	go run github.com/baptr/factory-solver games/dyson-sphere-program.textproto "Information matrix" 30

configpb/config.pb.go: config.proto
	protoc --go_out=/home/bpu/go/src config.proto

