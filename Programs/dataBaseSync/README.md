# goTezosServer Application

A goTezosServer Application that will add all tezos block information to a mongodb databse, and keep the database synchronized with the network.

A special thank you to Raleigh, Tingham and XTZ.com for supporting and funding the development of goTezosServer. XTZ.com is an upcoming Tezos brand dedicated to providing the Tezos community with future tools and resources. If you have questions, please feel free to reach out to us on riot, and send the XTZ.com team some love. 

## Installation
The goTezosServer application requires the use of an env variable called TEZOSPATH.

Example:

```
export TEZOSPATH=/home/tezosuser/tezos/
```

You will then need to install [MongoDB](https://www.mongodb.com/) and configure it to your preferences. 

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

You can find the goTezosServer application inside the `Programs/dataBaseSync/` directory in the root of this repo. 

Clone the goTezosServer repo:
```
git clone https://github.com/DefinitelyNotAGoat/goTezosServer.git

```

Build the application:
```
cd Programs/dataBaseSync/
go build goTezosServer.go
```

## Running
```
Usage of ./goTezosServer:
  -collection string
    	Use the blocks collection (default "blocks")
  -connection string
    	URL or IP to the MongoDB Database (default "127.0.0.1")
  -db string
    	Use the TEZOS Database (default "TEZOS")
  -init
    	Start synchronization of the database from cycle 0
  -pass string
    	If using authentication, set the password
  -user string
    	If using authentication, set the user
```
## Authors

* **DefinitelyNotAGoat**

See also the list of [contributors](https://github.com/DefinitelyNotAGoat/goTezosServer/graphs/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
