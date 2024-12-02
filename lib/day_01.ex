defmodule AdventOfCode.Day01 do
  def solve do
    {left_list, right_list} = parse_file()

    part_one =
      Enum.zip(left_list, right_list)
      |> Enum.reduce(0, fn {left_num, right_num}, acc -> abs(left_num - right_num) + acc end)

    right_list_frequencies = Enum.frequencies(right_list)

    part_two =
      left_list
      |> Enum.reduce(0, fn left_num, acc ->
        left_num * Map.get(right_list_frequencies, left_num, 0) + acc
      end)

    {part_one, part_two}
  end

  defp parse_file do
    {left_list, right_list} =
      "input/day_01.txt"
      |> File.stream!()
      |> Enum.reduce({[], []}, &parse_line/2)

    sorted_left_list = Enum.sort(left_list)
    sorted_right_list = Enum.sort(right_list)

    {sorted_left_list, sorted_right_list}
  end

  defp parse_line(line, {left_list, right_list}) do
    [left, right] = String.split(line)
    left_num = String.to_integer(left)
    right_num = String.to_integer(right)

    {[left_num | left_list], [right_num | right_list]}
  end
end
