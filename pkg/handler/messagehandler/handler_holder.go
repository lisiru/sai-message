package messagehandler

var Handlers = make(map[int]MessageHandler)

func SelectHandler(chanelCode int) MessageHandler {
	return Handlers[chanelCode]
}
