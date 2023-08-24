package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/kindlyfire/go-keylogger"
)

func gen() {
	fImg1, _ := os.Open("source.jpg")
	defer fImg1.Close()
	img1, _, _ := image.Decode(fImg1)

	//choose a random point between 0,0 and 145, 165
	var x1 = rand.Intn(145)
	var y1 = rand.Intn(165)

	p1 := image.Point{x1, y1}

	//choose a random point between 450, 280 and 581, 351
	var x2 = rand.Intn(131) + 450
	var y2 = rand.Intn(71) + 280

	p2 := image.Point{x2, y2}

	//crop the image
	r1 := image.Rectangle{p1, p2}

	img2 := image.NewRGBA(r1)
	draw.Draw(img2, img2.Bounds(), img1, p1, draw.Src)

	//save the image
	var out_rand_str = "out_" + fmt.Sprint(rand.Intn(10000)) + ".jpg"
	toimg, _ := os.Create(out_rand_str)
	defer toimg.Close()

	jpeg.Encode(toimg, img2, nil)

	//copy THE FILE to clipboard, possibly via os.exec
	//get output from exec.Command("py copyimg.py " + out_rand_str).Run()
	var cmd = exec.Command("py", "copyimg.py", out_rand_str)
	//pipe out to println
	_, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	// fmt.Println(string(out))

	// //get image bytes
	// fImg2, _ := os.Open(out_rand_str)
	// defer fImg2.Close()

	// //convert to png
	// img3, _, _ := image.Decode(fImg2)

	// //convert to bytes, Image.Encode does not exist
	// var buf []byte = make([]byte, 0)
	// png.Encode(bytes.NewBuffer(buf), img3)

	// //copy to clipboard
	// clipboard.Write(clipboard.FmtImage, buf)

}

const (
	delayKeyfetchMS = 20
)

func main() {
	kl := keylogger.NewKeylogger()
	// emptyCount := 0

	for {
		key := kl.GetKey()

		if !key.Empty {
			fmt.Printf("'%c' %d                     \n", key.Rune, key.Keycode)
			// rune := fmt.Sprintf("%c", key.Rune)

			//if string(rune) == string("♀") /*Ctrl+Shift+L*/ {
			if key.Keycode == 111 /*Numpad / */ {
				fmt.Println("Generating...")
				gen()
				fmt.Println("Generated!")
			}
		}

		// emptyCount++

		// fmt.Printf("Empty count: %d\r", emptyCount)

		time.Sleep(delayKeyfetchMS * time.Millisecond)
	}
}
