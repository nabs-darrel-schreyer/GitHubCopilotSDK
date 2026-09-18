"""Python twin of DotNetGhCopilot.Cli / TsGhCopilot.Cli — GitHub Copilot SDK sample."""

from __future__ import annotations

import asyncio
import logging
import sys
from typing import Any, cast

from copilot import CopilotClient, CopilotSession
from copilot.generated.session_events import AssistantMessageData
from copilot.session import PermissionHandler

GREEN = "\x1b[32m"
RESET = "\x1b[0m"


def _mark_disconnected(session: CopilotSession) -> None:
    cast(Any, session)._mark_disconnected()


async def main() -> None:
    print("Starting a GitHub Copilot session...")

    client: CopilotClient = CopilotClient()
    await client.start()

    session: CopilotSession | None = None
    try:
        session = await client.create_session(
            model="auto",
            on_permission_request=PermissionHandler.approve_all,
        )

        print("Sending message to GitHub Copilot...")
        response = await session.send_and_wait(
            "Calculate the product of 2 and 4."
        )
        print("Received response from GitHub Copilot...")
        content = ""
        if response is not None and isinstance(response.data, AssistantMessageData):
            content = response.data.content
        sys.stdout.write("Response content: ")
        sys.stdout.write(f"{GREEN}{content}{RESET}")
        sys.stdout.write("\n")
        print("GitHub Copilot session ended.")
    finally:
        # github-copilot-sdk 1.0.14 disconnects via JSON-RPC session.detach.
        # Some Copilot runtimes return -32601 (method not found). Mark the
        # session closed locally first so client.stop() skips that RPC.
        if session is not None:
            _mark_disconnected(session)
        logging.getLogger("copilot").setLevel(logging.CRITICAL)
        try:
            await client.stop()
        except* Exception:
            pass


if __name__ == "__main__":
    asyncio.run(main())