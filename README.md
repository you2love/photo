# 计算机图像处理教程

这是一个全面的计算机图像处理教程网站，从基础知识到实践应用，涵盖图像处理的核心概念、常见算法以及 Python 和 Golang 实现示例。

## 功能特点

- **首页**: 详细介绍图像处理的基本概念和理论，配有流程图和对比表格
- **常见算法**: 涵盖经典的图像处理算法及其原理，使用交互式图表展示
- **算法详解**: 每个算法都有详细的流程图、序列图和代码实现
- **多语言示例**: 提供 Python 和 Golang 图像处理代码示例
- **响应式设计**: 适配各种屏幕尺寸
- **现代 UI**: 采用 Material Design 原则设计
- **Mermaid 图表**: 使用 Mermaid.js 展示流程图、序列图和分类图

## 内容结构

1. **首页 (`index.html`)**: 图像表示、色彩模型、质量指标、文件格式等
2. **常见算法 (`algorithms.html`)**: 几何变换、滤波、边缘检测、形态学等
3. **算法详解 (`algorithms_detail.html`)**: 算法分类导航和学习路径
4. **Python 示例**: 使用 OpenCV、PIL 等库的实现
5. **Golang 示例**: 使用标准库和 x/image 包的实现

## 技术栈

- HTML5, CSS3, JavaScript
- Bootstrap 5 (本地化)
- Prism.js (代码高亮，本地化)
- **Mermaid.js 10.6.1** (图表库，需下载)
- TensorFlow.js (本地化)
- Material Icons (本地 + CDN 回退)
- Roboto 字体 (本地 + CDN 回退)

## 快速开始

### 1. 下载 Mermaid.js（必需）

Mermaid.js 用于渲染流程图和图表，需要单独下载：

```bash
# 方法一：使用下载脚本
./download_mermaid.sh

# 方法二：使用 curl
curl -L https://cdn.jsdelivr.net/npm/mermaid@10.6.1/dist/mermaid.min.js -o libs/mermaid.min.js

# 方法三：使用 npm
npm install mermaid@10.6.1
cp node_modules/mermaid/dist/mermaid.min.js libs/
```

### 2. 启动本地服务器

```bash
# Python 3
python3 -m http.server 8000

# 或使用其他静态服务器
```

### 3. 访问网站

在浏览器中访问 `http://localhost:8000`

## 文件结构

```
photo-tutorial/
├── index.html                    # 首页 - 基础知识
├── algorithms.html               # 常见算法概览
├── algorithms_detail.html        # 算法详解导航
├── python-examples.html          # Python 示例
├── golang-examples.html          # Golang 示例
├── css/
│   ├── style.css                 # 主样式文件
│   └── diagram.css               # Mermaid 图表样式
├── js/
│   ├── main.js                   # 主要 JavaScript 逻辑
│   └── diagram.js                # 图表初始化和交互
├── libs/                         # 本地依赖库
│   ├── bootstrap.min.css
│   ├── bootstrap.bundle.min.js
│   ├── prism.css
│   ├── prism.js
│   ├── tf.min.js
│   ├── mermaid.min.js            # 需手动下载
│   ├── material-icons.css
│   └── roboto-font.css
├── algorithms/                   # 算法详解页面
│   ├── edge/
│   │   ├── canny.html
│   │   └── ...
│   ├── filtering/
│   ├── geometric/
│   ├── line_detection/
│   │   └── hough_transform.html
│   ├── morphology/
│   └── thresholding/
├── images/                       # 图像资源
├── python-examples/              # Python 代码示例
├── golang-examples/              # Golang 代码示例
├── download_mermaid.sh           # Mermaid 下载脚本
├── MERMAID_DOWNLOAD.md           # Mermaid 下载说明
├── DOWNLOAD_DEPENDENCIES.md      # 其他依赖下载指南
└── README.md                     # 本文件
```

## 依赖资源状态

| 依赖项 | 状态 | 文件大小 | 说明 |
|--------|------|----------|------|
| Bootstrap 5 | ✅ 已本地化 | 227KB CSS + 79KB JS | 完整本地化 |
| Prism.js | ✅ 已本地化 | 3KB | 代码高亮库 |
| TensorFlow.js | ✅ 已本地化 | 1.4MB | 机器学习库 |
| **Mermaid.js 10.6.1** | ✅ 已本地化 | 2.8MB | 图表渲染库 |
| **Material Icons** | ✅ 已本地化 | 348KB | 字体文件：libs/MaterialIcons-Regular.ttf |
| Roboto 字体 | ✅ 系统字体 | - | 使用系统字体栈回退 |

## 页面重构说明

本次重构添加了以下可视化元素：

### 图表类型
- **流程图 (Flowchart)**: 展示算法流程、分类关系
- **序列图 (Sequence)**: 展示算法步骤和数据流
- **饼图 (Pie)**: 展示统计分布
- **对比表格**: 展示算法特性、参数、性能对比

### 重构页面
1. **index.html**: 添加图像处理流程图、色彩模型转换图、采样量化序列图等
2. **algorithms.html**: 添加算法分类体系图、各算法流程图和对比表
3. **algorithms_detail.html**: 添加算法导航流程图、学习路径图
4. **canny.html**: 添加算法流程图、步骤序列图、折叠面板
5. **hough_transform.html**: 添加点线对偶原理图、投票流程图、参数空间示意图

## 浏览器兼容性

- Chrome/Firefox/Safari/Edge (最新版本)
- 需要启用 JavaScript 以支持图表渲染

## 注意事项

1. **Mermaid.js 是必需的**: 下载 mermaid.min.js 之前，图表将无法显示
2. **CDN 回退**: Material Icons 和 Roboto 字体有 CDN 回退，需要网络连接
3. **静态网站**: 所有资源本地化后可完全离线使用（除 CDN 回退部分）

## 版权声明

本项目仅供学习和演示用途。
