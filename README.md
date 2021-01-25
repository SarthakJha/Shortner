# Shortner


# **URL shortner** api written in **Golang** which uses **Redis** for caching and **MongoDB** for persitance.

**URL shortner** api written in **Golang** which uses **Redis** for caching and **MongoDB** for persitance.

_Running with **Docker** requires no pre-requisite software installation except Docker!_

## To start with docker-compose:

1. Clone the repository
   `git clone https://github.com/SarthakJha/Shortner.git`
2. Run `docker-compose build`. It will build the go application create a docker image for it
3. Run `docker-compose up` to start the application

## To Start the server locally:

1.  Clone the repository
    `git clone https://github.com/SarthakJha/Shortner.git`
2.  Run `make build` to compile the code
3.  To start the application run `make run`

**Make sure to start your Redis and MongoDB instances**

_To build compile your code for other platforms run `make cross_build` and run the executable respectively by `./bin/<executable_name>`_
