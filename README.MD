# Clean Architecture Implementation in

<div align="center">
    <img src="./images/golang_logo.png" alt="drawing" width="300"/>
</div>

## About the Project

This is the example clean architecture written in GoLang using [Gin](https://gin-gonic.com/) as the web framework, [dig](https://github.com/uber-go/dig) as the dependency injector and [GORM](https://gorm.io/index.html) as the ORM.

## Why? Because DRY

With new microservices being created, having to re-create the boilerplate codes before implementating the actual business logics is time and energy consuming. So, I decided to create a boilerplate codes that I can reuse and thought to myself why not make it public. So, it is here.

## Future Plan

I am planning to make a CLI wrapper for this project where you can create a new module or service with one command (or more than that who knows?) and many other useful and commands that solves repetative tasks. The link to that project is [here - empty for now]().

## Project Architecture Explanations

!TODO

## Packages used in this project

- [Gin](https://gin-gonic.com/) - The fastest full-featured web framework for Go. Crystal clear.
- [dig](https://github.com/uber-go/dig) - A reflection based dependency injection toolkit for Go.
- [GORM](https://gorm.io/index.html) - The fantastic ORM library for Golang
- [BurntSushi/toml](https://github.com/BurntSushi/toml) - TOML parser for Golang with reflection.

## TODOs

- [ ] [Add Basic CRUD](https://github.com/Kybton/go-gin-clean-architecture/issues/1)
- [ ] [Add Validations](https://github.com/Kybton/go-gin-clean-architecture/issues/2)
- [ ] [Add Unit Testing](https://github.com/Kybton/go-gin-clean-architecture/issues/3)
- [ ] [Add Docs for Project structure](https://github.com/Kybton/go-gin-clean-architecture/issues/4)
- [ ] [Change configs Loading Package](https://github.com/Kybton/go-gin-clean-architecture/issues/5)

## API Documentation

Link to the API Documentation is [here](https://g5s5ah4cwz.apidog.io).

> [!IMPORTANT]
> Before running endpoints in that documentation, you have to run this server codes first.

> [!CAUTION]
> This project is using MySQL as the database. So, you need to change the dialect and driver of the GORM if you are planning to run with another RDB.

## Contributions

Everyone is welcomed to blame and contribute to this project if it is for the purpose of making the structural of the project better. But remember this is not a framework or library.
