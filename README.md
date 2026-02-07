# 计算机图像处理教程

这是一个全面的计算机图像处理教程网站，从基础知识到实践应用，涵盖图像处理的核心概念、常见算法以及Python和Golang实现示例。

## 功能特点

- **基础知识**: 详细介绍图像处理的基本概念和理论
- **常见算法**: 涵盖经典的图像处理算法及其原理
- **多语言示例**: 提供Python和Golang图像处理代码示例
- **响应式设计**: 适配各种屏幕尺寸
- **现代UI**: 采用Material Design原则设计
- **完全静态**: 所有资源本地化，无需外部依赖

## 内容结构

1. **基础知识**: 图像表示、色彩模型、质量指标等
2. **常见算法**: 几何变换、滤波、边缘检测、形态学操作等
3. **Python示例**: 使用OpenCV、PIL等库的实现
4. **Golang示例**: 使用标准库和x/image包的实现

## 技术栈

- HTML5, CSS3, JavaScript
- Bootstrap 5 (本地化)
- Prism.js (代码高亮，本地化)
- TensorFlow.js (本地化)

## 如何运行

1. 克隆或下载此项目
2. 下载外部依赖资源（参见 `DOWNLOAD_DEPENDENCIES.md`）
3. 启动本地服务器：
   ```bash
   python3 -m http.server 8000
   ```
4. 在浏览器中访问 `http://localhost:8000`

## 文件结构

```
photo-tutorial/
├── index.html              # 主页面
├── css/
│   └── style.css           # 样式文件
├── js/
│   └── main.js             # 主要JavaScript逻辑
├── libs/                   # 本地依赖库
├── images/                 # 图像资源
├── DOWNLOAD_DEPENDENCIES.md # 依赖下载指南
└── README.md
```

## 依赖资源

由于所有外部资源都已本地化，您需要手动下载以下依赖项：
- Bootstrap CSS 和 JS
- Google Fonts (Roboto)
- Material Icons
- TensorFlow.js
- Prism.js (代码高亮)

详情请参阅 `DOWNLOAD_DEPENDENCIES.md` 文件。

## 浏览器兼容性

- Chrome/Firefox/Safari/Edge (最新版本)

## 版权声明

本项目仅供学习和演示用途。