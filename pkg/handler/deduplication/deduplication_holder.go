package deduplication

var TypeMap = make(map[int]DeDuplicationService)


// 选择去重方式对应的具体执行逻辑
func SelectDeDuplicationService(deDuplicationType int) DeDuplicationService {
	return TypeMap[deDuplicationType]
}

