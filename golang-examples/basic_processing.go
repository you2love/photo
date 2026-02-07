// Golang图像处理示例
// 使用标准库和第三方包进行图像处理

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"

	// 使用golang.org/x/image进行高级图像操作
	"golang.org/x/image/draw"
	"golang.org/x/image/math/f64"
)

// resizeImage 调整图像大小
func resizeImage(inputPath, outputPath string, newWidth, newHeight int) error {
	// 打开输入文件
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("无法打开输入文件: %v", err)
	}
	defer file.Close()

	// 解码图像
	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("无法解码图像: %v", err)
	}

	// 创建新的RGBA图像
	dst := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	// 调整大小
	draw.NearestNeighbor.Scale(dst, dst.Bounds(), img, img.Bounds(), draw.Over, nil)

	// 创建输出文件
	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("无法创建输出文件: %v", err)
	}
	defer outFile.Close()

	// 编码为JPEG并写入文件
	err = jpeg.Encode(outFile, dst, &jpeg.Options{Quality: 90})
	if err != nil {
		return fmt.Errorf("无法编码图像: %v", err)
	}

	fmt.Printf("图像已调整大小并保存至: %s\n", outputPath)
	return nil
}

// applyGrayscale 应用灰度滤镜
func applyGrayscale(inputPath, outputPath string) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("无法打开输入文件: %v", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("无法解码图像: %v", err)
	}

	bounds := img.Bounds()
	grayImg := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			oldColor := img.At(x, y)
			grayColor := color.GrayModel.Convert(oldColor).(color.Gray)
			grayImg.SetGray(x, y, grayColor)
		}
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("无法创建输出文件: %v", err)
	}
	defer outFile.Close()

	err = jpeg.Encode(outFile, grayImg, &jpeg.Options{Quality: 90})
	if err != nil {
		return fmt.Errorf("无法编码图像: %v", err)
	}

	fmt.Printf("灰度滤镜已应用并保存至: %s\n", outputPath)
	return nil
}

// applyBrightnessContrast 调整亮度和对比度
func applyBrightnessContrast(inputPath, outputPath string, brightness float64, contrast float64) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("无法打开输入文件: %v", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("无法解码图像: %v", err)
	}

	bounds := img.Bounds()
	rgbaImg := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()

			// 将颜色值从uint32转换为float64进行计算
			rF := float64(r>>8) / 255.0
			gF := float64(g>>8) / 255.0
			bF := float64(b>>8) / 255.0

			// 应用亮度调整
			rF = rF + brightness
			gF = gF + brightness
			bF = bF + brightness

			// 应用对比度调整
			rF = ((rF - 0.5) * contrast) + 0.5
			gF = ((gF - 0.5) * contrast) + 0.5
			bF = ((bF - 0.5) * contrast) + 0.5

			// 确保值在有效范围内
			rF = mathMax(0, mathMin(1, rF))
			gF = mathMax(0, mathMin(1, gF))
			bF = mathMax(0, mathMin(1, bF))

			// 转换回uint32
			newR := uint32(rF * 255.0)
			newG := uint32(gF * 255.0)
			newB := uint32(bF * 255.0)

			rgbaImg.Set(x, y, color.RGBA{uint8(newR), uint8(newG), uint8(newB), uint8(a >> 8)})
		}
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("无法创建输出文件: %v", err)
	}
	defer outFile.Close()

	err = jpeg.Encode(outFile, rgbaImg, &jpeg.Options{Quality: 90})
	if err != nil {
		return fmt.Errorf("无法编码图像: %v", err)
	}

	fmt.Printf("亮度和对比度已调整并保存至: %s\n", outputPath)
	return nil
}

// mathMax 返回两个浮点数的最大值
func mathMax(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// mathMin 返回两个浮点数的最小值
func mathMin(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

// convertFormat 转换图像格式
func convertFormat(inputPath, outputPath string) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("无法打开输入文件: %v", err)
	}
	defer file.Close()

	img, format, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("无法解码图像: %v", err)
	}

	fmt.Printf("输入图像格式: %s\n", format)

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("无法创建输出文件: %v", err)
	}
	defer outFile.Close()

	// 根据输出文件扩展名确定格式
	if isJpeg(outputPath) {
		err = jpeg.Encode(outFile, img, &jpeg.Options{Quality: 90})
	} else if isPng(outputPath) {
		err = png.Encode(outFile, img)
	} else if isGif(outputPath) {
		err = gif.Encode(outFile, img, nil)
	} else {
		return fmt.Errorf("不支持的输出格式: %s", outputPath)
	}

	if err != nil {
		return fmt.Errorf("无法编码图像: %v", err)
	}

	fmt.Printf("图像已转换格式并保存至: %s\n", outputPath)
	return nil
}

// isJpeg 检查文件是否为JPEG格式
func isJpeg(path string) bool {
	return len(path) > 4 && path[len(path)-4:] == ".jpg" || path[len(path)-5:] == ".jpeg"
}

// isPng 检查文件是否为PNG格式
func isPng(path string) bool {
	return len(path) > 4 && path[len(path)-4:] == ".png"
}

// isGif 检查文件是否为GIF格式
func isGif(path string) bool {
	return len(path) > 4 && path[len(path)-4:] == ".gif"
}

// createThumbnail 创建缩略图
func createThumbnail(inputPath, outputPath string, thumbSize int) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("无法打开输入文件: %v", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("无法解码图像: %v", err)
	}

	// 计算缩略图尺寸，保持宽高比
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	var newWidth, newHeight int
	if width > height {
		newWidth = thumbSize
		newHeight = (height * thumbSize) / width
	} else {
		newHeight = thumbSize
		newWidth = (width * thumbSize) / height
	}

	dst := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	draw.NearestNeighbor.Scale(dst, dst.Bounds(), img, img.Bounds(), draw.Over, nil)

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("无法创建输出文件: %v", err)
	}
	defer outFile.Close()

	err = jpeg.Encode(outFile, dst, &jpeg.Options{Quality: 80})
	if err != nil {
		return fmt.Errorf("无法编码图像: %v", err)
	}

	fmt.Printf("缩略图已创建并保存至: %s\n", outputPath)
	return nil
}

func main() {
	fmt.Println("开始Golang图像处理演示...")

	inputPath := "input.jpg"    // 替换为实际输入路径
	outputPath := "output.jpg"  // 输出路径

	// 调整图像大小
	err := resizeImage(inputPath, "resized_output.jpg", 800, 600)
	if err != nil {
		fmt.Printf("调整大小失败: %v\n", err)
	}

	// 应用灰度滤镜
	err = applyGrayscale(inputPath, "grayscale_output.jpg")
	if err != nil {
		fmt.Printf("应用灰度滤镜失败: %v\n", err)
	}

	// 调整亮度和对比度
	err = applyBrightnessContrast(inputPath, "adjusted_output.jpg", 0.1, 1.2)
	if err != nil {
		fmt.Printf("调整亮度对比度失败: %v\n", err)
	}

	// 转换图像格式
	err = convertFormat(inputPath, "converted_output.png")
	if err != nil {
		fmt.Printf("转换格式失败: %v\n", err)
	}

	// 创建缩略图
	err = createThumbnail(inputPath, "thumbnail.jpg", 150)
	if err != nil {
		fmt.Printf("创建缩略图失败: %v\n", err)
	}

	fmt.Println("Golang图像处理演示完成!")
}