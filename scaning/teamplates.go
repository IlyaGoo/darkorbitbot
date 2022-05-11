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
	name  string
	Image *image.RGBA
	color color.RGBA
}

func GetTeamplates() []*teamplate {
	x := []*teamplate{
		{"box", loadTeamplate("box"), red},
		{"resbox", loadTeamplate("resbox"), red},
		{"mmo", loadTeamplate("mmo"), red},
		{"connect", loadTeamplate("connect"), red},
		{"close", loadTeamplate("close"), red},
		{"cargo", loadTeamplate("ship/cargo"), red},
		{"trade", loadTeamplate("trade"), red},
		{"sell", loadTeamplate("sell"), red},
		{"minimap_icon", loadTeamplate("windows/minimap_icon"), red},
		{"user_icon", loadTeamplate("windows/user_icon"), red},
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
