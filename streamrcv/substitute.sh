#!/bin/bash

pattern='${{=PASSKEY=}}'
replacement="${PASSKEY}"
path='/usr/local/nginx/conf/nginx.conf'

# substitutes values in config files before starting nginx.
