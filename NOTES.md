title 《GO Web》
# session和数据存储
- http协议是无状态的-->经典解决方案是cookie和session
  
     cookie是客户端机制，把用户的数据保存在客户端；session是服务端机制，服务器使用类似于散列的结构存储数据。

- session和cookie
     
     1.go 设置cookie：go 语言中通过net/http包中的SetCookie来设置  http.SetCookie(w ResponseWriter, cookie *Cookie)

     2.session通过cookie在客户端保存session ID ，而将用户的其他会话消息保存在服务端session对象中。

- Go如何使用session
  
     1. session的基本原理是由服务器为每个会话信息维护一份信息数据，客户端和服务器依靠一个全局唯一的标识来f访问这份数据。当用户访问web服务器时，服务器终端会随需创建sessionsession，分三个步骤：生成全局唯一标识符；开辟数据存储空间； 将session的全局唯一标识符发送给客户端（cookie和URL重写）
      
    2.go实现session管理

         session管理器
         全局唯一的session ID
         session 创建
         ssession重制
         session销毁

- session存储

- 预防session劫持
   
   1. cookieonly和token 
   2. 间隔生成SID



# 文本处理
- XML处理
   1. 解析XML：通过XML包的unmarshal 函数
   2. 生成XML
   3. 
- JSONc处理
     
   第三方包 go-jsonsimple

- 正则处理
  
  regexp包：使用复杂的正则首先是compile,他会解析正则表达式是否合法，如果正确就会返回一个regexp，然后利用返回的regexp在任意的字符串上面执行需要的操作


  strings包

- 模板处理
  1.使用template包处理

- 文件操作：os包
   1. 目录操作
   2. 文件操作：建立与打开文件、写文件、读文件、删除文件

- 字符串处理
   1. strings包：字符串操作
   2. strconv包：字符串转换
  






     


     


