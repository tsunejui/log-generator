package template

func TextTemplate() string {
	return `time="{{ .Time }}" level="{{ .Level }}" msg="Port 80 for service ingress-nginx-controller is already opened by another service"`
}
