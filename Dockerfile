FROM ubuntu:latest
EXPOSE 8000
WORKDIR /APP
ENV HOST=localhost PORT=5432
ENV DBNAME=root USER=root PASSWORD=root
COPY ./main main
CMD [ "./main" ]