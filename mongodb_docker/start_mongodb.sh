#!/bin/bash
## Simple bash shell script will start the mongodb and insert the data 
## For now hardcoded the path of mongodb bin and data folder
/mongodb/bin/mongod --fork --port 5000 --dbpath /mongodb/data --logpath /mongodb/db.log
/mongodb/bin/mongo --port 5000 --quiet /mongodb/bin/data.js


