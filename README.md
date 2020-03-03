# Bank Account API

This is an API that provider a bank account consults.

## Install Guide

### Prerequisites
Check if you have the follow tools:
1. [Docker Engine](https://docs.docker.com/engine/installation/)
2. [Docker Compose](https://docs.docker.com/compose/install/).
3. [Git](https://git-scm.com/downloads)

### Run application
1. First, clone this repository into your local machine using the follow command:
`git clone git@github.com:fernandochimi/bank-account.git`
2. Run `make execute` to build and run the application
3. Go! :rocket:

PS: The file `config/api.env` is for environment values from the application. The default values corresponds to host/port of application and the database pre defined values.

### Test API
You can test the API directly in the [browser](http://localhost:8000/). This link will works when you run application.

## `Makefile` tips
The `Makefile` provides some commands to run the application automatically in the Docker container.

* `make execute` - This command will execute the following commands:
	* `clean` - Clean useless files in the application structure before you up the container.
	* `build` - This command will be build the application in container.
	* `startd` - This command will execute the container in background.
	* `prune` - This command will clean all unused containers and images.
* `start` - This command will execute the container and exhibit logs.
* `stop` - This command will stop the container.