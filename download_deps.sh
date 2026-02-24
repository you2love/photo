#!/bin/bash
# 外部资源下载脚本 - 完整版
# 使用方法：./download_deps.sh

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
LIBS_DIR="$SCRIPT_DIR/libs"

echo "========================================"
echo "  外部资源下载脚本"
echo "========================================"
echo ""

mkdir -p "$LIBS_DIR"

SUCCESS_COUNT=0
TOTAL=2

# ================================
# 1. 检查 Material Icons 字体
# ================================
echo "1. 检查 Material Icons 字体..."

FONT_FILE="$LIBS_DIR/MaterialIcons-Regular.ttf"
if [ -f "$FONT_FILE" ]; then
    FILE_SIZE=$(stat -f%z "$FONT_FILE" 2>/dev/null || stat -c%s "$FONT_FILE" 2>/dev/null || echo "0")
    if [ "$FILE_SIZE" -gt 100000 ]; then
        echo "   ✓ Material Icons 字体已存在"
        ((SUCCESS_COUNT++))
    else
        echo "   ⚠ Material Icons 字体文件不完整"
    fi
else
    echo "   ⚠ Material Icons 字体不存在"
fi

# ================================
# 2. 检查 Mermaid.js
# ================================
echo ""
echo "2. 检查 Mermaid.js..."

MERMAID_FILE="$LIBS_DIR/mermaid.min.js"
if [ -f "$MERMAID_FILE" ]; then
    FILE_SIZE=$(stat -f%z "$MERMAID_FILE" 2>/dev/null || stat -c%s "$MERMAID_FILE" 2>/dev/null || echo "0")
    if [ "$FILE_SIZE" -gt 100000 ]; then
        echo "   ✓ Mermaid.js 已存在"
        ((SUCCESS_COUNT++))
    else
        echo "   ⚠ Mermaid.js 文件不完整"
    fi
else
    echo "   ⚠ Mermaid.js 不存在"
fi

# ================================
# 总结
# ================================
echo ""
echo "========================================"
echo "  检查结果：$SUCCESS_COUNT/$TOTAL"
echo "========================================"
echo ""

if [ $SUCCESS_COUNT -eq $TOTAL ]; then
    echo "✓ 所有资源已就绪！网站可以完全离线运行。"
    echo ""
    echo "启动服务器：python3 -m http.server 8000"
    exit 0
else
    echo "⚠ 部分资源缺失"
    echo ""
    if [ ! -f "$FONT_FILE" ]; then
        echo "下载 Material Icons:"
        echo "  curl -L https://github.com/google/material-design-icons/raw/master/font/MaterialIcons-Regular.ttf -o libs/MaterialIcons-Regular.ttf"
        echo ""
    fi
    if [ ! -f "$MERMAID_FILE" ]; then
        echo "下载 Mermaid.js:"
        echo "  curl -L https://cdnjs.cloudflare.com/ajax/libs/mermaid/10.6.1/mermaid.min.js -o libs/mermaid.min.js"
        echo ""
    fi
    exit 1
fi
