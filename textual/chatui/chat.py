import asyncio

import openai

import os


OPENAI_API_KEY = os.getenv("OPENAI_API_KEY")


class Conversation:

    model: str = "gpt-3.5-turbo"

    def __init__(self) -> None:
        self.messages: list[dict] = []

    async def send(self, message: str) -> list[str]:
        self.messages.append({"role": "user", "content": message})
        r = await openai.ChatCompletion.acreate(
            model=self.model,
            messages = self.messages
        )

        return [choice["message"]["content"] for choice in r["choices"]]

    def pick_response(self, choice: str) -> None:
        self.messages.append({"role": "assistant", "content": choice})

    def clear(self) -> None:
        self.messages = []


async def main() -> None:
    conversation = Conversation()

    while True:
        message = input("Type your message")
        choices = await conversation.send(message)
        print("Here are your choices:", choices)
        choice_index = input("Pick your choice:")
        conversation.pick_response(choices[int(choice_index)])


if __name__ == "__main__":
    openai.api_key = OPENAI_API_KEY
    asyncio.run(main())