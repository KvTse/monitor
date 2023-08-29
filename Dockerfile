# 表示依赖 alpine 最新版
FROM centos:centos7.9.2009
MAINTAINER Xie Tongqing<xietongqing@piesat.cn>
ENV VERSION 1.0

# 在容器根目录 创建一个 apps 目录
WORKDIR /apps

# 挂载容器目录
VOLUME ["/apps/conf"]

# 拷贝当前目录下 go_docker_demo1 可以执行文件
COPY monitor-agent /apps/monitor-agent

# 拷贝配置文件到容器中
COPY config.yaml /apps/conf/config.yaml

# 设置时区为上海
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

# 设置编码
ENV LANG C.UTF-8

# 暴露端口
EXPOSE 8089

# 运行golang程序的命令
ENTRYPOINT ["/apps/monitor-agent"]