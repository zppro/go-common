package data

import (
	"encoding/json"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type QueryRowParser interface {
	ParseRow(data []byte) (r Row, ok bool)
}

type mapRowParser struct {}
type mapRow map[string]interface{}
var defaultRowParser = mapRowParser{}

func (mr mapRow) Get(keyPath string) (interface{}, error) {
	keys := strings.Split(keyPath, ".")
	return GetMapValue(mr, keys), nil
}

func checkDrillDown (v interface{}, keys []string) (canDrillDown bool) {
	switch v.(type) {
	case map[string]interface{} :
		canDrillDown = true
	case []map[string]interface{} :
		canDrillDown = true
	default:
		canDrillDown = false
	}
	return canDrillDown && len(keys) > 0
}

func GetMapValue (mapv map[string]interface{}, keys []string) (result interface{}) {
	if len(keys) == 0 {
		result = mapv
		return
	}
	key := keys[0]

	idx := -1
	pattern := `^(\S+)\[(\d+)\]$`
	if ok, _ := regexp.MatchString(pattern, key); ok {
		re := regexp.MustCompile(pattern)
		idx, _ = strconv.Atoi(re.FindStringSubmatch(key)[2])
		key = re.FindStringSubmatch(key)[1]
	}
	result = mapv[key]
	keys = keys[1:]

	switch typeV := result.(type) {
	case map[string]interface{}:
		result = GetMapValue(typeV, keys)
	case []interface{}:
		if idx != -1 {
			v2, ok2 := typeV[idx].(map[string]interface{})
			if ok2 {
				result = GetMapValue(v2, keys)
			} else {
				result = typeV[idx]
			}
		}
	}

	return
}


func (pr mapRowParser) ParseRow(data []byte) (r Row, ok bool) {
	var row mapRow
	err := json.Unmarshal(data, &row)
	if err != nil {
		log.Println("Deserialization failed")
		ok = false
		return
	}
	r, ok = row, err == nil
	return
}

func DefaultRowParser () QueryRowParser {
	return defaultRowParser
}