# Wiki crawler

[![Go Report Card](https://goreportcard.com/badge/github.com/goujonbe/wiki-crawler)](https://goreportcard.com/report/github.com/goujonbe/wiki-crawler) ![CI](https://github.com/goujonbe/wiki-crawler/workflows/CI/badge.svg) [![codecov](https://codecov.io/gh/goujonbe/wiki-crawler/branch/main/graph/badge.svg)](https://codecov.io/gh/goujonbe/wiki-crawler)

This repository contains the code of one of my side projects. The aim of this side project is to explore the relationships between Wipipedia pages. It is inspired by a game, where, given a start page, you need to reach as fast possible another page only using links in the page body.

This project leverages [DGraph](https://dgraph.io/) as a Graph database to store the relationships between the pages.
