#! /usr/bin/env ruby

BASE58 = '123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz'

def perm a
  if a.size == 1
    [a]
  else
    rv = []
    a.each_with_index do |e, i|
      perm(a[0 ... i] + a[i + 1 ... a.size]).each do |p|
        rv << [e, *p]
      end
    end
    rv
  end
end

lst = (0 .. 6).map { [] }

ARGF.readlines.each do |l|
  b = l.chomp.chars.map(&:to_i)
  lst[0] << b[0] << b[1]
  lst[1] << b[2] << b[3]
  lst[2] << b[4] << b[5]
  lst[3] << b[6] << b[6]
  lst[4] << b[7] << b[8]
  lst[5] << b[9] << b[10]
  lst[6] << b[11] << b[12]
end

str = []
0.upto(50).each do |i|
  li = (i - 1) % 51
  ri = (i + 1) % 51

  bits = []
  (0 .. 2).each do |j|
    if lst[j][li * 2 + 1] + lst[j][ri * 2] != 1
      puts "-- #{li * 2 + 1} #{ri * 2}"
      raise "#{lst[j][li * 2 + 1]} + #{lst[j][ri * 2]}"
    end
    bits << lst[j][ri * 2]
  end

  #bits << (lst[3][i * 2] ^ 1)
  bits << lst[3][i * 2]

  (4 .. 6).each do |j|
    if lst[j][li * 2 + 1] + lst[j][ri * 2] != 1
      puts "-- #{li * 2 + 1} #{ri * 2}"
      raise "#{lst[j][li * 2 + 1]} + #{lst[j][ri * 2]}"
    end
    bits << lst[j][ri * 2]
  end

  if bits[3] == 1
    bits = bits.map { |x| x ^ 1 }
  else
    bits[3] = 1
  end

  str << (bits.map(&:to_s).join.to_i(2))
end

s = str.map(&:chr).join

i = s.index('5')

puts(s[i ... s.size] + s[0 ... i])
puts((s[i + 1 ... s.size] + s[0 .. i]).reverse)
