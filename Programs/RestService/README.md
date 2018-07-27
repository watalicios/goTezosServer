# goTezosServer Application: A Tezos Query Server Backed by MongoDB

The purpose of the goTezosServer application is to provide an easy and queryable database to get data from the Tezos network.

If you would like to send me some coffee money:
```
tz1hyaA2mLUQLqQo3TVk6cQHXDc7xoKcBSbN
```

If you would like to delegate to me to show your support (5% dynamic fee):
```
tz1SUgyRB8T5jXgXAwS33pgRHAKrafyg87Yc
```


More robust documentation will come soon.

## Installation
The goTezosServer application requires the use of an env variable called TEZOSPATH.


Example:

```
export TEZOSPATH=/home/tezosuser/tezos/
```

If you haven't already, you will first need to install MongoDB and configure it to your preferences. Then create a database called `TEZOS` and a collection called `blocks`. Or import an the exported database. Currently the library assumes there is no auth configuration for MongoDB. I will add this soon.
```
mongo
use TEZOS
db.createCollection("blocks")
```
OR:

If you want to use the most recent database you can import it, and skip the creation part. [Download](https://www.dropbox.com/s/hq14v696ed99997/tezosdb-block-34046.json?dl=0).
```
sudo mongoimport --db TEZOS --collection blocks --file tezosdb-block-34046.json

Checksum
MD5 (tezosdb-block-34046.json) = 448b06b76e33449de97c7dce0efd5deb
```

You can find the goTezosServer application inside the `Programs` directory

Install the goTezosServer Lib:
```
go get github.com/DefinitelyNotAGoat/goTezosServer

```

Build and run the server:
```
go build goTezosServer.go
./goTezosServer
```

## Authors

* **DefinitelyNotAGoat**

See also the list of [contributors](https://github.com/DefinitelyNotAGoat/goTezosServer/graphs/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
