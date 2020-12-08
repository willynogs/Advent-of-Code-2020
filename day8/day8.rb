def parse_instruction(instruction)
  op, val = instruction.split(" ")
  if val.include? "-"
    [op, (val.delete("-").to_i * -1)]
  elsif val.include? "+"
    [op, val.delete("+").to_i]
  end
end

def part_1
  input = File.read("input.txt").split("\n")
  acc = 0
  runs = []
  i = 0
  while i < input.length
    if runs.include? i
      p acc
      break
    end
    runs << i
    op, val = parse_instruction(input[i])
    if op == "nop"
      i += 1
    elsif op == "acc"
      acc += val
      i += 1
    elsif op == "jmp"
      i += val
    end
  end
end

part_1

def part_2
  input = File.read("input.txt").split("\n")
  switch_i = 0
  acc = 0
  runs = []
  i = 0
  while i < input.length
    op, val = parse_instruction(input[i])
    if runs.include?(i)
      runs = []
      i = 0
      acc = 0
      switch_i += 1
    else
      runs << i
      if op == "nop"
        i == switch_i ? i += val : i += 1
      elsif op == "acc"
        acc += val
        i += 1
      elsif op == "jmp"
        i == switch_i ? i += 1 : i += val
      end
    end
  end
  p acc
end

part_2
