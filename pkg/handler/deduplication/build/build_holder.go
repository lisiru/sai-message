package build

var BuildTyepMap = make(map[int]build)

// 选择具体的参数构建方式
func SelectParamBuild(deDuplicationType int) build {
	return BuildTyepMap[deDuplicationType]
}
