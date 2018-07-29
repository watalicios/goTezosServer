# goTezosServer: A Tezos Query Library Backed by MongoDB

The purpose of the goTezosServer is to build server applications around the information exposed in the Tezos protocol RPC. This library is used to create the `goTezosRestService`, which is a ReSTful server that exposes all block information in the Tezos network into a ReST API. The library is also used to create the `goTezosServer` which is a program that updates and synchronizes a MongoDB database with Tezos block information. These two programs are to be used in conjunction, because the rest service depends on the database. 

A special thank you to Raleigh, Tingham and XTZ.com for supporting and funding the development of goTezosServer. XTZ.com is an upcoming Tezos brand dedicated to providing the Tezos community with future tools and resources. If you have questions, please feel free to reach out to us on riot, and send the XTZ.com team some love. 

## Installation
You will first need to install [MongoDB](https://www.mongodb.com/) and configure it to your preferences. 

Second you will need to setup the database that the goTezosServer Library will use. If you would like to prefill your databse with block information you can download the exported version [tezosdb-block-34046.json](https://www.dropbox.com/s/hq14v696ed99997/tezosdb-block-34046.json?dl=0)

First get the MD5 checksum of the database and check to make sure the import file is authentic.
```
Checksum
MD5 (tezosdb-block-34046.json) = 448b06b76e33449de97c7dce0efd5deb
```

After verifying the checksum, import the database into database `TEZOS` and collection `blocks`. You can change the names to your preferences, but those are the defaults for the `goTezosRestService`, and `goTezosServer`.

```
sudo mongoimport --db TEZOS --collection blocks --file tezosdb-block-34046.json
```

If you want to start from scratch, create the database and collection like below.
```
mongo
use TEZOS
db.createCollection("blocks")
```

To install the library please download first the `mgo` dependency, and then the library itself. 
```
go get gopkg.in/mgo.v2
go get github.com/DefinitelyNotAGoat/goTezosServer

```

## Quick Start
The goTezosServer library requires the use of an env variable called TEZOSPATH.
```
export TEZOSPATH=/home/tezosuser/tezos/
```

Import the library into your program. 
```
import "github.com/DefinitelyNotAGoat/goTezosServer"
```

Initialize the connection to your database. 
```
goTezosServer.SetDatabaseConnection(database_connection_uri, database, collection)
```

## goTezosServer
Please see the goTezosServer Application [README.md](https://github.com/DefinitelyNotAGoat/goTezosServer/tree/master/Programs/dataBaseSync).

## goTezosRestService Application

Please see the goTezosRestService Application [README.md](https://github.com/DefinitelyNotAGoat/goTezosServer/tree/master/Programs/RestService).

## Authors

* **DefinitelyNotAGoat**

See also the list of [contributors](https://github.com/DefinitelyNotAGoat/goTezosServer/graphs/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
