package main

import (
	"fmt"
	"io/ioutil"

	"gocv.io/x/gocv"
	//"Viewer/cmd/GUI"
)

func image(image_source string) {
	img := gocv.IMRead(image_source, gocv.IMReadColor)
	if img.Empty() {
		fmt.Println("Image processing error", image_source)
	}
	window := gocv.NewWindow("Output")
	window.IMShow(img)
	window.WaitKey(0)
}

func read_dir(dir_source string) []string {
	files, err := ioutil.ReadDir(dir_source)
	if err != nil {
		fmt.Println("Read error!")
	}
	out_lst := []string{}
	for _, file := range files {
		out_lst = append(out_lst, file.Name())
	}
	return out_lst
}

func mass_read(source string) {
	img_lst := read_dir(source)
	for i := 0; i < len(img_lst); i++ {
		fmt.Println(img_lst[i])
	}
}

func main() {
	source := "b.jpg"
	//dir := "/home/vyacheslav/Projects/active_projects/PH_CNN_Model/data"
	//read_dir(dir)
	image(source)
	//mass_read(dir)

}
