package utils

import "github.com/golang/glog"

// CheckErr is a convinience function to throw a log a fatal effort in case of errors.
func CheckErr(err error) {
	if err != nil {
		glog.Fatal(err)
	}
}
