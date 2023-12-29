import random

from textual import events
from textual.app import App
from textual.containers import Container
from textual.widgets import Header, Label, Input, Button, Digits, Footer


def get_random_problem():
    a = random.randint(1, 20)
    b = random.randint(1, 20)
    answer = 0
    choice = random.choice(["+", "-", "*"])
    if choice == "+":
        answer = a + b
    elif choice == "-":
        a, b = max(a, b), min(a, b)
        answer = a - b
    elif choice == "*":
        answer = a * b
    if (a + b + random.randint(1, 10)) % 2 == 0:
        choice = "÷"
        answer = a
        a = a * b
    return a, b, choice, answer


class CoolMaths(App):
    CSS_PATH = "coolmaths.css"
    BINDINGS = [("ctrl-c","Close", "Close")]

    def __init__(self):
        super().__init__()
        self.question = Digits(id="numbers")
        self.answer = Label(id="feedback")
        self.input = Input(placeholder="Type your answer here:", id="userinput")
        a, b, choice, answer = get_random_problem()
        self.questions = f"{a} {choice} {b}"
        self.answers = str(answer)

    def compose(self):
        yield Header()
        self.question.update(self.questions)
        with Container(id="question"):
            yield self.question
            yield Container(classes="filler")
            yield self.input
            yield Container(classes="filler")
            yield self.answer
            yield Button("Submit")
        yield Footer()

    def on_key(self, event: events.Key):
        if event.key == "enter":
            self.validate()
        if event.key == "q" or event.key == "Q":
            self.quit()

    def on_button_pressed(self):
        self.validate()

    def validate(self):
        data = " ".join(input.value for input in self.query(Input))
        if data == self.answers:
            self.answer.update(f"Yes. {self.questions} = {self.answers} . You got it! Let' try another one.")
            self.answer.styles.color = "green"
        else:
            self.answer.update(f"Nope. {self.questions} is not {data}. It is {self.answers}. Please continue practicing.")
            self.answer.styles.color = "red"

        a, b, choice, answer = get_random_problem()
        self.questions = f"{a} {choice} {b}"
        self.answers = str(answer)
        self.question.update(self.questions)
        self.input.clear()


if __name__ == "__main__":
    CoolMaths().run()