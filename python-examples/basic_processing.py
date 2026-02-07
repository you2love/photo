"""
Python图像处理基础示例
使用Pillow和OpenCV进行基本图像操作
"""

from PIL import Image, ImageFilter, ImageEnhance
import cv2
import numpy as np
import matplotlib.pyplot as plt


def resize_image(input_path, output_path, size):
    """
    调整图像大小
    :param input_path: 输入图像路径
    :param output_path: 输出图像路径
    :param size: 新尺寸 (width, height)
    """
    with Image.open(input_path) as image:
        resized_image = image.resize(size, Image.Resampling.LANCZOS)
        resized_image.save(output_path)
        print(f"图像已调整大小并保存至: {output_path}")


def apply_filters(input_path, output_path):
    """
    应用多种滤镜效果
    :param input_path: 输入图像路径
    :param output_path: 输出图像路径
    """
    with Image.open(input_path) as image:
        # 应用模糊效果
        blurred = image.filter(ImageFilter.BLUR)
        
        # 应用轮廓效果
        contour = image.filter(ImageFilter.CONTOUR)
        
        # 应用细节增强
        detail = image.filter(ImageFilter.DETAIL)
        
        # 应用边缘增强
        edge_enhance = image.filter(ImageFilter.EDGE_ENHANCE)
        
        # 保存处理后的图像
        blurred.save(output_path.replace('.jpg', '_blurred.jpg'))
        contour.save(output_path.replace('.jpg', '_contour.jpg'))
        detail.save(output_path.replace('.jpg', '_detail.jpg'))
        edge_enhance.save(output_path.replace('.jpg', '_edge_enhance.jpg'))
        
        print("滤镜效果已应用并保存")


def adjust_brightness_contrast(input_path, output_path, brightness_factor=1.0, contrast_factor=1.0):
    """
    调整图像亮度和对比度
    :param input_path: 输入图像路径
    :param output_path: 输出图像路径
    :param brightness_factor: 亮度因子 (1.0为原图)
    :param contrast_factor: 对比度因子 (1.0为原图)
    """
    with Image.open(input_path) as image:
        # 调整亮度
        brightness_enhancer = ImageEnhance.Brightness(image)
        brightened_image = brightness_enhancer.enhance(brightness_factor)
        
        # 调整对比度
        contrast_enhancer = ImageEnhance.Contrast(brightened_image)
        enhanced_image = contrast_enhancer.enhance(contrast_factor)
        
        enhanced_image.save(output_path)
        print(f"亮度和对比度已调整并保存至: {output_path}")


def detect_edges_opencv(input_path, output_path):
    """
    使用OpenCV进行边缘检测
    :param input_path: 输入图像路径
    :param output_path: 输出图像路径
    """
    # 读取图像
    image = cv2.imread(input_path)
    
    # 转换为灰度图
    gray = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)
    
    # 应用高斯模糊以减少噪声
    blurred = cv2.GaussianBlur(gray, (5, 5), 0)
    
    # 应用Canny边缘检测
    edges = cv2.Canny(blurred, threshold1=50, threshold2=150)
    
    # 保存结果
    cv2.imwrite(output_path, edges)
    print(f"边缘检测已完成并保存至: {output_path}")
    
    # 可视化结果
    plt.figure(figsize=(12, 4))
    
    plt.subplot(1, 3, 1)
    plt.title('原图')
    plt.imshow(cv2.cvtColor(image, cv2.COLOR_BGR2RGB))
    plt.axis('off')
    
    plt.subplot(1, 3, 2)
    plt.title('灰度图')
    plt.imshow(gray, cmap='gray')
    plt.axis('off')
    
    plt.subplot(1, 3, 3)
    plt.title('边缘检测结果')
    plt.imshow(edges, cmap='gray')
    plt.axis('off')
    
    plt.tight_layout()
    plt.savefig(output_path.replace('.jpg', '_visualization.jpg'))
    print(f"可视化结果已保存至: {output_path.replace('.jpg', '_visualization.jpg')}")


def color_space_conversion(input_path, output_path):
    """
    颜色空间转换示例
    :param input_path: 输入图像路径
    :param output_path: 输出图像路径
    """
    image = cv2.imread(input_path)
    
    # RGB转HSV
    hsv = cv2.cvtColor(image, cv2.COLOR_BGR2HSV)
    
    # RGB转LAB
    lab = cv2.cvtColor(image, cv2.COLOR_BGR2LAB)
    
    # 保存转换后的图像
    cv2.imwrite(output_path.replace('.jpg', '_hsv.jpg'), hsv)
    cv2.imwrite(output_path.replace('.jpg', '_lab.jpg'), lab)
    
    print("颜色空间转换已完成")


def main():
    """
    主函数 - 演示各种图像处理技术
    """
    input_path = 'sample.jpg'  # 替换为你的输入图像路径
    output_path = 'output.jpg'
    
    print("开始图像处理演示...")
    
    # 调整图像大小
    resize_image(input_path, output_path, (800, 600))
    
    # 应用滤镜
    apply_filters(input_path, output_path)
    
    # 调整亮度和对比度
    adjust_brightness_contrast(input_path, output_path, brightness_factor=1.2, contrast_factor=1.1)
    
    # 边缘检测
    detect_edges_opencv(input_path, output_path.replace('.jpg', '_edges.jpg'))
    
    # 颜色空间转换
    color_space_conversion(input_path, output_path)
    
    print("图像处理演示完成!")


if __name__ == "__main__":
    main()