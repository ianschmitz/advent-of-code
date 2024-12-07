defmodule AdventOfCode.Day02 do
  def solve do
    reports = parse_file()

    {part_one(reports), part_two(reports)}
  end

  defp part_one(reports), do: reports |> Enum.count(&report_safe?/1)

  defp part_two(reports), do: reports |> Enum.count(&report_safe_with_dampener?/1)

  defp report_safe?([a, b | _] = list) when a < b, do: report_safe?(:asc, list)
  defp report_safe?([a, b | _] = list) when a > b, do: report_safe?(:desc, list)
  defp report_safe?([a, a | _]), do: false

  defp report_safe?(:asc, [a, b | tail]) when a < b and (b - a) in 1..3,
    do: report_safe?(:asc, [b | tail])

  defp report_safe?(:desc, [a, b | tail]) when a > b and (a - b) in 1..3,
    do: report_safe?(:desc, [b | tail])

  defp report_safe?(_, [_last]), do: true
  defp report_safe?(_, _), do: false

  defp report_safe_with_dampener?(report) do
    # Create every possible combination of a removed element in the sequence
    dampened =
      0..(length(report) - 1)
      |> Enum.map(&List.delete_at(report, &1))

    report_safe?(report) || Enum.any?(dampened, &report_safe?/1)
  end

  defp parse_file do
    {flags, _} = OptionParser.parse!(System.argv(), strict: [test: :boolean])
    file = if flags[:test], do: "input/day_02_test.txt", else: "input/day_02.txt"

    file
    |> File.stream!()
    |> Enum.map(fn line -> String.split(line) |> Enum.map(&String.to_integer/1) end)
  end
end

IO.inspect(AdventOfCode.Day02.solve())
