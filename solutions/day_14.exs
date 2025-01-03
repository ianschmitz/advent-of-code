defmodule AdventOfCode.Day14 do
  defp part_one(%{robots: robots, width: width, height: height}) do
    robots
    |> Enum.map(&move_robot(&1, 100, width, height))
    |> Enum.group_by(&get_position_quadrant(&1, width, height))
    # Throw away robots that were sitting on a quadrant boundary
    |> Map.delete(0)
    |> Map.values()
    |> Enum.reduce(1, fn robots, acc ->
      acc * length(robots)
    end)
  end

  defp part_two(%{robots: robots, width: width, height: height}) do
    # Arbitrarily large number to search for the solution
    {seconds, robots} =
      1..20_000
      |> Enum.map(fn seconds ->
        {seconds,
         robots
         |> Enum.map(&move_robot(&1, seconds, width, height))}
      end)
      |> Enum.min_by(fn {seconds, robots} ->
        calc_std_dev(robots)
      end)

    IO.puts("Answer: #{seconds}")
    print_robots(robots, width, height)
  end

  # I had to look up the formula to calc std dev of a set of coordinates
  defp calc_std_dev(robots) do
    {mean_x, mean_y} = calc_mean(robots)

    variance =
      robots
      |> Enum.map(fn {x, y} ->
        dx = x - mean_x
        dy = y - mean_y
        dx * dx + dy * dy
      end)
      |> Enum.sum()
      |> Kernel./(length(robots))

    :math.sqrt(variance)
  end

  defp calc_mean(robots) do
    {sum_x, sum_y} =
      robots
      |> Enum.reduce({0, 0}, fn {x, y}, {acc_x, acc_y} ->
        {acc_x + x, acc_y + y}
      end)

    n = length(robots)
    {sum_x / n, sum_y / n}
  end

  defp print_robots(robots, width, height) do
    empty_grid = for _ <- 1..height, do: for(_ <- 1..width, do: " ")

    Enum.reduce(robots, empty_grid, fn {x, y}, acc ->
      new_acc =
        List.update_at(acc, y, fn row ->
          List.update_at(row, x, fn _ -> "*" end)
        end)

      new_acc
    end)
    |> Enum.each(fn row -> row |> Enum.join("") |> IO.puts() end)
  end

  defp get_position_quadrant({x, y}, width, height) do
    middle_height = (height - 1) / 2
    middle_width = (width - 1) / 2

    cond do
      x < middle_width and y < middle_height -> 1
      x > middle_width and y < middle_height -> 2
      x < middle_width and y > middle_height -> 3
      x > middle_width and y > middle_height -> 4
      true -> 0
    end
  end

  defp wrap_coordinate(coord, limit) do
    Integer.mod(coord, limit)
  end

  defp move_robot(%{px: px, py: py, vx: vx, vy: vy} = robot, seconds, width, height) do
    spaces_moved_x = vx * seconds
    spaces_moved_y = vy * seconds

    new_px = wrap_coordinate(px + spaces_moved_x, width)
    new_py = wrap_coordinate(py + spaces_moved_y, height)

    {new_px, new_py}
  end

  def solve do
    test_solve()

    IO.puts("Solving real data...")
    input = measure_execution_time(&parse_file/0)

    IO.inspect(
      {part_one(%{robots: input, width: 101, height: 103}),
       part_two(%{robots: input, width: 101, height: 103})},
      label: "Real data answer"
    )
  end

  defp test_solve() do
    IO.puts("Solving test data...")
    input = true |> with_timing(&parse_file/1)

    solved =
      {
        part_one(%{robots: input, width: 11, height: 7}),
        nil
        # part_two(%{robots: input, width: 11, height: 7})
      }

    case solved do
      {12, nil} -> true
    end

    IO.puts("Test data passed")
  end

  defp parse_file(test_input \\ false) do
    file = if test_input, do: "input/day_14_test.txt", else: "input/day_14.txt"

    file
    |> File.read!()
    |> String.split("\n", trim: true)
    |> Enum.map(fn line ->
      [pos_text, velocity_text] = String.split(line, " ")
      ["p=" <> px, py] = String.split(pos_text, ",")
      ["v=" <> vx, vy] = String.split(velocity_text, ",")

      %{
        px: String.to_integer(px),
        py: String.to_integer(py),
        vx: String.to_integer(vx),
        vy: String.to_integer(vy)
      }
    end)
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

AdventOfCode.Day14.solve()
