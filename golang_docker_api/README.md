## Readme file for building the docker and run the docker
## This docker is for installing the golang and running in the app server in port 5050

Steps to Run the simple mongodb server

1. Clean up the existing the docker images (Be caution it will remove every images in the server)

  'docker rm $(docker ps -a -q)'
  'docker rmi $(docker images -q)'

2. Build the docker container with the below command
  
   'docker build -t golang-app-server .'

3. Run the docker container with the below command

   'docker run -it -p 5050:5050 -d golang-app-server'

# In the golang app mongodb ip is hardcoded please change the value wherever, ip will be in constant.go file
