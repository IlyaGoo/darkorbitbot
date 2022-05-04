package scaning

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log"
	"os"
)

type teamplate struct {
	name                 string
	image                *image.RGBA
	recommendedTolerance float64
	color                color.RGBA
}

func getTeamplates() []*teamplate {
	x := []*teamplate{
		{"box", loadTeamplate("box"), 0.19, red},
		{"resbox", loadTeamplate("resbox"), 0.1, red},
		{"mmo", loadTeamplate("mmo"), 0.04, red},
		{"connect", loadTeamplate("connect"), 0.04, red},
	}
	return x
}

func loadTeamplate(addPath string) *image.RGBA {
	//Открываем экзампл
	path := fmt.Sprintf("templates/%s.png", addPath)
	existingImageFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer existingImageFile.Close()
	//Возвращаем каретку в начало (фиксит некоторые баги либы(я не шучу, удалять нельзя))
	existingImageFile.Seek(0, 0)
	//Декодируем экзампл
	imageData, _, err := image.Decode(existingImageFile)
	if err != nil {
		log.Fatal(err)
	}
	//Конвертим экзампл в RGBA
	b := imageData.Bounds()
	example := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(example, example.Bounds(), imageData, b.Min, draw.Src)
	return example
}
