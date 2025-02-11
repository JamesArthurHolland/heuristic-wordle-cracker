# Wordle Solver

![workingscreenshot.png](docs%2Fworkingscreenshot.png)

## To build the project:

```bash
go build -o wordlesolver ./cmd/main.go
chmod +x wordlesolver
```

## To run:
```bash
./wordlesolver
```

## To use:

The program will give you an initial suggestion. It will always be "alert".

It will ask for input. The input is put in without any spaces.

You input the letter, followed by the number representing the feedback given by wordle.

```
Not in word (grey): 0
In word but not in position (orange): 1
In word and in position (green): 2
```

For example, the input for the included screenshot is:

```
a0l0e1r2t0
All the greys have a 0 following them, e is 1 because it is in the word but the wrong place
and r is 2 because it is in the correct place.
```


```
Suggestion: alert 
Enter input: a0l0e1r2t0
Suggestion: score 
Enter input: (we don't need to because it's been solved)
```

## How it works

I needed dictionary list to filter only the 5 letter words in the dictionary. This would be trivial to filter for.

Unfortunately this was extremely hard to find as many of the dictionary lists available online are filled with non-words 
and special characters, as they are typically used for password cracking.

I found a list of 5 letter words that wordle user here:

[Wordle list](https://gist.githubusercontent.com/cfreshman/a03ef2cba789d8cf00c08f767e0fad7b/raw/28804271b5a226628d36ee831b0e36adef9cf449/wordle-answers-alphabetical.txt)

### Word suggester

The game state struct starts with a WordList of all possible words.

As more input is processed, the WordList gets reduced.

##### The heuristic

The word suggester finds the best word by recurrently:
 - counting up how common each letter is in the WordList, repeated letters within a word are not counted.
 - then finding the word with the highest score based on the count of letters.
