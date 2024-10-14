# goqteternity2
A golang playing around with using Simmulated Anneling on Eternity2 puzzle.

## Introduction

This was another step on my obsession with the Eternity2 puzzle (there are more).
puzzle [https://en.wikipedia.org/wiki/Eternity_II_puzzle](https://en.wikipedia.org/wiki/Eternity_II_puzzle)
Essentially the puzzle is about placing tiles on a 16 by 16 grid in such a way that all edges to the tiles match.
There is also the criteria that one of the tiles must occupy a defined location on the grid.
All tiles may be rotated.

This application played around with Simulated Annealing the tiles energy of a tile being
given by its lack of fitness with its neighbours. As the system slowly cools the tiles that best fit with their 
neighbours will be the ones that are most likely to be moved. The application also has a
GUI built using a binding to qml. Unfortunately times have moved on the binding
is no longer maintained and the code is not compatible with later versions of QT5. 

It was fun (to watch the "tile" crystals grow) while it lasted. 

## Build and test

In the unlikely event that this code will interest anyone.

As stated above the code is no longer compatible with the current version of QT5.
So it will not compile or run out of the box.

## Contribute

I would be amazed if anyone was interested in contributing...
but if anyone were to feel so inclined I'm open to pull requests.
Maybe sometime I'll come back to this and update the code to work with a different UI.
[CogentCore](https://github.com/cogentcore) looks interesting.