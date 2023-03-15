#!/bin/bash

#设置变量，url为你需要检测的目标网站的网址（IP或域名）

url=$1
  
    #定义函数check_http：
    
    #使用curl命令检查http服务器的状态
    
    #-m设置curl不管访问成功或失败，最大消耗的时间为5秒，5秒连接服务为相应则视为无法连接
    
    #-s设置静>默连接，不显示连接时的连接速度、时间消耗等信息
    
    #-o将curl下载的页面内容导出到/dev/null(默认会在屏幕显示页面内容)
    
    #-w设置curl命令需要显示的内容%{http_code}，指定curl返回服务器的状态码
  
check_http(){
    
   status_code=$(curl   -m 5 -s -o /dev/null -w %{http_code} -H 'HTTP_BLUEKING_SUPPLIER_ID:0' -H 'BK_USER:migrate' $url)
  
}

while :
  do
  check_http
  echo $status_code
  if [[ $status_code -ne 200 ]];then
       sleep 5
   else
       echo "$url success 200"
       exit 0
   fi
done
