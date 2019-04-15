PROTOC = protoc
GO_PROTO_BUILD = ${PROTOC} -I../../../ -I. --go_out=plugins=grpc:.

GO:
	${GO_PROTO_BUILD} ./pbLogin/login.proto
	${GO_PROTO_BUILD} ./pbServer/server.proto
