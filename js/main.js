// 页面加载完成后初始化
document.addEventListener('DOMContentLoaded', function() {
    // 初始化代码高亮
    Prism.highlightAll();
});

// 滚动到指定区域
function scrollToSection(sectionId) {
    const element = document.getElementById(sectionId);
    if (element) {
        element.scrollIntoView({ behavior: 'smooth' });
    }
}

// 添加一些动画效果
window.addEventListener('scroll', function() {
    // 检查元素是否进入视口并添加动画类
    const elements = document.querySelectorAll('.card');
    elements.forEach(element => {
        const rect = element.getBoundingClientRect();
        const windowHeight = window.innerHeight;
        
        if (rect.top < windowHeight * 0.8) {
            element.classList.add('fade-in-up');
        }
    });
});