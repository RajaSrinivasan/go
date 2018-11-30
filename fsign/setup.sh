#!/bin/bash
openssl genrsa -aes128 -passout pass:RecoveringHearts -out fd.key 2048
openssl rsa -passin pass:RecoveringHearts -in fd.key -pubout -out fd-public.key
#openssl rsa -nodes -in fd.key -pubout -out fd-public.key