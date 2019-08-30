FROM golang:1.12 AS build
ADD . /app
RUN cd /app && go build -o app

FROM scratch
COPY --from=build /app/app /
ENTRYPOINT ["/app"]
