package utils

import (
	"image"
	"image/draw"
	"image/jpeg"
	"math"
	"os"

	"github.com/nfnt/resize"
)

func fixSize(img1W, img2W int, MaxWidth float64) (new1W, new2W int) {
	var ( //为了方便计算，将两个图片的宽转为 float64
		img1Width, img2Width = float64(img1W), float64(img2W)
		ratio1, ratio2       float64
	)
	minWidth := math.Min(img1Width, img2Width) // 取出两张图片中宽度最小的为基准

	if minWidth > MaxWidth { // 如果最小宽度大于MaxWidth，那么两张图片都需要进行缩放
		ratio1 = MaxWidth / img1Width // 图片1的缩放比例
		ratio2 = MaxWidth / img2Width // 图片2的缩放比例

		// 原宽度 * 比例 = 新宽度
		return int(img1Width * ratio1), int(img2Width * ratio2)
	}

	// 如果最小宽度小于MaxWidth，那么需要将较大的图片缩放，使得两张图片的宽度一致
	if minWidth == img1Width {
		ratio2 = minWidth / img2Width // 图片2的缩放比例
		return img1W, int(img2Width * ratio2)
	}

	ratio1 = minWidth / img1Width // 图片1的缩放比例
	return int(img1Width * ratio1), img2W
}

func Stitching(file1path, file2path, outpath string, MaxWidth float64) error {
	file1, _ := os.Open(file1path)
	file2, _ := os.Open(file2path)
	defer file1.Close()
	defer file2.Close()

	var (
		img1, img2 image.Image
		err        error
	)
	if img1, _, err = image.Decode(file1); err != nil {
		return err
	}
	if img2, _, err = image.Decode(file2); err != nil {
		return err
	}
	b1 := img1.Bounds()
	b2 := img2.Bounds()
	new1W, new2W := fixSize(b1.Max.X, b2.Max.X, MaxWidth)

	// 调用resize库进行图片缩放(高度填0，resize.Resize函数中会自动计算缩放图片的宽高比)
	m1 := resize.Resize(uint(new1W), 0, img1, resize.Lanczos3)
	m2 := resize.Resize(uint(new2W), 0, img2, resize.Lanczos3)

	// 将两个图片合成一张
	newWidth := m1.Bounds().Max.X                                                                          //新宽度 = 随意一张图片的宽度
	newHeight := m1.Bounds().Max.Y + m2.Bounds().Max.Y                                                     // 新图片的高度为两张图片高度的和
	newImg := image.NewNRGBA(image.Rect(0, 0, newWidth, newHeight))                                        //创建一个新RGBA图像
	draw.Draw(newImg, newImg.Bounds(), m1, m1.Bounds().Min, draw.Over)                                     //画上第一张缩放后的图片
	draw.Draw(newImg, newImg.Bounds(), m2, m2.Bounds().Min.Sub(image.Pt(0, m1.Bounds().Max.Y)), draw.Over) //画上第二张缩放后的图片（这里需要注意Y值的起始位置）
	// 保存文件
	imgfile, _ := os.Create(outpath)
	defer imgfile.Close()
	jpeg.Encode(imgfile, newImg, &jpeg.Options{Quality: 100})
	return nil
}
