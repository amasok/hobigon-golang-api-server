package infra

import (
	"github.com/google/wire"
	"github.com/yyh-gl/hobigon-golang-api-server/app/infra/dao"
	"github.com/yyh-gl/hobigon-golang-api-server/app/infra/db"
)

// APISet : infra層のWireSet（API用）
var APISet = wire.NewSet(
	db.NewDB,
	dao.NewBlog,
	dao.NewBirthday,
	dao.NewSlack,
	dao.NewTask,
)

// CLISet : infra層のWireSet（CLI用）
var CLISet = wire.NewSet(
	db.NewDB,
	dao.NewBirthday,
	dao.NewSlack,
	dao.NewTask,
)
