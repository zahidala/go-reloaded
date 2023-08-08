package main

import "project"

func main() {
	arr := project.DataHandler()
	project.CheckArgsAndRun(arr)
}
