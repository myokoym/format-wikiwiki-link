#!/usr/bin/env ruby

readlines.each do |line|
  /(.+) - YouTube (https.+)(&list=.*)/.match(line)
  puts "-2022/0x/xx &color(Red){■};[[#{$1}>#{$2}]]"
end
