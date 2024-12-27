package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"sync"
	"time"
)

func count(ch <-chan int) {
	for num := range ch {
		fmt.Printf("Квадрат числа %d: %d\n", num, num*num)
	}
}

func filter(img *image.RGBA) {
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalColor := img.RGBAAt(x, y)
			gray := uint8((uint16(originalColor.R) + uint16(originalColor.G) + uint16(originalColor.B)) / 3)
			img.SetRGBA(x, y, color.RGBA{R: gray, G: gray, B: gray, A: originalColor.A})
		}
	}
}

func filterParallel(img *image.RGBA, y int, wg *sync.WaitGroup) {
	defer wg.Done()
	bounds := img.Bounds()
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		originalColor := img.RGBAAt(x, y)
		gray := uint8((uint16(originalColor.R) + uint16(originalColor.G) + uint16(originalColor.B)) / 3)
		img.SetRGBA(x, y, color.RGBA{R: gray, G: gray, B: gray, A: originalColor.A})
	}
}

func main() {
	// Задание 1
	ch := make(chan int)

	go count(ch)

	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)

	time.Sleep(time.Second)
	fmt.Println("Задание 1 завершено.")

	// Задание 2
	inputFile, err := os.Open("image.png")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer inputFile.Close()

	img, _, err := image.Decode(inputFile)
	if err != nil {
		fmt.Println("Ошибка декодирования изображения:", err)
		return
	}

	bounds := img.Bounds()
	rgbaImg := image.NewRGBA(bounds)

	// Копирование пикселей из img в rgbaImg
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalColor := img.At(x, y)
			rgbaImg.Set(x, y, originalColor)
		}
	}

	start := time.Now()
	filter(rgbaImg)
	fmt.Println("Последовательная обработка заняла:", time.Since(start))

	outputFile, err := os.Create("output.png")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, rgbaImg)
	if err != nil {
		fmt.Println("Ошибка сохранения изображения:", err)
		return
	}
	
	fmt.Println("Задание 2 завершено.")
 
	// Задание 3
	start = time.Now()
	var wg sync.WaitGroup
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		wg.Add(1)
		go filterParallel(rgbaImg, y, &wg)
	}
	wg.Wait()
	fmt.Println("Параллельная обработка заняла:", time.Since(start))

	outputFile, err = os.Create("output_parallel.png")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, rgbaImg)
	if err != nil {
		fmt.Println("Ошибка сохранения изображения:", err)
		return
	}

	fmt.Println("Задание 3 завершено.")
}
