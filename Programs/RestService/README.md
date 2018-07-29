# goTezosRestService Application

The purpose of the goTezosRestService application is to expose a rest API to query your MongoDB Tezos database.You can use this application without the `goTezosServer` application, if your database has block information already in it. Otherwise, you first need to install and run `goTezosServer` so that you have a useful database to query. 

A special thank you to Raleigh, Tingham and XTZ.com for supporting and funding the development of goTezosRestService. XTZ.com is an upcoming Tezos brand dedicated to providing the Tezos community with future tools and resources. If you have questions, please feel free to reach out to us on riot, and send the XTZ.com team some love. 

More robust documentation will come soon.

## Installation
The goTezosRestService application requires the use of an env variable called TEZOSPATH.

Example:

```
export TEZOSPATH=/home/tezosuser/tezos/
```

If you haven't already, you will then need to install [MongoDB](https://www.mongodb.com/) and configure it to your preferences. 

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

You can find the goTezosServer application inside the `Programs/RestService/` directory in the root of this repo. 

Clone the goTezosServer repo:
```
git clone https://github.com/DefinitelyNotAGoat/goTezosServer.git

```

Build the application:
```
cd Programs/RestService/
go build goTezosRestService.go
```
## Running
```
Usage of ./goTezosRestService:
  -collection string
    	Use the blocks collection (default "blocks")
  -connection string
    	URL or IP to the MongoDB Database (default "127.0.0.1")
  -db string
    	Use the TEZOS Database (default "TEZOS")
  -pass string
    	If using authentication, set the password
  -user string
    	If using authentication, set the user
 ```
 
## Temporary REST API Documentation
```
{id} Can be either a block level or a block hash
{kind} Can be the kind of operation. Example transfer or delegation
GET /head 
GET /block/{id}
GET /block/{id}/protocol
GET /block/{id}/chain_id
GET /block/{id}/hash
GET /block/{id}/header
GET /block/{id}/level
GET /block/{id}/proto
GET /block/{id}/predecessor
GET /block/{id}/timestamp
GET /block/{id}/validation_pass
GET /block/{id}/operation_hash
GET /block/{id}/fitness
GET /block/{id}/context
GET /block/{id}/priority
GET /block/{id}/proof_of_work_nonce
GET /block/{id}/signature
GET /block/{id}/metadata
GET /block/{id}/metadata/protocol
GET /block/{id}/metadata/next_protocol
GET /block/{id}/metadata/test_chain_status
GET /block/{id}/metadata/max_operations_ttl
GET /block/{id}/metadata/max_operation_data_length
GET /block/{id}/metadata/max_block_header_length
GET /block/{id}/metadata/max_operation_list_length
GET /block/{id}/metadata/baker
GET /block/{id}/metadata/level
GET /block/{id}/metadata/level/level
GET /block/{id}/metadata/level/position
GET /block/{id}/metadata/level/cycle
GET /block/{id}/metadata/level/voting_period
GET /block/{id}/metadata/level/expected_commitment
GET /block/{id}/metadata/voting_period_kind
GET /block/{id}/metadata/nonce_hash
GET /block/{id}/metadata/consumed_gas
GET /block/{id}/metadata/deactivated
GET /block/{id}/metadata/balance_updates
GET /block/{id}/operations
GET /block/operation/{id}               -{id} represents an operation hash
GET /block/operation/{id}/protocol      -{id} represents an operation hash
GET /block/operation/{id}/branch        -{id} represents an operation hash
GET /block/operation/{id}/contents      -{id} represents an operation hash
GET /block/operation/{id}/signature     -{id} represents an operation hash
GET /block/{id}/operations/kind/{kind}  -{id} represents an operation hash
```

