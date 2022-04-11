FROM         		adoptopenjdk/openjdk8:latest
MAINTAINER 			tao.jun@neolink.com
RUN                 ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN                 echo 'Asia/Shanghai' >/etc/timezone
COPY                ./app.jar /usr/local/app.jar
COPY                ./entrypoint.sh /usr/local/entrypoint.sh
RUN					chmod +x /usr/local/entrypoint.sh
ENTRYPOINT  		["/bin/sh","/usr/local/entrypoint.sh"]