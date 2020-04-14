module google.golang.org/grpc

go 1.11

require (
	github.com/cncf/udpa/go v0.0.0-20191209042840-269d4d468f6f
	github.com/envoyproxy/go-control-plane v0.9.4
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/mock v1.1.1
	github.com/golang/protobuf v1.3.3
	github.com/google/go-cmp v0.2.0
	github.com/panjf2000/ants/v2 v2.3.1 // indirect
	golang.org/x/net v0.0.0-20190311183353-d8887717615a
	golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be
	golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd
	google.golang.org/genproto v0.0.0-20190819201941-24fa4b261c55
)

replace golang.org/x/net => github.com/hortor-games/net v0.0.0-20200414122753-12f6f5e7d50d
//replace golang.org/x/net => ../net
