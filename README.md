# email-blaster
Highly concurrent email sender client.


## Notes
- Tested with 10,000,000 records and 1,000,000 goroutines, thus stable at least up to this size
- There is no proper error handling mechanism for each go-routine. I would extend goroutine to report errors into
a dedicated channel, and set more workers to re-try in case of failure.
- I would also use postgresql instead of sqlite, as postgres offers much better and faster mutex-ing for highly concurrent
read/writes, where sqlite can get corrupted if the concurrency is too sever (> 1000 goroutines)
- A temp database is created and removed during the test or main run
- Tests could be better, I created relatively simple ones, as most of them need integration testing, I tried to cover
most of the unit test
## Install
  ```
  git clone https://github.com/CommoDor64/email-blaster.git
  cd email-blaster
  go get ./...
  ```
## Run
Run with default configurations (add -s to prevent debug print)
  ```
  go run . -s
  ```
## Usage
the application has couple of configurable flags
- -w - size of worker pool (default 10,000)
- -c - chunk size, i.e how many record are pulled from DB on each iteration (default 100,000)
- -s - is silent, if specified, removes debug output (default false)
- -r - rows number, defines how many rows the database will be populated with (default 1,000,000)
## Examples
This is an example how to blast with 1000 worker goroutines, chunking from database by 100000
with silent mode and populate/seed database with 1 milion rows
  ```
  cd email-blaster
  go run . -w 1000 -c 100000 -r 1000000 -s
  ```
## Test
  ```
  go test ./... -v -bench=.
  ```
## Structure
The structure itself is simple, as usually I start my project simple and make s clearer, 
more nested struture as I go.  
In this case there is a single package, namely, the root package **pkg** which consists of:

- emailblaster - the main component, it utilizes the database and worker nodes in order to blast
high thoughput email stream
- repo - all database related functions
- sender - email-sender client wrapper
- types - all types/models for the applications, namely contains a model of the table and one for the business logic
- workerpool - the workerpool model, helps in creating high throughput stream

Another side package is the **dev** package, contains development related functionalities.

- db - contains special database operations such as setup, teardown, seeds

## Design Patterns
1) concurrency, utilizing message passing to a large pool of worker nodes
2) dependency injection, and one way dependency, my main package requires all dependency and injects
it to each one accordingly - Clean Code by Robert C. Martin
3) functional configuration / options - https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
