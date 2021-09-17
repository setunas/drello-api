FROM golang:1.17

RUN mkdir /drello-api
WORKDIR /drello-api

COPY . .
RUN go get

# Connection to DB
ARG DB_USER
ARG DB_PASS
ARG DB_TCP_HOST
ARG DB_PORT
ARG DB_NAME

ENV DB_USER $DB_USER
ENV DB_PASS $DB_PASS
ENV DB_TCP_HOST $DB_TCP_HOST
ENV DB_PORT $DB_PORT
ENV DB_NAME $DB_NAME

# Database Migration
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
# RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest && \
#   migrate -path db/migrations -database "mysql://beba733eff51b5:${DB_PASS}@tcp(${DB_TCP_HOST}:${DB_PORT})/${DB_NAME}" up

# API server's endpoint
ARG PORT=8080
ENV PORT $PORT
EXPOSE $PORT

CMD ["go", "run", "."]