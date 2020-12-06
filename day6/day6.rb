# this is me giving up on go :(
f = File.open("input.txt")
input = f.read
groups = input.split("\n\n")
sums = groups.map do |group|
  group.split("\n").map { |member| member.split("") }.inject(:&).length
end
p sums.sum