/*
Golang图像处理综合示例
实现图像处理基础知识和常见算法
*/

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"math"
	"os"

	// 使用golang.org/x/image进行高级图像操作
	"golang.org/x/image/draw"
)

// BasicImageOperations 基础图像操作示例
func BasicImageOperations() {
	fmt.Println("=== 基础图像操作 ===")
	
	// 读取图像
	img, err := loadImage("sample.jpg")
	if err != nil {
		fmt.Println("无法读取图像，跳过此部分")
		return
	}
	
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	
	fmt.Printf("图像尺寸: %dx%d\n", width, height)
	fmt.Printf("图像边界: %v\n", bounds)
}

// GeometricTransformations 几何变换示例
func GeometricTransformations() {
	fmt.Println("\n=== 几何变换 ===")
	
	img, err := loadImage("sample.jpg")
	if err != nil {
		fmt.Println("无法读取图像，跳过此部分")
		return
	}
	
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	
	// 缩放图像
	scaledImg := resizeImage(img, width/2, height/2)
	fmt.Printf("缩放后尺寸: %dx%d\n", scaledImg.Bounds().Dx(), scaledImg.Bounds().Dy())
	
	// 裁剪图像
	croppedImg := cropImage(img, width/4, height/4, width/2, height/2)
	fmt.Printf("裁剪后尺寸: %dx%d\n", croppedImg.Bounds().Dx(), croppedImg.Bounds().Dy())
	
	fmt.Println("几何变换完成")
}

// FilteringAlgorithms 滤波算法示例
func FilteringAlgorithms() {
	fmt.Println("\n=== 滤波算法 ===")
	
	img, err := loadImage("sample.jpg")
	if err != nil {
		fmt.Println("无法读取图像，跳过此部分")
		return
	}
	
	// 均值滤波
	meanFiltered := convolve(img, meanKernel)
	
	// 锐化滤波
	sharpened := convolve(img, sharpenKernel)
	
	// 边缘检测滤波
	edgeDetected := convolve(img, edgeKernel)
	
	fmt.Printf("滤波处理完成，输出图像尺寸: %dx%d\n", meanFiltered.Bounds().Dx(), meanFiltered.Bounds().Dy())
}

// ColorSpaceConversions 色彩空间转换示例
func ColorSpaceConversions() {
	fmt.Println("\n=== 色彩空间转换 ===")
	
	img, err := loadImage("sample.jpg")
	if err != nil {
		fmt.Println("无法读取图像，跳过此部分")
		return
	}
	
	// RGB转灰度
	grayImg := rgbToGrayscale(img)
	fmt.Printf("灰度转换完成，输出图像尺寸: %dx%d\n", grayImg.Bounds().Dx(), grayImg.Bounds().Dy())
	
	// 演示HSV转换
	bounds := img.Bounds()
	x, y := bounds.Min.X, bounds.Min.Y
	r, g, b, _ := img.At(x, y).RGBA()
	h, s, v := rgbToHSV(uint8(r>>8), uint8(g>>8), uint8(b>>8))
	fmt.Printf("RGB(%d,%d,%d) -> HSV(%.2f,%.2f,%.2f)\n", r>>8, g>>8, b>>8, h, s, v)
	
	fmt.Println("色彩空间转换完成")
}

// MorphologicalOperations 形态学操作示例
func MorphologicalOperations() {
	fmt.Println("\n=== 形态学操作 ===")
	
	// 创建一个简单的二值图像用于演示
	binaryImg := createBinaryImage(100, 100)
	
	// 腐蚀操作
	eroded := erode(binaryImg, 3)
	fmt.Printf("腐蚀操作完成，输出图像尺寸: %dx%d\n", eroded.Bounds().Dx(), eroded.Bounds().Dy())
	
	// 膨胀操作
	dilated := dilate(binaryImg, 3)
	fmt.Printf("膨胀操作完成，输出图像尺寸: %dx%d\n", dilated.Bounds().Dx(), dilated.Bounds().Dy())
	
	fmt.Println("形态学操作完成")
}

