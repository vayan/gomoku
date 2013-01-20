# how to install

* `mkdir Go`

* ajoute ce dossier a l'env `export GOPATH='path/to/Go'

* `cd Go`

* `mkdir bin`

* Ajoute bin a ton PATH `export PATH=$PATH:/path/to/Go/bin`

* `mkdir src`

* `mkdir pkg`

* `go get code.google.com/p/go.net/websocket`

* `cd src`

* `git clone git@bitbucket.org:vayan/gomoku.git`

# Pour compiler / lancer le serv

* `cd gomoku/gomoserv`

* `go install gomoku/gomoserv`

* `gomoserv`

# Pour compiler / lancer l'AI

* bien avoir lancer le server + un joueur qui attend une AI

* `go install gomoku/gomokai`

* `gomokai`