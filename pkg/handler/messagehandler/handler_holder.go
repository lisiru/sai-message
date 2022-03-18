package messagehandler

var Handlers = make(map[int]MessageHandler)

// 选择具体执行的handler
func SelectHandler(chanelCode int) MessageHandler {
	return Handlers[chanelCode]
}