// HistogramProcessing 直方图处理示例
func HistogramProcessing() {
	fmt.Println("\n=== 直方图处理 ===")
	
	img, err := loadImage("sample.jpg")
	if err != nil {
		fmt.Println("无法读取图像，跳过此部分")
		return
	}
	
	// 计算直方图
	histogram := computeHistogram(img)
	fmt.Printf("计算直方图完成，共 %d 个像素\n", len(histogram))
	
	// 直方图均衡化
	equalized := histogramEqualization(img)
	fmt.Printf("直方图均衡化完成，输出图像尺寸: %dx%d\n", equalized.Bounds().Dx(), equalized.Bounds().Dy())
	
	fmt.Println("直方图处理完成")
}

// AdvancedTechniques 高级图像处理技术示例
func AdvancedTechniques() {
	fmt.Println("\n=== 高级图像处理技术 ===")
	
	img, err := loadImage("sample.jpg")
	if err != nil {
		fmt.Println("无法读取图像，跳过此部分")
		return
	}
	
	// 图像二值化
	binary := binarize(img, 128)
	fmt.Printf("二值化完成，输出图像尺寸: %dx%d\n", binary.Bounds().Dx(), binary.Bounds().Dy())
	
	// 图像亮度调整
	brightened := adjustBrightness(img, 1.2)
	fmt.Printf("亮度调整完成，输出图像尺寸: %dx%d\n", brightened.Bounds().Dx(), brightened.Bounds().Dy())
	
	// 图像对比度调整
	contrasted := adjustContrast(img, 1.2)
	fmt.Printf("对比度调整完成，输出图像尺寸: %dx%d\n", contrasted.Bounds().Dx(), contrasted.Bounds().Dy())
	
	fmt.Println("高级图像处理技术完成")
}

// loadImage 加载图像
func loadImage(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	img, _, err := image.Decode(file)
	return img, err
}

// resizeImage 图像缩放
func resizeImage(img image.Image, newWidth, newHeight int) image.Image {
	bounds := img.Bounds()
	dst := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	
	// 使用Catmull-Rom样条插值进行缩放
	draw.CatmullRom.Scale(dst, dst.Bounds(), img, bounds, draw.Over, nil)
	
	return dst
}

// cropImage 图像裁剪
func cropImage(img image.Image, x, y, width, height int) image.Image {
	bounds := img.Bounds()
	cropBounds := image.Rect(x, y, x+width, y+height)
	
	// 确保裁剪区域在原图范围内
	cropBounds = cropBounds.Intersect(bounds)
	
	dst := image.NewRGBA(cropBounds)
	draw.Draw(dst, cropBounds, img, cropBounds.Min, draw.Src)
	
	return dst
}

// convolve 卷积滤波
func convolve(img image.Image, kernel [][]float64) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	
	// 创建输出图像
	dst := image.NewRGBA(bounds)
	
	// 计算核的中心位置
	kernelSize := len(kernel)
	kernelCenter := kernelSize / 2
	
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			var rSum, gSum, bSum float64
			
			for ky := 0; ky < kernelSize; ky++ {
				for kx := 0; kx < kernelSize; kx++ {
					// 计算图像中的对应位置
					imgX := x + kx - kernelCenter
					imgY := y + ky - kernelCenter
					
					// 边界处理：使用镜像填充
					if imgX < bounds.Min.X {
						imgX = bounds.Min.X + (bounds.Min.X - imgX)
					} else if imgX >= bounds.Max.X {
						imgX = bounds.Max.X - 1 - (imgX - bounds.Max.X)
					}
					
					if imgY < bounds.Min.Y {
						imgY = bounds.Min.Y + (bounds.Min.Y - imgY)
					} else if imgY >= bounds.Max.Y {
						imgY = bounds.Max.Y - 1 - (imgY - bounds.Max.Y)
					}
					
					// 获取像素值
					r, g, b, _ := img.At(imgX, imgY).RGBA()
					
					// 应用核权重
					weight := kernel[ky][kx]
					rSum += float64(r>>8) * weight
					gSum += float64(g>>8) * weight
					bSum += float64(b>>8) * weight
				}
			}
			
			// 限制值范围并设置像素
			r := uint8(clamp(rSum, 0, 255))
			g := uint8(clamp(gSum, 0, 255))
			b := uint8(clamp(bSum, 0, 255))
			
			dst.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}
	
	return dst
}

