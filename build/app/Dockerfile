FROM ubuntu/nginx

RUN apt-get update -y
RUN apt-get upgrade -y

RUN apt-get install golang -y

RUN apt-get install curl -y

#Goose install
RUN curl -fsSL \https://raw.githubusercontent.com/pressly/goose/master/install.sh |\sh