## Follow the respective mongodb and data api steps and try the below API

1. Below API will list all the data from mongodb 

http://13.126.236.37:5050/data/getalldata

2. Below API will get the respective data using the ID

http://13.126.236.37:5050/data/getdatabyid/<id>

Example:
http://13.126.236.37:5050/data/getdatabyid/15f146d0-7030-490c-a8db-a1ebd0e900ad

3. Below API will retrieve the data w.r.t to search text in data title

http://13.126.236.37:5050/data/getdatabytitle/<search_text>

Example:
http://13.126.236.37:5050/data/getdatabytitle/esse

4. Below API will list all the data from mongodb either by acsending or descending order Date

http://13.126.236.37:5050/data/getdatasorted/<order_by>

Example:
Acsending order by date
http://13.126.236.37:5050/data/getdatasorted/asc

Descending order by date
http://13.126.236.37:5050/data/getdatasorted/desc


