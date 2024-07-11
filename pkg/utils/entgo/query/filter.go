package entgo

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/encoding"

	"github.com/realHoangHai/kratos/pkg/utils/stringcase"
)

type FilterOp int

const (
	FilterNot                   DatePart = iota // Not equal to
	FilterIn                                    // Check if value is in list
	FilterNotIn                                 // Not in list
	FilterGTE                                   // Greater than or equal to the passed value
	FilterGT                                    // Greater than passed value
	FilterLTE                                   // Less than or equal to the passed value
	FilterLT                                    // Less than passed value
	FilterRange                                 // Whether it is between the two given values
	FilterIsNull                                // Is it empty
	FilterNotIsNull                             // Whether it is not empty
	FilterContains                              // Whether to contain the specified substring
	FilterInsensitiveContains                   // Case-insensitive, whether to contain the specified substring
	FilterStartsWith                            // Start with value
	FilterInsensitiveStartsWith                 // Case-insensitive, starts with value
	FilterEndsWith                              // End with value
	FilterInsensitiveEndsWith                   // Case-insensitive, ends with value
	FilterExact                                 // Exact match
	FilterInsensitiveExact                      // Case-insensitive, exact match
	FilterRegex                                 // Regular expression
	FilterInsensitiveRegex                      // Case-insensitive, regular expression
	FilterSearch                                // Research all
)

var ops = [...]string{
	FilterNot:                   "not",
	FilterIn:                    "in",
	FilterNotIn:                 "not_in",
	FilterGTE:                   "gte",
	FilterGT:                    "gt",
	FilterLTE:                   "lte",
	FilterLT:                    "lt",
	FilterRange:                 "range",
	FilterIsNull:                "isnull",
	FilterNotIsNull:             "not_isnull",
	FilterContains:              "contains",
	FilterInsensitiveContains:   "icontains",
	FilterStartsWith:            "startswith",
	FilterInsensitiveStartsWith: "istartswith",
	FilterEndsWith:              "endswith",
	FilterInsensitiveEndsWith:   "iendswith",
	FilterExact:                 "exact",
	FilterInsensitiveExact:      "iexact",
	FilterRegex:                 "regex",
	FilterInsensitiveRegex:      "iregex",
	FilterSearch:                "search",
}

func hasOperations(str string) bool {
	str = strings.ToLower(str)
	for _, item := range ops {
		if str == item {
			return true
		}
	}
	return false
}

type DatePart int

const (
	DatePartDate        DatePart = iota // Date
	DatePartYear                        // Year
	DatePartISOYear                     // ISO 8601 Number of weeks in year
	DatePartQuarter                     // Quarter
	DatePartMonth                       // Month
	DatePartWeek                        // ISO 8601 Week number - The number of weeks in the year
	DatePartWeekDay                     // Day of the week
	DatePartISOWeekDay                  // Day of the week
	DatePartDay                         // Day
	DatePartTime                        // Hours: Minutes: Seconds
	DatePartHour                        // Hour
	DatePartMinute                      // Minutes
	DatePartSecond                      // Seconds
	DatePartMicrosecond                 // Microseconds
)

var dateParts = [...]string{
	DatePartDate:        "date",
	DatePartYear:        "year",
	DatePartISOYear:     "iso_year",
	DatePartQuarter:     "quarter",
	DatePartMonth:       "month",
	DatePartWeek:        "week",
	DatePartWeekDay:     "week_day",
	DatePartISOWeekDay:  "iso_week_day",
	DatePartDay:         "day",
	DatePartTime:        "time",
	DatePartHour:        "hour",
	DatePartMinute:      "minute",
	DatePartSecond:      "second",
	DatePartMicrosecond: "microsecond",
}

func hasDatePart(str string) bool {
	str = strings.ToLower(str)
	for _, item := range dateParts {
		if str == item {
			return true
		}
	}
	return false
}

