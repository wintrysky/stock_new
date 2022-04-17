FROM alpine

# ARG exename

RUN mkdir -p /data/app/
COPY ./bin/stocks /data/app
COPY ./bin/config.yaml /data/app

# RUN cd /data/app && echo `ls -la`
# RUN echo $exename

ENV CFG_DIR=/data/app/config/
# ENV EXE_NAME $exename

WORKDIR /data/app
ENTRYPOINT /data/app/stocks