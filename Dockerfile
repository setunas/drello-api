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
ARG HOGE=heroku_e60cb34a8aa0b6d

ENV DB_USER $DB_USER
ENV DB_PASS $DB_PASS
ENV DB_TCP_HOST $DB_TCP_HOST
ENV DB_PORT $DB_PORT
ENV DB_NAME $DB_NAME
ENV HOGE $HOGE

RUN echo "$DB_USER"
RUN echo "$DB_PASS"
RUN echo "$DB_TCP_HOST"
RUN echo "$DB_PORT"
RUN echo "$DB_NAME"
RUN echo "$HOGE"


# Database Migration
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN echo "mysql://$DB_USER:$DB_PASS@tcp($DB_TCP_HOST:$DB_PORT)/$DB_NAME"
# RUN migrate -path db/migrations -database "mysql://$DB_USER:$DB_PASS@tcp($DB_TCP_HOST:$DB_PORT)/$DB_NAME" up
RUN migrate -path db/migrations -database "mysql://beba733eff51b5:4b1ff93d@tcp(us-cdbr-east-04.cleardb.com:${DB_PORT})/heroku_e60cb34a8aa0b6d" up
# RUN migrate -path db/migrations -database "mysql://beba733eff51b5:4b1ff93d@tcp(us-cdbr-east-04.cleardb.com:3306)/heroku_e60cb34a8aa0b6d" up

# Environment variables for API server's endpoint.
ARG PORT=8080
ENV PORT $PORT
EXPOSE $PORT

CMD ["go", "run", "."]