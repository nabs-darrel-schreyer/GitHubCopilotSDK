import { CopilotClient, approveAll } from "@github/copilot-sdk";

console.log("Starting a GitHub Copilot session...");

const client = new CopilotClient();

try {
  const session = await client.createSession({
    model: "auto",
    onPermissionRequest: approveAll,
  });

  console.log("Sending message to GitHub Copilot...");
  const response = await session.sendAndWait({
    prompt: "Calculate the product of 2 and 4.",
  });
  console.log("Received response from GitHub Copilot...");
  process.stdout.write("Response content: ");
  process.stdout.write("\x1b[32m");
  process.stdout.write(response?.data.content ?? "");
  process.stdout.write("\x1b[0m");
  console.log();
  console.log("GitHub Copilot session ended.");
} finally {
  await client.stop();
}