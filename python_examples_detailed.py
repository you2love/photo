"""
Python图像处理综合示例
实现图像处理基础知识和常见算法
"""

import cv2
import numpy as np
from PIL import Image, ImageFilter, ImageEnhance
import matplotlib.pyplot as plt
from scipy import ndimage
from skimage import filters, feature, morphology
import math


def basic_image_operations():
    """
    基础图像操作示例
    """
    print("=== 基础图像操作 ===")
    
    # 读取图像
    img = cv2.imread('sample.jpg')
    if img is None:
        print("无法读取图像，跳过此部分")
        return
    
    print(f"图像尺寸: {img.shape}")
    print(f"图像数据类型: {img.dtype}")
    
    # 图像基本信息
    height, width, channels = img.shape
    print(f"高度: {height}, 宽度: {width}, 通道数: {channels}")


def geometric_transformations():
    """
    几何变换示例
    """
    print("\n=== 几何变换 ===")
    
    img = cv2.imread('sample.jpg')
    if img is None:
        print("无法读取图像，跳过此部分")
        return
    
    height, width = img.shape[:2]
    
    # 缩放
    scaled_img = cv2.resize(img, None, fx=0.5, fy=0.5, interpolation=cv2.INTER_LINEAR)
    
    # 旋转
    rotation_matrix = cv2.getRotationMatrix2D((width/2, height/2), 45, 1)
    rotated_img = cv2.warpAffine(img, rotation_matrix, (width, height))
    
    # 平移
    translation_matrix = np.float32([[1, 0, 50], [0, 1, 50]])
    translated_img = cv2.warpAffine(img, translation_matrix, (width, height))
    
    print("几何变换完成")


def filtering_algorithms():
    """
    滤波算法示例
    """
    print("\n=== 滤波算法 ===")
    
    img = cv2.imread('sample.jpg', 0)  # 读取为灰度图
    if img is None:
        print("无法读取图像，跳过此部分")
        return
    
    # 均值滤波
    mean_filtered = cv2.blur(img, (5, 5))
    
    # 高斯滤波
    gaussian_filtered = cv2.GaussianBlur(img, (5, 5), 0)
    
    # 中值滤波
    median_filtered = cv2.medianBlur(img, 5)
    
    # 双边滤波
    bilateral_filtered = cv2.bilateralFilter(img, 9, 75, 75)
    
    print("滤波算法完成")


def edge_detection_algorithms():
    """
    边缘检测算法示例
    """
    print("\n=== 边缘检测算法 ===")
    
    img = cv2.imread('sample.jpg', 0)  # 读取为灰度图
    if img is None:
        print("无法读取图像，跳过此部分")
        return
    
    # Sobel边缘检测
    sobel_x = cv2.Sobel(img, cv2.CV_64F, 1, 0, ksize=3)
    sobel_y = cv2.Sobel(img, cv2.CV_64F, 0, 1, ksize=3)
    sobel_combined = np.sqrt(sobel_x**2 + sobel_y**2)
    
    # Canny边缘检测
    canny_edges = cv2.Canny(img, threshold1=50, threshold2=150)
    
    # Laplacian边缘检测
    laplacian = cv2.Laplacian(img, cv2.CV_64F)
    
    print("边缘检测算法完成")


def morphological_operations():
    """
    形态学操作示例
    """
    print("\n=== 形态学操作 ===")
    
    # 创建一个二值图像用于演示
    img = np.zeros((100, 100), dtype=np.uint8)
    img[30:70, 30:70] = 255  # 白色正方形
    img[40:60, 40:60] = 0     # 黑色正方形（孔洞）
    
    # 定义结构元素
    kernel = np.ones((5, 5), np.uint8)
    
    # 腐蚀
    erosion = cv2.erode(img, kernel, iterations=1)
    
    # 膨胀
    dilation = cv2.dilate(img, kernel, iterations=1)
    
    # 开运算（先腐蚀后膨胀）
    opening = cv2.morphologyEx(img, cv2.MORPH_OPEN, kernel)
    
    # 闭运算（先膨胀后腐蚀）
    closing = cv2.morphologyEx(img, cv2.MORPH_CLOSE, kernel)
    
    print("形态学操作完成")


