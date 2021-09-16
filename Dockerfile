FROM golang:1.17

RUN mkdir /drello-api
WORKDIR /drello-api

COPY . .
RUN go get

# Environment variables about connection to DB.
ARG DB_USER=beba733eff51b5
ARG DB_PASS=4b1ff93d
ARG DB_TCP_HOST=us-cdbr-east-04.cleardb.com
ARG DB_PORT=4306
ARG DB_NAME=drello-stg

ENV DB_USER $DB_USER
ENV DB_PASS $DB_PASS
ENV DB_TCP_HOST $DB_TCP_HOST
ENV DB_PORT $DB_PORT
ENV DB_NAME $DB_NAME

# Environment variables for API server's endpoint.
ARG PORT=8080
ENV PORT $PORT
EXPOSE $PORT

CMD ["go", "run", "."]