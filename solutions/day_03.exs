defmodule AdventOfCode.Day03 do
  def solve do
    program = parse_file()

    {part_one(program), part_two(program)}
  end

  defp part_one(program) do
    Regex.scan(~r/mul\((\d{1,3}),(\d{1,3})\)/, program)
    |> Enum.map(fn [_, a, b] -> String.to_integer(a) * String.to_integer(b) end)
    |> Enum.sum()
  end

  defp part_two(program) do
    Regex.scan(~r/mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)/, program)
    |> Enum.reduce({0, true}, &parse_match/2)
    |> elem(0)
  end

  def parse_match(["do()"], {count, _}), do: {count, true}
  def parse_match(["don't()"], {count, _}), do: {count, false}

  def parse_match([_match, a, b], {count, true}),
    do: {String.to_integer(a) * String.to_integer(b) + count, true}

  def parse_match([_match, _a, _b], {_count, false} = acc), do: acc

  defp parse_file do
    {flags, _} = OptionParser.parse!(System.argv(), strict: [test: :boolean])
    file = if flags[:test], do: "input/day_03_test.txt", else: "input/day_03.txt"

    file
    |> File.read!()
  end
end

IO.inspect(AdventOfCode.Day03.solve())
