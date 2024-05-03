#!/bin/sh

fd -t f -x sed -i '' -E 's/(github\.com)\/.+\/(ollama)/\1\/ink-splatters\/\2/g'