// QueryCommandToWhereConditions Convert query commands to selection conditions
func QueryCommandToWhereConditions(strJson string, isOr bool) (error, func(s *sql.Selector)) {
	if len(strJson) == 0 {
		return nil, nil
	}

	codec := encoding.GetCodec("json")

	queryMap := make(map[string]string)
	var queryMapArray []map[string]string
	if err1 := codec.Unmarshal([]byte(strJson), &queryMap); err1 != nil {
		if err2 := codec.Unmarshal([]byte(strJson), &queryMapArray); err2 != nil {
			return err2, nil
		}
	}

	return nil, func(s *sql.Selector) {
		var ps []*sql.Predicate
		ps = append(ps, processQueryMap(s, queryMap)...)
		for _, v := range queryMapArray {
			ps = append(ps, processQueryMap(s, v)...)
		}

		if isOr {
			s.Where(sql.Or(ps...))
		} else {
			s.Where(sql.And(ps...))
		}
	}
}

func processQueryMap(s *sql.Selector, queryMap map[string]string) []*sql.Predicate {
	var ps []*sql.Predicate
	for k, v := range queryMap {
		key := stringcase.ToSnakeCase(k)

		keys := strings.Split(key, "__")

		if cond := oneFieldFilter(s, keys, v); cond != nil {
			ps = append(ps, cond)
		}
	}

	return ps
}

func BuildFilterSelector(andFilterJsonString, orFilterJsonString string) (error, []func(s *sql.Selector)) {
	var err error
	var queryConditions []func(s *sql.Selector)

	var andSelector func(s *sql.Selector)
	err, andSelector = QueryCommandToWhereConditions(andFilterJsonString, false)
	if err != nil {
		return err, nil
	}
	if andSelector != nil {
		queryConditions = append(queryConditions, andSelector)
	}

	var orSelector func(s *sql.Selector)
	err, orSelector = QueryCommandToWhereConditions(orFilterJsonString, true)
	if err != nil {
		return err, nil
	}
	if orSelector != nil {
		queryConditions = append(queryConditions, orSelector)
	}

	return nil, queryConditions
}

func oneFieldFilter(s *sql.Selector, keys []string, value string) *sql.Predicate {
	if len(keys) == 0 {
		return nil
	}
	if len(value) == 0 {
		return nil
	}

	field := keys[0]
	if len(field) == 0 {
		return nil
	}

	p := sql.P()

	switch len(keys) {
	case 1:
		return filterEqual(s, p, field, value)

	case 2:
		op := keys[1]
		if len(op) == 0 {
			return nil
		}

		var cond *sql.Predicate
		if hasOperations(op) {
			return processOp(s, p, op, field, value)
		} else if hasDatePart(op) {
			cond = filterDatePart(s, p, op, field).EQ("", value)
		} else {
			cond = filterJsonb(s, p, op, field).EQ("", value)
		}

		return cond

	case 3:
		op1 := keys[1]
		if len(op1) == 0 {
			return nil
		}

		op2 := keys[2]
		if len(op2) == 0 {
			return nil
		}

		// The second parameter is either the extraction date or a json field.

		//var cond *sql.Predicate
		if hasDatePart(op1) {
			str := filterDatePartField(s, op1, field)
			if hasOperations(op2) {
				return processOp(s, p, op2, str, value)
			}
			return nil
		} else {
			str := filterJsonbField(s, op1, field)

			if hasOperations(op2) {
				return processOp(s, p, op2, str, value)
			} else if hasDatePart(op2) {
				return filterDatePart(s, p, op2, str)
			}
			return nil
		}

	default:
		return nil
	}
}

