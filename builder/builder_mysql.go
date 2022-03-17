package builder

type BuilderMysql struct {
	BuilderStandard
}

// NewMysqlBuilder 构造新Builder
func NewMysqlBuilder() Builder {
	return BuilderMysql{}
}
