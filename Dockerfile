FROM busybox:glibc
LABEL maintainer="Shenle Lu <lushenle@gmail.com>" app="myapp"
COPY app /bin/app
RUN chmod +x /bin/app
CMD ["/bin/app"]
