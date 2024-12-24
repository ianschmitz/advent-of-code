defmodule AdventOfCode.Day05 do
  require IEx

  def solve do
    test_solve()

    input = parse_file()
    IO.inspect(input)

    {part_one(input), part_two(input)}
  end

  defp get_middle_number(update) do
    middle_index = update |> length() |> div(2)
    Enum.at(update, middle_index)
  end

  defp part_one({_rules, valid_updates, _invalid_updates}) do
    valid_updates
    |> Enum.map(&get_middle_number/1)
    |> Enum.sum()
  end

  defp part_two({rules, _valid_updates, invalid_updates}) do
    rules_graph =
      rules
      |> Enum.reduce(%{}, fn {before, follows}, acc ->
        # For each number, maintain a list of numbers that must come after it
        Map.update(acc, before, [follows], fn existing -> [follows | existing] end)
      end)

    invalid_updates
    |> Enum.map(&sort_numbers(&1, rules_graph))
    |> Enum.map(&get_middle_number/1)
    |> Enum.sum()
  end

  defp sort_numbers(numbers, rules) do
    numbers
    |> Enum.sort(fn n1, n2 ->
      rule = Map.get(rules, n1, [])
      Enum.member?(rule, n2)
    end)
  end

  defp parse_file(test \\ false) do
    file = if test, do: "input/day_05_test.txt", else: "input/day_05.txt"

    [rules_raw, updates_raw] =
      file
      |> File.read!()
      |> String.split("\n\n")
      |> Enum.map(&String.split(&1, "\n", trim: true))

    rules =
      rules_raw
      |> Enum.map(fn rule ->
        String.split(rule, "|", trim: true) |> Enum.map(&String.to_integer/1) |> List.to_tuple()
      end)

    # Parse all rules into a map that looks like:
    # Rules: [[10,20], [20, 42]]
    # Map: { 10: [[10, 20]], 20: [[10, 20], [20, 42]] }
    a_rules_map = Enum.group_by(rules, fn {a, _b} -> a end)
    b_rules_map = Enum.group_by(rules, fn {_a, b} -> b end)
    rules_map = Map.merge(a_rules_map, b_rules_map, fn _key, val1, val2 -> val1 ++ val2 end)

    updates =
      updates_raw
      |> Enum.map(fn update -> String.split(update, ",") |> Enum.map(&String.to_integer/1) end)

    {valid_updates, invalid_updates} =
      updates
      |> Enum.split_with(&update_is_valid(rules_map, &1))

    {rules, valid_updates, invalid_updates}
  end

  defp update_is_valid(rules_map, update) do
    update_with_index = update |> Enum.with_index()

    # Store the update in a map with number as key, and index as value
    # for more efficient lookup when checking rules: O(log n) vs O(n)
    # Is this necessary? Unlikely, but it's fun.
    update_map =
      update_with_index |> Enum.into(%{}, fn {value, index} -> {value, index} end)

    update_with_index
    |> Enum.all?(fn {element, index} ->
      rules_for_element = Map.get(rules_map, element)

      case rules_for_element do
        nil ->
          true

        _ ->
          Enum.all?(rules_for_element, fn rule ->
            element_passes_rule(rule, update_map, element, index)
          end)
      end
    end)
  end

  defp element_passes_rule(rule, update_map, element, index) do
    {a, b} = rule

    cond do
      a == element && Map.has_key?(update_map, b) -> index < Map.get(update_map, b)
      b == element && Map.has_key?(update_map, a) -> index > Map.get(update_map, a)
      true -> true
    end
  end

  defp test_solve() do
    input = parse_file(true)
    solved = {part_one(input), part_two(input)}

    case solved do
      {143, 123} -> nil
    end
  end
end

IO.inspect(AdventOfCode.Day05.solve())
