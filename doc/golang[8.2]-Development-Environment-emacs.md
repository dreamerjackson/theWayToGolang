# golang快速入门[6.2]-集成开发环境-emacs详解

## 前文
* [golang快速入门[1]-go语言导论](https://zhuanlan.zhihu.com/p/107658283)
* [golang快速入门[2.1]-go语言开发环境配置-windows](https://zhuanlan.zhihu.com/p/107659334)
* [golang快速入门[2.2]-go语言开发环境配置-macOS](https://zhuanlan.zhihu.com/p/107661202)
* [golang快速入门[2.3]-go语言开发环境配置-linux](https://zhuanlan.zhihu.com/p/107662649)
* [golang快速入门[3]-go语言helloworld](https://zhuanlan.zhihu.com/p/107664129)
* [golang快速入门[4]-go语言如何编译为机器码](https://zhuanlan.zhihu.com/p/107665043)
* [golang快速入门[5.1]-go语言是如何运行的-链接器](https://zhuanlan.zhihu.com/p/107665658)
* [golang快速入门[5.2]-go语言是如何运行的-内存概述](https://zhuanlan.zhihu.com/p/107807229)
* [golang快速入门[5.3]-go语言是如何运行的-内存分配](https://zhuanlan.zhihu.com/p/108598942)
* [golang快速入门[6.1]-集成开发环境-goland详解](https://zhuanlan.zhihu.com/p/109564120)

## 前言
* 在上一篇文章中介绍了go语言的集成开发环境`goland`的安装、配置、激活、以及基本的快捷键用法。
* 在本篇文章中，我们将介绍神之编辑器:`emacs` 如何集成go的开发环境
* 你想拥有如下图所示的工作环境吗？这就是`emacs`的强大功能,快来一起学习吧~
![image](../image/golang[8.2]-2.png)
![image](../image/golang[8.2]-1.png)

## emacs是什么
* Emacs（/ˈiːmæks/，源自Editor MACroS，宏编辑器），是一个文本编辑器家族，具有强大的可扩展性，在程序员和其他以技术工作为主的计算机用户中广受欢迎。
* Emacs在1970年代诞生于MIT人工智能实验室（MIT AI Lab）
* Emacs是当前世界上最具可移植性的重要软件之一，能够在当前绝大多数操作系统上运行，包括各种类Unix系统（GNU/Linux、FreeBSD、NetBSD、OpenBSD、Solaris、AIX、OS X等）
* Emacs不仅仅是一个编辑器，它是一个集成环境，或可称它为集成开发环境，这些功能如让用户置身于全功能的操作系统中。Emacs可以：
    + 收发电子邮件、上新闻组（Gnus）
    + 无缝直接编辑远程文件（Tramp）
    + 通过Telnet登录远程主机
    + 操作壳层（M-x EShell，Term）
    + 结合git, mercurial等版本控制系统并直接操作（Magit，VC）
    + 上Twitter（Twittering-mode）
    + 登陆IRC和朋友交流（M-x ERC，rcirc）
    + 电子数据表
    + 模拟其他编辑器，如vi（Evil-mode）、WordStar、EDT、TPU等
    + 编辑Wiki (Wikipedia-mode)
    + 对多种编程语言的编辑，如C/C++、Perl、Python、Lisp等等
    + 调试程序，结合GDB，EDebug等。
    + 玩游戏
    + 计算器
    + 心理咨询（M-x doctor）
    + 煮咖啡
    + 记笔记、日记（Org-mode）
    + 管理日程，Task，待办事项（ToDo），约会等GTD（Org-mode）
    + 写作与出版（Org-mode，Muse-mode）
    + 目录管理（Dired）
    + 文件比较、合并（Ediff）
    + 阅读info和man文档（M-x info，woman）
    + 浏览网站（M-x eww）
    + 为各种程序（TeX、LaTeX等）提供统一的操作界面
    + ……
* 所以有人说，你可以`住`在emacs里面。自诞生以来，Emacs演化出了众多分支，其中使用最广泛的两种分别是：1984年由理查·斯托曼发起并由他维护至2008年的GNU Emacs，以及1991年发起的XEmacs。XEmacs是GNU Emacs的分支，至今仍保持着相当的兼容性。它们都使用了Emacs Lisp这种有着极强扩展性的编程语言，从而实现了包括编程、编译乃至网络浏览等等功能的扩展。本文主要基于GNU Emacs进行讲解

## emacs 与 vim 的对比
* 我相信熟悉linux基本操作的同学对于vim这一款编辑器不会陌生，针对vim与emacs谁更好常常会引发一场论战
* 其实vim与emacs两种编辑器的设计哲学完全不同，风格迥异。vim的特点是组合性(Composability)，Emacs的特点是可扩展性(Extensibility)。vim使用键序列输入，Emacs则经常使用组合键（同时按）输入，跟弹钢琴一样。所以Emacs有个绰号`Esc + Meta + Alt + Ctrl + Shift`

## emacs安装
* 下面我们将分别介绍在windows、mac、linux平台安装`emacs`的方式
* 当以下方法遇到问题时，记得查看[GUN emacs官网](https://www.gnu.org/software/emacs/download.html)的最新介绍

#### windows用户
* 首先登陆GNU镜像下载页面[blog](http://mirror-hk.koddos.net/gnu/emacs/windows/emacs-26/)，下载最新版本为`emacs-26.3-x86_64.zip`
* 解压并重命名 emacs-26.3-x86_64.zip 到所需安装位置。如"D:\emacs-26.3"，后面均以此为例。
* 解压之后，创建一个指向文件`bin/runemacs.exe`的桌面快捷方式，然后双击该快捷方式的图标来启动Emacs
* 配置HOME目录:在注册表中添加`计算机`\HKEY_LOCAL_MACHINE\SOFTWARE\GNU\Emacs`项,为Emacs项添加字符串值.`HOME -> D:\emacs-26.3`
* 添加系统环境变量`D:\emacs-26.3`
#### linux用户
```
>> wget  http://mirror-hk.koddos.net/gnu/emacs/emacs-26.3.tar.gz
```
* 使用tar 命令解压刚刚下载的Go语言开发包到/usr/local目录
```
>> tar -C /usr/local -xzf emacs-26.3.tar.gz
```

* 编译
```
>> cd /usr/local/emacs-26.3
>> ./configure
>> make && make install
```
* 运行
直接在终端输入`emacs`

#### Ubuntu平台(16.04以上版本)
* 安装
```
>> sudo add-apt-repository ppa:kelleyk/emacs
>> sudo apt install emacs26
```
* 删除
```
>> sudo apt remove --autoremove emacs26 emacs26-nox
```
* 运行
直接在终端输入`emacs`
#### mac用户
* 可以使用[HomeBrew](https://brew.sh/)进行安装
```
brew cask install emacs
```
* 或者通过如下网站直接下载：
```
https://emacsformacosx.com/
```
* 运行
直接在终端输入`emacs`

## emacs快捷键
* emacs有一些基本的操作指令,是每一个`emacser`必备的
* 在我们查看emacs的快捷键时，要注意,大写字母C 与 M 代表的含义。在windows与linux中，`C` 表示 Ctrl , `M`表示 Alt. 但是在mac下,`C` 代表的是`Ctrl，`M`表示的是meta键。此键是可以在`系统偏好设置`中配置的。

* 开启
```
emacs
emacs -Q  开启不显示信息
```
* 关闭
```
C-x C-c
```
* 光标操作
```
C-b 向后移动一个字符
C-f 向前移动一个字符
M-b 向后移动一个单词
M-f 向前移动一个单词
C-p 向前移动一行
C-n 向后移动一行
C-a 当前行的开始位置
C-e 当前行的结束
M-a 向前移动一句话
M-e 向后移动一句话
M-< 移动到开始的位置
M-> 移动到结束的位置
注：< > 需要和 shift 按键 一块按
C-v  向后翻页
M-v 向前翻页
第一次 C-l(字符L不是数字1) 把光标所处的位置移动到中间 （center）
第二次 C-l(字符L不是数字1) 把光标所处的位置移动到上边（top）
第三次 C-l(字符L不是数字1) 把光标所处的位置移动到下边（bottom）
```

* 多窗口显示
```
C-x  + 数组（0–9）
C-x 1 取消所有的窗口，只保留一个原始窗口
C-x 2 垂直切割当前窗口，分成上下俩个窗口
C-x 3 水平切割当前窗口，分成左右俩个窗口
C-M v 下一个窗口翻页
C-x o 光标移动到下一个窗口
注:操作可叠加
```

* 删除和剪切
```
C-d  删除一个字符
M-d  剪切一个单词
C-k  剪切光标和该行末尾
M-k  剪切光标和标点符号
C- spaces(空格键)  先标记的地方为start  第二次标记的地方为end
C-w  剪切start 和 end 中间的内容
```

* 粘贴和复制
```
C-y  粘贴最近的内容
M-y  在C-y之后使用,往上追溯替换粘贴的内容
```
* 查找
```
C-s 向后查找
C-r 向前查找
```
* 打开、新建、保存、关闭文件
```
C-x C-f 打开文件，文件不存在则新建
C-x C-s 保存当前文件
C-x C-b 列出buffer文件列表
C-x b 切换buffer文件 （通过在面板最下面输入文件名字）
C-x C-c 关闭并且保存文件
```
* 撤销 & 反撤销
```
C-x u 撤销
C-- 撤销
C-/ 撤销
在撤销的时候如果撤销多了需要进行反撤销
具体是用C-f等打断当前撤销操作，接着进行的撤销动作就是反撤销。
C-x z 重复之前的操作 重复多次可以只按zzzz
```

## emacs 配置文件
* emacs 配置文件默认位于`Home`目录下(~/)，可以通过变量`user-emacs-directory`修改
* 配置文件夹可以是`Home`目录下的`.emacs.d`文件，通常会将所有配置放入其中
* 当启动Emacs时，通常会尝试从初始化文件加载Lisp程序。该文件（如果存在）指定如何为您初始化Emacs。初始文件为 `~/.emacs, ~/.emacs.el, or  ~/.emacs.d/init.el`  中的一个。
* 对于初学者，可以去查找网上一些有名的配置直接使用，大牛一般都是一个大的`.emacs.d`文件。如果想获取笔者对于emacs的配置，可以在后台留言。

## emacs 基本配置
* 添加官方与国内package的源
```

(setq package-archives '(("gnu"   . "http://elpa.emacs-china.org/gnu/")
                         ("melpa-stable" . "http://elpa.emacs-china.org/melpa-stable/")
                          ("melpa-stable2" . "https://stable.melpa.org/packages/")
                          ("melpa" . "http://elpa.emacs-china.org/melpa/")
                         ("marmalada" . "http://elpa.emacs-china.org/marmalade/")))
```
* 对于大量emacs 软件包的配置、管理、更新。我推荐使用[use-package](https://github.com/jwiegley/use-package)来管理包的配置与加载。
```
;;
;; use use-package
;;
(unless (package-installed-p 'use-package)
  (package-refresh-contents)
  (package-install 'use-package))
```

* 使用`ivy-mode`拓展套件完成快速搜索、快速查找、智能补全功能
```
;;
;; ivy mode
;;
(use-package ivy
  :ensure t
  :diminish (ivy-mode . "")
  :config
  (ivy-mode 1)
  (setq ivy-use-virutal-buffers t)
  (setq enable-recursive-minibuffers t)
  (setq ivy-height 10)
  (setq ivy-initial-inputs-alist nil)
  (setq ivy-count-format "%d/%d")
  (setq ivy-re-builders-alist
        `((t . ivy--regex-ignore-order)))
  )

;;
;; counsel
;;
(use-package counsel
  :ensure t
  :bind (("M-x" . counsel-M-x)
         ("\C-x \C-f" . counsel-find-file)))

;;
;; swiper
;;
(use-package swiper
  :ensure t
  :bind (("\C-s" . swiper))
  )
```

## emacs配置go开发环境
#### gomode + goimport自动导入 + godef跳转
* 首先安装包godef 与 goimports
```
go get -u github.com/rogpeppe/godef
go get -u golang.org/x/tools/cmd/goimports
```
* 配置
* `M-n` 查找函数应用，等价于（`M-x lsp-find-ref`）
* `M-.` 查找函数实现
```
(use-package go-mode
  ;; :load-path "~/.emacs.d/vendor/go-mode"
  :mode ("\\.go\\'" . go-mode)
  :ensure-system-package
  ((goimports . "go get -u golang.org/x/tools/cmd/goimports")
   (godef . "go get -u github.com/rogpeppe/godef"))
  :init
  (setq gofmt-command "goimports"
        indent-tabs-mode t)
  :config
  (add-hook 'before-save-hook 'gofmt-before-save)
  :bind (:map go-mode-map
              ("\C-c \C-c" . compile)
              ("\C-c \C-g" . go-goto-imports)
              ("\C-c \C-k" . godoc)
              ("M-j" . godef-jump)))
```

#### gocode+company 实现代码自动补全
* 安装gocode
```
go get -u github.com/stamblerre/gocode
```
gocode 是守护进程，查看是否在后台运行
```
ps -e | grep gocode
```
* 配置
```
;;
;; company
;;
(use-package company
  :ensure t
  :config
  (global-company-mode t)
  (setq company-idle-delay 0)
  (setq company-minimum-prefix-length 3)
  (setq company-backends
        '((company-files
           company-yasnippet
           company-keywords
           company-capf
           )
          (company-abbrev company-dabbrev))))

(add-hook 'emacs-lisp-mode-hook (lambda ()
                                  (add-to-list  (make-local-variable 'company-backends)
                                                '(company-elisp))))
```
#### flycheck 语法检查
* 配置
* `C-c ! l  , C-c ! v` 查看是否有语法错误
```
(use-package flycheck
  :ensure t
  :config
  (global-flycheck-mode t)
  )
```

#### gotest 进行各种测试
```
(use-package gotest
  :after go-mode
  :bind (:map go-mode-map
              ("C-c C-f" . go-test-current-file)
              ("C-c C-t" . go-test-current-test)
              ("C-c C-p" . go-test-current-project)
              ("C-c C-b" . go-test-current-benchmark)
              ("C-x x" . go-run))
  :config
  (setq go-test-verbose t))
```
#### errorcheck 进行程序错误检查
* 安装
```
go get -u github.com/kisielk/errcheck
```
* 配置
```
(use-package go-errcheck
  :after go-mode
  :ensure-system-package (errcheck . "go get -u github.com/kisielk/errcheck")
  :bind (:map go-mode-map
              ("C-c C-e" . go-errcheck)))
```
#### gtags 查找项目中的变量、函数等
* 安装源代码`GUN global`, 注意mac用户不要用brew安装，因为必须要配置sqlite3
```
wget http://tamacom.com/global/global-6.5.7.tar.gz
tar xvf global-6.5.7.tar.gz
cd global-6.5.7
./configure --with-sqlite3
make
sudo make install
```
* 安装`gtags`生成器
```
go get github.com/juntaki/gogtags
```
* 在项目目录中输入如下,会生成GTAGS等文件
```
gogtags -v
```
* 配置helm-gtags
* 开启helm-gtags-mode
```
M-x helm-gtags-mode
```
* 这时候在项目中即可使用`M-x helm-gtags find partern` 等查找tags
```
(use-package helm-gtags
  :config
  (setq helm-gtags-ignore-case t
        helm-gtags-auto-update t
        helm-gtags-use-input-at-cursor t
        helm-gtags-pulse-at-cursor t
        helm-gtags-prefix-key "\C-cg"
        helm-gtags-suggested-key-mapping t)
  :bind (:map helm-gtags-mode-map
              ("C-c g a" . helm-gtags-tags-in-this-function)
              ("C-j" . helm-gtags-select)
              ("M-." . helm-gtags-dwim)
              ("M-," . helm-gtags-pop-stack)
              ("C-c <" . helm-gtags-previous-history)
              ("C-c >" . helm-gtags-next-history))
  :hook ((dired-mode eshell-mode c-mode c++-mode asm-mode) . helm-gtags-mode))
```

#### projectile 在go项目中切换
* 配置
```
(use-package dumb-jump
  :bind (("M-g o" . dumb-jump-go-other-window)
         ("M-g j" . dumb-jump-go)
         ("M-g x" . dumb-jump-go-prefer-external)
         ("M-g z" . dumb-jump-go-prefer-external-other-window))
  :config
  ;; (setq dumb-jump-selector 'ivy) ;; (setq dumb-jump-selector 'helm)
:initny
(dumb-jump-mode)
  :ensure
)
```

## 总结
* 在本文中,我们介绍了神之编辑器`emacs`的安装、配置特别是对于go语言集成环境的配置。
* 本文使用了use-package来管理emacs的软件包，并详细介绍了go语言集成环境的配置细节和使用方法。
* 遗憾的是,本文不会带领大家从一个初学者完全入门emacs,用好emacs是值得一生努力的话题。学会emacs需要耐心，练习，更需要交流与指导
* 如果你有更好的配置建议，或者你在使用emacs-go的过程中有任何疑问，欢迎在后台留言,see you~

## 参考资料
* [项目链接](https://github.com/dreamerjackson/theWayToGolang)
* [作者知乎](https://www.zhihu.com/people/ke-ai-de-xiao-tu-ji-71)
* [blog](https://dreamerjonson.com/)
* [DotEmacsDotD](https://www.emacswiki.org/emacs/DotEmacsDotD)
* [emacs-helm-gtags](https://github.com/syohex/emacs-helm-gtags)
* [用Emacs来写Go设定篇](https://medium.com/@jerryhsieh/emacs-21-%E7%94%A8-emacs-%E4%BE%86%E5%AF%AB-go-%E8%A8%AD%E5%AE%9A%E7%AF%87-ce0e09f73c70)
