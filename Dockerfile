FROM golang:1.21.5-alpine3.18 AS build-stage

WORKDIR /app
# download the dependancy
COPY go.mod go.sum ./
RUN  go mod download
# copy the source code and html files
COPY cmd cmd/
COPY pkg pkg/
COPY view views/
# build the executable file
RUN go build -v -o ./build/api ./cmd/api
# final stage
FROM gcr.io/distroless/static-debian11

WORKDIR /app
# copy the binay file and html files
COPY --from=build-stage /app/build/api api
COPY --from=build-stage /app/views views/

CMD ["/app/api"]