def frequency_domain_processing():
    """
    频域处理示例
    """
    print("\n=== 频域处理 ===")
    
    img = cv2.imread('sample.jpg', 0)  # 读取为灰度图
    if img is None:
        print("无法读取图像，跳过此部分")
        return
    
    # 快速傅里叶变换
    f_transform = np.fft.fft2(img)
    f_shift = np.fft.fftshift(f_transform)
    
    # 计算幅度谱
    magnitude_spectrum = 20 * np.log(np.abs(f_shift) + 1)
    
    print("频域处理完成")


def feature_extraction():
    """
    特征提取示例
    """
    print("\n=== 特征提取 ===")
    
    img = cv2.imread('sample.jpg', 0)  # 读取为灰度图
    if img is None:
        print("无法读取图像，跳过此部分")
        return
    
    # 使用skimage进行特征提取
    # Canny边缘检测
    edges = feature.canny(img, sigma=1)
    
    # 角点检测
    coords = feature.corner_peaks(feature.corner_harris(img), min_distance=5)
    
    print(f"检测到 {len(coords)} 个角点")
    print("特征提取完成")


def color_space_conversions():
    """
    色彩空间转换示例
    """
    print("\n=== 色彩空间转换 ===")
    
    img = cv2.imread('sample.jpg')
    if img is None:
        print("无法读取图像，跳过此部分")
        return
    
    # BGR转RGB
    img_rgb = cv2.cvtColor(img, cv2.COLOR_BGR2RGB)
    
    # BGR转HSV
    img_hsv = cv2.cvtColor(img, cv2.COLOR_BGR2HSV)
    
    # BGR转GRAY
    img_gray = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)
    
    # BGR转LAB
    img_lab = cv2.cvtColor(img, cv2.COLOR_BGR2LAB)
    
    print("色彩空间转换完成")


def histogram_processing():
    """
    直方图处理示例
    """
    print("\n=== 直方图处理 ===")
    
    img = cv2.imread('sample.jpg', 0)  # 读取为灰度图
    if img is None:
        print("无法读取图像，跳过此部分")
        return
    
    # 计算直方图
    hist, bins = np.histogram(img.flatten(), 256, [0, 256])
    
    # 直方图均衡化
    equalized = cv2.equalizeHist(img)
    
    print("直方图处理完成")


def advanced_techniques():
    """
    高级图像处理技术示例
    """
    print("\n=== 高级图像处理技术 ===")
    
    img = cv2.imread('sample.jpg', 0)  # 读取为灰度图
    if img is None:
        print("无法读取图像，跳过此部分")
        return
    
    # 图像阈值处理
    ret, thresh = cv2.threshold(img, 127, 255, cv2.THRESH_BINARY)
    
    # 自适应阈值
    adaptive_thresh = cv2.adaptiveThreshold(img, 255, cv2.ADAPTIVE_THRESH_GAUSSIAN_C, cv2.THRESH_BINARY, 11, 2)
    
    # 轮廓检测
    contours, hierarchy = cv2.findContours(thresh, cv2.RETR_TREE, cv2.CHAIN_APPROX_SIMPLE)
    print(f"检测到 {len(contours)} 个轮廓")
    
    print("高级图像处理技术完成")


def main():
    """
    主函数 - 演示所有图像处理技术
    """
    print("开始Python图像处理示例演示...")
    
    # 执行所有示例
    basic_image_operations()
    geometric_transformations()
    filtering_algorithms()
    edge_detection_algorithms()
    morphological_operations()
    frequency_domain_processing()
    feature_extraction()
    color_space_conversions()
    histogram_processing()
    advanced_techniques()
    
    print("\nPython图像处理示例演示完成!")


if __name__ == "__main__":
    main()