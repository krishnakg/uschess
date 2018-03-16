package utils

import "github.com/golang/glog"

func CheckErr(err error) {
	if err != nil {
		glog.Fatal(err)
	}
}
