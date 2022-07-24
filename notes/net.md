# <center><font color="#006666">网络编程</font></center>
## __<font color="#006666">net包</font>__
- 服务端:
  - net.Listen
  - litener.Accept
  - conn.Read
  - conn.Write
- 客户端:
  - net.Dial
  - net.Write
  - net.Read
  
- 服务器判断关闭
  - Read读客户端，返回0 -- 对端关闭
