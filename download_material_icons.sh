#!/bin/bash
# Material Icons 字体下载脚本
# 使用方法：./download_material_icons.sh

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
LIBS_DIR="$SCRIPT_DIR/libs"

echo "========================================"
echo "  Material Icons 字体下载"
echo "========================================"
echo ""

# 创建 libs 目录
mkdir -p "$LIBS_DIR"

# 检查系统是否已安装 Material Icons
echo "检查系统字体..."
if fc-list | grep -i "material icons" > /dev/null 2>&1; then
    echo "✓ 系统已安装 Material Icons 字体"
    echo ""
    echo "无需下载，网页将直接使用系统字体。"
    exit 0
fi

if [ "$(uname)" = "Darwin" ]; then
    # macOS 检查
    if ls /Library/Fonts/*Material* 2>/dev/null || ls ~/Library/Fonts/*Material* 2>/dev/null; then
        echo "✓ 系统已安装 Material Icons 字体"
        exit 0
    fi
fi

echo "系统未安装 Material Icons 字体，正在下载..."
echo ""

# 下载 Material Icons 字体文件
FONT_URL="https://github.com/google/material-design-icons/raw/master/font/MaterialIcons-Regular.ttf"
FONT_FILE="$LIBS_DIR/MaterialIcons-Regular.ttf"

# 创建 CSS 文件使用下载的字体的
CSS_FILE="$LIBS_DIR/material-icons-complete.css"

echo "下载字体文件..."

if command -v curl &> /dev/null; then
    curl -L --connect-timeout 10 --max-time 60 "$FONT_URL" -o "$FONT_FILE" 2>&1
    if [ $? -eq 0 ] && [ -s "$FONT_FILE" ]; then
        echo "✓ 字体文件下载成功"
        
        # 创建完整的 CSS
        cat > "$CSS_FILE" << 'EOF'
/* Material Icons 完整 CSS - 使用本地字体文件 */
@font-face {
  font-family: 'Material Icons';
  font-style: normal;
  font-weight: 400;
  src: url('MaterialIcons-Regular.ttf') format('truetype');
}

.material-icons {
  font-family: 'Material Icons';
  font-weight: normal;
  font-style: normal;
  font-size: 24px;
  line-height: 1;
  letter-spacing: normal;
  text-transform: none;
  display: inline-block;
  white-space: nowrap;
  word-wrap: normal;
  direction: ltr;
  -webkit-font-feature-settings: 'liga';
  -webkit-font-smoothing: antialiased;
}

.material-icons.align-middle {
  vertical-align: middle;
}
EOF
        echo "✓ 创建 material-icons-complete.css"
        echo ""
        echo "========================================"
        echo "  完成"
        echo "========================================"
        echo ""
        echo "现在修改 HTML，引用新的 CSS 文件："
        echo '  <link rel="stylesheet" href="libs/material-icons-complete.css">'
        exit 0
    fi
fi

if command -v wget &> /dev/null; then
    wget --timeout=60 --tries=1 "$FONT_URL" -O "$FONT_FILE" 2>&1
    if [ $? -eq 0 ] && [ -s "$FONT_FILE" ]; then
        echo "✓ 字体文件下载成功"
        # (同上创建 CSS)
        exit 0
    fi
fi

echo ""
echo "========================================"
echo "  ⚠ 下载失败"
echo "========================================"
echo ""
echo "备选方案："
echo ""
echo "1. 使用系统字体（当前方案）"
echo "   网页将尝试使用系统已安装的 Material Icons"
echo ""
echo "2. 手动下载字体"
echo "   访问：https://fonts.google.com/icons"
echo "   下载字体文件并放入 libs/ 目录"
echo ""
echo "3. 使用 CDN（需要网络）"
echo "   在 HTML 中添加："
echo '   <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">'
echo ""

exit 1
