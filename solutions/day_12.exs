defmodule AdventOfCode.Day12 do
  def solve do
    test_solve()

    IO.puts("Solving real data...")
    input = measure_execution_time(&parse_file/0)

    IO.inspect({part_one(input), part_two(input)}, label: "Real data answer")
  end

  defp part_one(map) do
    map
    |> with_timing(&find_regions/1)
    |> with_timing(fn regions ->
      regions
      |> Enum.reduce(0, fn region, acc ->
        area = region |> length()
        perimeter = count_region_perimeters(region, map)

        area * perimeter + acc
      end)
    end)
  end

  defp part_two(map) do
    map
    |> with_timing(&find_regions/1)
    |> with_timing(fn regions ->
      regions
      |> Enum.reduce(0, fn region, acc ->
        area = region |> length()
        # sides == count of inside and outside corners
        sides = count_region_corners(region, map)

        area * sides + acc
      end)
    end)
  end

  defp count_region_perimeters(region, map) do
    region
    |> Enum.reduce(0, fn {x, y, crop}, acc ->
      cell_perimeters =
        if(get_cell_crop(map, x, y - 1) != crop, do: 1, else: 0) +
          if(get_cell_crop(map, x, y + 1) != crop, do: 1, else: 0) +
          if(get_cell_crop(map, x + 1, y) != crop, do: 1, else: 0) +
          if get_cell_crop(map, x - 1, y) != crop, do: 1, else: 0

      cell_perimeters + acc
    end)
  end

  defp count_region_corners(region, map) do
    region
    |> Enum.reduce(0, fn cell, acc ->
      cell_corners =
        outside_top_right(map, cell) +
          outside_bottom_right(map, cell) +
          outside_top_left(map, cell) +
          outside_bottom_left(map, cell) +
          inside_top_right(map, cell) +
          inside_bottom_right(map, cell) +
          inside_top_left(map, cell) +
          inside_bottom_left(map, cell)

      cell_corners + acc
    end)
  end

  defp outside_top_right(map, {x, y, crop}) do
    if get_cell_crop(map, x, y - 1) != crop and get_cell_crop(map, x + 1, y) != crop,
      do: 1,
      else: 0
  end

  defp outside_bottom_right(map, {x, y, crop}) do
    if get_cell_crop(map, x, y + 1) != crop and get_cell_crop(map, x + 1, y) != crop,
      do: 1,
      else: 0
  end

  defp outside_top_left(map, {x, y, crop}) do
    if get_cell_crop(map, x, y - 1) != crop and get_cell_crop(map, x - 1, y) != crop,
      do: 1,
      else: 0
  end

  defp outside_bottom_left(map, {x, y, crop}) do
    if get_cell_crop(map, x, y + 1) != crop and get_cell_crop(map, x - 1, y) != crop,
      do: 1,
      else: 0
  end

  defp inside_top_right(map, {x, y, crop}) do
    if get_cell_crop(map, x, y - 1) == crop and get_cell_crop(map, x + 1, y) == crop and
         get_cell_crop(map, x + 1, y - 1) != crop,
       do: 1,
       else: 0
  end

  defp inside_bottom_right(map, {x, y, crop}) do
    if get_cell_crop(map, x, y + 1) == crop and get_cell_crop(map, x + 1, y) == crop and
         get_cell_crop(map, x + 1, y + 1) != crop,
       do: 1,
       else: 0
  end

  defp inside_top_left(map, {x, y, crop}) do
    if get_cell_crop(map, x, y - 1) == crop and get_cell_crop(map, x - 1, y) == crop and
         get_cell_crop(map, x - 1, y - 1) != crop,
       do: 1,
       else: 0
  end

  defp inside_bottom_left(map, {x, y, crop}) do
    if get_cell_crop(map, x, y + 1) == crop and get_cell_crop(map, x - 1, y) == crop and
         get_cell_crop(map, x - 1, y + 1) != crop,
       do: 1,
       else: 0
  end

  defp find_regions(map) do
    coords =
      for y <- 0..(tuple_size(map) - 1),
          x <- 0..(tuple_size(elem(map, y)) - 1),
          into: [] do
        get_cell(map, x, y)
      end

    coords
    |> Enum.reduce({[], MapSet.new()}, fn coord, {regions, visited} = acc ->
      already_visited = visited |> MapSet.member?(coord)

      if already_visited do
        acc
      else
        region =
          find_region_coords(map, coord)

        new_visited = MapSet.union(visited, region)

        region_list =
          region
          |> MapSet.to_list()
          |> Enum.sort_by(fn {x, y, val} -> {y, x} end)

        {[region_list | regions], new_visited}
      end
    end)
    |> elem(0)
  end

  defp find_region_coords(map, coord, region \\ MapSet.new()) do
    {cells, region} = get_adjacent_coordinates_with_plot(map, coord, region)

    cells
    |> Enum.reduce(region, fn adj_coord, region ->
      find_region_coords(map, adj_coord, region)
    end)
  end

  defp get_adjacent_coordinates_with_plot(map, {x, y, val}, region) do
    adj_coords =
      [{x, y}, {x, y + 1}, {x + 1, y}, {x, y - 1}, {x - 1, y}]
      |> Enum.reduce({[], region}, fn {adj_x, adj_y}, {list, region} = acc ->
        cell = get_cell(map, adj_x, adj_y)

        if cell && !MapSet.member?(region, cell) do
          {_, _, adj_val} = cell

          cond do
            adj_val == val ->
              {[cell | list], MapSet.put(region, cell)}

            true ->
              acc
          end
        else
          acc
        end
      end)
  end

  defp get_cell(map, x, y) do
    try do
      val = map |> elem(y) |> elem(x)
      {x, y, val}
    rescue
      ArgumentError -> nil
    end
  end

  defp get_cell_crop(map, x, y) do
    cell = get_cell(map, x, y)

    case cell do
      nil -> nil
      _ -> elem(cell, 2)
    end
  end

  defp parse_file(test_input \\ false) do
    file = if test_input, do: "input/day_12_test.txt", else: "input/day_12.txt"

    file
    |> File.read!()
    |> String.split("\n", trim: true)
    |> Enum.map(&(String.graphemes(&1) |> List.to_tuple()))
    |> List.to_tuple()
  end

  defp test_solve() do
    IO.puts("Solving test data...")
    input = true |> with_timing(&parse_file/1)
    solved = {part_one(input), part_two(input)}

    case solved do
      {1930, 1206} -> true
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

AdventOfCode.Day12.solve()
