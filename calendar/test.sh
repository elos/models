#!/bin/bash

rm out.coverprofile
go test -race -cover -coverprofile=out.coverprofile && go tool cover -html=out.coverprofile
