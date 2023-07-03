package captcha

import (
	"errors"
	"fdutools-go/fdu"
	"fdutools-go/utils"
	"log"
)

func DoN(c *fdu.Client, n int) error {
	for i := 0; i < n; i++ {
		err := Do(c)
		if err == nil {
			return nil
		}
		log.Printf("captcha error: %s", err)
	}
	log.Printf("captcha failed after %d try", n)
	return errors.New("captcha failed")
}

func Do(c *fdu.Client) error {
	resp, err := c.Get("https://xk.fudan.edu.cn/xk/stdElectCourse!getImgSwipe.action")
	if err != nil {
		return err
	}
	var image ImageResponse
	err = utils.UnmarshalResponse(resp.Body, &image)
	if err != nil {
		return err
	}

	payload, err := process(image)
	if err != nil {
		return err
	}

	resp, err = c.Post(
		"https://xk.fudan.edu.cn/xk/stdElectCourse!rstImgSwipe.action",
		utils.StructToMap(payload),
	)
	if err != nil {
		return err
	}
	var res ValidateResponse
	err = utils.UnmarshalResponse(resp.Body, &res)
	if err != nil {
		return err
	}
	if !res.Success {
		return errors.New("captcha failed")
	}
	return nil
}

type ImageResponse struct {
	SrcImage string `json:"SrcImage,omitempty"`
	CutImage string `json:"CutImage,omitempty"`
	Width    int    `json:"SrcImageWidth,omitempty"`
	Height   int    `json:"SrcImageHeight,omitempty"`
}

type ValidateResponse struct {
	Success bool `json:"success"`
}

type ValidatePayload struct {
	Offset int     `json:"moveEnd_X"`
	Ratio  float64 `json:"wbili"`
}
