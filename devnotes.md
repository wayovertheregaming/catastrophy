# Dev Notes
Notes for devs.

## Running/building
I have created a run.sh file in the root of this project.  You should be able to
run the file from command line:
```bash
sh ./run.sh
```

On Linux you can first run `chmod +x run.sh` (you only ever need run this once)
then everytime you want to run (might work on Mac too...?):
```bash
./run.sh
```

This script installs all dependencies and makes sure you're running in debug
mode.

Nearer the end of the weekend I'll also create build script.

## Git workflow
We'll be using `master` as the release branch.  Use `devel` as the branch to
branch from.  Branch into your bug/feature branch from `devel`, make commits,
push then PR into develop.

Assign PR to both other devs.

### Assets
All assets (graphics, audio, CSV etc), should be stored in the `assets` directory.
This will be used by [go-bindata](github.com/kevinburke/go-bindata) (you should
go look at that - it's really simple and I'll do the setup for you).  It takes
files and turns them into `[]byte` so we can compile them and don't need to
distribute extra files.

## Game details
### Environment
The game takes place in a large house and garden.  This is made of three maps,
which have multiple rooms/areas each (this is the same map and will not reload,
but we will load individual maps as different levels):
 1. Upper floor
   - Bathroom
   - Master bedroom
   - Guest bedroom
   - Office (sort of library)
 2. Ground floor
   - Kitchen
   - Lounge
   - Dining room
   - Entrance hall
   - Conservatory
 3. Garden

### Gameplay
The player will control a cat.  The cat will get tokens (called "trophies") by
doing one of two things:
 1. Getting an achievement.  This is done by doing something 'annoying' for the
 first time; examples:
  - Getting in the humans way
  - Pissing on the sofa
  - Hiding in a shoe
 2. Completing a minigame.  These will be launched from specific areas of the
 house; examples:
  - Garden - duck hunt clone
  - Kitchen - catch falling food in a bowl

#### Minigames
We're using minigames so that we can easily add more gameplay to the game,
without it looking like we've missed bits.  Ideally we'll have at least one
minigame per room, this may not be possible due to time restaints.

#### Shadow realm
Having gained the first trophy, the player should be told they can sleep; this
will transport them to an identically laid out house, but looking different -
maybe less decorations, maybe it is all dark and purple (could use a shader to
reduce graphics needed).

In the shadow realm the player will meet the cat god, represented by a large
ball of wool.  The god will allow the player to exchange/sacrifice trophies to
gain abilities.  Abilities will be needed to progress through the game; i.e.
minigames/achievements/rooms will be locked off until certain sacrifices have
been made.

#### Trophies
Trophies will be unique.  Minigames are replayable, but you'll always get the
same trophy from playing it.  The cat god will accept repeat sacrifices (the
same trohpy twice) but will not unlock anything after the first one.

E.g. you gain a 'feather' trophy from catching birds in the garden; this, when
sacrificed gives access to the pond in the garden.  If the player goes back and
collects the 'feather' trophy again, the cat god will accept it but nothing will
be unlocked.

#### Winning the game
This is subject to having enough time.  If we do not get enough time, the boss
battle will be skipped and jump straight to the win screen.

When all trophies have been sacrificed, the player will be told something has
changed in the shadow realm.  They will find a cow, instead of the cat god.  The
dialogue will explain that while most people think of dogs as the enemy of cats,
it is infact cows.  Some final battle will occur; this could be:
 1. Classic side-view fight (minigame?)
 2. Combination of all preplayed minigames
 3. Dialogue choices, leading to slightly different outcomes
