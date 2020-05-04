module github.com/ohyo/revelmodules

go 1.13

replace github.com/swaggo/swag => ../swag

require (
	git.feneas.org/ganggo/gorm v1.9.3
	github.com/99designs/httpsignatures-go v0.0.0-20170731043157-88528bf4ca7e
	github.com/CloudyKit/jet/v3 v3.0.0 // indirect
	github.com/Joker/jade v1.0.0 // indirect
	github.com/KyleBanks/depth v1.2.1
	github.com/PuerkitoBio/goquery v1.5.1
	github.com/Shopify/goreferrer v0.0.0-20181106222321-ec9c9a553398 // indirect
	github.com/Zauberstuhl/go-xml v0.0.0-20180122175953-534ffa216723
	github.com/aymerick/raymond v2.0.2+incompatible // indirect
	github.com/certifi/gocertifi v0.0.0-20200211180108-c7c1fbc02894 // indirect
	github.com/cockroachdb/cockroach-go v0.0.0-20200411195601-6f5842749cfc // indirect
	github.com/dchest/captcha v0.0.0-20170622155422-6a29415a8364
	github.com/dgraph-io/badger/v2 v2.0.3
	github.com/dgraph-io/dgo/v2 v2.2.0 // indirect
	github.com/eknkc/amber v0.0.0-20171010120322-cdade1c07385 // indirect
	github.com/fatih/structs v1.1.0
	github.com/flosch/pongo2 v0.0.0-20190707114632-bbf5a6c351f4 // indirect
	github.com/getlantern/byteexec v0.0.0-20170405023437-4cfb26ec74f4
	github.com/getlantern/filepersist v0.0.0-20160317154340-c5f0cd24e799 // indirect
	github.com/getlantern/golog v0.0.0-20190830074920-4ef2e798c2d7 // indirect
	github.com/getsentry/raven-go v0.2.0
	github.com/ghodss/yaml v1.0.0
	github.com/gin-gonic/gin v1.6.2
	github.com/go-fed/activity v0.4.0 // indirect
	github.com/go-gomail/gomail v0.0.0-20160411212932-81ebce5c23df
	github.com/go-openapi/jsonreference v0.19.3
	github.com/go-openapi/spec v0.19.7
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gobuffalo/envy v1.9.0 // indirect
	github.com/gobuffalo/fizz v1.9.8 // indirect
	github.com/gobuffalo/flect v0.2.1 // indirect
	github.com/gobuffalo/genny v0.6.0 // indirect
	github.com/gobuffalo/nulls v0.4.0 // indirect
	github.com/gobuffalo/packd v1.0.0 // indirect
	github.com/gobuffalo/pop v4.13.1+incompatible
	github.com/gobuffalo/uuid v2.0.5+incompatible
	github.com/gobuffalo/validate v2.0.4+incompatible
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.3.3
	github.com/gomodule/redigo v2.0.0+incompatible // indirect
	github.com/gorilla/schema v1.1.0 // indirect
	github.com/inconshreveable/log15 v0.0.0-20200109203555-b30bc20e4fd1 // indirect
	github.com/iris-contrib/blackfriday v2.0.0+incompatible // indirect
	github.com/iris-contrib/formBinder v5.0.0+incompatible // indirect
	github.com/iris-contrib/go.uuid v2.0.0+incompatible // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/jpillora/overseer v1.1.4
	github.com/kataras/golog v0.0.10 // indirect
	github.com/kataras/iris v11.1.1+incompatible
	github.com/klauspost/compress v1.10.5 // indirect
	github.com/lib/pq v1.3.0
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/microcosm-cc/bluemonday v1.0.2
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pmezard/go-difflib v1.0.0
	github.com/revel/config v0.21.0
	github.com/revel/cron v0.21.0 // indirect
	github.com/revel/log15 v2.11.20+incompatible
	github.com/revel/modules v0.21.0
	github.com/revel/pathtree v0.0.0-20140121041023-41257a1839e9 // indirect
	github.com/revel/revel v0.21.0
	github.com/ryanuber/columnize v2.1.0+incompatible // indirect
	github.com/sec51/convert v0.0.0-20190309075348-ebe586d87951 // indirect
	github.com/sec51/cryptoengine v0.0.0-20180911112225-2306d105a49e // indirect
	github.com/sec51/gf256 v0.0.0-20160126143050-2454accbeb9e // indirect
	github.com/sec51/qrcode v0.0.0-20160126144534-b7779abbcaf1 // indirect
	github.com/sec51/twofactor v1.0.0
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/shamaton/msgpack v1.1.1
	github.com/shaoshing/train v0.0.0-20150517185910-c76b6fe70b70
	github.com/shopspring/decimal v0.0.0-20200419222939-1884f454f8ea // indirect
	github.com/shurcooL/github_flavored_markdown v0.0.0-20181002035957-2122de532470
	github.com/shurcooL/highlight_diff v0.0.0-20181222201841-111da2e7d480 // indirect
	github.com/shurcooL/highlight_go v0.0.0-20191220051317-782971ddf21b // indirect
	github.com/shurcooL/octicon v0.0.0-20191102190552-cbb32d6a785c // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/stretchr/testify v1.5.1
	github.com/twinj/uuid v1.0.0 // indirect
	github.com/xeonx/timeago v1.0.0-rc4 // indirect
	github.com/yvasiyarov/go-metrics v0.0.0-20150112132944-c25f46c4b940
	github.com/yvasiyarov/gorelic v0.0.7
	github.com/yvasiyarov/newrelic_platform_go v0.0.0-20160601141957-9c099fbc30e9 // indirect
	golang.org/x/crypto v0.0.0-20191205180655-e7c4368fe9dd
	golang.org/x/tools v0.0.0-20200420001825-978e26b7c37c
	google.golang.org/grpc v1.28.1 // indirect
	gopkg.in/AlecAivazis/survey.v1 v1.8.8
	gopkg.in/fsnotify/fsnotify.v1 v1.4.7 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/src-d/go-parse-utils.v1 v1.1.2 // indirect
	gopkg.in/stack.v0 v0.0.0-20141108040640-9b43fcefddd0 // indirect
	gopkg.in/urfave/cli.v1 v1.20.0 // indirect
	gopkg.in/yaml.v2 v2.2.8
)
