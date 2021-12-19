#Oratu
>一个用于隐藏C2的、开箱即用的反向代理服务器。旨在省去繁琐的配置Nginx服务的过程。

## 0x01 使用

1. 运行生成默认配置文件`config.ini`
2. 再次运行`Ortau`



此时Ortau监听在8091端口，判断发送至8091端口的请求UA中是否包含`Ortau`（默认配置，可修改），如果包含则转发至8092，如果不包含则转发至jd.com。

## 0x02 其他配置

个人使用的Profile配置文件为：

~~~shell
https://raw.githubusercontent.com/threatexpress/malleable-c2/master/jquery-c2.3.14.profile

#修改http-config中的x-forwarded-for头，以配合反向代理设置。
http-config {
    set trust_x_forwarded_for "true";
    header "Content-Type" "application/*; charset=utf-8";
}
#修改配置项set useragent
set useragent "Mozilla/5.0 (Windows NT 6.3; Trident/7.0; rv:11.0) like Ortau";
~~~



CS Listener配置为：

![image-20211219220450991](C:/Users/ZhXuting/AppData/Roaming/Typora/typora-user-images/image-20211219220450991.png)



iptables禁掉8092端口，仅允许127.0.0.1访问：

~~~shell
iptables -I INPUT -p tcp --dport 8092 -j DROP
iptables -I INPUT -s 127.0.0.1 -p tcp --dport 8092 -j ACCEPT
~~~



## 0x03

![image-20211219220911877](https://typora-mine.oss-cn-beijing.aliyuncs.com/typora/image-20211219220911877.png)

![image-20211219221006180](https://typora-mine.oss-cn-beijing.aliyuncs.com/typora/image-20211219221006180.png)

