package message

type MessageHandler interface {
	handle(message string, params ...interface{}) string
	test(message string) bool
}

type BaseMessageHandler struct {
}

var handlers = []MessageHandler{
	&ChooseHandler{},
	&YesOrNoHandler{},
}

func ProcessMessage(message string) string {
	for _, handler := range handlers {
		if handler.test(message) {
			return handler.handle(message)
		}
	}
	return ""
}
