# Assets
This directory contains all non .go assets that need to be used for the game.

Using the bin-asset application we can turn these into binary and not have to
distribute them with the compiled game.

This process is all done in the `run.sh` script.  When you want to access the
asset file in code, use `assets.Asset("assets/audio/test.mp3")` (example).  This
will the byte slice and the `os.FileInfo`.
