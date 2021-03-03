package main

// ParamType 参数类型
type ParamType int

const(
	INTEGER ParamType = 1
	FLOAT ParamType = 2
	STRING ParamType = 3
	VERSION ParamType = 4
)

// Param 匹配参数
type Param struct {
	Name			string
	Type			ParamType
	Priority		int
}
