package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func RemoveContents(dir string) error{
	d, err := os.Open(dir)
	if err != nil{
		return err
	}

	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil{
		return err
	}

	for _, name := range names{
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil{
			return err
		}
	}

	return nil
}

func StartProcess(Filename string) {
	// Filename = "cmd /C start \"" + Filename + "\""
	cmd := exec.Command(Filename)
	err := cmd.Start()
	if err!=nil{
		fmt.Println(err)
	}
}

func main(){
	err := RemoveContents("C:\\Program Files (x86)\\Steam\\appcache\\httpcache")
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

	StartProcess("C:\\Program Files (x86)\\Steam\\steam.exe")
}