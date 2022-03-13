package deduplication

var TypeMap = make(map[int]DeDuplicationService)


func SelectDeDuplicationService(deDuplicationType int) DeDuplicationService {
	return TypeMap[deDuplicationType]
}

