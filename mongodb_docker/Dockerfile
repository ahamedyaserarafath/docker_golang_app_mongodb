## Simple docker file which will copy the current directory
## and run the startup script which will insert the data 
## finally it will run the mongodb in background

FROM debian:jessie-slim

WORKDIR /mongodb/bin
COPY . /mongodb/bin/

RUN mkdir -p /mongodb/data/

RUN set -x && apt-get update && apt-get install -y libssl1.0.0 libssl-dev

RUN /mongodb/bin/start_mongodb.sh

CMD ["/mongodb/bin/mongod","--port","5000","--dbpath","/mongodb/data"]

## Commands useful for future
#CMD ["/mongodb/bin/mongod","--fork", "--port","5000","--dbpath","/mongodb/data","--logpath","/mongodb/db.log"]
#CMD ["/mongodb/bin/mongod","--port","5000","--dbpath","/mongodb/data"]
#CMD ["/mongodb/bin/mongo","--port","5000","--quiet","test.js"]
#CMD ["/mongodb/bin/mongo","--port","5000"]
#CMD ["/mongodb/bin/start_mongodb.sh"]


