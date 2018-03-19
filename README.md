# Angular + Golang Starter Kit

**Note:** *I have shifted my focus to development of a CLI tool which can automatically generate a project like this and a lot more things like serve and production build, if you are interested in testing it out or would like to contribute, you may find the project here: https://github.com/anshap1719/nggo. It's written in Golang*

## Prerequisites
- You have Angular CLI & npm installed.
- You have Go installed and GOPATH set up properly.

## Before You Begin
Instead of running `npm install` as you normally would, you need to run `npm run install-dependencies` which will run `npm install` for you in addition to installing the required Go dependencies for this boilerplate project.

## Run The Project
`npm start` will run both the Go server (by default at port 4201) and Webpack Dev Server for Angular (at port 4200). `npm start` basically runs the `serve.sh` bash script which in turn starts `ng serve` and Go live server parallely. You can access the app by visiting `http://localhost:4200`.

**Note:** *The console output for both the client and the server will be printed in the same console window.*

#### This boilerplate takes care of *Allow-Cross-Origin-Access* errors by default by explicitely allowing the client's address to make cross origin requests

This also means that you will have to edit the code to make it work for production builds. Currently, it is also required that both the client and server be built separately. This boilerplate currently provides a stage only for development purposes with no automatic combined builds. I am working on making it work properly with build and also on adding some customizabilty, so keep an eye out for those.