// clamp 限制值在指定范围内
func clamp(value float64, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// rgbToGrayscale RGB转灰度
func rgbToGrayscale(img image.Image) image.Image {
	bounds := img.Bounds()
	grayImg := image.NewGray(bounds)
	
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			
			// 使用加权平均转换为灰度
			grayValue := 0.299*float64(r>>8) + 0.587*float64(g>>8) + 0.114*float64(b>>8)
			
			grayImg.SetGray(x, y, color.Gray{uint8(grayValue)})
		}
	}
	
	return grayImg
}

// rgbToHSV RGB转HSV
func rgbToHSV(r, g, b uint8) (float64, float64, float64) {
	rf := float64(r) / 255.0
	gf := float64(g) / 255.0
	bf := float64(b) / 255.0
	
	maxVal := math.Max(math.Max(rf, gf), bf)
	minVal := math.Min(math.Min(rf, gf), bf)
	delta := maxVal - minVal
	
	// 计算明度V
	v := maxVal
	
	// 计算饱和度S
	var s float64
	if maxVal == 0 {
		s = 0
	} else {
		s = delta / maxVal
	}
	
	// 计算色调H
	var h float64
	if delta == 0 {
		h = 0
	} else if maxVal == rf {
		h = 60 * math.Mod((gf-bf)/delta, 6)
	} else if maxVal == gf {
		h = 60 * ((bf-rf)/delta + 2)
	} else { // maxVal == bf
		h = 60 * ((rf-gf)/delta + 4)
	}
	
	if h < 0 {
		h += 360
	}
	
	return h, s, v
}

// createBinaryImage 创建二值图像
func createBinaryImage(width, height int) image.Image {
	img := image.NewGray(image.Rect(0, 0, width, height))
	
	// 创建一些简单的图案
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			var grayValue uint8
			if (x+y)%20 < 10 {
				grayValue = 255 // 白色
			} else {
				grayValue = 0 // 黑色
			}
			img.SetGray(x, y, color.Gray{grayValue})
		}
	}
	
	return img
}

// erode 腐蚀操作
func erode(img image.Image, kernelSize int) image.Image {
	bounds := img.Bounds()
	dst := image.NewGray(bounds)
	
	kernelRadius := kernelSize / 2
	
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			minVal := uint8(255) // 初始化为最大值
			
			for ky := -kernelRadius; ky <= kernelRadius; ky++ {
				for kx := -kernelRadius; kx <= kernelRadius; kx++ {
					imgX := x + kx
					imgY := y + ky
					
					// 检查边界
					if imgX >= bounds.Min.X && imgX < bounds.Max.X &&
						imgY >= bounds.Min.Y && imgY < bounds.Max.Y {
						gray := color.GrayModel.Convert(img.At(imgX, imgY)).(color.Gray)
						if gray.Y < minVal {
							minVal = gray.Y
						}
					}
				}
			}
			
			dst.SetGray(x, y, color.Gray{minVal})
		}
	}
	
	return dst
}

// dilate 膨胀操作
func dilate(img image.Image, kernelSize int) image.Image {
	bounds := img.Bounds()
	dst := image.NewGray(bounds)
	
	kernelRadius := kernelSize / 2
	
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			maxVal := uint8(0) // 初始化为最小值
			
			for ky := -kernelRadius; ky <= kernelRadius; ky++ {
				for kx := -kernelRadius; kx <= kernelRadius; kx++ {
					imgX := x + kx
					imgY := y + ky
					
					// 检查边界
					if imgX >= bounds.Min.X && imgX < bounds.Max.X &&
						imgY >= bounds.Min.Y && imgY < bounds.Max.Y {
						gray := color.GrayModel.Convert(img.At(imgX, imgY)).(color.Gray)
						if gray.Y > maxVal {
							maxVal = gray.Y
						}
					}
				}
			}
			
			dst.SetGray(x, y, color.Gray{maxVal})
		}
	}
	
	return dst
}

// computeHistogram 计算直方图
func computeHistogram(img image.Image) []int {
	histogram := make([]int, 256)
	
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			// 使用加权平均计算灰度值
			gray := 0.299*float64(r>>8) + 0.587*float64(g>>8) + 0.114*float64(b>>8)
			histogram[int(gray)]++
		}
	}
	
	return histogram
}

