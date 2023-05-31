package captcha

import (
	"encoding/base64"
	"errors"
	"fmt"
	"gocv.io/x/gocv"
)

const width = 270

func process(payload ImageResponse) (ValidatePayload, error) {
	img, err := base64ToImage(payload.SrcImage)
	if err != nil {
		return ValidatePayload{}, err
	}

	offset, err := findOffset(img)
	if err != nil {
		return ValidatePayload{}, err
	}

	ratio := float64(width) / float64(payload.Width)
	return ValidatePayload{
		Offset: int(float64(offset)*ratio) - 4,
		Ratio:  ratio,
	}, nil
}

func findOffset(img gocv.Mat) (int, error) {
	contours := toContours(toMask(img))
	if contours.Size() != 1 {
		return 0, fmt.Errorf("there are %d contours", contours.Size())
	}
	rect := gocv.BoundingRect(contours.At(0))
	fmt.Println(rect.Size(), rect.Min)

	//// draw rect
	//gocv.Rectangle(&img, rect, color.RGBA{R: 255}, 1)
	//gocv.IMWrite("img.png", img)

	return rect.Min.X, nil
}

func base64ToImage(b64 string) (gocv.Mat, error) {
	imgData, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return gocv.Mat{}, err
	}
	img, err := gocv.IMDecode(imgData, gocv.IMReadColor)
	if err != nil {
		return gocv.Mat{}, err
	}
	if img.Empty() {
		return gocv.Mat{}, errors.New("empty image")
	}
	//hsvImg := gocv.NewMat()
	//gocv.CvtColor(img, &hsvImg, gocv.ColorBGRToHSV)
	return img, nil
}

func toEdge(img gocv.Mat) gocv.Mat {
	edge := gocv.NewMat()
	gocv.Canny(img, &edge, 100, 500)
	return edge
}

func toContours(img gocv.Mat) gocv.PointsVector {
	return gocv.FindContours(img, gocv.RetrievalExternal, gocv.ChainApproxSimple)
}

var grey = gocv.NewScalar(192.0, 192.0, 192.0, 0.0)

func toMask(img gocv.Mat) gocv.Mat {
	mat := gocv.NewMatWithSizeFromScalar(grey, img.Rows(), img.Cols(), gocv.MatTypeCV8UC3)
	mask := gocv.NewMat()
	gocv.InRange(img, mat, mat, &mask)
	return mask
}
