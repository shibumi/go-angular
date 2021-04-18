FROM node:12.11 AS ANGULAR_BUILD
RUN npm install -g @angular/cli@8.3.12
COPY webapp /webapp
WORKDIR webapp
RUN npm install && ng build --prod

FROM golang:1.16 as GO_BUILD
WORKDIR /go/src/app
ADD server /go/src/app
COPY --from=ANGULAR_BUILD /server/static /go/src/app
RUN go build -o /go/bin/app

FROM gcr.io/distroless/base
COPY --from=GO_BUILD /go/bin/app /
CMD ["/app"]