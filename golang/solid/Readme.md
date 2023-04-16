# SOLID

Harvested from chatGPT

## Single Responsibility Principle

A module or class should have only one reason to change. 
This means that a module or class should have only one responsibility or job to do, 
and all its functions and methods should be related to that responsibility.

## Open Close

Software entities (classes, modules, etc.) should be open for extension but closed for modification.
In other words, we should be able to add new functionality to a system without changing
the existing code, which can help to reduce the risk of introducing bugs and make it easier
to maintain and evolve the code over time.

## Liskov Substitution

objects of a superclass should be able to be replaced with objects of a subclass without 
affecting the correctness of the program. In other words, a subclass should be able to 
be used wherever its superclass is used, without causing any unexpected behavior.

## Interface Segregation

client code should not be forced to depend on interfaces that it does not use. 
In other words, interfaces should be designed to be as small and focused as possible, 
to minimize the impact of changes to the implementation and make it easier to 
write code that depends on them.


## Dependency Inversion Principle

High-level modules should not depend on low-level modules, both should depend on abstractions.
Additionally, abstractions should not depend on details, but details should depend on 
abstractions. This principle is closely related to the concept of dependency injection, 
which is a way of implementing the DIP in practice.