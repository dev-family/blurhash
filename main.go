package main

import (
    "bytes"
	"fmt"
	"github.com/buckket/go-blurhash"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"strconv"
	"encoding/json"
)

type BlurHashResponse struct {
	Message string `json:"blur_hash"`
}

func convertToInt(str string) (int, error) {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("error converting string to integer: %v", err)
	}
	return num, nil
}

func encode(c *routing.Context) error {
	imageURL := string(c.QueryArgs().Peek("image_url"))
	x := 9
	y := 9

	xConverted, err := convertToInt(string(c.QueryArgs().Peek("x")))
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
	} else {
		x = xConverted
	}

	yConverted, err := convertToInt(string(c.QueryArgs().Peek("y")))
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
	} else {
		y = yConverted
	}

	fmt.Println("Image URL:", imageURL)
	statusCode, body, err := fasthttp.Get(nil, imageURL)

	if err != nil {
		fmt.Println("Error downloading image:", err)
		return err
	}

	if statusCode != fasthttp.StatusOK {
		fmt.Println("Error downloading image, status code:", statusCode)
		return fmt.Errorf("error downloading image, status code: %d", statusCode)
	}

	img, _, err := image.Decode(bytes.NewReader(body))
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return err
	}

	blurHash, err := blurhash.Encode(x, y, img)
	if err != nil {
		fmt.Println("Error generating BlurHash:", err)
		return err
	}

	jsonResponse := BlurHashResponse{
		Message: blurHash,
	}

	jsonBytes, err := json.Marshal(jsonResponse)

	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	        return err
	}

	c.Response.Header.SetContentType("application/json")
	c.Response.SetBodyRaw(jsonBytes)

	return nil
}

func main() {
	router := routing.New()
	router.Get("/", encode)
	fasthttp.ListenAndServe(":3333", router.HandleRequest)
}
