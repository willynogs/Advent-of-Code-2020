def parsed_input
  o = {}
  File.read("input.txt").split("\n").map do |row|
    color, contents = row.split(" bags contain ")
    o[color] = contents.split(", ")
  end
  o
end

def color_possiblities(input, inspect_color = "shiny gold")
  found = input.select { |k, v| v.any? { |content| content.include?(inspect_color) } }.keys
  rec = found.map { |color| color_possiblities(input, color) }
  return [found + rec].flatten.uniq
end

def required_count(input, bag_count = 1, inspect_color = "shiny gold")
  total_count = 0
  contents = input[inspect_color]
  if !contents.first.include?("no other bags")
    contents.map do |content|
      count, color_1, color_2 = content.split(" ")
      total_count += required_count(input, count.to_i, "#{color_1} #{color_2}")
    end
  end
  (total_count * bag_count) + bag_count
end

p color_possiblities(parsed_input).count
p required_count(parsed_input) - 1
