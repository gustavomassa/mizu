module github.com/up9inc/mizu/tests

go 1.16

require (
	github.com/gorilla/websocket v1.4.2
	github.com/up9inc/mizu/shared v0.0.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

replace github.com/up9inc/mizu/shared v0.0.0 => ../shared

replace github.com/up9inc/mizu/tap/api v0.0.0 => ../tap/api
