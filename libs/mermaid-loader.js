/**
 * Mermaid.js 本地引用包装脚本
 * 如果本地 mermaid.min.js 不存在，则从 CDN 加载
 */
(function() {
  // 检查是否已经加载了 mermaid
  if (typeof mermaid !== 'undefined') {
    return;
  }

  // 尝试加载本地文件
  var script = document.createElement('script');
  script.src = 'libs/mermaid.min.js';
  script.onerror = function() {
    // 本地加载失败，使用 CDN
    console.log('本地 Mermaid 加载失败，使用 CDN...');
    var cdnScript = document.createElement('script');
    cdnScript.src = 'https://cdn.jsdelivr.net/npm/mermaid@10.6.1/dist/mermaid.min.js';
    document.head.appendChild(cdnScript);
  };
  document.head.appendChild(script);
})();
