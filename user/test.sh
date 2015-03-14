#!/bin/bash

go test -race -cover -coverprofile=out.coverprofile && go tool cover -html=out.coverprofile
