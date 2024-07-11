# Clean Architecture Implementation in Go

This is the example clean architecture written in GoLang using [Gin](https://gin-gonic.com/) as the web framework, [dig](https://github.com/uber-go/dig) as the dependency injector and [GORM](https://gorm.io/index.html) as the ORM.

### Why? Because DRY

With new microservices being created, having to create the boilerplate codes before implementating the actual business logics is time and energy consuming. So, I decided to create a boilerplate codes that I can reuse because DRY.

### Future Plan

I am planning to make a CLI wrapper for this project where you can create a new module or service with one command (or more than that who knows?) and many other useful and common commands. The link to that project is [here]().

### Project Architecture Explanations

!TODO

### TODOs

- [ ] #1
- [ ] #2
- [ ] #3
- [ ] #4

## API Documentation

Link to the API Documentation is [here](https://g5s5ah4cwz.apidog.io).

> [!IMPORTANT]
> Before running endpoints in that documentation, you have to run this server codes first.

> [!CAUTION]
> This project is using MySQL as the database. So, you need to change the dialect and driver of the GORM if you are planning to run with another RDB.

## Contributions

Everyone is welcomed to blame and contribute to this project if it is for the purpose of making the structural of the project better. But remember this is not a framework or library.
