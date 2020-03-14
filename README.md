# email-blaster
## Install
  ```
  git clone https://github.com/CommoDor64/email-blaster.git
  cd email-blaster
  go get ./...
  ```
## Run
Run with default configurations  
  ```
  cd email-blaster
  go run .
  ```
## Usage
the application has couple of configurable flags
- -w - size of worker pool (default 500)
- -c - chunk size, i.e how many record are pulled from DB on each iteration (default 10000)
- -s - is silent, if specified, removes debug output (default false)
- -r - rows number, defines how many rows the database will be populated with
## Examples
This is an example how to blast with 1000 worker goroutines, chunking from database by 100000
with silent mode and populate/seed database with 1 milion rows
  ```cd email-blaster
  go run . -w 1000 -c 100000 -s -r 1000000
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

## Notes
1) Each goroutine fetches a single payload to be send as email from the corresponding channel, I wanted to update the each row in the database whenever the email is successfuly sent, but in large numbers > 100,000 records and > 100 goroutines, the sqlite client doesn't functions properly. It has a poor support of concurrency.
### In hindsight
2) I would extend the model to have another worker nodes that update each record in the database in case of successful mail send

3) I would use postgresql in order to test real production grade throughput

## Design Patterns
1) concurrency, utilizing message passing to a large pool of worker nodes
2) dependency injection, and one way dependency, my main package requires all dependency and injects
it to each one accordingly - Clean Code by Robert C. Martin
3) functional configuration / options - https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
