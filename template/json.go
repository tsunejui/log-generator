package template

func JsonTemplate() string {
	return `{"level":"{{ .Level }}","ts":"{{ .Time }}","logger":"eventlistener","caller":"sink/sink.go:329","msg":"interceptor stopped trigger processing: rpc error: code = FailedPrecondition desc = event type Push Hook is not allowed","eventlistener":"test-pipeline","namespace":"default","eventlistenerUID":"123abcd4-1111-2222-3333-456a7b8c9d01","/triggers-eventid":"123abcd4-1111-2222-3333-456a7b8c9d01","/trigger":"trigger-cicd"}`
}
