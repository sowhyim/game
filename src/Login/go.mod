module Login

go 1.12

replace proto => ../proto

replace Server => ../Server

require (
	Server v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.19.1
	proto v0.0.0-00010101000000-000000000000
)
