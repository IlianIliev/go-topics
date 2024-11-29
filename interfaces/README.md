# Interfaces, Testing and Mocking


### What are interfaces?

Interfaces are abstraction types that define set of methods and their signatures without providing any implementation.
You can think of it as a contract that your module should follow.


### Why do we need interfaces?

By defining such contract, you allow different modules in your system to communicate without exposing their internal implementation details.
By using interfaces we can implement the Dependency Inversion Principe (the D in SOLID) and effectively decouple our codebase im a way that is not only easy to test but also easy to modify and refactor.


### How does interfaces improve testability?

By using interfaces to decouple our code, we can easily create unit tests that verify only the behaviour of a single module running in isolation. This way we not only depend on the implementation details anymore, but also there is no need recreate the whole environment around such implementations e.g. databases, networks etc.


### What about mocking?

Mocking is a technique that allows you to test the behaviour of a module by creating a fake implementation of the interface.


### Are unit tests enough?
With unit tests you can test module behavior even before the module is implemented. They are easy to write, fast to run and provide good coverage. They are also great help when refactoring.

However, they are not enough to guarantee that your code is bug free. To ensure a low risk of bugs, it is recommended to use a combination of unit tests and integration tests. 

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

