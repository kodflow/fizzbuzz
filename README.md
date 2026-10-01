# fizzbuzz [![Reference](https://pkg.go.dev/badge/github.com/kodflow/fizzbuzz.svg)](https://pkg.go.dev/github.com/kodflow/fizzbuzz) [![CI](https://img.shields.io/github/actions/workflow/status/kodflow/fizzbuzz/ci.yml?label=CI)](https://github.com/kodflow/fizzbuzz/actions/workflows/ci.yml) [![License](https://img.shields.io/github/license/kodflow/fizzbuzz?label=&cacheSeconds=1&color=blue)](https://github.com/kodflow/fizzbuzz/blob/main/LICENSE) [![Latest Stable Version](https://img.shields.io/github/v/tag/kodflow/fizzbuzz?label=&cacheSeconds=1&color=blue)](https://github.com/kodflow/fizzbuzz/releases/latest) [![Size](https://img.shields.io/docker/image-size/kodmain/fizzbuzz?label=&cacheSeconds=1&color=blue)](https://github.com/kodflow/fizzbuzz/pkgs/container/fizzbuzz)

## Assignment
Write a simple fizz-buzz REST server.

The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by "fizz", all multiples of 5 by "buzz", and all multiples of 15 by "fizzbuzz".

The output would look like this: "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,...".

1. Your goal is to implement a web server that will expose a REST API endpoint that:
    - ⁠Accepts five parameters: three integers int1, int2 and limit, and two strings str1 and str2.
    - ⁠Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.

2. The server needs to be:
    - ⁠Ready for production
    - ⁠Easy to maintain by other developers

3. Bonus: add a statistics endpoint allowing users to know what the most frequent request has been. This endpoint should:
    - ⁠Accept no parameter
    - ⁠Return the parameters corresponding to the most used request, as well as the number of hits for this request

## Prerequisites
To build and run this server, the following must be installed:
 - [Terraform](https://developer.hashicorp.com/terraform/install) *v1.7.5*
 - [Taskfile](https://taskfile.dev/installation/) *v3.35.1*
 - [Swag](https://github.com/swaggo/swag#getting-started) *v1.16.3*
 - [Docker](https://docs.docker.com/desktop/) *v25.0.3*

## Installation

````
git clone https://github.com/kodflow/fizzbuzz.git
cd fizzBuzz
````

## Run latest published release
````
task deploy:release
````

## Run from repository
````
task deploy:local
````

## Test
#### Swagger & API [http(s)://localhost](http(s)://localhost)
#### Grafana [http(s)://localhost:3000](http(s)://localhost:3000)
#### Prometheus [http://localhost:9090/targets](http://localhost:9090/targets)



## Benchmark
Perfect to see grafana
````
task test:api
````

## Preview
![FizzBuzz](preview.gif)
