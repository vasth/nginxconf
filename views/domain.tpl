<!DOCTYPE html> 
<html>
  	<head>
    	<title>domain</title>
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8"> 
	</head> 
  	<body>
 {{$option :=.option}}
 <p><h3>设置配置文件路径</h3></p>
  <form action="/addconfdir" method="post"> 
 当前nginx配置文件放置路径(最后面不加/或者\)：<input type="text" name="confdir" value="{{$option.Value}}" style="width:300px;">
  <input type="submit">
	 </form>
 <br>
	<table> 
	<tr >
	    <td>ID   </td>
	    <td>主机域名   </td>
	    <td>主机端口   </td>
	    <td>主机代理地址  </td>
	    <td>访问日志  </td>
	    <td>静态资源保留时间   </td>
	    <td>静态资源目录  </td> 
	    <td>操作</td>
	  </tr>

  	{{if .domians}}
  		  {{range $k, $v := .domians}}   
		
			<tr >
			    <td>  {{$v.Id}}  <br></td>
			    <td>  {{$v.Server_name}} <br></td>
			    <td>   {{$v.Port}}  <br></td>
			    <td>   {{$v.Proxy_pass}} <br></td>
			    <td> {{$v.Access_log}} <br></td>
			    <td>   {{$v.Expires}} <br></td>
			    <td>  {{$v.Root}} <br></td> 
			    <td><a href="/create/{{$v.Id}}">生效</a> <a href="/delete/{{$v.Id}}">删除</a><br></td>
			  </tr> 
			{{end}}
	{{else}}
		

		暂无主机配置 
	{{end}}
</table> 
<p><h3>添加虚拟主机</h3></p>
	 <form action="/addserver" method="post">

	主机域名(不需要添加http://)：<input type="text" name="Server_name"><br>
	主机端口：<input type="text" name="Port"><br>
	主机代理地址(需要添加http://)：<input type="text" name="Proxy_pass"><br>
	访问日志(不需要则设置为off)：<input type="text" name="Access_log"><br>
	静态资源保留时间(不需要则设置为例如1天：1d)：<input type="text" name="Expires"><br>
	静态资源目录(系统绝对路径)：<input type="text" name="Root"><br>
	<input type="submit">
	 </form>

	 {{if compare .isadd 1}}
	 添加域名成功
	 {{end}}


<pre>  
如果选择代理缓存则需要在nginx http 配置中添加下列代码
<code>  
proxy_connect_timeout 10;
proxy_read_timeout 180;
proxy_send_timeout 5;
proxy_buffer_size 16k;
proxy_buffers 4 64k;
proxy_busy_buffers_size 256k;
proxy_temp_file_write_size 256k;
proxy_temp_path /tmp/temp_dir;
proxy_cache_path /home/cache levels=1:2 keys_zone=cache_one:100m inactive=1d max_size=10g; 
</code>  
</pre>  
   
	</body>
</html>
