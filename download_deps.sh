#!/bin/bash

# 下载依赖资源脚本（带重试机制）
# 请在项目根目录下运行此脚本

echo "正在下载依赖资源..."

# 创建libs目录（如果不存在）
mkdir -p libs

# 定义下载函数（带重试）
download_with_retry() {
    local url=$1
    local output=$2
    local max_attempts=3
    local attempt=1

    while [ $attempt -le $max_attempts ]; do
        echo "尝试下载 $output (第 $attempt/$max_attempts 次)..."
        if curl -L --connect-timeout 30 --max-time 300 -o "$output" "$url"; then
            echo "✓ $output 下载成功"
            return 0
        else
            echo "✗ $output 下载失败，正在重试..."
            rm -f "$output"  # 删除可能损坏的文件
            attempt=$((attempt + 1))
            sleep 5
        fi
    done

    echo "✗ $output 下载失败，已达到最大重试次数"
    return 1
}

# 下载Bootstrap
echo "下载Bootstrap..."
download_with_retry "https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" "libs/bootstrap.min.css"
download_with_retry "https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js" "libs/bootstrap.bundle.min.js"

# 下载Prism.js
echo "下载Prism.js..."
download_with_retry "https://cdnjs.cloudflare.com/ajax.com/ajax/libs/prism/1.29.0/components/prism-core.min.js" "libs/prism.js"
download_with_retry "https://cdnjs.cloudflare.com/ajax/libs/prism/1.29.0/themes/prism.min.css" "libs/prism.css"

# 下载TensorFlow.js
echo "下载TensorFlow.js..."
download_with_retry "https://cdn.jsdelivr.net/npm/@tensorflow/tfjs@latest/dist/tf.min.js" "libs/tf.min.js"

echo ""
echo "下载完成！检查libs目录中的文件："
ls -la libs/

echo ""
echo "注意：Google Fonts和Material Icons需要特殊处理，因为它们可能包含跨域限制。"
echo "请手动创建以下文件："
echo ""
echo "1. libs/roboto-font.css - 包含Roboto字体定义"
echo "2. libs/material-icons.css - 包含Material Icons定义"
echo ""
echo "参考 DOWNLOAD_DEPENDENCIES.md 文件获取详细说明。"