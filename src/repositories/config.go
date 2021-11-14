package repositories

const (
	LIMIT  = "limit_"
	OFFSET = "offSet_"
)

func addWhere(where *string, sSql string) {
	if *where == "" {
		*where = " WHERE " + sSql
	} else {
		*where += " AND " + sSql
	}
}

func bindParamLikeFull(param string) string {
	return "%" + param + "%"
}