// histogramEqualization 直方图均衡化
func histogramEqualization(img image.Image) image.Image {
	bounds := img.Bounds()
	dst := image.NewRGBA(bounds)
	
	// 计算直方图
	hist := computeHistogram(img)
	
	// 计算累积分布函数
	cdf := make([]int, 256)
	cdf[0] = hist[0]
	for i := 1; i < 256; i++ {
		cdf[i] = cdf[i-1] + hist[i]
	}
	
	// 找到第一个非零CDF值
	var cdfMin int
	for i := 0; i < 256; i++ {
		if cdf[i] != 0 {
			cdfMin = cdf[i]
			break
		}
	}
	
	totalPixels := bounds.Dx() * bounds.Dy()
	
	// 应用直方图均衡化
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			
			// 计算灰度值
			gray := 0.299*float64(r>>8) + 0.587*float64(g>>8) + 0.114*float64(b>>8)
			
			// 应用直方图均衡化公式
			newGray := uint8(((cdf[int(gray)] - cdfMin) * 255) / (totalPixels - cdfMin))
			
			// 保持原始颜色比例
			ratio := float64(newGray) / float64(gray+1) // 避免除零
			newR := uint8(clamp(float64(r>>8)*ratio, 0, 255))
			newG := uint8(clamp(float64(g>>8)*ratio, 0, 255))
			newB := uint8(clamp(float64(b>>8)*ratio, 0, 255))
			
			dst.Set(x, y, color.RGBA{newR, newG, newB, uint8(a >> 8)})
		}
	}
	
	return dst
}

// binarize 图像二值化
func binarize(img image.Image, threshold uint8) image.Image {
	bounds := img.Bounds()
	dst := image.NewGray(bounds)
	
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			
			// 计算灰度值
			gray := 0.299*float64(r>>8) + 0.587*float64(g>>8) + 0.114*float64(b>>8)
			
			var newValue uint8
			if uint8(gray) > threshold {
				newValue = 255 // 白色
			} else {
				newValue = 0 // 黑色
			}
			
			dst.SetGray(x, y, color.Gray{newValue})
		}
	}
	
	return dst
}

// adjustBrightness 调整亮度
func adjustBrightness(img image.Image, factor float64) image.Image {
	bounds := img.Bounds()
	dst := image.NewRGBA(bounds)
	
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			
			newR := uint8(clamp(float64(r>>8)*factor, 0, 255))
			newG := uint8(clamp(float64(g>>8)*factor, 0, 255))
			newB := uint8(clamp(float64(b>>8)*factor, 0, 255))
			
			dst.Set(x, y, color.RGBA{newR, newG, newB, uint8(a >> 8)})
		}
	}
	
	return dst
}

// adjustContrast 调整对比度
func adjustContrast(img image.Image, factor float64) image.Image {
	bounds := img.Bounds()
	dst := image.NewRGBA(bounds)
	
	factor = (factor - 1.0) * 128.0
	
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			
			newR := uint8(clamp(float64(r>>8)+factor, 0, 255))
			newG := uint8(clamp(float64(g>>8)+factor, 0, 255))
			newB := uint8(clamp(float64(b>>8)+factor, 0, 255))
			
			dst.Set(x, y, color.RGBA{newR, newG, newB, uint8(a >> 8)})
		}
	}
	
	return dst
}

// 预定义的滤波核
var (
	// 均值滤波核
	meanKernel = [][]float64{
		{1.0 / 9, 1.0 / 9, 1.0 / 9},
		{1.0 / 9, 1.0 / 9, 1.0 / 9},
		{1.0 / 9, 1.0 / 9, 1.0 / 9},
	}

	// 锐化滤波核
	sharpenKernel = [][]float64{
		{-1, -1, -1},
		{-1, 9, -1},
		{-1, -1, -1},
	}

	// 边缘检测滤波核
	edgeKernel = [][]float64{
		{-1, -1, -1},
		{-1, 8, -1},
		{-1, -1, -1},
	}
)

func main() {
	fmt.Println("开始Golang图像处理示例演示...")
	
	// 执行所有示例
	BasicImageOperations()
	GeometricTransformations()
	FilteringAlgorithms()
	ColorSpaceConversions()
	MorphologicalOperations()
	HistogramProcessing()
	AdvancedTechniques()
	
	fmt.Println("\nGolang图像处理示例演示完成!")
}