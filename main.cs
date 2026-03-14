using System;

class Program
{
	static void Main()
	{
		Console.WriteLine("Welcome to my quiz game!!!");

		string name = "";
		int age = 0;
		Console.Write("Please enter your name: ");
		name = Console.ReadLine();
		Console.WriteLine($"Hello {name}, let's start the quiz!");
		Console.Write("Please enter your age: ");
		age = int.Parse(Console.ReadLine());
		if (age < 18)
		{
			Console.WriteLine($"Sorry {name}, you are not old enough to play the quiz.");
			return;
		}
		Console.WriteLine("Great! Let's get started.");
		string[] questions = {
			"What is the capital of France?",
			"What is the largest planet in our solar system?",
			"What is the smallest country in the world?",
			"What is the chemical symbol for water?"
		};
		string[] answers = {
			"Paris",
			"Jupiter",
			"City",
			"H2O"
		};
		int score = 0;
		for (int i = 0; i < questions.Length; i++)
		{
			Console.WriteLine($"{i + 1}: {questions[i]}");
			string answer = Console.ReadLine();
			if (string.Equals(answer.Trim(), answers[i], StringComparison.OrdinalIgnoreCase))
			{
				Console.WriteLine("Correct!");
				score++;
			}
			else
			{
				Console.WriteLine($"Wrong! The correct answer is {answers[i]}.");
			}
		}
		Console.WriteLine($"Your final score is {score} out of {questions.Length}.");
	}
}
