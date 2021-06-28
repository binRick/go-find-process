nodemon -w ../../. -w . -e go,sh --signal SIGKILL -x sh -- -c "./run.sh||true"