func processOp(s *sql.Selector, p *sql.Predicate, op, field, value string) *sql.Predicate {
	var cond *sql.Predicate

	switch op {
	case ops[FilterNot]:
		cond = filterNot(s, p, field, value)
	case ops[FilterIn]:
		cond = filterIn(s, p, field, value)
	case ops[FilterNotIn]:
		cond = filterNotIn(s, p, field, value)
	case ops[FilterGTE]:
		cond = filterGTE(s, p, field, value)
	case ops[FilterGT]:
		cond = filterGT(s, p, field, value)
	case ops[FilterLTE]:
		cond = filterLTE(s, p, field, value)
	case ops[FilterLT]:
		cond = filterLT(s, p, field, value)
	case ops[FilterRange]:
		cond = filterRange(s, p, field, value)
	case ops[FilterIsNull]:
		cond = filterIsNull(s, p, field, value)
	case ops[FilterNotIsNull]:
		cond = filterIsNotNull(s, p, field, value)
	case ops[FilterContains]:
		cond = filterContains(s, p, field, value)
	case ops[FilterInsensitiveContains]:
		cond = filterInsensitiveContains(s, p, field, value)
	case ops[FilterStartsWith]:
		cond = filterStartsWith(s, p, field, value)
	case ops[FilterInsensitiveStartsWith]:
		cond = filterInsensitiveStartsWith(s, p, field, value)
	case ops[FilterEndsWith]:
		cond = filterEndsWith(s, p, field, value)
	case ops[FilterInsensitiveEndsWith]:
		cond = filterInsensitiveEndsWith(s, p, field, value)
	case ops[FilterExact]:
		cond = filterExact(s, p, field, value)
	case ops[FilterInsensitiveExact]:
		cond = filterInsensitiveExact(s, p, field, value)
	case ops[FilterRegex]:
		cond = filterRegex(s, p, field, value)
	case ops[FilterInsensitiveRegex]:
		cond = filterInsensitiveRegex(s, p, field, value)
	case ops[FilterSearch]:
		cond = filterSearch(s, p, field, value)
	default:
		return nil
	}

	return cond
}

// filterEqual = Equality operation
// SQL: WHERE "name" = "tom"
func filterEqual(s *sql.Selector, p *sql.Predicate, field, value string) *sql.Predicate {
	return p.EQ(s.C(field), value)
}

// filterNot NOT Inequality operation
// SQL: WHERE NOT ("name" = "tom")
// or： WHERE "name" <> "tom"
// You can use NOT to filter out NULL, but you can't use <> and !=
func filterNot(s *sql.Selector, p *sql.Predicate, field, value string) *sql.Predicate {
	return p.Not().EQ(s.C(field), value)
}

// filterIn IN operation
// SQL: WHERE name IN ("tom", "jimmy")
func filterIn(s *sql.Selector, p *sql.Predicate, field, value string) *sql.Predicate {
	var values []any
	if err := json.Unmarshal([]byte(value), &values); err == nil {
		return p.In(s.C(field), values...)
	}
	return nil
}

// filterNotIn NOT IN operation
// SQL: WHERE name NOT IN ("tom", "jimmy")`
func filterNotIn(s *sql.Selector, p *sql.Predicate, field, value string) *sql.Predicate {
	var values []any
	if err := json.Unmarshal([]byte(value), &values); err == nil {
		return p.NotIn(s.C(field), values...)
	}
	return nil
}

// filterGTE GTE (Greater Than or Equal)
// SQL: WHERE "create_time" >= "2023-10-25"
func filterGTE(s *sql.Selector, p *sql.Predicate, field, value string) *sql.Predicate {
	return p.GTE(s.C(field), value)
}

// filterGT GT (Greater than)
// SQL: WHERE "create_time" > "2023-10-25"
func filterGT(s *sql.Selector, p *sql.Predicate, field, value string) *sql.Predicate {
	return p.GT(s.C(field), value)
}

// filterLTE LTE (Less Than or Equal)
// SQL: WHERE "create_time" <= "2023-10-25"
func filterLTE(s *sql.Selector, p *sql.Predicate, field, value string) *sql.Predicate {
	return p.LTE(s.C(field), value)
}

// filterLT LT (Less than)
// SQL: WHERE "create_time" < "2023-10-25"
func filterLT(s *sql.Selector, p *sql.Predicate, field, value string) *sql.Predicate {
	return p.LT(s.C(field), value)
}

// filterRange BETWEEN operation in the range of values
// SQL: WHERE "create_time" BETWEEN "2023-10-25" AND "2024-10-25"
// 或者： WHERE "create_time" >= "2023-10-25" AND "create_time" <= "2024-10-25"
func filterRange(s *sql.Selector, _ *sql.Predicate, field, value string) *sql.Predicate {
	var values []any
	if err := json.Unmarshal([]byte(value), &values); err == nil {
		if len(values) != 2 {
			return nil
		}

		return sql.And(
			sql.GTE(s.C(field), values[0]),
			sql.LTE(s.C(field), values[1]),
		)
	}

	return nil
}

