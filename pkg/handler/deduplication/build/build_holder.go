package build

var BuildTyepMap = make(map[int]build)

func SelectParamBuild(deDuplicationType int) build {
	return BuildTyepMap[deDuplicationType]
}
