package main

import (
	"errors"
	"math/rand"
)

// Quote represents a motivational quote with its author
type Quote struct {
	Text   string
	Author string
}

// ErrNoQuotes is returned when attempting to select from an empty quote list
var ErrNoQuotes = errors.New("no quotes available")

// defaultQuotes contains the built-in collection of quotes
// Used as fallback when ~/.quotes.json doesn't exist or is invalid
var defaultQuotes = []Quote{
	{Text: "Be the change you wish to see in the world", Author: "Gandhi"},
	{Text: "Code is poetry", Author: "Unknown"},
	{Text: "The only way to do great work is to love what you do", Author: "Steve Jobs"},
	{Text: "Innovation distinguishes between a leader and a follower", Author: "Steve Jobs"},
	{Text: "Stay hungry, stay foolish", Author: "Steve Jobs"},
	{Text: "Simplicity is the ultimate sophistication", Author: "Leonardo da Vinci"},
	{Text: "Talk is cheap. Show me the code", Author: "Linus Torvalds"},
	{Text: "First, solve the problem. Then, write the code", Author: "John Johnson"},
	{Text: "Any fool can write code that a computer can understand. Good programmers write code that humans can understand", Author: "Martin Fowler"},
	{Text: "The best error message is the one that never shows up", Author: "Thomas Fuchs"},
	{Text: "Programs must be written for people to read, and only incidentally for machines to execute", Author: "Harold Abelson"},
	{Text: "Experience is the name everyone gives to their mistakes", Author: "Oscar Wilde"},
	{Text: "In order to be irreplaceable, one must always be different", Author: "Coco Chanel"},
	{Text: "Java is to JavaScript what car is to Carpet", Author: "Chris Heilmann"},
	{Text: "Knowledge is power", Author: "Francis Bacon"},
	{Text: "Sometimes it pays to stay in bed on Monday, rather than spending the rest of the week debugging Monday's code", Author: "Dan Salomon"},
	{Text: "Perfection is achieved not when there is nothing more to add, but rather when there is nothing more to take away", Author: "Antoine de Saint-Exupery"},
	{Text: "Code never lies, comments sometimes do", Author: "Ron Jeffries"},
	{Text: "Before software can be reusable it first has to be usable", Author: "Ralph Johnson"},
	{Text: "Simplicity is the soul of efficiency", Author: "Austin Freeman"},
	{Text: "Make it work, make it right, make it fast", Author: "Kent Beck"},
	{Text: "Walking on water and developing software from a specification are easy if both are frozen", Author: "Edward V. Berard"},
	{Text: "If debugging is the process of removing software bugs, then programming must be the process of putting them in", Author: "Edsger Dijkstra"},
	{Text: "Measuring programming progress by lines of code is like measuring aircraft building progress by weight", Author: "Bill Gates"},
	{Text: "Controlling complexity is the essence of computer programming", Author: "Brian Kernighan"},
	{Text: "The most disastrous thing that you can ever learn is your first programming language", Author: "Alan Kay"},
	{Text: "There are only two hard things in Computer Science: cache invalidation and naming things", Author: "Phil Karlton"},
	{Text: "The function of good software is to make the complex appear to be simple", Author: "Grady Booch"},
	{Text: "Testing leads to failure, and failure leads to understanding", Author: "Burt Rutan"},
	{Text: "Truth can only be found in one place: the code", Author: "Robert C. Martin"},
	{Text: "Optimism is an occupational hazard of programming: feedback is the treatment", Author: "Kent Beck"},
	{Text: "It's not a bug – it's an undocumented feature", Author: "Anonymous"},
	{Text: "One of my most productive days was throwing away 1000 lines of code", Author: "Ken Thompson"},
	{Text: "Deleted code is debugged code", Author: "Jeff Sickel"},
	{Text: "The best way to predict the future is to invent it", Author: "Alan Kay"},
	{Text: "Every great developer you know got there by solving problems they were unqualified to solve until they actually did it", Author: "Patrick McKenzie"},
	{Text: "The most important property of a program is whether it accomplishes the intention of its user", Author: "C.A.R. Hoare"},
	{Text: "Hofstadter's Law: It always takes longer than you expect, even when you take into account Hofstadter's Law", Author: "Douglas Hofstadter"},
	{Text: "The only way to learn a new programming language is by writing programs in it", Author: "Dennis Ritchie"},
	{Text: "It's harder to read code than to write it", Author: "Joel Spolsky"},
	{Text: "Software and cathedrals are much the same – first we build them, then we pray", Author: "Anonymous"},
	{Text: "Give someone a program, you frustrate them for a day; teach them how to program, you frustrate them for a lifetime", Author: "David Leinweber"},
	{Text: "The computer was born to solve problems that did not exist before", Author: "Bill Gates"},
	{Text: "Documentation is a love letter that you write to your future self", Author: "Damian Conway"},
	{Text: "There are two ways to write error-free programs; only the third one works", Author: "Alan J. Perlis"},
	{Text: "You might not think that programmers are artists, but programming is an extremely creative profession. It's logic-based creativity", Author: "John Romero"},
	{Text: "The best programmers are not marginally better than merely good ones. They are an order-of-magnitude better", Author: "Randall E. Stross"},
	{Text: "People think that computer science is the art of geniuses but the actual reality is the opposite", Author: "Edsger Dijkstra"},
	{Text: "The value of a prototype is in the education it gives you, not in the code itself", Author: "Alan Cooper"},
	{Text: "Fix the cause, not the symptom", Author: "Steve Maguire"},
	{Text: "It's not at all important to get it right the first time. It's vitally important to get it right the last time", Author: "Andrew Hunt"},
	{Text: "Simplicity is prerequisite for reliability", Author: "Edsger Dijkstra"},
	{Text: "Don't comment bad code – rewrite it", Author: "Brian Kernighan"},
	{Text: "Low-level programming is good for the programmer's soul", Author: "John Carmack"},
	{Text: "We have to stop optimizing for programmers and start optimizing for users", Author: "Jeff Atwood"},
	{Text: "The best thing about a boolean is even if you are wrong, you are only off by a bit", Author: "Anonymous"},
	{Text: "Without requirements or design, programming is the art of adding bugs to an empty text file", Author: "Louis Srygley"},
	{Text: "The trouble with programmers is that you can never tell what a programmer is doing until it's too late", Author: "Seymour Cray"},
	{Text: "Most good programmers do programming not because they expect to get paid or get adulation by the public, but because it is fun to program", Author: "Linus Torvalds"},
	{Text: "Always code as if the guy who ends up maintaining your code will be a violent psychopath who knows where you live", Author: "John Woods"},
	{Text: "In theory, theory and practice are the same. In practice, they're not", Author: "Yogi Berra"},
	{Text: "Programming today is a race between software engineers striving to build bigger and better idiot-proof programs, and the universe trying to produce bigger and better idiots", Author: "Rick Cook"},
	{Text: "A language that doesn't affect the way you think about programming is not worth knowing", Author: "Alan J. Perlis"},
	{Text: "Good code is its own best documentation", Author: "Steve McConnell"},
	{Text: "Weeks of coding can save you hours of planning", Author: "Unknown"},
	{Text: "Debugging is twice as hard as writing the code in the first place", Author: "Brian Kernighan"},
	{Text: "If you can't explain it simply, you don't understand it well enough", Author: "Albert Einstein"},
	{Text: "The sooner you start to code, the longer the program will take", Author: "Roy Carlson"},
	{Text: "Quality is never an accident; it is always the result of intelligent effort", Author: "John Ruskin"},
	{Text: "The greatest enemy of knowledge is not ignorance, it is the illusion of knowledge", Author: "Stephen Hawking"},
}

// SelectRandom returns a random quote from the provided slice using the given seed
// for reproducible randomness. Returns ErrNoQuotes if the quotes slice is empty.
func SelectRandom(quotes []Quote, seed int64) (Quote, error) {
	if len(quotes) == 0 {
		return Quote{}, ErrNoQuotes
	}

	source := rand.NewSource(seed)
	rng := rand.New(source)
	index := rng.Intn(len(quotes))

	return quotes[index], nil
}
