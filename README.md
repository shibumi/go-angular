[![Build Status](https://travis-ci.com/Shpota/go-angular.svg?branch=master)](https://travis-ci.com/Shpota/go-angular)

# Go-Angular

This project is a modified fork of https://github.com/shpota/go-angular.

The goal of shpotas' project was to showcase Go + Angular.
In this modified version I have concentrated on showcasing Go 1.16's
new `embed` module.

I have modified the following parts:

1. I have substitutes the postgres DB against a sqlite DB
2. I have removed the static web directory and introduced the `embed` module instead.
3. I have modified the outputPath for angular `ng build --prod` moves the assets to `../server/static` now.
4. The Dockerfile makes use of Google's distroless images now.

Below is the original introduction of Shpota's README:

## A simple web application written with Go and Angular

I implemented this application while evaluating Go. 
On the back end side, I used 
[gorilla/mux](https://github.com/gorilla/mux) for 
routing, [Gorm](https://github.com/jinzhu/gorm) as an 
ORM engine and 
[google/uuid](https://github.com/google/uuid) 
for UUID generation. On the front end side, I used 
[Angular 8](https://angular.io/) and 
[Angular Material](https://material.angular.io/).

![Showcase](showcase/showcase.gif)

