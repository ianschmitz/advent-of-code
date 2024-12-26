defmodule AdventOfCode.Day09 do
  def solve do
    test_solve()

    IO.puts("Solving real data...")
    input = parse_file()

    IO.inspect({part_one(input), part_two(input)}, label: "Real data answer")
  end

  defp part_one(input) do
    input
    |> with_timing(&expand_disk_map/1)
    |> with_timing(&move_file_blocks_v1/1)
    |> with_timing(&calc_checksum/1)
  end

  defp expand_disk_map(disk_map) do
    disk_map
    |> Enum.with_index()
    |> Enum.reduce({[], 0}, fn {number, idx}, {blocks, file_id} ->
      value_to_add =
        case rem(idx, 2) do
          # File
          0 -> file_id
          # Free space
          1 -> nil
        end

      values = List.duplicate(value_to_add, number)

      next_file_id = if value_to_add == nil, do: file_id, else: file_id + 1

      {[values | blocks], next_file_id}
    end)
    |> elem(0)
    |> Enum.reverse()
    |> List.flatten()
  end

  defp move_file_blocks_v1(disk) do
    head = disk |> Enum.with_index()
    tail = head |> Enum.reverse()

    do_move_file_blocks_v1(head, tail, [])
  end

  defp do_move_file_blocks_v1(
         [{head_val, head_idx} | rest_head] = head_list,
         [{tail_val, tail_idx} | rest_tail] = tail_list,
         acc
       ) do
    cond do
      # All done!
      tail_idx < head_idx -> acc |> Enum.reverse()
      # Head is data, add it to acc and move head along
      head_val != nil -> do_move_file_blocks_v1(rest_head, tail_list, [head_val | acc])
      # Head is free space and tail is data, add tail to acc and move both along
      tail_val != nil -> do_move_file_blocks_v1(rest_head, rest_tail, [tail_val | acc])
      # Both are free space, move tail along
      true -> do_move_file_blocks_v1(head_list, rest_tail, acc)
    end
  end

  defp calc_checksum(disk) do
    disk
    |> Enum.with_index()
    |> Enum.map(fn {file_id, idx} -> file_id * idx end)
    |> Enum.sum()
  end

  defp part_two(input) do
    input
    |> with_timing(&expand_disk_map/1)
    |> with_timing(&move_file_blocks_v2/1)
    |> with_timing(&sort_blocks_by_start_idx/1)
    |> with_timing(&expand_blocks/1)
    |> with_timing(&calc_checksum/1)
  end

  defp move_file_blocks_v2(disk) do
    head = disk |> Enum.with_index()
    tail = head |> Enum.reverse()

    free_blocks = find_free_blocks(head)
    data_blocks = find_data_blocks(tail)

    do_move_file_blocks_v2(data_blocks, free_blocks, [])
  end

  defp do_move_file_blocks_v2(
         data_blocks,
         free_blocks,
         acc
       )
       when free_blocks == [] or data_blocks == [],
       do: Enum.concat([free_blocks, data_blocks, acc])

  defp do_move_file_blocks_v2(
         [{data_id, data_start_idx, data_end_idx} = data_block | rest_data_blocks],
         free_blocks,
         acc
       ) do
    matching_free_block_idx =
      free_blocks
      |> Enum.find_index(fn {nil, start_idx, end_idx} ->
        end_idx - start_idx >= data_end_idx - data_start_idx and data_start_idx > start_idx
      end)

    matching_free_block =
      if matching_free_block_idx, do: Enum.at(free_blocks, matching_free_block_idx)

    cond do
      # We've found a free block that can accomodate our data block
      matching_free_block ->
        data_length = data_end_idx - data_start_idx + 1

        # Update the free block by moving the starting point to account for data that is being placed here
        new_free_block =
          {nil, elem(matching_free_block, 1) + data_length, elem(matching_free_block, 2)}

        # Add a new free block in place of the data we're moving

        # Is there still room left in the free block after moving data?
        new_free_blocks =
          if elem(new_free_block, 1) <= elem(new_free_block, 2),
            # Update free block to remove space now replaced by data
            do: List.replace_at(free_blocks, matching_free_block_idx, new_free_block),
            # Remove free block
            else: List.delete_at(free_blocks, matching_free_block_idx)

        new_free_blocks =
          [{nil, data_start_idx, data_end_idx} | new_free_blocks] |> sort_blocks_by_start_idx()

        # Move data to start of free block
        {nil, new_start_idx, _} = matching_free_block
        new_acc = [{data_id, new_start_idx, new_start_idx + data_length - 1} | acc]

        do_move_file_blocks_v2(rest_data_blocks, new_free_blocks, new_acc)

      # We can't move this data as there's no free block big enough left
      true ->
        new_acc = [data_block | acc]
        do_move_file_blocks_v2(rest_data_blocks, free_blocks, new_acc)
    end
  end

  defp find_free_blocks(list_with_idx) do
    list_with_idx
    |> Enum.chunk_by(fn {val, _idx} -> is_nil(val) end)
    |> Enum.filter(fn [{val, _} | _] -> is_nil(val) end)
    |> Enum.map(fn chunk ->
      {_, first_idx} = List.first(chunk)
      {_, last_idx} = List.last(chunk)
      {nil, first_idx, last_idx}
    end)
  end

  defp find_data_blocks(list_with_idx) do
    list_with_idx
    |> Enum.chunk_by(fn {val, _idx} -> val end)
    |> Enum.filter(fn [{val, _} | _] -> !is_nil(val) end)
    |> Enum.map(fn chunk ->
      # Since we're processing the list reversed,
      # List.last is the first occurence of the data when reading from the left
      {val, first_idx} = List.last(chunk)
      {_, last_idx} = List.first(chunk)
      {val, first_idx, last_idx}
    end)
  end

  defp sort_blocks_by_start_idx(list), do: list |> List.keysort(1)

  defp expand_blocks(blocks), do: do_expand_blocks(blocks, [])

  defp do_expand_blocks([], acc), do: acc |> Enum.reverse() |> List.flatten()

  defp do_expand_blocks([{id, start_idx, end_idx} | rest_blocks], acc) do
    parsed_id = if id == nil, do: 0, else: id
    new_acc = [List.duplicate(parsed_id, end_idx - start_idx + 1) | acc]
    do_expand_blocks(rest_blocks, new_acc)
  end

  defp parse_file(test_input \\ false) do
    file = if test_input, do: "input/day_09_test.txt", else: "input/day_09.txt"

    file
    |> File.read!()
    |> String.trim()
    |> String.graphemes()
    |> Enum.map(&String.to_integer/1)
  end

  defp test_solve() do
    IO.puts("Solving test data...")
    input = parse_file(true)
    solved = {part_one(input), part_two(input)}

    case solved do
      {1928, 2858} -> true
    end

    IO.puts("Test data passed")
  end

  defp with_timing(input, func) do
    measure_execution_time(func, [input])
  end

  require Logger

  def measure_execution_time(function, args) do
    {time_in_microseconds, result} =
      :timer.tc(function, args)

    time_in_milliseconds = time_in_microseconds / 1_000
    IO.puts("{#{Function.info(function)[:name]}} Execution time: #{time_in_milliseconds} ms")
    result
  end
end

AdventOfCode.Day09.solve()
