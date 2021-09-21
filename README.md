# Alisazavr

## Overview

This project is used to play youtube video on Yandex.Station

## How to use

- Clone repo [alisazavr](https://github.com/outoffcontrol/alisazavr/tree/master)
- Edit **env.list** file
- ```docker build -t alisazavr:latest .```
- ```docker run -ti --rm --env-file ./env.list alisazavr:latest```
- Send youtube video link to telegram bot
- Awesome!