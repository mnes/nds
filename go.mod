module github.com/qedus/nds/v2

require (
	cloud.google.com/go v0.61.0 // indirect
	cloud.google.com/go/datastore v1.2.0
	github.com/google/go-cmp v0.5.1 // indirect
	github.com/mnes/logger v0.0.0-00010101000000-000000000000
	github.com/opencensus-integrations/redigo v2.0.1+incompatible
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	go.opencensus.io v0.22.4
	google.golang.org/appengine/v2 v2.0.1
	google.golang.org/genproto v0.0.0-20200722002428-88e341933a54 // indirect
)

go 1.11

replace github.com/mnes/logger => ./logger
