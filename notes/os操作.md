# <center>os操作</center>

## __打开创建文件__
- 创建文件 Create: 文件不存在创建,文件存在清空
- 打开文件 Open: 以只读方式打开文件
- 打开文件 OpenFile: 以只读,只写,只读的方式打开文件
    - 参数1: name 打开文件路径,绝对相对都可
    - 参数2:打开文件权限  os.O_RDONLY os.WRONLY os.RDWR
    - 参数3:一般传6 对目录来说一般是os.ModeDir
    - 返回值 返回一个可以读写目录的指针

## __写文件__
- 按字符串写 WriteString()
- 按位置写 f.Seek() 修改文件的读写指针位置
    - 参数1:偏移量(正数->向文件尾,负数->向文件头)
    - 参数2:偏移起始位置
        - io.SeekStart:文件起始位置
        - ioSeekCurrent:文件当前位置
        - ioSeekEnd:文件结尾位置
    - 返回值(ret):表示从文件开始到当前文件读写位置的偏移量
- 按字节写
    - WriteAt([]byte("xxxxxx"), offset)
    

## __读文件__
- 按行读
    - 创建一个带有缓冲区的Reader(读写器)
        - bufio.NewReader(打开的文件指针)
        - 从reader缓冲区中读取指定长度的数据,数据长度取决于参数dlime
        - 文件结束标记是要单独读一次读出来的
- 按字节读(read和write)
    - read
    - write 

## __目录__
- 打开目录 OpenFile

## __缓冲区__
- 内存中的一块区域,用来减少物理磁盘访问操作

## __<font color="#006666">标准库io包</font>__
- Writer接口
- Reader接口
- os.File同时实现了Writer和Reader接口

