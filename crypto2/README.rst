Zden's crypto puzzles, #2
=========================

See http://crypto.haluska.sk/ for the task details.

Here's my solver. I didn't find the solution, though -- the finishing 
touches are written after Zden has published the solution. Oh well.

`crypto2.go` parses the puzzle image and prints decoded raw bits from the 
puzzle picture. It also writes a png file with a visualization of the 
decoding, for manual verification.

`d.rb` reads the the raw bits and prints two candidate values for the 
private key.

`c.py` is a private key validity checker, pulled from stackoverflow.

So, ``for i in `./crypto2 crypto2.png o.png | ./d.rb`; do ./c.py $i; done`` 
prints the private key.
