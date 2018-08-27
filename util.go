package main

import "os"

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func removeTempFile(path string) {
	if exists(path) {
		if err := os.Remove(path); err != nil {
			panic(err)
		}
	}
}
