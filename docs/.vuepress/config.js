module.exports = {
    title: 'Panda8z`Blog',
    description: 'Panda8z developer engineer designer',
    port: 8011,
    themeConfig: {
        logo: 'https://golang.google.cn/lib/godoc/images/go-logo-blue.svg',
        nav: [
            { text: 'Home', link: '/' },
            { text: 'About', link: '/about.md' },
            {
                text: '文集',
                ariaLabel: 'Kinds',
                items: [
                    { text: 'Rust', link: '/rust/', target: 'self' },
                    { text: 'CMake', link: '/cmake/', target: 'self' },
                    { text: '技术面试', link: '/技术面试/', target: 'self' },
                ]
            },

        ],
        sidebar: {
            '/cmake/': [
                ['./Cmake教程01-入门和第一个项目.md', 'Cmake教程01-入门和第一个项目'],
                ['./Cmake教程02-配置版本号.md', 'Cmake教程02-配置版本号'],
                ['./Cmake教程03-指定C++11编译标准.md', 'Cmake教程03-指定C++11编译标准'],
                ['./Cmake教程04-生成lib或so库.md', 'Cmake教程04-生成lib或so库'],
                ['./Cmake教程05-install打包.md', 'Cmake教程05-install打包'],
            ],
            '/技术面试/': [
                ['./字节面试01-后端.md', '字节面试01-后端'],
                ['./字节面试02-java后端已入职.md', '字节面试02-java后端已入职'],
                ['./字节面试03-后端.md', '字节面试03-后端'],
                ['./字节面试04-后端失败了.md', '字节面试04-后端失败了'],
                ['./字节面试05-后端算法题.md', '字节面试05-后端算法题'],
                ['./技术面试刷题指南.md', '技术面试刷题指南'],
            ],
            '/rust/': [
                ['./2020-08-28-rust-day01-笔记.md', 'rust-day01-笔记']
            ],
            '/': [
                './about.md'
            ],
        }

    },
    configureWebpack: {
        resolve: {
          alias: {
            '@alias': './public/pdfs'
          }
        }
      }
}