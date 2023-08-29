from textual.app import App, ComposeResult
from textual.containers import ScrollableContainer
from textual.widgets import Header, Footer, Button, Static


class TimeDisplay(Static):
    """A widget to display the elapsed time."""


class Stopwatch(Static):
    """A stopwatch widget."""
    def compose(self) -> ComposeResult:
        """Create child widgets of a stopwatch."""
        yield Button("Start", id="start", variant="success")
        yield Button("Stop", id="stop", variant="error")
        yield Button("Reset", id="reset")
        yield TimeDisplay("00:00:00:00")


class StopwatchApp(App):
    """A Textual app to manage stopwatches."""

    BINDINGS = [("d", "toggle_dark", "Toggle Dark Mode")]

    def compose(self) -> ComposeResult:
        """Create child widgets for the App."""
        yield Header()
        yield Footer()
        yield ScrollableContainer(Stopwatch(), Stopwatch(), Stopwatch())

    def action_toggle_dark(self) -> None:
        """Toggle dark mode."""
        self.dark = not self.dark


if __name__ == "__main__":
    app = StopwatchApp()
    app.run()
