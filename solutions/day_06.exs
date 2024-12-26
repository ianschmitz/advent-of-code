defmodule AdventOfCode.Day06 do
  @visited_cell "X"
  @empty_cell "."
  @obstacle_cell "#"
  @starting_cell "^"

  def solve do
    test_solve()

    input = parse_file()

    {part_one(input), part_two(input)}
  end

  defp part_one(map) do
    guard_pos = find_guard_index(map)

    run_simulation(map, guard_pos)
    |> Enum.reduce(0, fn row, acc ->
      # Count up all the visited spaces
      acc + (Enum.filter(row, fn space -> space == @visited_cell end) |> length())
    end)
  end

  defp part_two(map) do
    guard_pos = find_guard_index(map)
    {x, y, _rotation} = guard_pos

    # Solve part_one to get all the possible cells we could place an obstacle
    visited_map = run_simulation(map, guard_pos)

    possible_positions =
      for {row, row_idx} <- Enum.with_index(visited_map),
          {value, col_idx} <- Enum.with_index(row),
          value == @visited_cell && {x, y} != {col_idx, row_idx} do
        {col_idx, row_idx}
      end

    possible_positions
    |> Task.async_stream(fn pos ->
      # Add an obstacle at the position and then run the simulation
      # to see if the guard gets in a loop
      result = map |> update_map(pos, @obstacle_cell) |> run_simulation(guard_pos)

      case result do
        :loop -> true
        _ -> false
      end
    end)
    |> Enum.map(fn {:ok, result} -> result end)
    |> Enum.filter(& &1)
    |> length()
  end

  defp find_guard_index(map) do
    Enum.with_index(map)
    |> Enum.reduce_while(nil, fn {row, row_idx}, _acc ->
      case Enum.find_index(row, fn item -> item == @starting_cell end) do
        # Continue searching
        nil -> {:cont, nil}
        # Found the item
        col_idx -> {:halt, {col_idx, row_idx, :up}}
      end
    end)
  end

  defp run_simulation(map, {x, y, _rotation} = guard_pos, visited \\ MapSet.new()) do
    {next_pos_value, next_pos} = get_next_cell(map, guard_pos)

    cond do
      # Next position would move guard off the map, we're done
      next_pos_value == nil ->
        update_map(map, {x, y}, @visited_cell)

      # Encountered an obstacle, rotate guard and continue
      next_pos_value == @obstacle_cell ->
        run_simulation(map, rotate_90_deg(guard_pos), visited)

      # We've already been here with the same rotation, this means we're in a loop
      MapSet.member?(visited, next_pos) ->
        :loop

      # Safe to keep moving
      next_pos_value in [@empty_cell, @visited_cell] ->
        update_map(map, {x, y}, @visited_cell)
        |> run_simulation(next_pos, MapSet.put(visited, next_pos))
    end
  end

  defp update_map(map, {x, y}, value) do
    List.update_at(map, y, fn row ->
      List.update_at(row, x, fn _old_value -> value end)
    end)
  end

  defp get_next_cell(map, {x, y, rotation}) do
    next_pos =
      case rotation do
        :up -> {x, y - 1, :up}
        :right -> {x + 1, y, :right}
        :down -> {x, y + 1, :down}
        :left -> {x - 1, y, :left}
      end

    {get_cell_value(map, next_pos), next_pos}
  end

  defp get_cell_value(map, {x, y, _rotation}) do
    map |> Enum.at(y, []) |> Enum.at(x)
  end

  defp rotate_90_deg({x, y, rotation}) do
    new_rotation =
      case rotation do
        :up -> :right
        :right -> :down
        :down -> :left
        :left -> :up
      end

    {x, y, new_rotation}
  end

  defp parse_file(test_input \\ false) do
    file = if test_input, do: "input/day_06_test.txt", else: "input/day_06.txt"

    file
    |> File.read!()
    |> String.split("\n", trim: true)
    |> Enum.map(&String.graphemes/1)
  end

  defp test_solve() do
    input = parse_file(true)
    solved = {part_one(input), part_two(input)}

    case solved do
      {41, 6} -> nil
    end
  end
end

IO.inspect(AdventOfCode.Day06.solve())
