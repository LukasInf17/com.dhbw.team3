#!/bin/bash

git pull
buffalo build
systemctl restart invit
