#!/bin/bash

foo=$(go run . '%m月%d日 %Ja %JH:%M')
date +"$foo"
