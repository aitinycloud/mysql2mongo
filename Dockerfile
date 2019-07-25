FROM centos
RUN mkdir -p /opt/mysql2mongo/
ADD ./release /opt/mysql2mongo/
EXPOSE 8000
HEALTHCHECK --interval=5s --timeout=3s CMD curl --fail http://localhost:8000/ping || exit 1
ENV PATH $PATH:/opt/mysql2mongo/bin
WORKDIR /opt/mysql2mongo/
CMD ["/opt/mysql2mongo/bin/start.sh"]
