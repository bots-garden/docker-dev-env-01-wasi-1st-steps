#!/bin/bash
clear
bat $0 --line-range 5:
echo ""
extism call host-functions.wasm say_hello \
  --input "😀 Hello World 🌍! (from TinyGo)" \
  --log-level info \
  --allow-host "*" \
  --set-config '{"route":"https://jsonplaceholder.typicode.com/todos/1"}' \
  --wasi true
echo ""
