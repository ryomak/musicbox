package action

import (
	"fmt"
	"github.com/ryomak/musicbox/sound"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
)

var mp3Path string = os.Getenv("GOPATH") + "/src/github.com/ryomak/musicbox/sound/mp3"

func List(c *cli.Context) {
	list, err := ioutil.ReadDir(mp3Path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	writeText := ""
	for i, fileInfo := range list {
		c.Set(fmt.Sprint(i),fileInfo.Name())
		fmt.Printf("[%d]:%s\n", i+1,fileInfo.Name())
		writeText += fileInfo.Name()+"\n"
	}
	err = WriteList(writeText)
	if err != nil{
		fmt.Printf("list err %v",err)
	}
}

func Play(c *cli.Context) {
	fmt.Printf("music list....\n")
	musicList := []string{}
	for i,v := range c.Args(){
		str ,err := ReadList(v)
		if err != nil{
			fmt.Printf("%v", err)
		}
		fmt.Printf("[%d]%s\n",i,str)
		musicList = append(musicList,str)
	}
	fmt.Printf("music start....\n")
	for _,v := range musicList{
		err := sound.Sound(mp3Path+"/"+v)
		if err != nil {
			fmt.Printf("%v", err)
		}
	}
	fmt.Printf("music end...")
}

/*
func Play(c *cli.Context) {
	fmt.Printf("music list....\n")
	for i,v := range c.Args(){
		fmt.Printf("%d:%s\n",i,v)
	}
	fmt.Printf("music start....\n")
	for _,v := range c.Args(){
		err := sound.Sound(mp3Path+"/"+v)
		if err != nil {
			fmt.Printf("%v", err)
		}
	}
	fmt.Printf("music end...")
}
*/