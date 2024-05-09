FROM nginx:alpine


RUN apk --no-cache add tree
COPY ./src /src

RUN tree /src

COPY /src/index.html /usr/share/nginx/html/index.html