// filterIsNull IS NULL operation
// SQL: WHERE name IS NULL
func filterIsNull(s *sql.Selector, p *sql.Predicate, field, _ string) *sql.Predicate {
	return p.IsNull(s.C(field))
}

// filterIsNotNull IS NOT NULL operation
// SQL: WHERE name IS NOT NULL
func filterIsNotNull(s *sql.Selector, p *sql.Predicate, field, _ string) *sql.Predicate {
	return p.Not().IsNull(s.C(field))
}

// filterContains LIKE Fuzzy query before and after
// SQL: WHERE name LIKE '%L%';
func filterContains(s *sql.Selector, p *sql.Predicate, field, value string) *sql.Predicate {
	return p.Contains(s.C(field), value)
}

// filterInsensitiveContains ILIKE Fuzzy query before and after
// SQL: WHERE name ILIKE '%L%';
func filterInsensitiveContains(s *sql.Selector, p *sql.Predicate, field, value string) *sql.Predicate {
	return p.ContainsFold(s.C(field), value)
}

// filterStartsWith LIKE Prefix + fuzzy query
// SQL: WHERE name LIKE 'La%';
func filterStartsWith(s *sql.Selector, p *sql.Predicate, field, value string) *sql.Predicate {
	return p.HasPrefix(s.C(field), value)
}

// filterInsensitiveStartsWith ILIKE Prefix + fuzzy query
// SQL: WHERE name ILIKE 'La%';
func filterInsensitiveStartsWith(s *sql.Selector, p *sql.Predicate, field, value string) *sql.Predicate {
	return p.EqualFold(s.C(field), value+"%")
}

// filterEndsWith LIKE Suffix + fuzzy query
// SQL: WHERE name LIKE '%a';
func filterEndsWith(s *sql.Selector, p *sql.Predicate, field, value string) *sql.Predicate {
	return p.HasSuffix(s.C(field), value)
}

// filterInsensitiveEndsWith ILIKE Suffix + fuzzy query
// SQL: WHERE name ILIKE '%a';
func filterInsensitiveEndsWith(s *sql.Selector, p *sql.Predicate, field, value string) *sql.Predicate {
	return p.EqualFold(s.C(field), "%"+value)
}

// filterExact LIKE Operations Accurate comparison
// SQL: WHERE name LIKE 'a';
func filterExact(s *sql.Selector, p *sql.Predicate, field, value string) *sql.Predicate {
	return p.Like(s.C(field), value)
}

// filterInsensitiveExact ILIKE Operation case-insensitive, accurate comparison
// SQL: WHERE name ILIKE 'a';
func filterInsensitiveExact(s *sql.Selector, p *sql.Predicate, field, value string) *sql.Predicate {
	return p.EqualFold(s.C(field), value)
}

// filterRegex Regular search
// MySQL: WHERE title REGEXP BINARY '^(An?|The) +'
// Oracle: WHERE REGEXP_LIKE(title, '^(An?|The) +', 'c');
// PostgreSQL: WHERE title ~ '^(An?|The) +';
// SQLite: WHERE title REGEXP '^(An?|The) +';
func filterRegex(s *sql.Selector, p *sql.Predicate, field, value string) *sql.Predicate {
	p.Append(func(b *sql.Builder) {
		switch s.Builder.Dialect() {
		case dialect.Postgres:
			b.Ident(s.C(field)).WriteString(" ~ ")
			b.Arg(value)
			break

		case dialect.MySQL:
			b.Ident(s.C(field)).WriteString(" REGEXP BINARY ")
			b.Arg(value)
			break

		case dialect.SQLite:
			b.Ident(s.C(field)).WriteString(" REGEXP ")
			b.Arg(value)
			break

		case dialect.Gremlin:
			break
		}
	})
	return p
}

