module chenblog

go 1.16

replace (
	chenblog/models v0.0.1 => ../chenblog/models
	chenblog/routers v0.0.1 => ../chenblog/routers
)

require (
	github.com/astaxie/beego v1.12.3
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/smartystreets/goconvey v1.6.4
	github.com/unknwon/com v1.0.1
)
