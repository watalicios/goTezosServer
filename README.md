# goTezosServer: A Tezos Query Library Backed by MongoDB

The purpose of the goTezosServer is to build server applications around the tezos blockchain's data. In the
process of developing this library, I will create a usable server using this lib.

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
You will first need to install MongoDB and configure it to your preferences. Currently the library assumes there is no auth configuration for MongoDB. I will add this soon.


Create a database called `TEZOS` and a collection called `blocks`.
```
mongo
use TEZOS
db.createCollection("blocks")
```


Then install the goTezosServer:
```
go get gopkg.in/mgo.v2
go get github.com/DefinitelyNotAGoat/goTezosServer

```

Connection to the database is hardcoded for now. You can find it in tezosMongo.go
```
func SynchronizeTezosMongo(){
  blocks, err := GetAllBlocks()
  if (err != nil){
    fmt.Println(err)
  }

  session, err := mgo.Dial("127.0.0.1")
  c := session.DB("TEZOS").C("blocks")


  for _, block := range blocks{
  //  fmt.Println(block)
    err = c.Insert(block)
    if (err != nil){
      fmt.Println(err)
    }
  }
}
```


## goTezosServer Lib Documentation
The goTezosServer requires the use of an env variable called TEZOSPATH.


Example:

```
export TEZOSPATH=/home/tezosuser/tezos/
```

I will create a wiki shortly describing the functions available.


## Server Application

See the application [README.md](https://github.com/DefinitelyNotAGoat/goTezosServer/blob/master/Programs/README.md) for more information.

```
go build goTezosServer.go
./goTezosServer
```

## Authors

* **DefinitelyNotAGoat**

See also the list of [contributors](https://github.com/DefinitelyNotAGoat/goTezosServer/graphs/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
