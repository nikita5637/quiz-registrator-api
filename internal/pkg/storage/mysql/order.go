package mysql

import "strings"

func getOrderStmt(sort string) string {
	var order string
	if sort != "" {
		var orders []string
		for _, expr := range strings.Split(sort, ",") {
			if strings.HasPrefix(expr, "-") {
				orders = append(orders, expr[1:]+" DESC ")
			} else {
				orders = append(orders, expr+" ASC ")
			}
		}
		order = "ORDER BY " + strings.Join(orders, ", ")
	}
	return order
}
