# Send email 
    
    type SendEmailParams struct {
		UserEmail string	//发送邮件的邮箱
		Password string		//邮箱密码或授权码
		Host string		//邮箱服务器
		Port int		//邮箱服务器端口
		ContentType string  	//"Content-Type: text/plain; charset=UTF-8"
		} 

