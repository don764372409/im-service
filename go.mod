module IM-Service

go 1.21.3

require (
	github.com/antonfisher/nested-logrus-formatter v1.3.1
	github.com/aws/aws-sdk-go v1.47.4
	github.com/go-netty/go-netty v1.6.4
	github.com/go-netty/go-netty-transport v1.7.6
	github.com/google/uuid v1.5.0
	github.com/h2non/filetype v1.1.3
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5
	github.com/sirupsen/logrus v1.9.3
	google.golang.org/protobuf v1.32.0
	gorm.io/driver/sqlite v1.5.4
	gorm.io/gorm v1.25.5
	im-sdk v0.0.0-00010101000000-000000000000
)

replace im-sdk => github.com/don764372409/im-sdk v0.4.4

require (
	github.com/go-audio/audio v1.0.0 // indirect
	github.com/go-audio/riff v1.0.0 // indirect
	github.com/go-audio/wav v1.1.0 // indirect
	github.com/gobwas/httphead v0.1.0 // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/gobwas/ws v1.3.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/jonboulle/clockwork v0.4.0 // indirect
	github.com/lestrrat-go/strftime v1.0.6 // indirect
	github.com/mattn/go-sqlite3 v1.14.17 // indirect
	golang.org/x/mobile v0.0.0-20240112133503-c713f31d574b // indirect
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/tools v0.17.0 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
