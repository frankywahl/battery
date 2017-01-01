#!/usr/bin/env ruby
require "csv"
require "observer"

class Datapoint
  include Observable

  module Refinements
    module Time
      refine ::Time do
        def human
          strftime("%H:%M%P")
        end
      end
    end

    module Numeric
      refine ::Numeric do
        def to_time
          ::Time.at(self).utc.strftime("%H:%M:%S")
        end
      end
    end
  end

  using Refinements::Time
  using Refinements::Numeric

  attr_reader :percentage, :change_time

  def initialize(start_time:)
    @percentage = 101
    @start_time = start_time
  end

  def percentage=(new_percentage)
    return @percentage if @percentage == new_percentage
    @percentage = new_percentage
    @change_time = Time.now
    changed
    notify_observers(self)
  end

  def elapsed_time
    @change_time - @start_time
  end

  def human_elapsed_time
    elapsed_time.to_time
  end
end

class StandardOutObserver
  def initialize
    puts ["Percentage", "Elapsed Time", "Time: Human readable"].join(",")
  end

  def update(datapoint)
    puts "#{datapoint.percentage}, #{datapoint.elapsed_time.to_i}, #{datapoint.human_elapsed_time}"
  end
end

class FileObserver
  attr_reader :filename
  def initialize(filename)
    require "csv"
    @filename = filename
    CSV.open(filename, "w") do |csv|
      csv << ["Percentage", "Time", "Time: Human readable"]
    end
  end

  def update(datapoint)
    CSV.open(filename, "a+") do |f|
      f << [datapoint.percentage, datapoint.elapsed_time.to_i, datapoint.human_elapsed_time]
    end
  end
end

datapoint = Datapoint.new(start_time: Time.now)
datapoint.add_observer(StandardOutObserver.new)
datapoint.add_observer(FileObserver.new(ARGV[0])) if ARGV[0]

loop do
  new_percentage = `pmset -g batt`.strip.match(/[0-9]*%/).to_s.to_i
  datapoint.percentage = new_percentage
  sleep 1
  break if datapoint.percentage <= 2
end
