#!/usr/bin/env bash

pushd ../ui/static/javascripts && npm install --save-dev webpack webpack-cli && npm install --save react react-dom && npm install --save-dev @types/react @types/react-dom && npm install --save-dev typescript ts-loader source-map-loader && npm install style-loader css-loader --save