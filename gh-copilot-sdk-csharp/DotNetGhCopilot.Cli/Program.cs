using GitHub.Copilot;

Console.WriteLine("Starting a GitHub Copilot session...");

var sessionConfig = new SessionConfig
{
    Model = "auto",
    OnPermissionRequest = PermissionHandler.ApproveAll
};

var messageOptions = new MessageOptions 
{ 
    Prompt = "Calculate the product of 2 and 4."
};

await using var client = new CopilotClient();
await using var session = await client.CreateSessionAsync(sessionConfig);
Console.WriteLine("Sending message to GitHub Copilot...");
var response = await session.SendAndWaitAsync(messageOptions);
Console.WriteLine("Received response from GitHub Copilot...");
Console.Write("Response content: ");
var originalForegroundColor = Console.ForegroundColor;
Console.ForegroundColor = ConsoleColor.Green;
Console.Write(response?.Data.Content);
Console.ForegroundColor = originalForegroundColor;
Console.WriteLine();
Console.WriteLine("GitHub Copilot session ended.");