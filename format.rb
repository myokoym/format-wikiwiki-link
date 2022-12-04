#!/usr/bin/env ruby

readlines.each do |line|
  /(.+) - YouTube (https.+)(\?list=.*)/.match(line)
  puts "-2022/0x/xx &color(Red){■};[[#{$1}>#{$2}#{$3}]]"
  #/(.+) - YouTube (https.+)/.match(line)
  #puts "-&color(Red){■};[[#{$1}>#{$2}]]"
end
