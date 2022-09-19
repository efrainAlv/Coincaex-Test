FROM golang

#ARG APP_LOCATION

WORKDIR /usr/src/app

#COPY ./${APP_LOCATION} .

COPY ./src .

EXPOSE 8080

CMD [ "go", "run", "main.go" ]

#sudo docker build -t golang-native --no-cache=true --build-arg APP_LOCATION=p1 .
#sudo docker build -t golang-native --no-cache=true .

#sudo docker run --name p1 -it golang-native -env APP_LOCATION=./p1