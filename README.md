# municipalproduct1-service

#DB SCRIPTS

#Start Servers
mongod --dbpath "C:\database\mongodb\sample\data1" --logpath "C:\database\mongodb\sample\data1\log\mongod.log" --port 27018 --storageEngine=wiredTiger --journal --replSet rsSample
mongod --dbpath "C:\database\mongodb\sample\data2" --logpath "C:\database\mongodb\sample\data2\log\mongod.log" --port 27019 --storageEngine=wiredTiger --journal --replSet rsSample
mongod --dbpath "C:\database\mongodb\sample\data3" --logpath "C:\database\mongodb\sample\data3\log\mongod.log" --port 27020 --storageEngine=wiredTiger --journal --replSet rsSample

#Start Clients
mongo --port 27018
mongo --port 27019
mongo --port 27020


#One Time Commands

#On Primary
rsconf={_id:"rsSample",members:[{_id:0,host:"localhost:27018"}]}
rs.initiate(rsconf)
rs.add("localhost:27019")
rs.add("localhost:27020")
rs.status()

#On Secondary
rs.slaveOk()
