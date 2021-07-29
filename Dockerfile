FROM node:lts-alpine as client-builder
COPY client/package.json /client/package.json
WORKDIR /client
RUN npm install
COPY client /client
RUN npm run build

FROM golang:latest as server-builder
COPY server /server
WORKDIR /server

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/web/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=client-builder /client/build static
COPY --from=server-builder /server/app .

EXPOSE 8080
CMD ["./app"]