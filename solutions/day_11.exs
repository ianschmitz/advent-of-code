defmodule AdventOfCode.Day11 do
  @cache_name :stone_cache

  def solve do
    test_solve()

    IO.puts("Solving real data...")
    input = measure_execution_time(&parse_file/0)

    IO.inspect({part_one(input), part_two(input)}, label: "Real data answer")
  end

  defp part_one(stones) do
    init_cache()
    stones |> with_timing(&count_stones(&1, 25))
  end

  defp part_two(stones) do
    init_cache()
    stones |> with_timing(&count_stones(&1, 75))
  end

  defp count_stones(stones, num_blinks) do
    stones |> Enum.reduce(0, fn stone, acc -> acc + blink(stone, num_blinks) end)
  end

  defp blink(_stone, 0), do: 1

  defp blink(stone, blinks_remaining) do
    memoize({stone, blinks_remaining}, fn ->
      case(do_blink(stone)) do
        {s1, s2} -> blink(s1, blinks_remaining - 1) + blink(s2, blinks_remaining - 1)
        stone -> blink(stone, blinks_remaining - 1)
      end
    end)
  end

  defp do_blink(0), do: 1

  defp do_blink(stone) do
    digits = Integer.digits(stone)
    digits_length = length(digits)
    is_even_num_digits = rem(digits_length, 2) == 0

    if is_even_num_digits do
      {first_half, second_half} =
        digits |> Enum.split(div(digits_length, 2))

      {Integer.undigits(first_half), Integer.undigits(second_half)}
    else
      stone * 2024
    end
  end

  defp init_cache do
    if :ets.whereis(@cache_name) != :undefined do
      :ets.delete(@cache_name)
    end

    :ets.new(@cache_name, [:named_table])
  end

  def memoize(key, fun) do
    case :ets.lookup(@cache_name, key) do
      [{^key, val}] -> val
      [] -> fun.() |> tap(&:ets.insert(@cache_name, {key, &1}))
    end
  end

  defp parse_file(test_input \\ false) do
    file = if test_input, do: "input/day_11_test.txt", else: "input/day_11.txt"

    file
    |> File.read!()
    |> String.trim()
    |> String.split(" ")
    |> Enum.map(&String.to_integer/1)
  end

  defp test_solve() do
    IO.puts("Solving test data...")
    input = true |> with_timing(&parse_file/1)
    solved = {part_one(input), part_two(input)}

    case solved do
      {55312, 65_601_038_650_482} -> true
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

AdventOfCode.Day11.solve()
