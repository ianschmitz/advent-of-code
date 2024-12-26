defmodule AdventOfCode.Day07 do
  def solve do
    test_solve()

    input = parse_file()

    {part_one(input), part_two(input)}
  end

  defp part_one({input, longest_numbers_length}) do
    # There's a much simpler way that doesn't involve precomputing combinations
    # See https://github.com/bjorng/advent-of-code/blob/main/2024/day07/lib/day07.ex
    combinations =
      Map.new(
        for num <- 1..(longest_numbers_length - 1) do
          {num, generate_combinations(num, ["+", "*"])}
        end
      )

    valid_examples =
      input
      |> Task.async_stream(fn {test_value, numbers} ->
        found =
          combinations
          |> Map.get(length(numbers) - 1)
          |> Enum.find(fn combination ->
            evaluate(numbers, combination) == test_value
          end)

        {test_value, numbers, found}
      end)
      |> Stream.filter(fn {:ok, {_, _, found}} -> not is_nil(found) end)
      |> Stream.map(fn {:ok, {test_value, numbers, _}} -> {test_value, numbers} end)

    valid_examples |> Enum.reduce(0, fn {value, _numbers}, acc -> acc + value end)
  end

  defp part_two({input, longest_numbers_length}) do
    combinations =
      Map.new(
        for num <- 1..(longest_numbers_length - 1) do
          {num, generate_combinations(num, ["+", "*", "||"])}
        end
      )

    valid_examples =
      input
      |> Enum.filter(fn {test_value, numbers} ->
        combinations
        |> Map.get(length(numbers) - 1)
        |> Enum.find(fn combination -> evaluate(numbers, combination) == test_value end)
      end)

    valid_examples |> Enum.reduce(0, fn {value, _numbers}, acc -> acc + value end)
  end

  # Generate all possible combinations of "+" and "*" for a given length
  defp generate_combinations(0, _operators), do: [[]]

  defp generate_combinations(n, operators) do
    for op <- operators, rest <- generate_combinations(n - 1, operators), do: [op | rest]
  end

  # Evaluate the expression left-to-right using the given numbers and operators
  defp evaluate([num | nums], operators) do
    Enum.zip(nums, operators)
    |> Enum.reduce(num, fn {next_num, op}, acc ->
      case op do
        "+" -> acc + next_num
        "*" -> acc * next_num
        "||" -> concat(acc, next_num)
      end
    end)
  end

  # A much more efficient way to concat ints.
  # Went from 9s using string interpolation and integer conversion to 2s
  defp concat(a, b) do
    shift(a, b) + b
  end

  defp shift(a, 0), do: a
  defp shift(a, b), do: shift(a * 10, div(b, 10))

  defp parse_file(test_input \\ false) do
    file = if test_input, do: "input/day_07_test.txt", else: "input/day_07.txt"

    input =
      file
      |> File.read!()
      |> String.split("\n", trim: true)
      |> Enum.map(fn line ->
        [test_value, raw_numbers] = String.split(line, ":")

        numbers = raw_numbers |> String.split(" ", trim: true) |> Enum.map(&String.to_integer/1)

        {String.to_integer(test_value), numbers}
      end)

    # Precompute the possible combinations for each size of list
    longest_numbers_length =
      input
      |> Enum.map(fn {_test_value, numbers} -> length(numbers) end)
      |> Enum.max()

    {input, longest_numbers_length}
  end

  defp test_solve() do
    input = parse_file(true)
    solved = {part_one(input), part_two(input)}

    case solved do
      {3749, 11387} -> true
    end
  end
end

IO.inspect(AdventOfCode.Day07.solve())
