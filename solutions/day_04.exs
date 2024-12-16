defmodule AdventOfCode.Day04 do
  require IEx

  def solve do
    grid = parse_file()

    {part_one(grid), part_two(grid)}
  end

  defp part_one(grid) do
    grid
    |> find_coords("X")
    |> Enum.flat_map(&find_xmas_words(grid, &1))
    |> length()
  end

  defp part_two(grid) do
    grid
    |> find_coords("A")
    |> Enum.flat_map(&find_x_mas_words(grid, &1))
    |> length()
  end

  defp find_coords(grid, search_char) do
    grid
    |> Enum.with_index(fn row, row_idx ->
      Enum.with_index(row, fn char, col_idx ->
        if char == search_char, do: {row_idx, col_idx}, else: false
      end)
    end)
    |> List.flatten()
    |> Enum.filter(& &1)
  end

  defp find_xmas_words(grid, start) do
    # 8 coordinates relative to current char
    [{0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}]
    |> Enum.filter(fn direction ->
      matches?(start, direction, 1, "M", grid) &&
        matches?(start, direction, 2, "A", grid) &&
        matches?(start, direction, 3, "S", grid)
    end)
  end

  defp find_x_mas_words(grid, start) do
    [[[{1, -1}, {1, 1}], [{-1, 1}, {-1, -1}]], [[{-1, -1}, {1, -1}], [{-1, 1}, {1, 1}]]]
    |> Enum.filter(fn [side1, side2] ->
      (Enum.all?(side1, &matches?(start, &1, 1, "M", grid)) &&
         Enum.all?(side2, &matches?(start, &1, 1, "S", grid))) ||
        (Enum.all?(side1, &matches?(start, &1, 1, "S", grid)) &&
           Enum.all?(side2, &matches?(start, &1, 1, "M", grid)))
    end)
  end

  def matches?({row1, col1}, {row_direction, col_direction}, offset, match, grid) do
    row_idx = offset * row_direction + row1
    col_idx = offset * col_direction + col1

    char = get_char_at(grid, row_idx, col_idx)

    if row1 == 9 && col1 == 1 && row_direction == -1 && col_direction == -1 && match == "A",
      do: IEx.pry()

    char == match
  end

  def get_char_at(grid, row_idx, col_idx) do
    row = if row_idx >= 0 && row_idx < length(grid), do: Enum.at(grid, row_idx)
    if row && col_idx >= 0 && col_idx < length(row), do: Enum.at(row, col_idx)
  end

  defp parse_file do
    {flags, _} = OptionParser.parse!(System.argv(), strict: [test: :boolean])
    file = if flags[:test], do: "input/day_04_test.txt", else: "input/day_04.txt"

    file
    |> File.stream!()
    |> Enum.map(&(&1 |> String.trim() |> String.graphemes()))
  end
end

IO.inspect(AdventOfCode.Day04.solve())
