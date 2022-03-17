package builder

type BuilderSqlite struct {
	BuilderStandard
}

// NewSqliteBuilder 构造新Builder
func NewSqliteBuilder() Builder {
	return BuilderSqlite{}
}
