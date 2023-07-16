# syntax=docker/dockerfile:1

# specify the base image to  be used for the application, alpine or ubuntu
FROM golang:1.17-alpine
RUN apk add build-base
# create a working directory inside the image
WORKDIR /app
RUN mkdir -p /app/logs

ENV TWILIO_SID_DREWS_MAIN='SKe7d704d9913b40109edcfcc2741ef3d9'
ENV TWILIO_ACCOUNT_SID='AC88fb22e78ad73cda3cd91d903ab7b0c5'
ENV TWILIO_SECRET_DREWS_MAIN='sqiKnSTLADGBaY2cOPZ67Wea1ZsJ5RXR'
ENV TWILIO_PHONE_NUMBER='13149124603'
ENV RECIPIENTS_PHONE_NUMBER='14802971704'
# ENV SQLITE_DB='/app/Craigslist.db'
ENV selectQuery='`\
	SELECT \
	  c.city || s.search_string \
	FROM \
	searches s, \
	cities c`' 
# copy Go modules and dependencies to image
COPY go.mod ./

# download Go modules and dependencies
RUN go mod download

# copy directory files i.e all files ending with .go
COPY . ./



# compile application
RUN go build -o /scraper .

# tells Docker that the container listens on specified network ports at runtime
EXPOSE 8080

ENTRYPOINT ["tail", "-f", "/dev/null"]