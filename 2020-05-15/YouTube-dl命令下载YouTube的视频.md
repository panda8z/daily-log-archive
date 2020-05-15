# YouTube-dl 命令下载 YouTube 的视频

## youtube-dl 安装

```bash
# curl 安装
$ sudo curl -L https://yt-dl.org/downloads/latest/youtube-dl -o /usr/local/bin/youtube-dl

# pip 安装
$ pip install youtube-dl

# wget 安装
$ sudo wget https://yt-dl.org/downloads/latest/youtube-dl -O /usr/local/bin/youtube-dl

# Homebrew 安装
$ brew install youtube-dl

# MacPorts 安装
$ sudo port install youtube-dl

# 给执行权限
$ sudo chmod a+rx /usr/local/bin/youtube-dl

# 升级
$ sudo -H pip install --upgrade youtube-dl
```

## youtube-dl 命令格式

```bash
$ youtube-dl [OPTIONS] URL [URL...]
```

## youtube-dl 下载视频

```bash
# 指定文件格式下载
$ youtube-dl --format mp4 https://www.youtube.com/watch?v=IcrbM1l_BoI
```

如果每次都需要下载到 `~/Downloads` 文件夹下，可以设置配置文件如：

```bash
$ mkdir -p ~/.config/youtube-dl
$ cd ~/.config/youtube-dl
$ touch config
$ echo "-o ~/Downloads/%(title)s-%(id)s.%(ext)s" > ~/.config/youtube-dl/config
$ youtube-dl "ytsearch:~/.config/youtube-dl/config"
```

查看视频文件可被下载的格式：

```bash
$ youtube-dl --list-formats https://www.youtube.com/watch?v=IcrbM1l_BoI
# option: -F = --list-formats
$ youtube-dl -F https://www.youtube.com/watch?v=IcrbM1l_BoI
# 得到格式列表之后第一列 format code 代表格式代码
# -f 指定代码格式下载
$ youtube-dl -f 18 https://www.youtube.com/watch\?v=IcrbM1l_BoI
```

如果没有指定 format code，是直接指定如 3gp、aac、flv、m4a、mp3、mp4、ogg、wav、webm 等格式，将会下载最优质的该格式文件：

```bash
# -f 指定后缀格式下载
$ youtube-dl -f mp4 https://www.youtube.com/watch\?v=IcrbM1l_BoI
# -f 还能指定特殊名词来选择特殊的边缘情况格式
# best 带有视频和音频的单个文件代表的最佳质量格式
# worst 带有视频和音频的单个文件代表的质量最差的格式
# bestvideo 选择最佳质量的纯视频格式（例如 DASH 视频）可能不可用
# worstvideo 选择质量最差的纯视频格式 可能不可用
# bestaudio 选择最佳质量的纯音频格式 可能不可用
# worstaudio 选择质量最差的音频纯格式 可能不可用
# -f 22/17/18 下载多个视频，优先级从左到右，如果优先的没找到可下载格式
```

[更多说明参考](https://javascript.ctolib.com/rg3-youtube-dl.html)