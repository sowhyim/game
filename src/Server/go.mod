module Server

replace github.com/xormplus/xorm => gitlab.followme.com/miaolizhao/xorm v0.0.0-20170804044433-806b7631fdfb

replace proto => ../proto

go 1.12

require (
	cloud.google.com/go v0.37.2 // indirect
	github.com/armon/circbuf v0.0.0-20190214190532-5111143e8da2 // indirect
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-xorm/xorm v0.7.1
	github.com/mattn/go-colorable v0.1.1 // indirect
	github.com/micro/go-micro v1.1.0
	google.golang.org/appengine v1.5.0 // indirect
	proto v0.0.0-00010101000000-000000000000
)
