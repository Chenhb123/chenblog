FROM centos:7.2.1511
# 假设已经完成编译，取名blog
COPY main /opt/blog/
CMD ["/opt/blog/main"]
