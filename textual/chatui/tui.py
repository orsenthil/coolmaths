from textual.app import App, ComposeResult
from textual.widgets import Footer, Header, Placeholder, Input, Button, Static
from textual.containers import Container, Horizontal
from textual.widget import Widget
from textual.binding import Binding

from chat import Conversation


class FocusableContainer(Container, can_focus=True):
    ...


class MessageBox(Widget, can_focus=True):
    def __init__(self, text:str, role: str) -> None:
        self.text = text
        self.role = role
        super().__init__()

    def compose(self) -> ComposeResult:
        yield Static(self.text, classes=f"message {self.role}")


class ChatApp(App):
    TITLE = "chatui"
    SUB_TITLE = "ChatGPT directly in your terminal."
    CSS_PATH = "static/styles.css"

    BINDINGS = [
        Binding("q", "quit", "Quit", key_display="Q / CTRL-C"),
        ("ctrl-x", "clear", "Clear"),
    ]

    def compose(self) -> ComposeResult:
        yield Header()
        with FocusableContainer(id="conversation_box"):
            yield MessageBox(
                "Welcome to ChatGPT\\n"
                "Type your question, click enter or 'send' button "
                "and wait for the response.\\n"
                "At the botton you can find few more helpful commands.",
                role="info",
            )

        with Horizontal(id="input_box"):
            yield Input(placeholder="Enter your message", id="message_input")
            yield Button(label="Send", variant="success", id="send_button")

        yield Footer()

    def action_clear(self) -> None:
        self.conversation.clear()
        conversation_box = self.query_one("#conversation_box")
        conversation_box.remove()
        self.mount(FocusableContainer(id="conversation_box"))

    def on_mount(self) -> None:
        self.conversation = Conversation()
        self.query_one("#message_input", Input).focus()

    async def on_button_pressed(self) -> None:
        await self.process_conversation()

    async def on_input_submitted(self) -> None:
        await self.process_conversation()

    async def process_conversation(self) -> None:
        message_input = self.query_one("#message_input", Input)
        if message_input.value == "":
            return
        button = self.query_one("#send_button")
        conversation_box = self.query_one("#conversation_box")

        self.toggle_widgets(message_input, button)

        # Create a question message, and add it to the conversation and scroll down
        message_box = MessageBox(message_input.value, "question")
        conversation_box.mount(message_box)
        conversation_box.scroll_end(animate=False)

        # clean up the input without triggering events
        with message_input.prevent(Input.Changed):
            message_input.value = ""

        choices = await self.conversation.send(message_box.text)
        self.conversation.pick_response(choices[0])

        conversation_box.mount(
            MessageBox(choices[0].removeprefix("\\n").removeprefix("\\n"),
                       "answer")
        )

        self.toggle_widgets(message_input, button)
        conversation_box.scroll_end(animate=False)

    def toggle_widgets(self, *widgets: Widget) -> None:
        for widget in widgets:
            widget.disabled = not widget.disabled