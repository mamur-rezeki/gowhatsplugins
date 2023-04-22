package helper

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// Fungsi untuk mengembalikan objek sebagai string json
func JsonMe(me interface{}) (string, error) {
	if b, err := ByteMe(me); err != nil {
		return "", err
	} else {
		return string(b), nil
	}
}

func ByteMe(me interface{}) ([]byte, error) {
	if json_me, err := json.MarshalIndent(me, "", "  "); err != nil {
		return nil, err
	} else {
		return json_me, nil
	}
}

func MapMe(me interface{}) (map[string]interface{}, error) {
	if b, err := ByteMe(me); err != nil {
		return nil, err
	} else {
		var res map[string]interface{}
		if err := json.Unmarshal(b, &res); err != nil {
			return nil, err
		} else {
			return res, nil
		}
	}
}

func GetType(i interface{}) string {
	return reflect.TypeOf(i).Elem().Name()
}

func GetKeys(amap map[string]interface{}) []string {
	var result []string
	for key := range amap {
		result = append(result, key)
	}

	return result
}

// Mengcek value ada dalam array
func ContainA(element any, array []any) (result bool) {
	result = false
	for _, v := range array {
		if v == element {
			result = true
			break
		}
	}

	return
}

func ContainS(element string, array []string) (result bool) {
	result = false
	for _, v := range array {
		if v == element {
			result = true
			break
		}
	}

	return
}
func ContainSInsensitives(element string, array []string) (result bool) {
	result = false
	for _, v := range array {
		if strings.EqualFold(v, element) {
			result = true
			break
		}
	}

	return
}

func Int64String(integer int64) string {
	return strconv.Itoa(int(integer))
}

func Int32String(integer int32) string {
	return strconv.Itoa(int(integer))
}

func IntString(integer int) string {
	return strconv.Itoa(int(integer))
}

func AnyString(i interface{}) string {
	return fmt.Sprintf(`%s`, i)
}

func Nothing() bool {
	return true
}

func SliceSAny(slices []string) []any {
	var newany []any
	for _, an := range slices {
		newany = append(newany, an)
	}

	return newany
}

func GetOnlyNumber(text string) []string {

	var result []string

	if rgx, err := regexp.Compile(`(\+\d{1,3}|0)?\s?[\d\s-]+`); err == nil {

		if rgx.MatchString(text) {

			for _, numb := range rgx.FindAllString(text, -1) {
				numb = strings.ReplaceAll(numb, "+", "")
				numb = strings.ReplaceAll(numb, "-", "")
				numb = strings.ReplaceAll(numb, " ", "")

				// if strings.HasPrefix(numb, "08") {
				// 	numb = "62" + numb[1:]
				// }

				result = append(result, numb)
			}
		}
	}

	return result
}
