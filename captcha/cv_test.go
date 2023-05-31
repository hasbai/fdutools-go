package captcha

import (
	"encoding/json"
	"gocv.io/x/gocv"
	"os"
	"testing"
)

func TestCV(t *testing.T) {
	_, err := findOffset(getImage(t))
	if err != nil {
		t.Fatal(err)
	}
}

func TestCanny(t *testing.T) {
	img := getImage(t)
	edge := toEdge(img)
	gocv.IMWrite("src.png", img)
	gocv.IMWrite("edge.png", edge)
}

func TestMask(t *testing.T) {
	img := getImage(t)
	mask := toMask(img)
	gocv.IMWrite("mask.png", mask)
}

func getPayload(t *testing.T) ImageResponse {
	data, err := os.ReadFile("resp.json")
	if err != nil {
		t.Fatal(err)
	}

	var payload ImageResponse
	err = json.Unmarshal(data, &payload)
	if err != nil {
		t.Fatal(err)
	}
	return payload
}

func getImage(t *testing.T) gocv.Mat {
	payload := getPayload(t)
	img, err := base64ToImage(payload.SrcImage)
	if err != nil {
		t.Fatal(err)
	}
	return img
}
