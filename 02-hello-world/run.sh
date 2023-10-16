#!/bin/bash
clear
bat $0 --line-range 5:
echo ""
extism call hello.wasm hello \
  --input "Bob Morane" \
  --wasi true
echo ""
