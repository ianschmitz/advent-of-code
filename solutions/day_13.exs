defmodule AdventOfCode.Day13 do
  def solve do
    test_solve()

    IO.puts("Solving real data...")
    input = measure_execution_time(&parse_file/0)

    IO.inspect({part_one(input), part_two(input)}, label: "Real data answer")
  end

  defp part_one(input) do
    input
    |> Enum.map(&solve_prize/1)
    |> Enum.reduce(0, fn result, acc ->
      calc_game_tokens(result) + acc
    end)
  end

  defp part_two(input) do
    input
    |> Enum.map(fn game ->
      game
      |> Map.update!(:px, &(&1 + 10_000_000_000_000))
      |> Map.update!(:py, &(&1 + 10_000_000_000_000))
      |> solve_prize()
    end)
    |> Enum.reduce(0, fn result, acc ->
      calc_game_tokens(result) + acc
    end)
  end

  defp calc_game_tokens(nil), do: 0
  defp calc_game_tokens(%{a: a, b: b}), do: a * 3 + b

  defp solve_prize(%{ax: ax, ay: ay, bx: bx, by: by, px: px, py: py}) do
    # Example game:
    # Button A: X+94, Y+34
    # Button B: X+22, Y+67
    # Prize: X=8400, Y=5400
    #
    # Equations we need to solve:
    # (ax * a) + (bx * b) = px
    # (ay * a) + (by * b) = py
    #
    # 94a + 22b = 8400
    # 34a + 67b = 5400
    #
    # Eliminate b:
    # 67(94a + 22b = 8400) -> 6298a + 1474b = 562800 
    # 22(34a + 67b = 5400) -> 748a + 1474b = 118800
    #
    # Subtract both equations to elimate b:
    # 6298a + 1474b = 562800 
    # 748a + 1474b = 118800
    # ----------------------
    # 5550a = 444000
    #
    # a = 80
    #
    # Substitute a back in either orig equation to find b
    # 94(80) + 22b = 8400
    # 7520 + 22b = 8400
    # 22b = 880
    # b = 40
    # 
    # Answer = {a: 80, b:40}

    a = (px * by - py * bx) / (ax * by - bx * ay)
    b = (px - ax * a) / bx

    if is_float_integer?(a) and is_float_integer?(b) do
      %{a: trunc(a), b: trunc(b)}
    else
      nil
    end
  end

  defp is_float_integer?(number) do
    number == Float.floor(number)
  end

  defp parse_file(test_input \\ false) do
    file = if test_input, do: "input/day_13_test.txt", else: "input/day_13.txt"

    file
    |> File.read!()
    |> String.split("\n\n", trim: true)
    |> Enum.map(fn game_text ->
      [
        "Button A: X+" <> <<ax::2-binary>> <> ", Y+" <> <<ay::2-binary>>,
        "Button B: X+" <> <<bx::2-binary>> <> ", Y+" <> <<by::2-binary>>,
        "Prize: " <> prize_text
      ] = String.split(game_text, "\n", trim: true)

      ["X=" <> px, "Y=" <> py] = String.split(prize_text, ", ")

      %{
        ax: String.to_integer(ax),
        ay: String.to_integer(ay),
        bx: String.to_integer(bx),
        by: String.to_integer(by),
        px: String.to_integer(px),
        py: String.to_integer(py)
      }
    end)
  end

  defp test_solve() do
    IO.puts("Solving test data...")
    input = true |> with_timing(&parse_file/1)
    solved = {part_one(input), part_two(input)}

    case solved do
      {480, 875_318_608_908} -> true
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

AdventOfCode.Day13.solve()