// filterInsensitiveRegex Regular search is not case sensitive
// MySQL: WHERE title REGEXP '^(an?|the) +'
// Oracle: WHERE REGEXP_LIKE(title, '^(an?|the) +', 'i');
// PostgreSQL: WHERE title ~* '^(an?|the) +';
// SQLite: WHERE title REGEXP '(?i)^(an?|the) +';
func filterInsensitiveRegex(s *sql.Selector, p *sql.Predicate, field, value string) *sql.Predicate {
	p.Append(func(b *sql.Builder) {
		switch s.Builder.Dialect() {
		case dialect.Postgres:
			b.Ident(s.C(field)).WriteString(" ~* ")
			b.Arg(strings.ToLower(value))
			break

		case dialect.MySQL:
			b.Ident(s.C(field)).WriteString(" REGEXP ")
			b.Arg(strings.ToLower(value))
			break

		case dialect.SQLite:
			b.Ident(s.C(field)).WriteString(" REGEXP ")
			if !strings.HasPrefix(value, "(?i)") {
				value = "(?i)" + value
			}
			b.Arg(strings.ToLower(value))
			break

		case dialect.Gremlin:
			break
		}
	})
	return p
}

// filterSearch research all
// SQL:
func filterSearch(s *sql.Selector, p *sql.Predicate, _, _ string) *sql.Predicate {
	p.Append(func(b *sql.Builder) {
		switch s.Builder.Dialect() {

		}
	})
	return p
}

// filterDatePart timestamp extraction date
// SQL: select extract(quarter from timestamp '2018-08-15 12:10:10');
func filterDatePart(s *sql.Selector, p *sql.Predicate, datePart, field string) *sql.Predicate {
	p.Append(func(b *sql.Builder) {
		switch s.Builder.Dialect() {
		case dialect.Postgres:
			str := fmt.Sprintf("EXTRACT('%s' FROM %s)", strings.ToUpper(datePart), s.C(field))
			b.WriteString(str)
			//b.Arg(strings.ToLower(value))
			break

		case dialect.MySQL:
			str := fmt.Sprintf("%s(%s)", strings.ToUpper(datePart), s.C(field))
			b.WriteString(str)
			//b.Arg(strings.ToLower(value))
			break
		}
	})
	return p
}

func filterDatePartField(s *sql.Selector, datePart, field string) string {
	p := sql.P()
	switch s.Builder.Dialect() {
	case dialect.Postgres:
		str := fmt.Sprintf("EXTRACT('%s' FROM %s)", strings.ToUpper(datePart), s.C(field))
		p.WriteString(str)
		break

	case dialect.MySQL:
		str := fmt.Sprintf("%s(%s)", strings.ToUpper(datePart), s.C(field))
		p.WriteString(str)
		break
	}

	return p.String()
}

// filterJsonb Extract JSONB fields
// Postgresql: WHERE ("app_profile"."preferences" -> daily_email) = 'true'
func filterJsonb(s *sql.Selector, p *sql.Predicate, jsonbField, field string) *sql.Predicate {
	p.Append(func(b *sql.Builder) {
		switch s.Builder.Dialect() {
		case dialect.Postgres:
			b.Ident(s.C(field)).WriteString(" -> ").WriteString(jsonbField)
			//b.Arg(strings.ToLower(value))
			break

		case dialect.MySQL:
			str := fmt.Sprintf("JSON_EXTRACT(%s, '$.%s')", s.C(field), jsonbField)
			b.WriteString(str)
			//b.Arg(strings.ToLower(value))
			break
		}
	})
	return p
}

func filterJsonbField(s *sql.Selector, jsonbField, field string) string {
	p := sql.P()
	switch s.Builder.Dialect() {
	case dialect.Postgres:
		p.Ident(s.C(field)).WriteString(" -> ").WriteString(jsonbField)
		//b.Arg(strings.ToLower(value))
		break

	case dialect.MySQL:
		str := fmt.Sprintf("JSON_EXTRACT(%s, '$.%s')", s.C(field), jsonbField)
		p.WriteString(str)
		//b.Arg(strings.ToLower(value))
		break
	}

	return p.String()
}
