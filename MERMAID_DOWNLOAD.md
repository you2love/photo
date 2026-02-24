# 外部资源下载说明

## ✅ 所有资源已本地化

本网站所有外部资源已完全本地化，可以离线运行。

## 已本地化的资源

| 资源 | 文件大小 | 位置 |
|------|----------|------|
| Mermaid.js 10.6.1 | 2.8MB | libs/mermaid.min.js |
| Material Icons | 348KB | libs/MaterialIcons-Regular.ttf |
| Bootstrap 5 | 227KB + 79KB | libs/bootstrap.* |
| TensorFlow.js | 1.4MB | libs/tf.min.js |
| Prism.js | 3KB | libs/prism.* |

## 验证安装

运行检查脚本：

```bash
./download_deps.sh
```

如果显示 "✓ 所有资源已就绪！"，说明所有资源都已正确安装。

## 手动下载（如果需要）

### Mermaid.js

```bash
curl -L https://cdnjs.cloudflare.com/ajax/libs/mermaid/10.6.1/mermaid.min.js -o libs/mermaid.min.js
```

### Material Icons 字体

```bash
curl -L https://github.com/google/material-design-icons/raw/master/font/MaterialIcons-Regular.ttf -o libs/MaterialIcons-Regular.ttf
```

## 启动网站

```bash
python3 -m http.server 8000
```

然后访问 http://localhost:8000
