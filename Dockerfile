FROM golang:1.17

RUN mkdir /drello-api
WORKDIR /drello-api

COPY . .
RUN go get

# Environment variables about connection to DB.
ARG DB_USER
ARG DB_PASS
ARG DB_TCP_HOST=us-cdbr-east-04.cleardb.com
ARG DB_PORT
ARG DB_NAME=heroku_e60cb34a8aa0b6d

ENV DB_USER $DB_USER
ENV DB_PASS $DB_PASS
ENV DB_TCP_HOST $DB_TCP_HOST
ENV DB_PORT $DB_PORT
ENV DB_NAME $DB_NAME

# RUN go get -tags 'mysql' -u github.com/golang-migrate/migrate/cmd/migrate
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN migrate -path db/migrations -database "mysql://beba733eff51b5:4b1ff93d@tcp(us-cdbr-east-04.cleardb.com:3306)/heroku_e60cb34a8aa0b6d" up

# Environment variables for API server's endpoint.
ARG PORT=8080
ENV PORT $PORT
EXPOSE $PORT

CMD ["go", "run", "."]