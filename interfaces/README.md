# Interfaces, Testing and Mocking

### Overview

In this session, I demonstrate how to use interfaces and the gomock tool in Go. Screencast is available at https://youtu.be/lzOhQ_xaIds


### What are interfaces?

Interfaces are abstraction types that define a set of methods and their signatures without providing any implementation.
You can think of it as a contract that your module should follow.


### Why do we need interfaces?

By defining such a contract, you allow different modules in your system to communicate without exposing their internal implementation details.
By using interfaces we can implement the Dependency Inversion Principe (the D in SOLID) and effectively decouple our codebase in a way that is not only easy to test but also easy to modify and refactor.


### How do interfaces improve testability?

By using interfaces to decouple our code, we can easily create unit tests that verify only the behavior of a single module running in isolation. This way we not only depend on the implementation details anymore, but also there is no need to recreate the whole environment around such implementations e.g. databases, networks, etc.


### What about mocking?

Mocking is a technique that allows you to test the behavior of a module by creating a fake implementation of the interface.


### Are unit tests enough?
With unit tests, you can test module behavior even before the module is implemented. They are easy to write, fast to run, and provide good coverage. They are also a great help when refactoring.

However, they are not enough to guarantee that your code is free of bugs. To ensure a low risk of bugs, a combination of unit tests and integration tests is recommended.

----

Related articles:

- [Interfaces](https://golang.org/doc/effective_go.html#interfaces)
- [gomock](https://github.com/uber-go/mock)

The gomock alias I am using:

```bash
gomock() {
    if [[ $# -gt 0 ]]
      then
        fileName=$(echo $1 | sed -e "s/.go$//")
        packageName=$(grep '^package .*$' $1 | sed -E 's/package //')
        mockgen -source=$fileName.go -package=$packageName -destination=${fileName}_mock.go
        echo "Done."
      else
        echo No file name provided
    fi
}
```

