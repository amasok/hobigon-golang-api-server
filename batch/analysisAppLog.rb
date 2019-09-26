file = File.open("#{__dir__}/../logs/app.log", mode = "r")

INDEX_PREFIX = 2
INDEX_METHOD = 3
INDEX_ENDPOINT = 4

requests = {}

file.each_line do |line|
  req = line.split(' ')

  if req[INDEX_PREFIX] == '[AccessLog]'
    if requests[req[INDEX_ENDPOINT]] == nil
      requests[req[INDEX_ENDPOINT]] = 1
    else
      requests[req[INDEX_ENDPOINT]] += 1
    end
  end
end

file.close

output = ""
output <<  "\nアクセスランキング"
requests.sort.reverse.each_with_index do |req, i|
  output << "\n#{i + 1}位： #{req[0]}"
end

puts output
