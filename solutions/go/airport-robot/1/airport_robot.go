package airportrobot
import "fmt"
// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
const Def = "I can speak"
type Greeter interface {
    LanguageName() string
    Greet(string) string
}


func SayHello(name string, greeter Greeter) string {
    return greeter.Greet(name)
}
// We're nowhere mentioning that the type structs germanGreeter, italianGreeter and portugueseGreeter are using the Greeter interface, that is done implicitely by GoLang as soon as we define functions with names of the functions defined in the Greeter interface.
// German Greeter
type germanGreeter struct {}
func (g germanGreeter) LanguageName () string {
    return "german"
}

func (g germanGreeter) Greet(name string) string {
    return fmt.Sprintf("Hallo %s!", name)
}
// Italian Greeter
type Italian struct {}
func (i Italian) LanguageName () string{
    return "Italian"
}
func (i Italian) Greet(name string) string {
    return fmt.Sprintf("%s %s: Ciao %s!",Def, i.LanguageName(), name)
}

// Portuguese Greeter
type Portuguese struct{}
func (p Portuguese) LanguageName() string{
    return "Portuguese"
}
func (p Portuguese) Greet(name string) string{
	return fmt.Sprintf("%s %s: Olá %s!", Def, p.LanguageName(), name)
    }