Example: 
```
GET /block/10000

{
    "protocol": "PtCJ7pwoxe8JasnHY8YonnLYjcVHmhiARPJvqcC6VfHT5s8k8sY",
    "chain_id": "NetXdQprcVkpaWU",
    "hash": "BLc7tKfzia9hnaY1YTMS6RkDniQBoApM4EjKFRLucsuHbiy3eqt",
    "header": {
        "level": 10000,
        "proto": 1,
        "Predecessor": "BMG7bSzAh1is2896bUkK7RnUREqqN4BjcH4J7YgkFKcNHWNe4cM",
        "timestamp": "2018-07-07T12:06:27-05:00",
        "validation_pass": 4,
        "operations_hash": "LLob71uMBRtLaKGj3sDJmAT7VEdGTtEoogrbFFnPjxXiYfDmUQrgr",
        "fitness": [
            "00",
            "000000000004fff6"
        ],
        "context": "CoUnq1qGxUtidFCdcaCWXEQdefFDSdBTpjnYVcrHJ1cKYqL6HLiA",
        "priority": 0,
        "proof_of_work_nonce": "d4dac173a3904dfc",
        "signature": "sigRg6mM8oEt5y7nzSwi34P3UEoNDYjHF2Nik9s2f7xFGzMbbgmVYrc3uXdAKPF3ayDLv7vaEN4U2ZeDC69EJp4keYphw9WQ"
    },
    "metadata": {
        "protocol": "PtCJ7pwoxe8JasnHY8YonnLYjcVHmhiARPJvqcC6VfHT5s8k8sY",
        "next_protocol": "PtCJ7pwoxe8JasnHY8YonnLYjcVHmhiARPJvqcC6VfHT5s8k8sY",
        "test_chain_status": {
            "status": "not_running"
        },
        "max_operations_ttl": 60,
        "max_operation_data_length": 16384,
        "max_block_header_length": 238,
        "max_operation_list_length": [
            {
                "max_size": 32768,
                "max_op": 32
            },
            {
                "max_size": 32768
            },
            {
                "max_size": 135168,
                "max_op": 132
            },
            {
                "max_size": 524288
            }
        ],
        "baker": "tz3RDC3Jdn4j15J7bBHZd29EUee9gVB1CxD9",
        "level": {
            "level": 10000,
            "level_position": 9999,
            "cycle": 2,
            "cycle_position": 1807,
            "voting_period": 0,
            "voting_period_position": 9999,
            "expected_commitment": false
        },
        "voting_period_kind": "proposal",
        "nonce_hash": null,
        "consumed_gas": "0",
        "deactivated": [],
        "balance_updates": [
            {
                "kind": "contract",
                "contract": "tz3RDC3Jdn4j15J7bBHZd29EUee9gVB1CxD9",
                "change": "-16000000"
            },
            {
                "kind": "freezer",
                "change": "16000000",
                "category": "deposits",
                "delegate": "tz3RDC3Jdn4j15J7bBHZd29EUee9gVB1CxD9",
                "level": 2
            }
        ]
    },
    "operations": [
        [
            {
                "protocol": "PtCJ7pwoxe8JasnHY8YonnLYjcVHmhiARPJvqcC6VfHT5s8k8sY",
                "chain_id": "NetXdQprcVkpaWU",
                "hash": "oo5otsjk5Rrs3dvSemoV6Ni3bVJSVt9tM51edyiPLTyzwjMJBZK",
                "branch": "BMG7bSzAh1is2896bUkK7RnUREqqN4BjcH4J7YgkFKcNHWNe4cM",
                "contents": [
                    {
                        "kind": "endorsement",
                        "level": 9999,
                        "metadata": {
                            "balance_updates": [
                                {
                                    "kind": "contract",
                                    "contract": "tz3bTdwZinP8U1JmSweNzVKhmwafqWmFWRfk",
                                    "change": "-10000000"
                                },
                                {
                                    "kind": "freezer",
                                    "change": "10000000",
                                    "category": "deposits",
                                    "delegate": "tz3bTdwZinP8U1JmSweNzVKhmwafqWmFWRfk",
                                    "level": 2
                                }
                            ],
                            "delegate": "tz3bTdwZinP8U1JmSweNzVKhmwafqWmFWRfk",
                            "slots": [
                                30,
                                26,
                                22,
                                17,
                                9
                            ]
                        }
                    }
                ],
                "signature": "siggHkiwKvLDRmA7UUgWJewieSr42Dj1Xzgo2JqafxCtkoeXae8B2c8vokJbfhKzaqpftMoPvvDjXzx8k1zvSDmazC9W4nMn"
            },
            {
                "protocol": "PtCJ7pwoxe8JasnHY8YonnLYjcVHmhiARPJvqcC6VfHT5s8k8sY",
                "chain_id": "NetXdQprcVkpaWU",
                "hash": "onuw2K8sTenxbVG857d64JJ9oYFm5GiUKSQGBqwPBqAvhz3jQe9",
                "branch": "BMG7bSzAh1is2896bUkK7RnUREqqN4BjcH4J7YgkFKcNHWNe4cM",
                "contents": [
                    {
                        "kind": "endorsement",
                        "level": 9999,
                        "metadata": {
                            "balance_updates": [
                                {
                                    "kind": "contract",
                                    "contract": "tz3RB4aoyjov4KEVRbuhvQ1CKJgBJMWhaeB8",
                                    "change": "-6000000"
                                },
                                {
                                    "kind": "freezer",
                                    "change": "6000000",
                                    "category": "deposits",
                                    "delegate": "tz3RB4aoyjov4KEVRbuhvQ1CKJgBJMWhaeB8",
                                    "level": 2
                                }
                            ],
                            "delegate": "tz3RB4aoyjov4KEVRbuhvQ1CKJgBJMWhaeB8",
                            "slots": [
                                24,
                                15,
                                4
                            ]
                        }
                    }
                ],
                "signature": "sigSJC4HoiA5UJo5yHeTSgT4kSAortGj6KNdRFUUmUPWUXucfgGSUEzQmaSL4JaNds6MQN1tVPGZBTkMYonxBvJB5fiH8Psa"
            },
            {
                "protocol": "PtCJ7pwoxe8JasnHY8YonnLYjcVHmhiARPJvqcC6VfHT5s8k8sY",
                "chain_id": "NetXdQprcVkpaWU",
                "hash": "ooC5JaC1nXfvaT89WiJKjEonWvGeWnmRTdQc1wYQ3okHx4kkGvu",
                "branch": "BMG7bSzAh1is2896bUkK7RnUREqqN4BjcH4J7YgkFKcNHWNe4cM",
                "contents": [
                    {
                        "kind": "endorsement",
                        "level": 9999,
                        "metadata": {
                            "balance_updates": [
                                {
                                    "kind": "contract",
                                    "contract": "tz3RDC3Jdn4j15J7bBHZd29EUee9gVB1CxD9",
                                    "change": "-4000000"
                                },
                                {
                                    "kind": "freezer",
                                    "change": "4000000",
                                    "category": "deposits",
                                    "delegate": "tz3RDC3Jdn4j15J7bBHZd29EUee9gVB1CxD9",
                                    "level": 2
                                }
                            ],
                            "delegate": "tz3RDC3Jdn4j15J7bBHZd29EUee9gVB1CxD9",
                            "slots": [
                                28,
                                10
                            ]
                        }
                    }
                ],
                "signature": "sigmyorF4narTBSNMXemm6M6ggbNpWsCeFZGhYfR4QEutkBNXgKV3c33uiMx649zHCFsoKK6PX5SRbv8y1SMn94F53tANkEL"
            },
            {
                "protocol": "PtCJ7pwoxe8JasnHY8YonnLYjcVHmhiARPJvqcC6VfHT5s8k8sY",
                "chain_id": "NetXdQprcVkpaWU",
                "hash": "ooFBHKYXySSzanGtLhzT5TzyvSZUfA2mFXAYjJidhXdujgJvThn",
                "branch": "BMG7bSzAh1is2896bUkK7RnUREqqN4BjcH4J7YgkFKcNHWNe4cM",
                "contents": [
                    {
                        "kind": "endorsement",
                        "level": 9999,
                        "metadata": {
                            "balance_updates": [
                                {
                                    "kind": "contract",
                                    "contract": "tz3bvNMQ95vfAYtG8193ymshqjSvmxiCUuR5",
                                    "change": "-10000000"
                                },
                                {
                                    "kind": "freezer",
                                    "change": "10000000",
                                    "category": "deposits",
                                    "delegate": "tz3bvNMQ95vfAYtG8193ymshqjSvmxiCUuR5",
                                    "level": 2
                                }
                            ],
                            "delegate": "tz3bvNMQ95vfAYtG8193ymshqjSvmxiCUuR5",
                            "slots": [
                                29,
                                16,
                                5,
                                3,
                                0
                            ]
                        }
                    }
                ],
                "signature": "sigSJm6WS5rxkFdZFsuBboDMNeubGE5bAScK2rVVfx9XXSmCXAAVN7vxm9LaAMAXkzunvkLK8SkQhTynGCC7PiK9UAFwchex"
            },
            {
                "protocol": "PtCJ7pwoxe8JasnHY8YonnLYjcVHmhiARPJvqcC6VfHT5s8k8sY",
                "chain_id": "NetXdQprcVkpaWU",
                "hash": "onfFpbNQzmkrtsFk8tGgo8uXYRuEhjonni7paNkirMG4P47G8fQ",
                "branch": "BMG7bSzAh1is2896bUkK7RnUREqqN4BjcH4J7YgkFKcNHWNe4cM",
                "contents": [
                    {
                        "kind": "endorsement",
                        "level": 9999,
                        "metadata": {
                            "balance_updates": [
                                {
                                    "kind": "contract",
                                    "contract": "tz3UoffC7FG7zfpmvmjUmUeAaHvzdcUvAj6r",
                                    "change": "-14000000"
                                },
                                {
                                    "kind": "freezer",
                                    "change": "14000000",
                                    "category": "deposits",
                                    "delegate": "tz3UoffC7FG7zfpmvmjUmUeAaHvzdcUvAj6r",
                                    "level": 2
                                }
                            ],
                            "delegate": "tz3UoffC7FG7zfpmvmjUmUeAaHvzdcUvAj6r",
                            "slots": [
                                27,
                                25,
                                23,
                                20,
                                12,
                                8,
                                6
                            ]
                        }
                    }
                ],
                "signature": "sigufzbcHLNgdUgfg2PkQbVrJZTpE5JhgDtHZRUsSGNdxDg2zj53g2LzyVyX6nK88xDwMFh626auHgtPNssfSueEypAjoDdc"
            },
            {
                "protocol": "PtCJ7pwoxe8JasnHY8YonnLYjcVHmhiARPJvqcC6VfHT5s8k8sY",
                "chain_id": "NetXdQprcVkpaWU",
                "hash": "ooHTyojDYarDp12mqujgcgszaGNZjjSR8m5NTukEqxtWktv9Hpt",
                "branch": "BMG7bSzAh1is2896bUkK7RnUREqqN4BjcH4J7YgkFKcNHWNe4cM",
                "contents": [
                    {
                        "kind": "endorsement",
                        "level": 9999,
                        "metadata": {
                            "balance_updates": [
                                {
                                    "kind": "contract",
                                    "contract": "tz3NExpXn9aPNZPorRE4SdjJ2RGrfbJgMAaV",
                                    "change": "-4000000"
                                },
                                {
                                    "kind": "freezer",
                                    "change": "4000000",
                                    "category": "deposits",
                                    "delegate": "tz3NExpXn9aPNZPorRE4SdjJ2RGrfbJgMAaV",
                                    "level": 2
                                }
                            ],
                            "delegate": "tz3NExpXn9aPNZPorRE4SdjJ2RGrfbJgMAaV",
                            "slots": [
                                13,
                                2
                            ]
                        }
                    }
                ],
                "signature": "sigY3X2mjS6nRGVGhDaUdZeJQeLcqbcp5ZGrMoK8EHW19k5V2xchaxp3FVF4sbrhFfZQyz4jzsU1Q8TqXWf1VtfFznZ2v9wN"
            },
            {
                "protocol": "PtCJ7pwoxe8JasnHY8YonnLYjcVHmhiARPJvqcC6VfHT5s8k8sY",
                "chain_id": "NetXdQprcVkpaWU",
                "hash": "oniGQzstxfFeqfvAfSwEH632R6rP41thv8zUKKpYEBEAPezEKh7",
                "branch": "BMG7bSzAh1is2896bUkK7RnUREqqN4BjcH4J7YgkFKcNHWNe4cM",
                "contents": [
                    {
                        "kind": "endorsement",
                        "level": 9999,
                        "metadata": {
                            "balance_updates": [
                                {
                                    "kind": "contract",
                                    "contract": "tz3VEZ4k6a4Wx42iyev6i2aVAptTRLEAivNN",
                                    "change": "-8000000"
                                },
                                {
                                    "kind": "freezer",
                                    "change": "8000000",
                                    "category": "deposits",
                                    "delegate": "tz3VEZ4k6a4Wx42iyev6i2aVAptTRLEAivNN",
                                    "level": 2
                                }
                            ],
                            "delegate": "tz3VEZ4k6a4Wx42iyev6i2aVAptTRLEAivNN",
                            "slots": [
                                31,
                                19,
                                7,
                                1
                            ]
                        }
                    }
                ],
                "signature": "sigePU7QxfEJUo4LnpWWMaGf5eTb95VP3BXjR5NZPNJxoKemkriftuRBZ2jfTh67xUkitKF4beNbXPDx4qvpTjmVbYtJb8Gm"
            },
            {
                "protocol": "PtCJ7pwoxe8JasnHY8YonnLYjcVHmhiARPJvqcC6VfHT5s8k8sY",
                "chain_id": "NetXdQprcVkpaWU",
                "hash": "ooX8hcwA2Nx8fhgGWDqaT5ZWtUeVm77TGMnGxWXM5c1CVkGALou",
                "branch": "BMG7bSzAh1is2896bUkK7RnUREqqN4BjcH4J7YgkFKcNHWNe4cM",
                "contents": [
                    {
                        "kind": "endorsement",
                        "level": 9999,
                        "metadata": {
                            "balance_updates": [
                                {
                                    "kind": "contract",
                                    "contract": "tz3WMqdzXqRWXwyvj5Hp2H7QEepaUuS7vd9K",
                                    "change": "-8000000"
                                },
                                {
                                    "kind": "freezer",
                                    "change": "8000000",
                                    "category": "deposits",
                                    "delegate": "tz3WMqdzXqRWXwyvj5Hp2H7QEepaUuS7vd9K",
                                    "level": 2
                                }
                            ],
                            "delegate": "tz3WMqdzXqRWXwyvj5Hp2H7QEepaUuS7vd9K",
                            "slots": [
                                21,
                                18,
                                14,
                                11
                            ]
                        }
                    }
                ],
                "signature": "sigw9g4z8mRd9JMUpxxaJRxUTbHq4evpVAHFqsTY4foh83853pmbuNXwUgRNwtRgnzBQV4iHyQ4xeUjbwk8aWDSr1FVG1qGk"
            }
        ],
        [],
        [],
        []
    ]
}
```

## Authors

* **DefinitelyNotAGoat**

See also the list of [contributors](https://github.com/DefinitelyNotAGoat/goTezosServer/graphs/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
