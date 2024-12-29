defmodule AdventOfCode.Day10 do
  def solve do
    test_solve()

    IO.puts("Solving real data...")
    input = measure_execution_time(&parse_file/0)

    IO.inspect({part_one(input), part_two(input)}, label: "Real data answer")
  end

  defp part_one(map) do
    get_trailheads(map)
    |> Enum.map(fn {x, y} ->
      get_trailhead_max_height_peaks(map, x, y, 1)
      |> List.flatten()
      # Only grab the unique peaks that are possible to visit from this trailhead
      |> Enum.uniq()
      |> length()
    end)
    |> List.flatten()
    |> Enum.sum()
  end

  defp part_two(map) do
    get_trailheads(map)
    |> Enum.map(fn {x, y} ->
      get_trailhead_max_height_peaks(map, x, y, 1)
      |> List.flatten()
      |> length()
    end)
    |> List.flatten()
    |> Enum.sum()
  end

  defp get_trailheads(map) do
    for y <- 0..(tuple_size(map) - 1),
        x <- 0..(tuple_size(elem(map, y)) - 1),
        # Pattern match to get val
        {^x, ^y, val} = get_cell(map, x, y),
        val == 0,
        into: [] do
      {x, y}
    end
  end

  defp get_trailhead_max_height_peaks(map, x, y, match_height) do
    get_adjacent_coordinates_with_height(map, x, y, match_height)
    |> Enum.map(fn {adj_x, adj_y, _val} = cell ->
      case match_height do
        9 -> cell
        _ -> get_trailhead_max_height_peaks(map, adj_x, adj_y, match_height + 1)
      end
    end)
  end

  defp get_adjacent_coordinates_with_height(map, x, y, height) do
    [{x, y + 1}, {x + 1, y}, {x, y - 1}, {x - 1, y}]
    |> Enum.map(fn {x, y} ->
      cell = get_cell(map, x, y)

      if cell do
        {_, _, cell_height} = cell

        cond do
          cell_height == height -> cell
          true -> nil
        end
      end
    end)
    |> Enum.filter(& &1)
  end

  defp get_cell(map, x, y) do
    try do
      val = map |> elem(y) |> elem(x)
      {x, y, val}
    rescue
      ArgumentError -> nil
    end
  end

  defp parse_file(test_input \\ false) do
    file = if test_input, do: "input/day_10_test.txt", else: "input/day_10.txt"

    file
    |> File.read!()
    |> String.split("\n", trim: true)
    |> Enum.map(fn line ->
      line |> String.graphemes() |> Enum.map(&String.to_integer/1) |> List.to_tuple()
    end)
    |> List.to_tuple()
  end

  defp test_solve() do
    IO.puts("Solving test data...")
    input = true |> with_timing(&parse_file/1)
    solved = {part_one(input), part_two(input)}

    case solved do
      {36, 81} -> true
    end

    IO.puts("Test data passed")
  end

  defp with_timing(input, func) do
    measure_execution_time(func, [input])
  end

  require Logger

  def measure_execution_time(function, args \\ []) do
    {time_in_microseconds, result} =
      :timer.tc(function, args)

    time_in_milliseconds = time_in_microseconds / 1_000
    IO.puts("{#{Function.info(function)[:name]}} Execution time: #{time_in_milliseconds} ms")
    result
  end
end

AdventOfCode.Day10.solve()
