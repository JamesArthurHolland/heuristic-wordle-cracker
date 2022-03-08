We need a wordlist. //TODO dictionary words aren't dictionary. They're from the web.
I found one on github. Let's wget it.

`wget https://gist.githubusercontent.com/cfreshman/a03ef2cba789d8cf00c08f767e0fad7b/raw/28804271b5a226628d36ee831b0e36adef9cf449/wordle-answers-alphabetical.txt`

// TODO remove
We need to filter all the words that are 5 letters long.

DiskUsage shows the file is 5MB big:

`du -shm words.txt`

We're going to hold a fraction of this in memory (the 5 letter words) before we flush the 
buffer. This is more than manageable for a bufferedwriter in golang. Your IDE or text editor might shit 
itself trying to open words.txt, but don't worry about that. Printing text on screen is slow.
Reading text in memory and filtering (without output) is fast.

Debug script to take words.txt and spit out 5 letter words into five_letter_words.txt:

`cmd/debug/five_letter_filter/main.go`

We need lookup tables. This will allow us to input an individual character and get all the 
words that have that character in them. We will expand this letter to remove letters that 
aren't at certain positions, based on the feedback we get from the wordle puzzle.

We need to generate the lookup table, write it to a file for later use, then reread it as part
of our main program.


