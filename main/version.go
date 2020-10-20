package main

import (
	"strconv"
	"strings"
)

// Version 版本
type Version string

// GetMajor 获取主版本号
func (v Version) GetMajor() int {
	return v.parse()[0]
}

// GetMinor 获取次版本号
func (v Version) GetMinor() int {
	return v.parse()[1]
}

// GetBuild 获取构建号
func (v Version) GetBuild() int {
	return v.parse()[2]
}

// Eq 等于
func (v Version) Eq(o Version) bool {
	partsV := v.parse()
	partsO := o.parse()
	return partsV[0]==partsO[0] && partsV[1]==partsO[1] && partsV[2]==partsO[2]
}

// Gt 大于
func (v Version) Gt(o Version) bool {
	partsV := v.parse()
	partsO := o.parse()
	// 比较major
	if partsV[0] > partsO[0] {
		return true
	} 
	if partsV[0] < partsO[0] {
		return false
	}
	// 比较minor
	if partsV[1] > partsO[1] {
		return true
	} 
	if partsV[1] < partsO[1] {
		return false
	}
	// 比较build
	return partsV[2] > partsO[2]
}

// Ge 大于等于
func (v Version) Ge(o Version) bool {
	partsV := v.parse()
	partsO := o.parse()
	// 比较major
	if partsV[0] > partsO[0] {
		return true
	} 
	if partsV[0] < partsO[0] {
		return false
	}
	// 比较minor
	if partsV[1] > partsO[1] {
		return true
	} 
	if partsV[1] < partsO[1] {
		return false
	}
	// 比较build
	return partsV[2] >= partsO[2]
}

// Lt 小于
func (v Version) Lt(o Version) bool {
	partsV := v.parse()
	partsO := o.parse()
	// 比较major
	if partsV[0] < partsO[0] {
		return true
	} 
	if partsV[0] > partsO[0] {
		return false
	}
	// 比较minor
	if partsV[1] < partsO[1] {
		return true
	} 
	if partsV[1] > partsO[1] {
		return false
	}
	// 比较build
	return partsV[2] < partsO[2]
}

// Le 小于等于
func (v Version) Le(o Version) bool {
	partsV := v.parse()
	partsO := o.parse()
	// 比较major
	if partsV[0] < partsO[0] {
		return true
	} 
	if partsV[0] > partsO[0] {
		return false
	}
	// 比较minor
	if partsV[1] < partsO[1] {
		return true
	} 
	if partsV[1] > partsO[1] {
		return false
	}
	// 比较build
	return partsV[2] <= partsO[2]
}

func (v Version) parse() []int {
	major := 0
	minor := 0
	build := 0
	var err error
	fields := strings.Split(string(v), ":")
	switch len(fields) {
	case 3:
		build, err = strconv.Atoi(fields[2])
		if err != nil {
			panic("Invalid version: " + v)
		}
		fallthrough
	case 2:
		minor, err = strconv.Atoi(fields[1])
		if err != nil {
			panic("Invalid version: " + v)
		}
		fallthrough
	case 1:
		major, err = strconv.Atoi(fields[0])
		if err != nil {
			panic("Invalid version: " + v)
		}
	default:
		panic("Invalid empty version")
	}
	return []int{major, minor, build}
}