// Mermaid 图表初始化和交互功能
document.addEventListener('DOMContentLoaded', function() {
  // 初始化 Mermaid
  if (typeof mermaid !== 'undefined') {
    try {
      mermaid.initialize({
        startOnLoad: true,
        theme: 'default',
        fontFamily: '-apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif',
        flowchart: {
          useMaxWidth: true,
          htmlLabels: true,
          curve: 'basis'
        },
        sequence: {
          useMaxWidth: true,
          boxMargin: 10,
          boxTextMargin: 5,
          noteMargin: 10,
          messageMargin: 35,
          messageAlign: 'center'
        },
        class: {
          useMaxWidth: true
        }
      });
      console.log('Mermaid 初始化成功');
    } catch (e) {
      console.error('Mermaid 初始化失败:', e);
    }
  } else {
    console.error('Mermaid 未加载');
  }

  // 可折叠图表功能
  const collapsibleDiagrams = document.querySelectorAll('.collapsible-diagram .diagram-header');
  collapsibleDiagrams.forEach(header => {
    header.addEventListener('click', function() {
      const body = this.nextElementSibling;
      const icon = this.querySelector('.toggle-icon');
      
      if (body.style.display === 'none') {
        body.style.display = 'block';
        if (icon) icon.textContent = '−';
      } else {
        body.style.display = 'none';
        if (icon) icon.textContent = '+';
      }
    });
  });

  // 图表缩放功能
  const zoomableDiagrams = document.querySelectorAll('.mermaid-container');
  zoomableDiagrams.forEach(container => {
    container.addEventListener('wheel', function(e) {
      if (e.ctrlKey) {
        e.preventDefault();
        const scale = this.style.transform ? parseFloat(this.style.transform.replace('scale(', '').replace(')', '')) : 1;
        const newScale = e.deltaY < 0 ? scale + 0.1 : scale - 0.1;
        this.style.transform = `scale(${Math.max(0.5, Math.min(2, newScale))})`;
        this.style.transformOrigin = 'center center';
      }
    });
  });

  // 图表切换功能
  const diagramSwitches = document.querySelectorAll('.diagram-switch');
  diagramSwitches.forEach(switchEl => {
    switchEl.addEventListener('change', function() {
      const targetId = this.dataset.target;
      const targetEl = document.getElementById(targetId);
      if (targetEl) {
        targetEl.style.display = this.checked ? 'block' : 'none';
      }
    });
  });

  // 高亮表格行功能
  const comparisonTables = document.querySelectorAll('.table-comparison tbody tr');
  comparisonTables.forEach(row => {
    row.addEventListener('mouseenter', function() {
      // 可以添加更多交互逻辑
    });
  });

  // 添加工提示
  const diagramNodes = document.querySelectorAll('.mermaid .node');
  diagramNodes.forEach(node => {
    if (node.dataset.tooltip) {
      node.setAttribute('title', node.dataset.tooltip);
    }
  });
});

// 动态加载 Mermaid
function loadMermaid() {
  return new Promise((resolve, reject) => {
    if (typeof mermaid !== 'undefined') {
      resolve(mermaid);
      return;
    }
    
    const script = document.createElement('script');
    script.src = 'libs/mermaid.min.js';
    script.onload = () => resolve(mermaid);
    script.onerror = reject;
    document.head.appendChild(script);
  });
}

// 导出函数供其他脚本使用
window.diagramUtils = {
  refreshDiagrams: function() {
    if (typeof mermaid !== 'undefined') {
      mermaid.init();
    }
  },
  
  toggleTheme: function(theme) {
    if (typeof mermaid !== 'undefined') {
      mermaid.initialize({ theme: theme });
      this.refreshDiagrams();
    }
  }
};
