#!/usr/bin/env ruby
require 'minitest/autorun'
require 'shellwords'

class Test < Minitest::Test
  def test_split
    assert_equal "a test".shellsplit, ["a", "test"]
    assert_equal %(a "quoted string" test).shellsplit, ["a", "quoted string", "test"]
    assert_equal %(a "quoted string" test).shellsplit, ["a", "quoted string", "test"]
    assert_equal %(a "quoted \"with escape\" string" test).shellsplit, ["a", "quoted string", "test"]
  end
end
