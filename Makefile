proto:
	@protoc -I $(HOME)/go/src \
		-I api/vector3 \
		--go_out=api/vector3 \
		api/vector3/vector3.proto
	@protoc -I $(HOME)/go/src \
		-I api/quaternion \
		--go_out=api/quaternion \
		api/quaternion/quaternion.proto
	@protoc -I $(HOME)/go/src \
		-I api/volume \
		--go_out=api/volume \
		api/volume/volume.proto
	@echo 'Protobuf built'

bench:
	go test ./... -bench=. -benchmem -benchtime 1000000x

test:
	go test ./...
	@echo 'Test passed'
