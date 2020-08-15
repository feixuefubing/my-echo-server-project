# my-echo-server-project
my-echo-server-project----工程目录  
&emsp;app----程序启动类  
&emsp;&emsp;main.go----程序入口  
&emsp;domain----领域层  
&emsp;&emsp;wallet.go----领域模型定义和业务接口定义  
&emsp;go.mod----项目管理指导文件  
&emsp;go.sum----项目管理验证文件  
&emsp;wallet----钱包模块  
&emsp;&emsp;handler----处理器层  
&emsp;&emsp;&emsp;dto----数据传输对象  
&emsp;&emsp;&emsp;&emsp;extract_message_req.go----（测试）解签地址请求  
&emsp;&emsp;&emsp;&emsp;new_account_req.go----（测试）创建新账户请求  
&emsp;&emsp;&emsp;&emsp;sign_in_resp.go----登录随机消息返回  
&emsp;&emsp;&emsp;&emsp;sign_message_req.go----（测试）签名随机消息请求  
&emsp;&emsp;&emsp;&emsp;verify_req.go----验证地址与私钥匹配性请求  
&emsp;&emsp;&emsp;&emsp;verify_resp.go----验证地址与私钥匹配性返回  
&emsp;&emsp;&emsp;wallet_handler.go----Api处理方法和Routing定义  
&emsp;&emsp;usecase----业务层  
&emsp;&emsp;&emsp;wallet_ucase.go----业务接口实现  
&emsp;&emsp;&emsp;wallet_ucase_test.go----业务单元测试  
&emsp;&emsp;util----工具类  
&emsp;&emsp;&emsp;utils.go----工具类  

使用go语言echo框架和geth模块实现用户拥有私钥地址的验证功能。  
对于两个要求的Api，使用去-号的uuid作为登录的随机字符串，使用cookie保存。  
初学golang，代码风格和项目工程结构设计仍在学习和尝试过程当中，不足之处请您指出。  