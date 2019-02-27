# Trigram Text Generator
Tech test written for Geckoboard.

A Go program that will "learn" from the text it receives, and generate random text using trigrams.

## Endpoints

The application expose an HTTP interface with two endpoints:

- POST `/learn` will "teach" the application about a body of text, sent as the POST body with Content-Type: text/plain.

- GET `/generate` will return randomly-generated text based on all the trigrams that have been learned since starting the program.

## Notes
- The application will generate the number of sentences as defined in the `config.json` file.

## Next Steps
- Implement Go routines for quicker processing (mutex support has already been added)
- Investigate and implement E2E tests and middleware tests

## Running The Application
Two scripts have been created to make the process of running this application easier. 

- `run.sh` will run the application with the `--race` flag. This is intended to be for development purposes. 
- `build.sh` will create a `build` folder, which will have `darwin` and `linux` compiled versions of the application. This is intended to be for production purposes. 

The application will run according to the `listenAddress` specified in the `config.json` file, in the project root (defaults to `localhost:3005`).

## Tests
A script (`test.sh`) has been added which will run all of the tests in the project. 