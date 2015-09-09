# Ruby
	
## Remove empty go files
	Dir.glob("**/*.go").select { |f| File.read(f).split("\n").length == 1 }.each { |f| `rm #{f} ` }
