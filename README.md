﻿# AutoDefence
一个自动化对抗闭环的服务器相关部分<br>
<br>
这个解决方案看起来似乎很简单，<br>
真的很简单，因为复杂的部分没有在这个里面，<br>
这里只是流程框架部分。<br>
<br>
我本人是做二进制的攻防对抗的，<br>
我们有云控，我们的云控也是需要做各种事情，<br>
包括数据收集，数据整理，数据转化，数据测试，数据部署，数据上线，数据执行，数据回报，<br>
之后回到前面的数据收集部分。<br>
<br>
这个工程里面，目前处理的是数据整理部分，<br>
大致工作流程是这样的，<br>
在某台机器上搭建了一个数据处理服务器A，使用工程Server，<br>
在某台机器上搭建了一个数据接收服务器B，使用工程HttpServer，<br>
在B 服务器上，同时放置了一个客户端工具，使用工程Client。<br>
<br>
工作步骤：<br>
1：当前置服务器有数据过来的话，那么前置服务器会向我们的HttpServer发送一条Post请求，<br>
由于我们的前置服务器可能是一个PHP的服务器，所以用Http带Post请求，我感觉很人道。<br>
2：根据我的要求，前置服务器发送过来一个结构体，体现为JSON，HttpServer解析这个JSON，之后，<br>
得到了一块数据，这个数据其实就是我攻防对抗要处理的数据。<br>
3：HttpServer把数据保存到本地之后，调起同一台服务器上的Client，并且传入数据文件路径，<br>
之后HttpServer当前阶段工作结束，继续等待前置服务器发来数据。<br>
4：Client获取数据文件之后，把数据抛给Server，让Server去处理数据，并且等待反馈。<br>
5：在Server所在的服务器上，部署了一套专门用做数据整理的一套工具（不在此解决方案中），
Server在得到Client上报上来的数据之后，会去调用数据整理工具，并且等待结束。<br>
6：数据整理工具组会开始整理数据，然后把整理结果返回给Server。<br>
7：Server拿到整理结果，并且拿到返回值，把返回值以及整理结果发送给Client，
之后Server当前阶段的工作结束，关闭当前链接，等待下次信息。<br>
8：Client根据结果使用一个EMail工具（不在此解决方案中）发送EMail给相关责任人。<br>
<br>
目前工作即此。<br>
<br>
后续开发相关的工作包括数据转化，没有了，所以后续我的工作就是把数据转化整合到自动化流程中，<br>
至于自动化测试流程与我无关，我不再弄它。<br>
<br>