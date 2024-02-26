package main

import (

)

type Merger interface {

}

type merger struct {

}

func NewMerger() (Merger) {
	return &merger{}
}
