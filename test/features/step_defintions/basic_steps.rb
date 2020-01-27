require 'fileutils'
require 'tmpdir'

BIN_LOCATION = ENV.fetch('ENVGARDE_TEST_BINARY_LOCATION',
                         File.join(Dir.pwd, '..', 'envgarde'))
WORKING_DIR = ENV.fetch('ENVGARDE_TEST_WORKSPACE_DIR', Dir.tmpdir)
FIXTURE_DIR = ENV.fetch('ENVGARDE_TEST_FIXTURE_DIR',
                        File.join(Dir.pwd, 'fixtures'))

Before do
  reset_working_dir
end

def reset_working_dir
  Dir.chdir(WORKING_DIR)
  if File.exist?(workspace_dir)
    FileUtils.rm_rf(workspace_dir)
  end
  Dir.mkdir(workspace_dir)
end

def workspace_dir
  File.join(WORKING_DIR, 'workspace')
end

def use_config_file(file_name)
  src = File.join(FIXTURE_DIR, 'config', file_name)
  ext = File.extname(src)
  dest = File.join(workspace_dir, ".envgarde#{ext}")
  FileUtils.copy_file(src, dest)
end

def run_envgarde(args = '')
  Dir.chdir(workspace_dir)
  @command_result = `#{BIN_LOCATION} #{args}`
  puts @command_result
  @exit_code = $?.exitstatus
end

Given(/^there is no \.envgarde file$/) do
  reset_working_dir
end

When(/^envgarde is run without arguments$/) do
  run_envgarde
end

Then(/^it's exit code is (\d+)$/) do |exit_code|
  expect(@exit_code).to eq(exit_code.to_i)
end

Given(/^there is a (.*) config file$/) do |file_name|
  reset_working_dir
  use_config_file(file_name)
end

Given(/^the (.*) environment variable is not set$/) do |env_var_name|
  ENV.delete(env_var_name) if ENV[env_var_name]
end

Given(/^the (.*) environment variable is set to (.*)$/) do |env_var_name, env_var_value|
  ENV[env_var_name] = env_var_value
end

When(/^envgarde is run with the arguments '(.*)'$/) do |args|
  run_envgarde(args)
end

Then(/^it's output includes '(.*)'$/) do |included_output|
  expect(@command_result).to include(included_output)
end
