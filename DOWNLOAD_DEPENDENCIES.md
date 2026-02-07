# 下载外部依赖资源指南

要使网站完全离线运行，您需要下载以下外部依赖项并放置在 `libs/` 目录中：

## 1. Bootstrap CSS 和 JS

下载地址：
- CSS: https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css
- JS: https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js

## 2. Google Fonts

### Roboto 字体
创建文件 `libs/roboto-font.css`，内容如下：
```css
/* roboto-font.css */
@font-face {
  font-family: 'Roboto';
  font-style: normal;
  font-weight: 300;
  src: url(https://fonts.gstatic.com/s/roboto/v30/KFOlCnqEu92Fr1MmSU5fBBc4AMP6lQ.woff2) format('woff2');
  unicode-range: U+0000-00FF, U+0131, U+0152-0153, U+02BB-02BC, U+02C6, U+02DA, U+02DC, U+2000-206F, U+2074, U+20AC, U+2122, U+2191, U+2193, U+2212, U+2215, U+FEFF, U+FFFD;
}
@font-face {
  font-family: 'Roboto';
  font-style: normal;
  font-weight: 400;
  src: url(https://fonts.gstatic.com/s/roboto/v30/KFOmCnqEu92Fr1Mu72xKOzY.woff2) format('woff2');
  unicode-range: U+0000-00FF, U+0131, U+0152-0153, U+02BB-02BC, U+02C6, U+02DA, U+02DC, U+2000-206F, U+2074, U+20AC, U+2122, U+2191, U+2193, U+2212, U+2215, U+FEFF, U+FFFD;
}
@font-face {
  font-family: 'Roboto';
  font-style: normal;
  font-weight: 500;
  src: url(https://fonts.gstatic.com/s/roboto/v30/KFOlCnqEu92Fr1MmEU9fBBc4AMP6lQ.woff2) format('woff2');
  unicode-range: U+0000-00FF, U+0131, U+0152-0153, U+02BB-02BC, U+02C6, U+02DA, U+02DC, U+2000-206F, U+2074, U+20AC, U+2122, U+2191, U+2193, U+2212, U+2215, U+FEFF, U+FFFD;
}
@font-face {
  font-family: 'Roboto';
  font-style: normal;
  font-weight: 700;
  src: url(https://fonts.gstatic.com/s/roboto/v30/KFOlCnqEu92Fr1MmWUlfBBc4AMP6lQ.woff2) format('woff2');
  unicode-range: U+0000-00FF, U+0131, U+0152-0153, U+02BB-02BC, U+02C6, U+02DA, U+02DC, U+2000-206F, U+2074, U+20AC, U+2122, U+2191, U+2193, U+2212, U+2215, U+FEFF, U+FFFD;
}
```

### Material Icons
创建文件 `libs/material-icons.css`，内容如下：
```css
/* material-icons.css */
@font-face {
  font-family: 'Material Icons';
  font-style: normal;
  font-weight: 400;
  src: url(https://fonts.gstatic.com/s/materialicons/v149/flUhRq6tzZclQEJ-Vdg-IuiaDsNZ.ttf) format('truetype');
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
```

## 3. Prism.js (代码高亮)

下载地址：
- JS: https://cdnjs.cloudflare.com/ajax/libs/prism/1.29.0/components/prism-core.min.js
- CSS: https://cdnjs.cloudflare.com/ajax/libs/prism/1.29.0/themes/prism.min.css

## 4. TensorFlow.js

下载地址：
- JS: https://cdn.jsdelivr.net/npm/@tensorflow/tfjs@latest/dist/tf.min.js

## 完整的 libs 目录结构

```
libs/
├── bootstrap.min.css
├── bootstrap.bundle.min.js
├── roboto-font.css
├── material-icons.css
├── prism.js
├── prism.css
└── tf.min.js
```

## 手动下载步骤

1. 创建 `libs` 目录：
   ```bash
   mkdir -p libs
   ```

2. 使用浏览器或命令行工具下载上述文件到 `libs` 目录

3. 或者运行下载脚本（可能因网络问题而失败）：
   ```bash
   ./download_deps.sh
   ```

下载完成后，网站就可以完全离线运行了。