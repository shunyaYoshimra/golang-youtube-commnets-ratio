# https://github.com/shunyaYoshimra/golang-youtube-commnets-ratio
 
This application is supposed to be used for knowing how much ratio of specific videos's comments are written in except for Japanese.

# Requirement

* Docker v20.10.7
 
# Usage
 
```bash
git clone https://github.com/shunyaYoshimra/golang-youtube-commnets-ratio
docker-compose build
docker-compose up -d
docker-compose exec app /bin/bash
cd golang
go run main.go
```
 
# Note
 
Configure of Database & YouTube API Key must be matched with your environment.

.env file is at  ./golang/.env
 
# Author
 
* Shunya Yoshimura
* Kobe City University of Foreign Studies

 

