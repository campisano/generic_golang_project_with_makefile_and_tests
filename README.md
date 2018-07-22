###### *Tests:* linux: [![Build Status](https://travis-ci.org/campisano/golang-sample-project-with-makefile-and-test.svg?branch=master "Linux build")](https://travis-ci.org/campisano/golang-sample-project-with-makefile-and-test)
---


# golang-sample-project-with-makefile-and-test

A minimal golang sample project configured to use a Makefile to build and test.

The Makefile is configured to run the test as its main target. The result of such run is the same reported in Travis-CI site [![Build Status](https://travis-ci.org/campisano/golang-sample-project-with-makefile-and-test.svg?branch=master "Linux build")](https://travis-ci.org/campisano/golang-sample-project-with-makefile-and-test).

Before the test, the make command install Gometalinter and Testify that are used to check the code and tests.
