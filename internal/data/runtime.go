package data

type Runtime int32

//todo 报错了，暂时无法找到解决方案
// func (r Runtime) MarshalJSON([]byte, error) {
// 	// jsonValue := fmt.Sprintf("%d mins", r)
// 	jsonValue := fmt.Sprintf("%d mins", r)
// 	// Use the strconv.Quote() function on the string to wrap it in double quotes. It
// 	// needs to be surrounded by double quotes in order to be a valid *JSON string*.
// 	// quotedJSONValue := strconv.Quote(jsonValue)
// 	quotedJSONValue := strconv.Quote(jsonValue)
// 	// Convert the quoted string value to a byte slice and return it.
// 	return []byte(quotedJSONValue), nil

// }
