# Wordle Solver

![workingscreenshot.png](docs%2Fworkingscreenshot.png)

## Run as docker container

The wordle solver runs as an docker container.

So you can use it even if you don't have the Go toolchain installed.

```bash
./script/build_and_run.sh
```

## For goland / terminal users:

You can just build and run the main file in the run configuration. There's no special environment variables needed.

```bash
./cmd/main.go
```

## To use:

The program will give you an initial suggestion. It will always be "alert".

It will ask for input. The input is to be entered without any spaces.

You will input numbers which represent the feedback given by wordle.

```
Not in word (grey): 0
In word but not in position (orange): 1
In word and in position (green): 2
```

For example, the input for the included screenshot is:

```
Suggestion: alert 
Enter input: 10010
Suggestion: sonar 
Enter input: 00011
Suggestion: braid 
Enter input: 11100
Suggestion: rumba 
...

No need to enter any more input as it's been solved!

```

## How it works

I needed a dictionary list, to filter only the 5 letter words in the dictionary.

This would be trivial to filter for, using a script.

Unfortunately this was extremely hard to find, as many of the dictionary lists available online are filled with non-words 
and special characters, as they are typically used for password cracking.

I found a list of 5 letter words that wordle uses here, I'm not sure whether this is all of the 5 letter words in the English
dictionary or not:

[Wordle list](https://gist.githubusercontent.com/cfreshman/a03ef2cba789d8cf00c08f767e0fad7b/raw/28804271b5a226628d36ee831b0e36adef9cf449/wordle-answers-alphabetical.txt)

## Word suggester

The game state struct starts with a WordList of all possible words.

As more input is processed, the WordList gets reduced.

## The heuristic

The word suggester finds the best word by recurrently:
 - counting up how common each letter is in the WordList, repeated letters within a word are not counted.
 - then finding the word with the highest score based on the count of letters.
 - parsing the input and removing all the words that don't correspond to the input.

This approach maximises the amount of words removed from the WordList with each guess, until you're left with the answer.   

