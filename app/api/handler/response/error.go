package response

func ErrorResponse(message string) map[string]string {
	return map[string]string{
		"error": message,
	}
}