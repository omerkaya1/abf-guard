# TODO list:

## Necessary
1. Add Prometheus metrics to the project;
2. Add more generalised DB transactions;
3. Add a script that allows for a safe start of integration tests (using Sleep to wait for Postgres is embarrassing);

## Not so necessary
1. Add error wrapping capabilities to the project;

## Proposed
1. Let's re-evaluate the necessity of some imports to this project. By removing some of them we can significantly reduce 
the project's build time and the size of the binary (or at least let's experiment with that);
2. Should we restrict the user rights within the run application? If so, make sure we do that when we create a Docker
container.
