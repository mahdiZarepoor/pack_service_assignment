package logging

func mapToZapParams(keys map[ExtraKey]interface{}) []interface{} {
	params := make([]interface{}, 0, len(keys))
	for k, v := range keys {
		params = append(params, string(k), v)
	}
	return params
}
